package page

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type page struct {
	number int
	total  int
	rows   []core.Row
	config *config.Config
	prop   props.Page
}

func New(ps ...props.Page) core.Page {
	prop := props.Page{}
	if len(ps) > 0 {
		prop = ps[0]
	}

	return &page{
		prop: prop,
	}
}

func (p *page) Render(provider core.Provider, cell core.Cell) {
	innerCell := cell.Copy()

	for _, row := range p.rows {
		row.Render(provider, innerCell)
		innerCell.Y += row.GetHeight(provider, innerCell.Width)
	}

	if p.prop.Pattern != "" {
		provider.AddText(p.prop.GetPageString(p.number, p.total), cell, p.prop.GetNumberTextProp(cell.Height))
	}
}

func (p *page) SetConfig(config *config.Config) {
	p.config = config
	for _, row := range p.rows {
		row.SetConfig(config)
	}
}

func (p *page) SetNumber(number int, total int) {
	p.number = number
	p.total = total
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
