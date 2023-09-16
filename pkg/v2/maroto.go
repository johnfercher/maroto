package v2

import (
	"bytes"
	"errors"
	"github.com/f-amaral/go-async/pool"
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/v2/cache"
	"github.com/johnfercher/maroto/pkg/v2/config"
	"github.com/johnfercher/maroto/pkg/v2/context"
	"github.com/johnfercher/maroto/pkg/v2/domain"
	"github.com/johnfercher/maroto/pkg/v2/grid/col"
	"github.com/johnfercher/maroto/pkg/v2/grid/page"
	"github.com/johnfercher/maroto/pkg/v2/grid/row"
	"github.com/johnfercher/maroto/pkg/v2/provider"
	"github.com/johnfercher/maroto/pkg/v2/provider/gofpdf"
	"github.com/johnfercher/maroto/pkg/v2/provider/html"
	"github.com/johnfercher/maroto/pkg/v2/providers"
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
	headerHeight  float64
	currentHeight float64
}

func NewMaroto(config ...*config.Maroto) domain.Maroto {
	cache := cache.New()
	cfg := getConfig(config...)
	provider := getProvider(cache, cfg)

	width, height := provider.GetDimensions()
	left, top, right, bottom := provider.GetMargins()

	return &maroto{
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
}

func (m *maroto) ForceAddPage(pages ...domain.Page) {
	m.pages = append(m.pages, pages...)
}

func (m *maroto) Add(rows ...domain.Row) {
	m.addRows(rows...)
}

func (m *maroto) RegisterHeader(rows ...domain.Row) error {
	var headerHeight float64
	for _, row := range rows {
		headerHeight += row.GetHeight()
	}
	if headerHeight > m.config.Dimensions.Height {
		return errors.New("header height is greater than page")
	}

	m.headerHeight = headerHeight
	m.header = rows

	for _, headerRow := range rows {
		m.addRow(headerRow)
	}

	return nil
}

func (m *maroto) Generate() (domain.Document, error) {
	m.fillPageToAddNew()

	if m.config.ThreadPool > 0 {
		return m.generate()
	}

	return m.generate()
}

func (m *maroto) GetStructure() *tree.Node[domain.Structure] {
	m.fillPageToAddNew()

	str := domain.Structure{
		Type: "maroto",
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

	height := r.GetHeight()
	sumHeight := height + m.currentHeight

	// Row smaller than the remain space on page
	if sumHeight < maxHeight {
		m.currentHeight += height
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

	// Add row on the new page
	m.currentHeight += height
	m.rows = append(m.rows, r)
}

func (m *maroto) fillPageToAddNew() {
	space := m.cell.Height - m.currentHeight

	p := page.New()
	p.SetNumber(len(m.pages))
	p.Add(m.rows...)

	c := col.New(12)
	row := row.New(space, color.Color{255, 0, 0})
	row.Add(c)
	p.Add(row)

	m.pages = append(m.pages, p)
	m.rows = nil
	m.currentHeight = 0
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

func (m *maroto) generate() (domain.Document, error) {
	innerCtx := m.cell.Copy()

	for _, page := range m.pages {
		page.Render(m.provider, innerCtx, m.config)
	}

	bytes, err := m.provider.GenerateBytes()
	if err != nil {
		return nil, err
	}

	return domain.NewDocument(bytes, nil), nil
}

func (m *maroto) generateConcurrently() (domain.Document, error) {
	innerCtx := m.cell.Copy()

	p := pool.NewPool(10, func(i domain.Page) ([]byte, error) {
		innerProvider := getProvider(m.imageCache, m.config)
		i.Render(innerProvider, innerCtx, nil)
		return innerProvider.GenerateBytes()
	})

	processed := p.Process(m.pages)
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
	if m.config.ProviderType == provider.Gofpdf {
		err := mergePdfs(readers, writer)
		if err != nil {
			return nil, err
		}
	} else {
		for _, reader := range readers {
			r := io.TeeReader(reader, &buf)
			_, err := io.ReadAll(r)
			if err != nil {
				return nil, err
			}
		}
	}

	return domain.NewDocument(buf.Bytes(), nil), nil
}

func mergePdfs(readers []io.ReadSeeker, writer io.Writer) error {
	conf := api.LoadConfiguration()
	conf.WriteXRefStream = false
	return api.MergeRaw(readers, writer, conf)
}
