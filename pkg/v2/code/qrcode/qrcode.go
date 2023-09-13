package qrcode

import (
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/internal/fpdf"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/context"
)

type qrCode struct {
	code       string
	_type      v2.DocumentType
	components []v2.Component
	prop       props.Rect
}

func New(code string, barcodeProps ...props.Rect) *qrCode {
	prop := props.Rect{}
	if len(barcodeProps) > 0 {
		prop = barcodeProps[0]
	}
	prop.MakeValid()

	return &qrCode{
		_type: v2.Leaf,
		code:  code,
		prop:  prop,
	}
}

func (q *qrCode) Render(fpdf fpdf.Fpdf, ctx context.Context) {
	math := internal.NewMath(fpdf)

	code := internal.NewCode(fpdf, math)
	code.AddQr(q.code, internal.Cell{fpdf.GetX() - ctx.Margins.Left,
		fpdf.GetY() - ctx.Margins.Top,
		ctx.Dimensions.Width,
		ctx.Dimensions.Height}, q.prop)
}

func (q *qrCode) GetType() string {
	return q._type.String()
}

func (q *qrCode) Add(component ...v2.Component) v2.Component {
	return q
}
