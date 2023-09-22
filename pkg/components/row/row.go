package row

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"

	"github.com/johnfercher/go-tree/tree"
)

type row struct {
	height float64
	cols   []core.Col
	style  *props.Cell
	config *config.Config
}

func (r *row) CalculateHeight(provider core.Provider, cellWidth float64) float64 {
	for _, c := range r.cols {
		h := c.CalculateHeight(provider, cellWidth)
		if h > r.height {
			r.height = h
		}
	}

	return r.height
}

func New(height float64) core.Row {
	return &row{
		height: height,
	}
}

func NewAdaptive(cols ...core.Col) core.Row {
	return &row{
		cols: cols,
	}
}

func Empty(height float64) core.Row {
	r := New(height)
	r.Add(col.New())
	return r
}

func (r *row) SetConfig(config *config.Config) {
	r.config = config
	for _, cols := range r.cols {
		cols.SetConfig(config)
	}
}

func (r *row) Add(cols ...core.Col) core.Row {
	r.cols = append(r.cols, cols...)
	return r
}

func (r *row) GetHeight(provider core.Provider, cellWidth float64) float64 {
	if r.height == 0 {
		r.height = r.CalculateHeight(provider, cellWidth)
	}

	return r.height
}

func (r *row) GetStructure() *tree.Node[core.Structure] {
	str := core.Structure{
		Type:  "row",
		Value: fmt.Sprintf("%2.f", r.height),
	}

	node := tree.NewNode(str)

	for _, c := range r.cols {
		inner := c.GetStructure()
		node.AddNext(inner)
	}

	return node
}

func (r *row) Render(provider core.Provider, cell core.Cell) {
	cell.Height = r.height
	innerCell := cell.Copy()

	if r.style != nil {
		provider.CreateCol(cell.Width, cell.Height, r.config, r.style)
	}

	for _, c := range r.cols {
		innerCell.Width = c.GetWidth(cell.Width)

		c.Render(provider, innerCell, r.style == nil)
		innerCell.X += innerCell.Width
	}

	provider.CreateRow(cell.Height)
}

func (r *row) WithStyle(style *props.Cell) core.Row {
	r.style = style
	return r
}
