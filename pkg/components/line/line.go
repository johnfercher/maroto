// Package line implements creation of lines.
package line

import (
	"github.com/johnfercher/go-tree/node"

	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type Line struct {
	config *entity.Config
	prop   props.Line
}

// New is responsible to create an instance of a Line.
func New(ps ...props.Line) core.Component {
	lineProp := props.Line{}
	if len(ps) > 0 {
		lineProp = ps[0]
	}
	lineProp.MakeValid()

	return &Line{
		prop: lineProp,
	}
}

// NewCol is responsible to create an instance of a Line wrapped in a Col.
func NewCol(size int, ps ...props.Line) core.Col {
	r := New(ps...)
	return col.New(size).Add(r)
}

// NewRow is responsible to create an instance of a Line wrapped in a Row.
func NewRow(height float64, ps ...props.Line) core.Row {
	r := New(ps...)
	c := col.New().Add(r)
	return row.New(height).Add(c)
}

// NewRow is responsible to create an instance of a Line wrapped in a automatic Row.
func NewAutoRow(ps ...props.Line) core.Row {
	r := New(ps...)
	c := col.New().Add(r)
	return row.New().Add(c)
}

// GetStructure returns the Structure of a Line.
func (l *Line) GetStructure() *node.Node[core.Structure] {
	str := core.Structure{
		Type:    "line",
		Details: l.prop.ToMap(),
	}

	return node.New(str)
}

// SetConfig sets the config.
func (l *Line) SetConfig(config *entity.Config) {
	l.config = config
}

// GetHeight returns the height that the line will have in the PDF
func (b *Line) GetHeight(provider core.Provider, cell *entity.Cell) float64 {
	return b.prop.Thickness
}

// Render renders a Line into a PDF context.
func (l *Line) Render(provider core.Provider, cell *entity.Cell) {
	provider.AddLine(cell, &l.prop)
}
