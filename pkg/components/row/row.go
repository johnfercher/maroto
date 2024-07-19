// Package row implements creation of rows.
package row

import (
	"github.com/johnfercher/maroto/v2/pkg/core/entity"

	"github.com/johnfercher/go-tree/node"

	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type Row struct {
	height     float64
	autoHeight bool
	cols       []core.Col
	style      *props.Cell
	config     *entity.Config
}

// New is responsible to create a core.Row.
//
// Height is an optional parameter that, if not sent, will be calculated automatically
// Height is defined in mm.
func New(height ...float64) core.Row {
	autoHeight := false
	if len(height) == 0 {
		height = append(height, 0)
		autoHeight = true
	}
	return &Row{
		height:     height[0],
		autoHeight: autoHeight,
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
	if r.autoHeight {
		r.resetHeight()
	}
	return r
}

// GetColumns returns the columns of a core.Row.
func (r *Row) GetColumns() []core.Col {
	return r.cols
}

// Returns the height of the row content
func (r *Row) getBiggestCol(provider core.Provider, cell *entity.Cell) float64 {
	greaterHeight := 0.0
	for _, col := range r.cols {
		height := col.GetHeight(provider, cell)
		if greaterHeight < height {
			greaterHeight = height
		}
	}
	return greaterHeight
}

// GetHeight returns the height of a core.Row.
func (r *Row) GetHeight(provider core.Provider, cell *entity.Cell) float64 {
	if r.height == 0 {
		r.height = r.getBiggestCol(provider, cell)
	}
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
	cell.Height = r.GetHeight(provider, &cell)
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

// resetHeight resets the line height to 0
func (r *Row) resetHeight() {
	r.height = 0
}
