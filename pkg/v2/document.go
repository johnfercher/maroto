package v2

import (
	"fmt"
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
	file     string
	cell     internal.Cell
	_type    types.DocumentType
	provider domain.Provider
	pages    []domain.Page
	rows     []domain.Row
}

func NewDocument(provider domain.Provider, file string) *document {
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
	d.rows = append(d.rows, rows...)
}

func (d *document) Generate() error {
	//d.ctx.Print(d._type)

	maxHeight := d.cell.Height
	currentHeight := 0.0
	var buf []domain.Row
	for _, dRow := range d.rows {
		height := dRow.GetHeight()
		sumHeight := height + currentHeight
		if sumHeight >= maxHeight {
			p := page.New()
			p.Add(buf...)

			c := col.New(12)
			lastRowHeight := maxHeight - currentHeight
			r := row.New(lastRowHeight, color.Color{255, 0, 0})
			r.Add(c)
			p.Add(r)

			d.pages = append(d.pages, p)
			buf = nil
			currentHeight = 0
		}

		currentHeight += height
		buf = append(buf, dRow)
	}

	p := page.New()
	p.Add(buf...)
	d.pages = append(d.pages, p)

	innerCtx := d.cell.Copy()
	for _, page := range d.pages {
		page.Render(d.provider, innerCtx)
	}

	fmt.Println(len(d.pages))

	/*for _, dRow := range d.rows {
		dRow.Render(d.fpdf, ctx)
	}

	for _, page := range d.pages {
		page.Render(d.fpdf, ctx)
	}*/

	return d.provider.Generate(d.file)
}

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
