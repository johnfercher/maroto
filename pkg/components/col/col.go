package col

import (
	"fmt"

	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type col struct {
	size       int
	isMax      bool
	components []core.Component
	config     *config.Config
	style      *props.Cell
}

func New(size ...int) core.Col {
	if len(size) == 0 {
		return &col{isMax: true}
	}

	return &col{size: size[0]}
}

func (c *col) Add(components ...core.Component) core.Col {
	c.components = append(c.components, components...)
	return c
}

func (c *col) GetSize() int {
	if c.isMax {
		return c.config.MaxGridSize
	}

	return c.size
}

func (c *col) GetStructure() *tree.Node[core.Structure] {
	str := core.Structure{
		Type:  "col",
		Value: fmt.Sprintf("%d", c.size),
	}

	node := tree.NewNode(str)

	for _, c := range c.components {
		inner := c.GetStructure()
		node.AddNext(inner)
	}

	return node
}

func (c *col) Render(provider core.Provider, cell core.Cell, createCell bool) {
	if createCell {
		provider.CreateCol(cell.Width, cell.Height, c.config, c.style)
	}

	for _, component := range c.components {
		component.Render(provider, cell)
	}
}

func (c *col) SetConfig(config *config.Config) {
	c.config = config
	for _, component := range c.components {
		component.SetConfig(config)
	}
}

func (c *col) WithStyle(style *props.Cell) core.Col {
	c.style = style
	return c
}
