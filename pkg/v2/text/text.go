package text

import (
	"github.com/johnfercher/maroto/internal/fpdf"
	v2 "github.com/johnfercher/maroto/pkg/v2"
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

func (t *text) Render(fpdf fpdf.Fpdf, ctx v2.Context) {
	ctx.Print(t.value)
	for _, component := range t.components {
		component.Render(fpdf, ctx)
	}
}

func (t *text) GetType() string {
	return t._type.String()
}

func (t *text) Add(_ ...v2.Component) {
	return
}
