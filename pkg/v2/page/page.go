package page

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/pkg/v2/domain"
	"github.com/johnfercher/maroto/pkg/v2/types"
)

type page struct {
	number int
	_type  types.DocumentType
	rows   []domain.Row
}

func New() domain.Page {
	return &page{
		_type: types.Page,
	}
}

func (p *page) Render(provider domain.Provider, cell internal.Cell) {
	innerCell := cell.Copy()
	for _, row := range p.rows {
		row.Render(provider, innerCell)
		innerCell.Y += row.GetHeight()
	}
}

func (p *page) GetType() string {
	return p._type.String()
}

func (p *page) SetNumber(number int) {
	p.number = number
}

func (p *page) GetNumber() int {
	return p.number
}

func (p *page) Add(rows ...domain.Row) domain.Page {
	p.rows = append(p.rows, rows...)
	return p
}

func (p *page) GetStructure() *tree.Node[domain.Structure] {
	str := domain.Structure{
		Type: string(p._type),
	}

	node := tree.NewNode(str)

	for _, r := range p.rows {
		inner := r.GetStructure()
		node.AddNext(inner)
	}

	return node
}
