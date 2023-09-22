package pkg

import (
	"bytes"
	"errors"
	"io"
	"log"

	"github.com/johnfercher/maroto/v2/pkg/cache"

	"github.com/f-amaral/go-async/async"
	"github.com/f-amaral/go-async/pool"
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/page"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/provider"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/providers"
	"github.com/johnfercher/maroto/v2/pkg/providers/gofpdf"
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
	headerHeight  float64
	footerHeight  float64
	currentHeight float64

	// Processing
	pool async.Processor[[]core.Page, []byte]
}

func NewMaroto(config ...*config.Config) core.Maroto {
	cache := cache.New()
	cfg := getConfig(config...)
	provider := getProvider(cache, cfg)

	width, height := provider.GetDimensions()
	left, top, right, bottom := provider.GetMargins()

	m := &maroto{
		provider: provider,
		cell: core.NewRootContext(width, height, core.Margins{
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

func (m *maroto) ForceAddPage(pages ...core.Page) {
	m.pages = append(m.pages, pages...)
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
	for _, r := range rows {
		r.SetConfig(m.config)
	}
	height := m.getRowsHeight(rows...)
	if height > m.config.Dimensions.Height {
		return errors.New("footer height is greater than page useful area")
	}

	m.footerHeight = height
	m.footer = rows
	return nil
}

func (m *maroto) Generate() (core.Document, error) {
	//m.fillPageToAddNew()
	m.splitPages()

	if m.config.Workers > 0 && m.config.ProviderType != provider.HTML {
		return m.generateConcurrently()
	}

	return m.generate()
}

func (m *maroto) GetStructure() *tree.Node[core.Structure] {
	m.fillPageToAddNew()

	str := core.Structure{
		Type: "pkg",
	}
	node := tree.NewNode(str)

	for _, p := range m.pages {
		inner := p.GetStructure()
		node.AddNext(inner)
	}

	return node
}

func (m *maroto) addRows(rows ...core.Row) {
	for _, r := range rows {
		m.addRow(r)
	}
}

func (m *maroto) addRow(r core.Row) {
	m.rows = append(m.rows, r)
}

func (m *maroto) fillPageToAddNew() {
	space := m.cell.Height - m.currentHeight - m.footerHeight

	c := col.New(m.config.MaxGridSize)
	spaceRow := row.New(space)
	spaceRow.Add(c)

	m.rows = append(m.rows, spaceRow)
	m.rows = append(m.rows, m.footer...)

	p := page.New()
	p.SetNumber(len(m.pages))
	p.Add(m.rows...)

	m.pages = append(m.pages, p)
	m.rows = nil
	m.currentHeight = 0
}

func (m *maroto) setConfig() {
	for _, page := range m.pages {
		page.SetConfig(m.config)
	}
}

func (m *maroto) createPage() core.Page {
	p := page.New()
	p.SetNumber(len(m.pages))

	m.pages = append(m.pages, p)
	m.currentHeight = 0

	for _, headerRow := range m.header {
		m.currentHeight += headerRow.GetHeight(m.provider, m.cell.Width)
		p.Add(headerRow)
	}

	p.SetConfig(m.config)

	return p
}

func (m *maroto) addPageFooter(p core.Page) {
	space := m.cell.Height - m.currentHeight - m.footerHeight

	er := row.New(space).Add(col.New(m.config.MaxGridSize))
	er.SetConfig(m.config)

	p.Add(er)
	p.Add(m.footer...)
}

func (m *maroto) splitPages() {
	p := m.createPage()

	maxHeight := m.cell.Height
	for _, r := range m.rows {
		r.SetConfig(m.config)
		rowHeight := r.GetHeight(m.provider, m.cell.Width)

		sumHeight := rowHeight + m.currentHeight + m.footerHeight

		// Row smaller than the remain space on page
		if sumHeight > maxHeight {
			m.addPageFooter(p)
			p = m.createPage()
		}

		m.currentHeight += rowHeight
		p.Add(r)
	}

	m.addPageFooter(p)
}

func (m *maroto) generate() (core.Document, error) {
	innerCtx := m.cell.Copy()

	for _, page := range m.pages {
		page.SetConfig(m.config)
		page.Render(m.provider, innerCtx)
	}

	bytes, err := m.provider.GenerateBytes()
	if err != nil {
		return nil, err
	}

	return core.NewDocument(bytes, nil), nil
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

	readers := make([]io.ReadSeeker, len(processed.Results))
	for i, result := range processed.Results {
		b := result.Output.([]byte)
		readers[i] = bytes.NewReader(b)
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
		height += r.GetHeight(m.provider, m.cell.Width)
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

	return gofpdf.New(cfg, providers.WithCache(cache))
}

func mergePdfs(readers []io.ReadSeeker, writer io.Writer) error {
	conf := api.LoadConfiguration()
	conf.WriteXRefStream = false
	return api.MergeRaw(readers, writer, conf)
}
