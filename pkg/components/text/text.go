package text

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type text struct {
	value  string
	prop   props.Text
	config *config.Config
}

func New(value string, ps ...props.Text) core.Component {
	textProp := props.Text{
		Color: &props.Color{
			Red:   0,
			Green: 0,
			Blue:  0,
		},
	}

	if len(ps) > 0 {
		textProp = ps[0]
	}

	return &text{
		value: value,
		prop:  textProp,
	}
}

func NewCol(size int, value string, ps ...props.Text) core.Col {
	text := New(value, ps...)
	return col.New(size).Add(text)
}

func NewRow(height float64, value string, ps ...props.Text) core.Row {
	r := New(value, ps...)
	c := col.New().Add(r)
	return row.New(height).Add(c)
}

func (t *text) GetStructure() *tree.Node[core.Structure] {
	str := core.Structure{
		Type:  "text",
		Value: t.value,
	}

	return tree.NewNode(str)
}

func (t *text) SetConfig(config *config.Config) {
	t.config = config
}

func (t *text) Render(provider core.Provider, cell *core.Cell) {
	t.prop.MakeValid(t.config.DefaultFont)
	provider.AddText(t.value, cell, t.prop)
}
