package text

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/johnfercher/maroto/pkg/v2/domain"
	"github.com/johnfercher/maroto/pkg/v2/types"
)

type text struct {
	value      string
	_type      types.DocumentType
	components []domain.Node
	prop       props.Text
}

func New(value string, prop ...props.Text) domain.Component {
	textProp := props.Text{
		Color: color.Color{
			Red:   0,
			Green: 0,
			Blue:  0,
		},
	}

	if len(prop) > 0 {
		textProp = prop[0]
	}
	textProp.MakeValid(consts.Arial)

	return &text{
		_type: types.Text,
		value: value,
		prop:  textProp,
	}
}

func (t *text) GetType() string {
	return t._type.String()
}

func (t *text) Add(_ ...domain.Node) domain.Node {
	return t
}

func (t *text) GetStructure() *tree.Node[domain.Structure] {
	str := domain.Structure{
		Type:  string(t._type),
		Value: t.value,
	}

	return tree.NewNode(str)
}

func (t *text) Render(provider domain.Provider, cell internal.Cell) {
	t.render(provider, cell)
	return
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
