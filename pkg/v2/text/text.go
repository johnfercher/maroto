package text

import (
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/internal/fpdf"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/johnfercher/maroto/pkg/v2/context"
	"github.com/johnfercher/maroto/pkg/v2/types"
)

type text struct {
	value string
	_type types.DocumentType
}

func New(value string) *text {
	return &text{
		_type: types.Text,
		value: value,
	}
}

func (t *text) Render(fpdf fpdf.Fpdf, ctx context.Context) {
	t.render(fpdf, ctx)
	return
}

func (t *text) render(fpdf fpdf.Fpdf, ctx context.Context) {
	font := internal.NewFont(fpdf, 2, consts.Arial, consts.Normal)
	math := internal.NewMath(fpdf)
	text := internal.NewText(fpdf, math, font)

	props := props.Text{}
	props.MakeValid(consts.Arial)

	text.Add(t.value, internal.Cell{ctx.Coordinate.X, ctx.Coordinate.Y, ctx.Dimensions.Width, ctx.Dimensions.Height}, props)
}
