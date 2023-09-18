package text

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/v2/internal"
	"github.com/johnfercher/maroto/v2/pkg/color"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/domain"
	"github.com/johnfercher/maroto/v2/pkg/grid/col"
	"github.com/johnfercher/maroto/v2/pkg/grid/row"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type text struct {
	value  string
	prop   props.Text
	config *config.Maroto
}

func New(value string, ps ...props.Text) domain.Component {
	textProp := props.Text{
		Color: &color.Color{
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

func NewCol(size int, value string, ps ...props.Text) domain.Col {
	text := New(value, ps...)
	return col.New(size).Add(text)
}

func NewRow(height float64, value string, ps ...props.Text) domain.Row {
	r := New(value, ps...)
	c := col.New().Add(r)
	return row.New(height).Add(c)
}

func (t *text) GetStructure() *tree.Node[domain.Structure] {
	str := domain.Structure{
		Type:  "text",
		Value: t.value,
	}

	return tree.NewNode(str)
}

func (t *text) SetConfig(config *config.Maroto) {
	t.config = config
}

func (t *text) GetValue() string {
	return t.value
}

func (t *text) Render(provider domain.Provider, cell internal.Cell) {
	t.prop.MakeValid(t.config.Font)
	provider.AddText(t.value, cell, t.prop)
}
