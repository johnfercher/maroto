package col

import (
	"fmt"

	"github.com/johnfercher/maroto/internal"

	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/johnfercher/maroto/pkg/v2/config"
	"github.com/johnfercher/maroto/pkg/v2/domain"
)

type col struct {
	size       int
	isMax      bool
	components []domain.Component
	config     *config.Maroto
	style      *props.Style
}

func New(size ...int) domain.Col {
	if len(size) == 0 {
		return &col{isMax: true}
	}

	return &col{size: size[0]}
}

func (c *col) Add(components ...domain.Component) domain.Col {
	c.components = append(c.components, components...)
	return c
}

func (c *col) GetSize() int {
	if c.isMax {
		return c.config.MaxGridSize
	}

	return c.size
}

func (c *col) GetStructure() *tree.Node[domain.Structure] {
	str := domain.Structure{
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

func (c *col) Render(provider domain.Provider, cell internal.Cell, createCell bool) {
	if createCell {
		provider.CreateCol(cell.Width, cell.Height, c.config, c.style)
	}

	for _, component := range c.components {
		component.Render(provider, cell)
	}
}

func (c *col) SetConfig(config *config.Maroto) {
	c.config = config
	for _, component := range c.components {
		component.SetConfig(config)
	}
}

func (c *col) WithStyle(style *props.Style) domain.Col {
	c.style = style
	return c
}
