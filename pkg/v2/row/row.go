package row

import (
	"github.com/johnfercher/maroto/internal/fpdf"
	"github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/context"
	"github.com/johnfercher/maroto/pkg/v2/types"
)

type row struct {
	height float64
	_type  types.DocumentType
	cols   []v2.Col
}

func (r *row) Add(cols ...v2.Col) {
	for _, col := range cols {
		r.cols = append(r.cols, col)
	}
}

func New(height float64) *row {
	return &row{
		_type:  types.Row,
		height: height,
	}
}

func (r *row) GetHeight() float64 {
	return r.height
}

func (r *row) Render(fpdf fpdf.Fpdf, ctx context.Context) {
	ctx.Print(r.height)

	ctx.Dimensions.Height = r.height
	innerCtx := ctx.Copy()
	for _, col := range r.cols {
		col.Render(fpdf, innerCtx)

		size := col.GetSize()
		parentWidth := ctx.Dimensions.Width
		percent := float64(size) / 12
		colDimension := parentWidth * percent

		innerCtx.Coordinate.X += colDimension
	}

	r.render(fpdf, ctx)
	return
}

func (r *row) render(fpdf fpdf.Fpdf, ctx context.Context) {
	fpdf.Ln(ctx.Dimensions.Height)
}
