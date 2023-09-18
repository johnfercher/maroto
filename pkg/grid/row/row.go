package row

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/internal"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/domain"
	"github.com/johnfercher/maroto/v2/pkg/grid/col"
	"github.com/johnfercher/maroto/v2/pkg/props"

	"github.com/johnfercher/go-tree/tree"
)

type row struct {
	height float64
	cols   []domain.Col
	style  *props.Style
	config *config.Maroto
}

func New(height float64) domain.Row {
	return &row{
		height: height,
	}
}

func Empty(height float64) domain.Row {
	r := New(height)
	r.Add(col.New())
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

func (r *row) WithStyle(style *props.Style) domain.Row {
	r.style = style
	return r
}
