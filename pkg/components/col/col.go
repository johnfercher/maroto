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

func (c *col) CalculateHeight(provider core.Provider, rowWidth float64) float64 {
	var h float64
	for _, cmp := range c.components {
		if ahc, ok := cmp.(core.AdaptiveHeightComponent); ok {
			h += ahc.GetComputedHeight(provider, c.GetWidth(rowWidth))
		}
	}

	return h
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

func (c *col) GetWidth(rowWidth float64) float64 {
	return rowWidth * float64(c.GetSize()) / float64(c.config.MaxGridSize)
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
	innerCell := cell.Copy()

	if createCell {
		provider.CreateCol(cell.Width, cell.Height, c.config, c.style)
	}

	for _, component := range c.components {
		component.Render(provider, innerCell)
		h := 0.0
		if ccc, ok := component.(core.AdaptiveHeightComponent); ok {
			h += ccc.GetComputedHeight(provider, cell.Width)
		}
		innerCell.Y += h
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
