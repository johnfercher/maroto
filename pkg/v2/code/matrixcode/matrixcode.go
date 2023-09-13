package matrixcode

import (
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/internal/fpdf"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/context"
)

type matrixCode struct {
	code       string
	_type      v2.DocumentType
	components []v2.Component
	prop       props.Rect
}

func New(code string, barcodeProps ...props.Rect) *matrixCode {
	prop := props.Rect{}
	if len(barcodeProps) > 0 {
		prop = barcodeProps[0]
	}
	prop.MakeValid()

	return &matrixCode{
		_type: v2.Leaf,
		code:  code,
		prop:  prop,
	}
}

func (m *matrixCode) Render(fpdf fpdf.Fpdf, ctx context.Context) {
	math := internal.NewMath(fpdf)

	code := internal.NewCode(fpdf, math)
	code.AddDataMatrix(m.code, internal.Cell{fpdf.GetX() - ctx.Margins.Left,
		fpdf.GetY() - ctx.Margins.Top,
		ctx.Dimensions.Width,
		ctx.Dimensions.Height}, m.prop)
}

func (m *matrixCode) GetType() string {
	return m._type.String()
}

func (m *matrixCode) Add(component ...v2.Component) v2.Component {
	return m
}
