// Package page implements creation of pages.
package page

import (
	"github.com/johnfercher/go-tree/node"

	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type page struct {
	number int
	total  int
	rows   []core.Row
	config *entity.Config
	prop   props.Page
}

// New is responsible to create a core.Page.
func New(ps ...props.Page) core.Page {
	prop := props.Page{}
	if len(ps) > 0 {
		prop = ps[0]
	}

	return &page{
		prop: prop,
	}
}

// Render renders a Page into a PDF context.
func (p *page) Render(provider core.Provider, cell entity.Cell) {
	innerCell := cell.Copy()

	prop := &props.Rect{}
	prop.MakeValid()

	if p.config.BackgroundImage != nil {
		provider.AddBackgroundImageFromBytes(p.config.BackgroundImage.Bytes, &innerCell, prop, p.config.BackgroundImage.Extension)
	}

	for _, row := range p.rows {
		row.Render(provider, innerCell)
		innerCell.Y += row.GetHeight()
	}

	if p.prop.Pattern != "" {
		provider.AddText(p.prop.GetPageString(p.number, p.total), &cell, p.prop.GetNumberTextProp(cell.Height))
	}
}

// SetConfig sets the page configuration.
func (p *page) SetConfig(config *entity.Config) {
	p.config = config
	for _, row := range p.rows {
		row.SetConfig(config)
	}
}

// SetNumber sets the page number and total.
func (p *page) SetNumber(number int, total int) {
	p.number = number
	p.total = total
}

// GetNumber returns the page number.
func (p *page) GetNumber() int {
	return p.number
}

// Add adds one or more rows to the page.
func (p *page) Add(rows ...core.Row) core.Page {
	p.rows = append(p.rows, rows...)
	return p
}

// GetRows returns the rows of the page.
func (p *page) GetRows() []core.Row {
	return p.rows
}

// GetStructure returns the Structure of a Page.
func (p *page) GetStructure() *node.Node[core.Structure] {
	str := core.Structure{
		Type: "page",
	}

	n := node.New(str)
	for _, r := range p.rows {
		inner := r.GetStructure()
		n.AddNext(inner)
	}

	return n
}
