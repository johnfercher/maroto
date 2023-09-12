package text

import (
	"github.com/johnfercher/maroto/internal/fpdf"
	v2 "github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/context"
)

type text struct {
	value      string
	_type      v2.DocumentType
	components []v2.Component
}

func New(value string) *text {
	return &text{
		_type: v2.Text,
		value: value,
	}
}

func (t *text) GetType() string {
	return t._type.String()
}

func (t *text) Add(_ ...v2.Component) {
	return
}

func (t *text) Render(fpdf fpdf.Fpdf, ctx context.Context) {
	ctx.Print(t.value)
	t.render(fpdf, ctx)
	return
}

func (t *text) render(fpdf fpdf.Fpdf, ctx context.Context) {
	fpdf.Text(ctx.GetXOffset(), ctx.GetYOffset(), t.value)
}
