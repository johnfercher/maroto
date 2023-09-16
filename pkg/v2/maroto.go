package v2

import (
	"bytes"
	"github.com/f-amaral/go-async/async"
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

func (d *maroto) ForceAddPage(pages ...domain.Page) {
	d.pages = append(d.pages, pages...)
}

func (d *maroto) Add(rows ...domain.Row) {
	d.addRows(rows...)
}

func (d *maroto) Generate() (domain.Document, error) {
	d.fillPage()

	if d.config.Workers > 0 && d.config.ProviderType != provider.HTML {
		return d.generateConcurrently()
	}

	return d.generate()
}

func (d *maroto) GetStructure() *tree.Node[domain.Structure] {
	d.fillPage()

	str := domain.Structure{
		Type: "maroto",
	}
	node := tree.NewNode(str)

	for _, p := range d.pages {
		inner := p.GetStructure()
		node.AddNext(inner)
	}

	return node
}

func (d *maroto) addRows(rows ...domain.Row) {
	for _, row := range rows {
		d.addRow(row)
	}
}

func (d *maroto) addRow(r domain.Row) {
	maxHeight := d.cell.Height

	height := r.GetHeight()
	sumHeight := height + d.currentHeight

	// Row smaller than the remain space on page
	if sumHeight < maxHeight {
		d.currentHeight += height
		d.rows = append(d.rows, r)
		return
	}

	// As row will extrapolate page, we will the empty space
	// on the page to force a new page
	d.fillPage()

	// Add row on the new page
	d.currentHeight += height
	d.rows = append(d.rows, r)
}

func (d *maroto) fillPage() {
	space := d.cell.Height - d.currentHeight

	p := page.New()
	p.SetNumber(len(d.pages))
	p.Add(d.rows...)

	c := col.New(12)
	row := row.New(space, color.Color{255, 0, 0})
	row.Add(c)
	p.Add(row)

	d.pages = append(d.pages, p)
	d.rows = nil
	d.currentHeight = 0
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

func (d *maroto) generate() (domain.Document, error) {
	innerCtx := d.cell.Copy()

	for _, page := range d.pages {
		page.Render(d.provider, innerCtx)
	}

	bytes, err := d.provider.GenerateBytes()
	if err != nil {
		return nil, err
	}

	return domain.NewDocument(bytes, nil), nil
}

func (d *maroto) generateConcurrently() (domain.Document, error) {
	chunks := len(d.pages) / d.config.Workers
	if chunks == 0 {
		chunks = 1
	}
	pageGroups := make([][]domain.Page, 0)
	for i := 0; i < len(d.pages); i += chunks {
		end := i + chunks

		if end > len(d.pages) {
			end = len(d.pages)
		}

		pageGroups = append(pageGroups, d.pages[i:end])
	}

	processed := d.pool.Process(pageGroups)
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

func mergePdfs(readers []io.ReadSeeker, writer io.Writer) error {
	conf := api.LoadConfiguration()
	conf.WriteXRefStream = false
	return api.MergeRaw(readers, writer, conf)
}

func (d *maroto) processPage(pages []domain.Page) ([]byte, error) {
	innerCtx := d.cell.Copy()
	innerProvider := getProvider(d.imageCache, d.config)
	for _, page := range pages {
		page.Render(innerProvider, innerCtx)
	}

	return innerProvider.GenerateBytes()
}
