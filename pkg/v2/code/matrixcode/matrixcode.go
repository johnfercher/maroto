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

func (i *matrixCode) Render(fpdf fpdf.Fpdf, ctx context.Context) {
	math := internal.NewMath(fpdf)

	code := internal.NewCode(fpdf, math)
	code.AddDataMatrix(i.code, internal.Cell{fpdf.GetX() - ctx.Margins.Left,
		fpdf.GetY() - ctx.Margins.Top,
		ctx.Dimensions.Width,
		ctx.Dimensions.Height}, i.prop)
}

func (i *matrixCode) GetType() string {
	return i._type.String()
}

func (i *matrixCode) Add(_ ...v2.Component) {
	return
}
