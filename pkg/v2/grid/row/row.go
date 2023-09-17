package row

import (
	"fmt"
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/v2/config"
	"github.com/johnfercher/maroto/pkg/v2/domain"
	"github.com/johnfercher/maroto/pkg/v2/grid/col"
)

type row struct {
	height float64
	cols   []domain.Col
	color  color.Color
	config *config.Maroto
}

func New(height float64, c ...color.Color) domain.Row {
	cx := color.NewBlack()
	if len(c) > 0 {
		cx = c[0]
	}

	return &row{
		height: height,
		color:  cx,
	}
}

func Empty(height float64) domain.Row {
	r := New(height)
	r.Add(col.Empty())
	return r
}

func (r *row) SetConfig(config *config.Maroto) {
	r.config = config
	for _, cols := range r.cols {
		cols.SetConfig(config)
	}
}

func (r *row) Add(cols ...domain.Col) domain.Row {
	r.cols = append(r.cols, cols...)
	return r
}

func (r *row) GetHeight() float64 {
	return r.height
}

func (r *row) GetStructure() *tree.Node[domain.Structure] {
	str := domain.Structure{
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

func (r *row) Render(provider domain.Provider, cell internal.Cell) {
	cell.Height = r.height
	innerCell := cell.Copy()
	for _, col := range r.cols {
		size, isMax := col.GetSize()
		parentWidth := cell.Width

		percent := float64(size) / float64(r.config.MaxGridSize)
		if isMax {
			percent = 1
		}

		colDimension := parentWidth * percent
		innerCell.Width = colDimension

		col.Render(provider, innerCell)
		innerCell.X += colDimension
	}

	provider.CreateRow(cell.Height)
	return
}
