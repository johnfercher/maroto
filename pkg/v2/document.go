package v2

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/v2/col"
	"github.com/johnfercher/maroto/pkg/v2/context"
	"github.com/johnfercher/maroto/pkg/v2/domain"
	"github.com/johnfercher/maroto/pkg/v2/page"
	"github.com/johnfercher/maroto/pkg/v2/row"
	"github.com/johnfercher/maroto/pkg/v2/types"
)

type document struct {
	file          string
	cell          internal.Cell
	_type         types.DocumentType
	provider      domain.Provider
	pages         []domain.Page
	rows          []domain.Row
	currentHeight float64
}

func NewMaroto(provider domain.Provider, file string) *document {
	width, height := provider.GetDimensions()
	left, top, right, bottom := provider.GetMargins()

	return &document{
		file:     file,
		provider: provider,
		_type:    types.Document,
		cell: context.NewRootContext(width, height, context.Margins{
			Left:   left,
			Top:    top,
			Right:  right,
			Bottom: bottom,
		}),
	}
}

func (d *document) ForceAddPage(pages ...domain.Page) {
	d.pages = append(d.pages, pages...)
}

func (d *document) Add(rows ...domain.Row) {
	d.addRows(rows...)
}

func (d *document) Generate() error {
	d.fillPage()
	innerCtx := d.cell.Copy()

	for _, page := range d.pages {
		//fmt.Printf("render page: %d\n", page.GetNumber())
		page.Render(d.provider, innerCtx)
	}

	return d.provider.Generate(d.file)
}

/*func (d *document) Generate() error {
	d.fillPage()
	innerCtx := d.cell.Copy()

	p := pool.NewPool(10, func(i domain.Page) (bytes.Buffer, error) {
		innerProvider := providers.NewGofpdf(size.A4)
		i.Render(innerProvider, innerCtx)
		return d.provider.GenerateAndOutput()
	})

	processed := p.Process(d.pages)
	if processed.HasError {
		log.Fatal("error on generating pages")
	}
	readers := make([]io.ReadSeeker, len(processed.Results))
	for i, result := range processed.Results {
		buffer := result.Output.(bytes.Buffer)
		readers[i] = bytes.NewReader(buffer.Bytes())
	}
	writer, _ := os.Create(d.file)
	conf := api.LoadConfiguration()
	conf.CreateBookmarks = false
	conf.WriteXRefStream = false
	err := api.MergeRaw(readers, writer, conf)
	if err != nil {
		return err
	}
	return d.provider.Generate(d.file)
}*/

func (d *document) GetStructure() *tree.Node[domain.Structure] {
	str := domain.Structure{
		Type:  string(d._type),
		Value: d.file,
	}
	node := tree.NewNode(str)

	for _, r := range d.rows {
		inner := r.GetStructure()
		node.AddNext(inner)
	}

	return node
}

func (d *document) addRows(rows ...domain.Row) {
	for _, row := range rows {
		d.addRow(row)
	}
}

func (d *document) addRow(r domain.Row) {
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

func (d *document) fillPage() {
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
