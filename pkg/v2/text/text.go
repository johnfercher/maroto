package text

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/johnfercher/maroto/pkg/v2/config"
	"github.com/johnfercher/maroto/pkg/v2/domain"
	"github.com/johnfercher/maroto/pkg/v2/grid/col"
	"github.com/johnfercher/maroto/pkg/v2/grid/row"
)

type text struct {
	value  string
	prop   props.Text
	config *config.Maroto
}

func New(value string, ps ...props.Text) domain.Component {
	textProp := props.Text{
		Color: color.Color{
			Red:   0,
			Green: 0,
			Blue:  0,
		},
	}

	if len(ps) > 0 {
		textProp = ps[0]
	}
	textProp.MakeValid(consts.Arial)

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
	t.render(provider, cell)
}

func (t *text) render(provider domain.Provider, cell internal.Cell) {
	if t.prop.Top > cell.Height {
		t.prop.Top = cell.Height
	}

	if t.prop.Left > cell.Width {
		t.prop.Left = cell.Width
	}

	if t.prop.Right > cell.Width {
		t.prop.Right = cell.Width
	}

	provider.AddText(t.value, cell, t.prop)
}
