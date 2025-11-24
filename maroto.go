package maroto

import (
	"errors"
	"math"

	"github.com/f-amaral/go-async/pool"
	"github.com/johnfercher/maroto/v2/pkg/consts/generation"

	"github.com/johnfercher/maroto/v2/internal/cache"

	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf"

	"github.com/johnfercher/maroto/v2/pkg/merge"

	"github.com/johnfercher/maroto/v2/pkg/core/entity"

	"github.com/johnfercher/go-tree/node"

	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/page"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/core"
)

type Maroto struct {
	config   *entity.Config
	provider core.Provider
	cache    cache.Cache

	// Building
	cell          entity.Cell
	pages         []core.Page
	rows          []core.Row
	header        []core.Row
	footer        []core.Row
	headerHeight  float64
	footerHeight  float64
	currentHeight float64
}

// GetCurrentConfig is responsible for returning the current settings from the file
func (m *Maroto) GetCurrentConfig() *entity.Config {
	return m.config
}

// New is responsible for create a new instance of core.Maroto.
// It's optional to provide an *entity.Config with customizations
// those customization are created by using the config.Builder.
func New(cfgs ...*entity.Config) core.Maroto {
	cache := cache.New()
	cfg := getConfig(cfgs...)
	provider := getProvider(cache, cfg)

	m := &Maroto{
		provider: provider,
		cell: entity.NewRootCell(cfg.Dimensions.Width, cfg.Dimensions.Height, entity.Margins{
			Left:   cfg.Margins.Left,
			Top:    cfg.Margins.Top,
			Right:  cfg.Margins.Right,
			Bottom: cfg.Margins.Bottom,
		}),
		cache:  cache,
		config: cfg,
	}

	return m
}

// AddPages is responsible for add pages directly in the document.
// By adding a page directly, the current cursor will reset and the
// new page will appear as the next. If the page provided have
// more rows than the maximum useful area of a page, maroto will split
// that page in more than one.
func (m *Maroto) AddPages(pages ...core.Page) {
	for _, page := range pages {
		if m.currentHeight != m.headerHeight {
			m.fillPageToAddNew()
			m.addHeader()
		}
		m.addRows(page.GetRows()...)
	}
}

// AddRows is responsible for add rows in the current document.
// By adding a row, if the row will extrapolate the useful area of a page,
// maroto will automatically add a new page. Maroto use the information of
// PageSize, PageMargin, FooterSize and HeaderSize to calculate the useful
// area of a page.
func (m *Maroto) AddRows(rows ...core.Row) {
	m.addRows(rows...)
}

// AddRow is responsible for add one row in the current document.
// By adding a row, if the row will extrapolate the useful area of a page,
// maroto will automatically add a new page. Maroto use the information of
// PageSize, PageMargin, FooterSize and HeaderSize to calculate the useful
// area of a page.
func (m *Maroto) AddRow(rowHeight float64, cols ...core.Col) core.Row {
	r := row.New(rowHeight).Add(cols...)
	m.addRow(r)
	return r
}

// AddAutoRow is responsible for adding a line with automatic height to the
// current document.
// The row height will be calculated based on its content.
func (m *Maroto) AddAutoRow(cols ...core.Col) core.Row {
	r := row.New().Add(cols...)
	m.addRow(r)
	return r
}

// FitlnCurrentPage is responsible to validating whether a line fits on
// the current page.
func (m *Maroto) FitlnCurrentPage(heightNewLine float64) bool {
	contentSize := m.getRowsHeight(m.rows...) + m.footerHeight + m.headerHeight
	return contentSize+heightNewLine < m.cell.Height
}

// RegisterHeader is responsible to define a set of rows as a header
// of the document. The header will appear in every new page of the document.
// The header cannot occupy an area greater than the useful area of the page,
// it this case the method will return an error.
func (m *Maroto) RegisterHeader(rows ...core.Row) error {
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

// RegisterFooter is responsible to define a set of rows as a footer
// of the document. The footer will appear in every new page of the document.
// The footer cannot occupy an area greater than the useful area of the page,
// it this case the method will return an error.
func (m *Maroto) RegisterFooter(rows ...core.Row) error {
	height := m.getRowsHeight(rows...)
	if height > m.config.Dimensions.Height {
		return errors.New("footer height is greater than page useful area")
	}

	m.footerHeight = height
	m.footer = rows
	return nil
}

// Generate is responsible to compute the component tree created by
// the usage of all other Maroto methods, and generate the PDF document.
func (m *Maroto) Generate() (core.Document, error) {
	m.fillPageToAddNew()
	m.setConfig()

	if m.config.GenerationMode == generation.Concurrent {
		return m.generateConcurrently()
	}

	if m.config.GenerationMode == generation.SequentialLowMemory {
		return m.generateLowMemory()
	}

	return m.generate()
}

// GetStructure is responsible for return the component tree, this is useful
// on unit tests cases.
func (m *Maroto) GetStructure() *node.Node[core.Structure] {
	m.fillPageToAddNew()

	str := core.Structure{
		Type:    "maroto",
		Details: m.config.ToMap(),
	}
	node := node.New(str)

	for _, p := range m.pages {
		inner := p.GetStructure()
		node.AddNext(inner)
	}

	return node
}

func (m *Maroto) addRows(rows ...core.Row) {
	for _, row := range rows {
		m.addRow(row)
	}
}

func (m *Maroto) addRow(r core.Row) {
	if len(r.GetColumns()) == 0 {
		r.Add(col.New())
	}

	maxHeight := m.cell.Height

	r.SetConfig(m.config)
	rowHeight := r.GetHeight(m.provider, &m.cell)
	sumHeight := rowHeight + m.currentHeight + m.footerHeight

	// Row smaller than the remain space on page
	if sumHeight <= maxHeight {
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

func (m *Maroto) addHeader() {
	for _, headerRow := range m.header {
		m.currentHeight += headerRow.GetHeight(m.provider, &m.cell)
		m.rows = append(m.rows, headerRow)
	}
}

func (m *Maroto) fillPageToAddNew() {
	space := m.cell.Height - m.currentHeight - m.footerHeight

	// Truncate space to 9 decimal places to avoid rounding errors
	space = math.Floor(space*math.Pow10(9)) / math.Pow10(9)

	c := col.New(m.config.MaxGridSize)
	spaceRow := row.New(space)
	spaceRow.Add(c)

	m.rows = append(m.rows, spaceRow)
	m.rows = append(m.rows, m.footer...)

	var p core.Page
	if m.config.PageNumber != nil {
		p = page.New(*m.config.PageNumber)
	} else {
		p = page.New()
	}

	p.SetConfig(m.config)
	p.Add(m.rows...)

	m.pages = append(m.pages, p)
	m.rows = nil
	m.currentHeight = 0
}

func (m *Maroto) setConfig() {
	for i, page := range m.pages {
		page.SetConfig(m.config)
		page.SetNumber(i+1, len(m.pages))
	}
}

func (m *Maroto) generate() (core.Document, error) {
	innerCtx := m.cell.Copy()

	for _, page := range m.pages {
		page.Render(m.provider, innerCtx)
	}

	documentBytes, err := m.provider.GenerateBytes()
	if err != nil {
		return nil, err
	}

	return core.NewPDF(documentBytes, nil), nil
}

func (m *Maroto) generateConcurrently() (core.Document, error) {
	p := pool.NewPool[[]core.Page, []byte](m.config.ChunkWorkers, m.processPage,
		pool.WithSortingOutput[[]core.Page, []byte]())
	defer p.Close()
	chunks := len(m.pages) / m.config.ChunkWorkers
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

	processed := p.Process(pageGroups)
	if processed.HasError {
		return nil, errors.New("an error has occurred while trying to generate PDFs concurrently")
	}

	pdfs := make([][]byte, len(processed.Results))
	for i, result := range processed.Results {
		bytes := result.Output.([]byte)
		pdfs[i] = bytes
	}

	mergedBytes, err := merge.Bytes(pdfs...)
	if err != nil {
		return nil, err
	}

	return core.NewPDF(mergedBytes, nil), nil
}

func (m *Maroto) generateLowMemory() (core.Document, error) {
	chunks := len(m.pages) / m.config.ChunkWorkers
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

	var pdfResults [][]byte
	for _, pageGroup := range pageGroups {
		bytes, err := m.processPage(pageGroup)
		if err != nil {
			return nil, errors.New("an error has occurred while trying to generate PDFs in low memory mode")
		}

		pdfResults = append(pdfResults, bytes)
	}

	mergedBytes, err := merge.Bytes(pdfResults...)
	if err != nil {
		return nil, err
	}

	return core.NewPDF(mergedBytes, nil), nil
}

func (m *Maroto) processPage(pages []core.Page) ([]byte, error) {
	innerCtx := m.cell.Copy()

	innerProvider := getProvider(cache.NewMutexDecorator(cache.New()), m.config)
	for _, page := range pages {
		page.Render(innerProvider, innerCtx)
	}

	return innerProvider.GenerateBytes()
}

func (m *Maroto) getRowsHeight(rows ...core.Row) float64 {
	var height float64
	for _, r := range rows {
		r.SetConfig(m.config)
		height += r.GetHeight(m.provider, &m.cell)
	}

	return height
}

func getConfig(configs ...*entity.Config) *entity.Config {
	if len(configs) > 0 {
		return configs[0]
	}

	return config.NewBuilder().Build()
}

func getProvider(cache cache.Cache, cfg *entity.Config) core.Provider {
	deps := gofpdf.NewBuilder().Build(cfg, cache)
	provider := gofpdf.New(deps)
	provider.SetMetadata(cfg.Metadata)
	provider.SetCompression(cfg.Compression)
	provider.SetProtection(cfg.Protection)
	return provider
}
