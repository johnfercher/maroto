package qrcode

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/internal/fpdf"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/johnfercher/maroto/pkg/v2/context"
	"github.com/johnfercher/maroto/pkg/v2/domain"
	"github.com/johnfercher/maroto/pkg/v2/types"
)

type qrCode struct {
	code       string
	_type      types.DocumentType
	components []domain.Node
	prop       props.Rect
}

func New(code string, barcodeProps ...props.Rect) domain.Component {
	prop := props.Rect{}
	if len(barcodeProps) > 0 {
		prop = barcodeProps[0]
	}
	prop.MakeValid()

	return &qrCode{
		_type: types.QrCode,
		code:  code,
		prop:  prop,
	}
}

func (q *qrCode) Render(fpdf fpdf.Fpdf, ctx context.Context) {
	math := internal.NewMath(fpdf)

	code := internal.NewCode(fpdf, math)
	x := fpdf.GetX() - ctx.Margins.Left - ctx.Dimensions.Width
	y := fpdf.GetY() - ctx.Margins.Top
	code.AddQr(q.code,
		internal.Cell{x, y, ctx.Dimensions.Width, ctx.Dimensions.Height},
		q.prop)
}

func (q *qrCode) GetType() string {
	return q._type.String()
}

func (q *qrCode) Add(component ...domain.Node) domain.Node {
	return q
}

func (q *qrCode) GetStructure() *tree.Node[domain.Structure] {
	str := domain.Structure{
		Type:  string(q._type),
		Value: q.code,
	}

	return tree.NewNode(0, str)
}
