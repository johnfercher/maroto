package pkg

import (
	"bytes"
	"errors"
	"github.com/johnfercher/go-tree/node"
	"io"
	"log"

	"github.com/johnfercher/maroto/v2/pkg/props"

	"github.com/johnfercher/maroto/v2/pkg/cache"

	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/page"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/consts/provider"
	"github.com/johnfercher/maroto/v2/pkg/providers/gofpdf"
	"github.com/johnfercher/maroto/v2/pkg/providers/html"

	"github.com/f-amaral/go-async/async"
	"github.com/f-amaral/go-async/pool"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/providers"
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

type maroto struct {
	config     *config.Config
	provider   core.Provider
	imageCache cache.Cache

	// Building
	cell          core.Cell
	pages         []core.Page
	rows          []core.Row
	header        []core.Row
	footer        []core.Row
	pdfs          [][]byte
	headerHeight  float64
	footerHeight  float64
	currentHeight float64

	// Processing
	pool async.Processor[[]core.Page, []byte]
}

func NewMaroto(cfgs ...*config.Config) core.Maroto {
	cache := cache.New()
	cfg := getConfig(cfgs...)
	provider := getProvider(cache, cfg)

	width, height := provider.GetDimensions()
	left, top, right, bottom := provider.GetMargins()

	m := &maroto{
		provider: provider,
		cell: core.NewRootContext(width, height, config.Margins{
			Left:   left,
			Top:    top,
			Right:  right,
			Bottom: bottom,
		}),
		imageCache: cache,
		config:     cfg,
	}

	if cfg.Workers > 0 {
		p := pool.NewPool(cfg.Workers, m.processPage)
		m.pool = p
	}
	return m
}

func (m *maroto) AddPages(pages ...core.Page) {
	for _, page := range pages {
		if m.currentHeight != m.headerHeight {
			m.fillPageToAddNew()
			m.addHeader()
		}
		m.addRows(page.GetRows()...)
	}
}

func (m *maroto) AddRows(rows ...core.Row) {
	m.addRows(rows...)
}

func (m *maroto) AddRow(rowHeight float64, cols ...core.Col) core.Row {
	r := row.New(rowHeight).Add(cols...)
	m.addRow(r)
	return r
}

func (m *maroto) RegisterHeader(rows ...core.Row) error {
	height := m.getRowsHeight(rows...)
	if height+m.footerHeight > m.config.Dimensions.Height {
		return errors.New("header height is greater than page useful area")
	}

	m.headerHeight = height
	m.header = rows

	for _, headerRow := range rows {
		m.addRow(headerRow)
	}

	return nil
}

func (m *maroto) RegisterFooter(rows ...core.Row) error {
	height := m.getRowsHeight(rows...)
	if height > m.config.Dimensions.Height {
		return errors.New("footer height is greater than page useful area")
	}

	m.footerHeight = height
	m.footer = rows
	return nil
}

func (m *maroto) Generate() (core.Document, error) {
	m.provider.SetProtection(m.config.Protection)
	m.provider.SetCompression(m.config.Compression)
	m.provider.SetMetadata(m.config.Metadata)

	m.fillPageToAddNew()
	m.setConfig()

	if m.config.Workers > 0 && m.config.ProviderType != provider.HTML {
		return m.generateConcurrently()
	}

	return m.generate()
}

func (m *maroto) GetStructure() *node.Node[core.Structure] {
	m.fillPageToAddNew()

	str := core.Structure{
		Type: "pkg",
	}
	node := node.New(str)

	for _, p := range m.pages {
		inner := p.GetStructure()
		node.AddNext(inner)
	}

	return node
}

func (m *maroto) AddPDFs(pdfs ...[]byte) {
	m.pdfs = append(m.pdfs, pdfs...)
}

func (m *maroto) addRows(rows ...core.Row) {
	for _, row := range rows {
		m.addRow(row)
	}
}

func (m *maroto) addRow(r core.Row) {
	maxHeight := m.cell.Height

	rowHeight := r.GetHeight()
	sumHeight := rowHeight + m.currentHeight + m.footerHeight

	// Row smaller than the remain space on page
	if sumHeight < maxHeight {
		m.currentHeight += rowHeight
		m.rows = append(m.rows, r)
		return
	}

	// As row will extrapolate page, we will add empty space
	// on the page to force a new page
	m.fillPageToAddNew()

	m.addHeader()

	// AddRows row on the new page
	m.currentHeight += rowHeight
	m.rows = append(m.rows, r)
}

func (m *maroto) addHeader() {
	for _, headerRow := range m.header {
		m.currentHeight += headerRow.GetHeight()
		m.rows = append(m.rows, headerRow)
	}
}

func (m *maroto) fillPageToAddNew() {
	space := m.cell.Height - m.currentHeight - m.footerHeight

	c := col.New(m.config.MaxGridSize)
	spaceRow := row.New(space)
	spaceRow.Add(c)

	m.rows = append(m.rows, spaceRow)
	m.rows = append(m.rows, m.footer...)

	prop := props.Page{
		Pattern: m.config.PageNumberPattern,
		Place:   m.config.PageNumberPlace,
		Family:  m.config.DefaultFont.Family,
		Style:   m.config.DefaultFont.Style,
		Size:    m.config.DefaultFont.Size,
		Color:   m.config.DefaultFont.Color,
	}
	p := page.New(prop)
	p.Add(m.rows...)

	m.pages = append(m.pages, p)
	m.rows = nil
	m.currentHeight = 0
}

func (m *maroto) setConfig() {
	for i, page := range m.pages {
		page.SetConfig(m.config)
		page.SetNumber(i+1, len(m.pages))
	}
}

func (m *maroto) generate() (core.Document, error) {
	innerCtx := m.cell.Copy()

	for _, page := range m.pages {
		page.Render(m.provider, innerCtx)
	}

	documentBytes, err := m.provider.GenerateBytes()
	if err != nil {
		return nil, err
	}

	if len(m.pdfs) == 0 {
		return core.NewDocument(documentBytes, nil), nil
	}

	readers := []io.ReadSeeker{}
	readers = append(readers, bytes.NewReader(documentBytes))
	for _, pdf := range m.pdfs {
		readers = append(readers, bytes.NewReader(pdf))
	}

	var buf bytes.Buffer
	writer := io.Writer(&buf)
	err = mergePdfs(readers, writer)
	if err != nil {
		return nil, err
	}

	return core.NewDocument(buf.Bytes(), nil), nil
}

func (m *maroto) generateConcurrently() (core.Document, error) {
	chunks := len(m.pages) / m.config.Workers
	if chunks == 0 {
		chunks = 1
	}
	pageGroups := make([][]core.Page, 0)
	for i := 0; i < len(m.pages); i += chunks {
		end := i + chunks

		if end > len(m.pages) {
			end = len(m.pages)
		}

		pageGroups = append(pageGroups, m.pages[i:end])
	}

	processed := m.pool.Process(pageGroups)
	if processed.HasError {
		log.Fatal("error on generating pages")
	}

	readers := make([]io.ReadSeeker, len(processed.Results)+len(m.pdfs))
	for i, result := range processed.Results {
		b := result.Output.([]byte)
		readers[i] = bytes.NewReader(b)
	}

	for i, pdf := range m.pdfs {
		readers[i+len(processed.Results)] = bytes.NewReader(pdf)
	}

	var buf bytes.Buffer
	writer := io.Writer(&buf)
	err := mergePdfs(readers, writer)
	if err != nil {
		return nil, err
	}

	return core.NewDocument(buf.Bytes(), nil), nil
}

func (m *maroto) processPage(pages []core.Page) ([]byte, error) {
	innerCtx := m.cell.Copy()
	innerProvider := getProvider(m.imageCache, m.config)
	for _, page := range pages {
		page.Render(innerProvider, innerCtx)
	}

	return innerProvider.GenerateBytes()
}

func (m *maroto) getRowsHeight(rows ...core.Row) float64 {
	var height float64
	for _, r := range rows {
		height += r.GetHeight()
	}

	return height
}

func getConfig(configs ...*config.Config) *config.Config {
	if len(configs) > 0 {
		return configs[0]
	}

	return config.NewBuilder().Build()
}

func getProvider(cache cache.Cache, cfg *config.Config) core.Provider {
	if cfg.ProviderType == provider.HTML {
		return html.New(cfg, providers.WithCache(cache))
	}

	return gofpdf.New(cfg, providers.WithCache(cache))
}

func mergePdfs(readers []io.ReadSeeker, writer io.Writer) error {
	conf := api.LoadConfiguration()
	conf.WriteXRefStream = false
	return api.MergeRaw(readers, writer, conf)
}
