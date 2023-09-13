package text

import (
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/internal/fpdf"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
	v2 "github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/context"
)

type text struct {
	value      string
	_type      v2.DocumentType
	components []v2.Component
	prop       props.Text
}

func New(value string, textProps ...props.Text) *text {
	prop := props.Text{}
	if len(textProps) > 0 {
		prop = textProps[0]
	}
	prop.MakeValid(consts.Arial)

	return &text{
		_type: v2.Text,
		value: value,
		prop:  prop,
	}
}

func (t *text) GetType() string {
	return t._type.String()
}

func (t *text) Add(_ ...v2.Component) v2.Component {
	return t
}

func (t *text) Render(fpdf fpdf.Fpdf, ctx context.Context) {
	ctx.Print(t.value)
	t.render(fpdf, ctx)
	return
}

func (t *text) render(fpdf fpdf.Fpdf, ctx context.Context) {
	font := internal.NewFont(fpdf, 2, consts.Arial, consts.Normal)
	math := internal.NewMath(fpdf)
	text := internal.NewText(fpdf, math, font)

	text.Add(
		t.value,
		internal.Cell{fpdf.GetX() - ctx.Margins.Left,
			fpdf.GetY() - ctx.Margins.Top,
			ctx.Dimensions.Width,
			ctx.Dimensions.Height},
		t.prop)
}
