package col

import (
	"github.com/johnfercher/maroto/internal/fpdf"
	"github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/context"
	"github.com/johnfercher/maroto/pkg/v2/types"
)

const (
	defaultGridSize = 12.0
)

type col struct {
	size        int
	_type       types.DocumentType
	renderables []v2.Renderable
}

func New(size int) *col {
	return &col{
		_type: types.Col,
		size:  size,
	}
}

func (c *col) GetSize() int {
	return c.size
}

func (c *col) Add(renderables ...v2.Renderable) {
	c.renderables = append(renderables, renderables...)
}

func (c *col) Render(fpdf fpdf.Fpdf, ctx context.Context) {
	ctx.Print(c.size)

	parentWidth := ctx.Dimensions.Width
	percent := float64(c.size) / defaultGridSize
	colDimension := parentWidth * percent
	ctx.Dimensions.Width = colDimension
	for _, renderable := range c.renderables {
		renderable.Render(fpdf, ctx)
	}

	c.render(fpdf, ctx)
	return
}

func (c *col) render(fpdf fpdf.Fpdf, ctx context.Context) {
	fpdf.SetDrawColor(255, 0, 0)
	fpdf.CellFormat(ctx.Dimensions.Width, ctx.Dimensions.Height, "", "1", 0, "C", false, 0, "")
}

func (c *col) setRelativeDimension(ctx context.Context) context.Context {
	parentWidth := ctx.Dimensions.Width
	percent := float64(c.size) / defaultGridSize
	colDimension := parentWidth * percent
	return ctx.WithDimension(colDimension, ctx.Dimensions.Height)
}
