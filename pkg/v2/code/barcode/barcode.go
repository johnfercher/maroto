package barcode

import (
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/internal/fpdf"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/context"
	"github.com/johnfercher/maroto/pkg/v2/text"
)

type barcode struct {
	code       string
	_type      v2.DocumentType
	components []v2.Component
	prop       props.Barcode
}

func New(code string, barcodeProps ...props.Barcode) *barcode {
	prop := props.Barcode{}
	if len(barcodeProps) > 0 {
		prop = barcodeProps[0]
	}
	prop.MakeValid()

	return &barcode{
		_type: v2.Leaf,
		code:  code,
		prop:  prop,
	}
}

func (b *barcode) Render(fpdf fpdf.Fpdf, ctx context.Context) {
	math := internal.NewMath(fpdf)

	code := internal.NewCode(fpdf, math)
	err := code.AddBar(b.code, internal.Cell{fpdf.GetX() - ctx.Margins.Left,
		fpdf.GetY() - ctx.Margins.Top,
		ctx.Dimensions.Width,
		ctx.Dimensions.Height}, b.prop)

	if err != nil {
		fpdf.ClearError()
		txt := text.New("Failed to render code")
		txt.Render(fpdf, ctx)
	}
}

func (b *barcode) GetType() string {
	return b._type.String()
}

func (b *barcode) Add(component ...v2.Component) v2.Component {
	return b
}
