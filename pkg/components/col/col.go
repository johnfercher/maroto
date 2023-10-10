package col

import (
	"github.com/johnfercher/maroto/v2/pkg/core/entity"

	"github.com/johnfercher/go-tree/node"

	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type col struct {
	size       int
	isMax      bool
	components []core.Component
	config     *entity.Config
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

func (c *col) GetStructure() *node.Node[core.Structure] {
	str := core.Structure{
		Type:    "col",
		Value:   c.size,
		Details: c.style.ToMap(),
	}

	if c.isMax {
		if len(str.Details) == 0 {
			str.Details = make(map[string]interface{})
		}
		str.Details["is_max"] = true
	}

	node := node.New(str)

	for _, c := range c.components {
		inner := c.GetStructure()
		node.AddNext(inner)
	}

	return node
}

func (c *col) Render(provider core.Provider, cell entity.Cell, createCell bool) {
	if createCell {
		provider.CreateCol(cell.Width, cell.Height, c.config, c.style)
	}

	for _, component := range c.components {
		component.Render(provider, &cell)
	}
}

func (c *col) SetConfig(config *entity.Config) {
	c.config = config
	for _, component := range c.components {
		component.SetConfig(config)
	}
}

func (c *col) WithStyle(style *props.Cell) core.Col {
	c.style = style
	return c
}
