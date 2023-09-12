package col

import (
	"github.com/johnfercher/maroto/internal/fpdf"
	"github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/context"
)

const (
	defaultGridSize = 12.0
)

type col struct {
	size       int
	_type      v2.DocumentType
	components []v2.Component
}

func New(size int) *col {
	return &col{
		_type: v2.Col,
		size:  size,
	}
}

func (c *col) GetType() string {
	return c._type.String()
}

func (c *col) Add(components ...v2.Component) {
	for _, component := range components {
		if c._type.Accept(component.GetType()) {
			c.components = append(c.components, component)
		}
	}
}

func (c *col) Render(fpdf fpdf.Fpdf, ctx context.Context) {
	ctx.Print(c.size)
	ctx = c.setRelativeDimension(ctx)
	for _, component := range c.components {
		component.Render(fpdf, ctx)
	}
	c.render(fpdf, ctx)
	return
}

func (c *col) render(fpdf fpdf.Fpdf, ctx context.Context) {
	fpdf.SetDrawColor(255, 0, 0)
	fpdf.CellFormat(ctx.GetXOffset(), ctx.GetYOffset(), "", "1", 0, "C", false, 0, "")
}

func (c *col) setRelativeDimension(ctx context.Context) context.Context {
	parentWidth := ctx.Dimensions.Width
	percent := float64(c.size) / defaultGridSize
	colDimension := parentWidth * percent
	return ctx.WithDimension(colDimension, ctx.Dimensions.Height)
}
