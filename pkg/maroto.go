package pkg

import (
	"bytes"
	"errors"
	"github.com/f-amaral/go-async/async"
	"github.com/f-amaral/go-async/pool"
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/v2/internal"
	"github.com/johnfercher/maroto/v2/pkg/cache"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/context"
	"github.com/johnfercher/maroto/v2/pkg/domain"
	"github.com/johnfercher/maroto/v2/pkg/grid/col"
	"github.com/johnfercher/maroto/v2/pkg/grid/page"
	"github.com/johnfercher/maroto/v2/pkg/grid/row"
	"github.com/johnfercher/maroto/v2/pkg/provider"
	"github.com/johnfercher/maroto/v2/pkg/provider/gofpdf"
	"github.com/johnfercher/maroto/v2/pkg/provider/html"
	"github.com/johnfercher/maroto/v2/pkg/providers"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"io"
	"log"
)

type maroto struct {
	config     *config.Maroto
	provider   domain.Provider
	imageCache cache.Cache

	// Building
	cell          internal.Cell
	pages         []domain.Page
	rows          []domain.Row
	header        []domain.Row
	footer        []domain.Row
	headerHeight  float64
	footerHeight  float64
	currentHeight float64

	// Processing
	pool async.Processor[[]domain.Page, []byte]
}

func NewMaroto(config ...*config.Maroto) domain.Maroto {
	cache := cache.New()
	cfg := getConfig(config...)
	provider := getProvider(cache, cfg)

	width, height := provider.GetDimensions()
	left, top, right, bottom := provider.GetMargins()

	m := &maroto{
		provider: provider,
		cell: context.NewRootContext(width, height, context.Margins{
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

func (m *maroto) ForceAddPage(pages ...domain.Page) {
	m.pages = append(m.pages, pages...)
}

func (m *maroto) AddRows(rows ...domain.Row) {
	m.addRows(rows...)
}

func (m *maroto) AddRow(rowHeight float64, cols ...domain.Col) domain.Row {
	r := row.New(rowHeight).Add(cols...)
	m.addRow(r)
	return r
}

func (m *maroto) RegisterHeader(rows ...domain.Row) error {
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

func (m *maroto) RegisterFooter(rows ...domain.Row) error {
	height := m.getRowsHeight(rows...)
	if height > m.config.Dimensions.Height {
		return errors.New("footer height is greater than page useful area")
	}

	m.footerHeight = height
	m.footer = rows
	return nil
}

func (m *maroto) Generate() (domain.Document, error) {
	m.fillPageToAddNew()
	m.setConfig()

	if m.config.Workers > 0 && m.config.ProviderType != provider.HTML {
		return m.generateConcurrently()
	}

	return m.generate()
}

func (m *maroto) GetStructure() *tree.Node[domain.Structure] {
	m.fillPageToAddNew()

	str := domain.Structure{
		Type: "pkg",
	}
	node := tree.NewNode(str)

	for _, p := range m.pages {
		inner := p.GetStructure()
		node.AddNext(inner)
	}

	return node
}

func (m *maroto) addRows(rows ...domain.Row) {
	for _, row := range rows {
		m.addRow(row)
	}
}

func (m *maroto) addRow(r domain.Row) {
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

	for _, headerRow := range m.header {
		m.currentHeight += headerRow.GetHeight()
		m.rows = append(m.rows, headerRow)
	}

	// AddRows row on the new page
	m.currentHeight += rowHeight
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

func (m *maroto) generate() (domain.Document, error) {
	innerCtx := m.cell.Copy()

	for _, page := range m.pages {
		page.Render(m.provider, innerCtx)
	}

	bytes, err := m.provider.GenerateBytes()
	if err != nil {
		return nil, err
	}

	return domain.NewDocument(bytes, nil), nil
}

func (m *maroto) generateConcurrently() (domain.Document, error) {
	chunks := len(m.pages) / m.config.Workers
	if chunks == 0 {
		chunks = 1
	}
	pageGroups := make([][]domain.Page, 0)
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

	return domain.NewDocument(buf.Bytes(), nil), nil
}

func (m *maroto) processPage(pages []domain.Page) ([]byte, error) {
	innerCtx := m.cell.Copy()
	innerProvider := getProvider(m.imageCache, m.config)
	for _, page := range pages {
		page.Render(innerProvider, innerCtx)
	}

	return innerProvider.GenerateBytes()
}

func (m *maroto) getRowsHeight(rows ...domain.Row) float64 {
	var height float64
	for _, r := range rows {
		height += r.GetHeight()
	}

	return height
}

func getConfig(configs ...*config.Maroto) *config.Maroto {
	if len(configs) > 0 {
		return configs[0]
	}

	return config.NewBuilder().Build()
}

func getProvider(cache cache.Cache, cfg *config.Maroto) domain.Provider {
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
