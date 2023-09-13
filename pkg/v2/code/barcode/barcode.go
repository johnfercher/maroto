package barcode

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/internal/fpdf"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/context"
	"github.com/johnfercher/maroto/pkg/v2/text"
	"github.com/johnfercher/maroto/pkg/v2/types"
)

type barcode struct {
	code       string
	_type      types.DocumentType
	components []v2.Node
	prop       props.Barcode
}

func New(code string, barcodeProps ...props.Barcode) v2.Component {
	prop := props.Barcode{}
	if len(barcodeProps) > 0 {
		prop = barcodeProps[0]
	}
	prop.MakeValid()

	return &barcode{
		_type: types.Barcode,
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

func (b *barcode) Add(component ...v2.Node) v2.Node {
	return b
}

func (b *barcode) GetStructure() *tree.Node[v2.Structure] {
	str := v2.Structure{
		Type:  string(b._type),
		Value: b.code,
	}

	return tree.NewNode(0, str)
}
