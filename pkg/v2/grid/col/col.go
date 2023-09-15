package col

import (
	"fmt"
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/pkg/v2/domain"
)

const (
	defaultGridSize = 12.0
)

type col struct {
	size  int
	nodes []domain.Node
	rows  []domain.Row
}

func New(size int) domain.Col {
	return &col{
		size: size,
	}
}

func (c *col) Add(node ...domain.Node) domain.Col {
	c.nodes = append(c.nodes, node...)
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
		Type:  "col",
		Value: fmt.Sprintf("%d", c.size),
	}

	node := tree.NewNode(str)

	for _, c := range c.nodes {
		inner := c.GetStructure()
		node.AddNext(inner)
	}

	return node
}

func (c *col) Render(provider domain.Provider, cell internal.Cell) {
	c.render(provider, cell)

	for _, component := range c.nodes {
		component.Render(provider, cell)
	}

	return
}

func (c *col) render(provider domain.Provider, cell internal.Cell) {
	provider.CreateCol(cell.Width, cell.Height)
}
