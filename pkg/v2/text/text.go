package text

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/internal/fpdf"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/johnfercher/maroto/pkg/v2/context"
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

	return tree.NewNode(0, str)
}

func (t *text) Render(fpdf fpdf.Fpdf, ctx context.Context) {
	t.render(fpdf, ctx)
	return
}

func (t *text) render(fpdf fpdf.Fpdf, ctx context.Context) {
	font := internal.NewFont(fpdf, 2, consts.Arial, consts.Normal)
	math := internal.NewMath(fpdf)
	text := internal.NewText(fpdf, math, font)

	if t.prop.Top > ctx.Dimensions.Height {
		t.prop.Top = ctx.Dimensions.Height
	}

	if t.prop.Left > ctx.Dimensions.Width {
		t.prop.Left = ctx.Dimensions.Width
	}

	if t.prop.Right > ctx.Dimensions.Width {
		t.prop.Right = ctx.Dimensions.Width
	}

	text.Add(
		t.value,
		internal.Cell{ctx.Coordinate.X, ctx.Coordinate.Y, ctx.Dimensions.Width, ctx.Dimensions.Height},
		t.prop)
}
