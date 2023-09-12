package col

import (
	"github.com/johnfercher/maroto/internal/fpdf"
	"github.com/johnfercher/maroto/pkg/v2"
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

func (c *col) Render(fpdf fpdf.Fpdf, ctx v2.Context) {
	ctx.Print(c.size)
	ctx = c.setRelativeDimension(ctx)
	for _, component := range c.components {
		component.Render(fpdf, ctx)
	}
	c.render(fpdf, ctx)
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

func (c *col) render(fpdf fpdf.Fpdf, ctx v2.Context) {
	fpdf.SetDrawColor(255, 0, 0)
	fpdf.CellFormat(ctx.GetX(), ctx.GetY(), "", "1", 0, "C", false, 0, "")
}

func (c *col) setRelativeDimension(ctx v2.Context) v2.Context {
	parentWidth := ctx.Dimensions.Width
	colDimension := parentWidth / float64(c.size)
	return ctx.WithDimension(colDimension, ctx.Dimensions.Height)
}
