package page

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/internal/fpdf"
	v2 "github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/context"
	"github.com/johnfercher/maroto/pkg/v2/types"
)

type page struct {
	_type types.DocumentType
	rows  []v2.Row
}

func New() v2.Page {
	return &page{
		_type: types.Page,
	}
}

func (p *page) Render(fpdf fpdf.Fpdf, ctx context.Context) {
	fpdf.AddPage()
	ctx = ctx.NewPage(fpdf.PageNo())
	for _, row := range p.rows {
		row.Render(fpdf, ctx)
	}
}

func (p *page) GetType() string {
	return p._type.String()
}

func (p *page) Add(rows ...v2.Row) v2.Page {
	p.rows = append(p.rows, rows...)
	return p
}

func (p *page) GetStructure() *tree.Node[v2.Structure] {
	str := v2.Structure{
		Type: string(p._type),
	}

	node := tree.NewNode(0, str)

	for _, r := range p.rows {
		inner := r.GetStructure()
		node.AddNext(inner)
	}

	return node
}
