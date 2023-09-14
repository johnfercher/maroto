package col

import (
	"fmt"
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/internal/fpdf"
	"github.com/johnfercher/maroto/pkg/v2/context"
	"github.com/johnfercher/maroto/pkg/v2/domain"
	"github.com/johnfercher/maroto/pkg/v2/types"
)

const (
	defaultGridSize = 12.0
)

type col struct {
	size       int
	_type      types.DocumentType
	components []domain.Component
	rows       []domain.Row
}

func New(size int) domain.Col {
	return &col{
		_type: types.Col,
		size:  size,
	}
}

func (c *col) GetType() string {
	return c._type.String()
}

func (c *col) Add(components ...domain.Component) domain.Col {
	c.components = append(c.components, components...)
	return c
}

func (c *col) AddInner(rows ...domain.Row) domain.Col {
	c.rows = append(c.rows, rows...)
	return c
}

func (c *col) GetSize() int {
	return c.size
}

func (c *col) GetStructure() *tree.Node[domain.Structure] {
	str := domain.Structure{
		Type:  string(c._type),
		Value: fmt.Sprintf("%d", c.size),
	}

	node := tree.NewNode(0, str)

	for _, c := range c.components {
		inner := c.GetStructure()
		node.AddNext(inner)
	}

	return node
}

func (c *col) Render(fpdf fpdf.Fpdf, ctx context.Context) {
	parentWidth := ctx.Dimensions.Width
	percent := float64(c.size) / defaultGridSize
	colDimension := parentWidth * percent
	ctx.Dimensions.Width = colDimension

	c.render(fpdf, ctx)

	for _, component := range c.components {
		component.Render(fpdf, ctx)
	}

	return
}

func (c *col) render(fpdf fpdf.Fpdf, ctx context.Context) {
	fpdf.CellFormat(ctx.Dimensions.Width, ctx.Dimensions.Height, "", "1", 0, "C", false, 0, "")
}

func (c *col) setRelativeDimension(ctx context.Context) context.Context {
	parentWidth := ctx.Dimensions.Width
	percent := float64(c.size) / defaultGridSize
	colDimension := parentWidth * percent
	return ctx.WithDimension(colDimension, ctx.Dimensions.Height)
}
