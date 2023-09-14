package page

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/internal/fpdf"
	"github.com/johnfercher/maroto/pkg/v2/context"
	"github.com/johnfercher/maroto/pkg/v2/domain"
	"github.com/johnfercher/maroto/pkg/v2/types"
)

type page struct {
	_type types.DocumentType
	rows  []domain.Row
}

func New() domain.Page {
	return &page{
		_type: types.Page,
	}
}

func (p *page) Render(fpdf fpdf.Fpdf, ctx context.Context) {
	//ctx = ctx.NewPage(fpdf.PageNo())

	innerCtx := ctx.Copy()
	for _, row := range p.rows {
		row.Render(fpdf, innerCtx)
		innerCtx.Coordinate.Y += row.GetHeight()
	}
	//fpdf.AddPage()

}

func (p *page) GetType() string {
	return p._type.String()
}

func (p *page) Add(rows ...domain.Row) domain.Page {
	p.rows = append(p.rows, rows...)
	return p
}

func (p *page) GetStructure() *tree.Node[domain.Structure] {
	str := domain.Structure{
		Type: string(p._type),
	}

	node := tree.NewNode(0, str)

	for _, r := range p.rows {
		inner := r.GetStructure()
		node.AddNext(inner)
	}

	return node
}
