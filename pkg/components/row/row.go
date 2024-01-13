// Package row implements creation of rows.
package row

import (
	"github.com/johnfercher/maroto/v2/pkg/core/entity"

	"github.com/johnfercher/go-tree/node"

	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type Row struct {
	height float64
	cols   []core.Col
	style  *props.Cell
	config *entity.Config
}

// New is responsible to create a core.Row.
func New(height float64) core.Row {
	return &Row{
		height: height,
	}
}

// SetConfig sets the Row configuration.
func (r *Row) SetConfig(config *entity.Config) {
	r.config = config
	for _, cols := range r.cols {
		cols.SetConfig(config)
	}
}

// Add is responsible to add one or more core.Col to a core.Row.
func (r *Row) Add(cols ...core.Col) core.Row {
	r.cols = append(r.cols, cols...)
	return r
}

// GetHeight returns the height of a core.Row.
func (r *Row) GetHeight() float64 {
	return r.height
}

// GetStructure returns the Structure of a core.Row.
func (r *Row) GetStructure() *node.Node[core.Structure] {
	detailsMap := r.style.ToMap()

	str := core.Structure{
		Type:    "row",
		Value:   r.height,
		Details: detailsMap,
	}

	node := node.New(str)

	for _, c := range r.cols {
		inner := c.GetStructure()
		node.AddNext(inner)
	}

	return node
}

// Render renders a Row into a PDF context.
func (r *Row) Render(provider core.Provider, cell entity.Cell) {
	cell.Height = r.height
	innerCell := cell.Copy()

	if r.style != nil {
		provider.CreateCol(cell.Width, cell.Height, r.config, r.style)
	}

	for _, col := range r.cols {
		size := col.GetSize()
		parentWidth := cell.Width

		percent := float64(size) / float64(r.config.MaxGridSize)

		colDimension := parentWidth * percent
		innerCell.Width = colDimension

		col.Render(provider, innerCell, r.style == nil)
		innerCell.X += colDimension
	}

	provider.CreateRow(cell.Height)
}

// WithStyle sets the style of a Row.
func (r *Row) WithStyle(style *props.Cell) core.Row {
	r.style = style
	return r
}
