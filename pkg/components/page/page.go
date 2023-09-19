package page

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/core"
)

type page struct {
	number int
	rows   []core.Row
	config *config.Config
}

func New() core.Page {
	return &page{}
}

func (p *page) Render(provider core.Provider, cell core.Cell) {
	innerCell := cell.Copy()
	for _, row := range p.rows {
		row.Render(provider, innerCell)
		innerCell.Y += row.GetHeight()
	}
}

func (p *page) SetConfig(config *config.Config) {
	p.config = config
	for _, row := range p.rows {
		row.SetConfig(config)
	}
}

func (p *page) SetNumber(number int) {
	p.number = number
}

func (p *page) GetNumber() int {
	return p.number
}

func (p *page) Add(rows ...core.Row) core.Page {
	p.rows = append(p.rows, rows...)
	return p
}

func (p *page) GetStructure() *tree.Node[core.Structure] {
	str := core.Structure{
		Type: "page",
	}

	node := tree.NewNode(str)

	for _, r := range p.rows {
		inner := r.GetStructure()
		node.AddNext(inner)
	}

	return node
}
