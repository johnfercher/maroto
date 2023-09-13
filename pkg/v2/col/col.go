package col

import (
	"fmt"
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/internal/fpdf"
	"github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/context"
	"github.com/johnfercher/maroto/pkg/v2/types"
)

const (
	defaultGridSize = 12.0
)

type col struct {
	size       int
	_type      types.DocumentType
	components []v2.Component
	rows       []v2.Row
}

func New(size int) v2.Col {
	return &col{
		_type: types.Col,
		size:  size,
	}
}

func (c *col) GetType() string {
	return c._type.String()
}

func (c *col) Add(components ...v2.Component) v2.Col {
	c.components = append(c.components, components...)
	return c
}

func (c *col) AddInner(rows ...v2.Row) v2.Col {
	c.rows = append(c.rows, rows...)
	return c
}

func (c *col) GetStructure() *tree.Node[v2.Structure] {
	str := v2.Structure{
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
