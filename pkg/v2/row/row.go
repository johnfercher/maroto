package row

import (
	"fmt"
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/internal/fpdf"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/v2/domain"
	"github.com/johnfercher/maroto/pkg/v2/types"
)

type row struct {
	height float64
	_type  types.DocumentType
	cols   []domain.Col
	color  color.Color
}

func (r *row) GetHeight() float64 {
	return r.height
}

func (r *row) Add(cols ...domain.Col) {
	r.cols = append(r.cols, cols...)
}

func New(height float64, c ...color.Color) domain.Row {
	cx := color.NewBlack()
	if len(c) > 0 {
		cx = c[0]
	}

	return &row{
		_type:  types.Row,
		height: height,
		color:  cx,
	}
}

func (r *row) GetType() string {
	return r._type.String()
}

func (r *row) GetStructure() *tree.Node[domain.Structure] {
	str := domain.Structure{
		Type:  string(r._type),
		Value: fmt.Sprintf("%2.f", r.height),
	}

	node := tree.NewNode(0, str)

	for _, c := range r.cols {
		inner := c.GetStructure()
		node.AddNext(inner)
	}

	return node
}

func (r *row) Render(fpdf fpdf.Fpdf, cell internal.Cell) {
	fpdf.SetDrawColor(r.color.Red, r.color.Green, r.color.Blue)

	cell.Height = r.height
	innerCell := cell.Copy()
	for _, col := range r.cols {
		size := col.GetSize()
		parentWidth := cell.Width
		percent := float64(size) / 12
		colDimension := parentWidth * percent

		col.Render(fpdf, innerCell)
		innerCell.X += colDimension
	}

	r.render(fpdf, cell)
	return
}

func (r *row) render(fpdf fpdf.Fpdf, cell internal.Cell) {
	fpdf.Ln(cell.Height)
}
