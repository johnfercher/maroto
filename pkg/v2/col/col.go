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

func (c *col) Render(fpdf fpdf.Fpdf, ctx *v2.Context) {
	ctx.Print(c.size)
	for _, component := range c.components {
		component.Render(fpdf, ctx)
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
