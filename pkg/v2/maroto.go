package v2

import (
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
)

type maroto struct {
	cell          internal.Cell
	provider      domain.Provider
	pages         []domain.Page
	rows          []domain.Row
	currentHeight float64
	imageCache    cache.Cache
	config        []config.Builder
}

func NewMaroto(config ...config.Builder) config.Maroto {
	cache := cache.New()
	provider := getProvider(cache, config...)

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
		config:     config,
	}
}

func (d *maroto) ForceAddPage(pages ...domain.Page) {
	d.pages = append(d.pages, pages...)
}

func (d *maroto) Add(rows ...domain.Row) {
	d.addRows(rows...)
}

func (d *maroto) Generate() (*domain.Document, error) {
	d.fillPage()
	innerCtx := d.cell.Copy()

	for _, page := range d.pages {
		page.Render(d.provider, innerCtx)
	}

	bytes, err := d.provider.GenerateBytes()
	if err != nil {
		return nil, err
	}

	return &domain.Document{
		Bytes: bytes,
	}, nil
}

/*func (d *maroto) GenerateConcurrently() error {
	d.fillPage()
	innerCtx := d.cell.Copy()

	p := pool.NewPool(10, func(i domain.Page) ([]byte, error) {
		innerProvider := getProvider(d.imageCache, d.config...)
		i.Render(innerProvider, innerCtx)
		return innerProvider.GenerateBytes()
	})

	processed := p.Process(d.pages)
	if processed.HasError {
		log.Fatal("error on generating pages")
	}
	readers := make([]io.ReadSeeker, len(processed.Results))
	for i, result := range processed.Results {
		b := result.Output.([]byte)
		readers[i] = bytes.NewReader(b)
	}
	writer, _ := os.Create(d.file)
	defer writer.Close()

	if len(d.config) == 0 || d.config[0].GetConfig().ProviderType == provider.Gofpdf {
		err := mergePdfs(readers, writer)
		if err != nil {
			return err
		}
	}

	for _, reader := range readers {
		_, err := io.Copy(writer, reader)
		if err != nil {
			return err
		}
	}

	return writer.Close()
}*/

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

func getProvider(cache cache.Cache, builders ...config.Builder) domain.Provider {
	builder := config.NewBuilder()
	if len(builders) > 0 {
		builder = builders[0]
	}

	cfg := builder.GetConfig()

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
