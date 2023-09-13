package barcode

import (
	"fmt"
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/internal/fpdf"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/johnfercher/maroto/pkg/v2/context"
	"github.com/johnfercher/maroto/pkg/v2/domain"
	"github.com/johnfercher/maroto/pkg/v2/text"
	"github.com/johnfercher/maroto/pkg/v2/types"
)

type barcode struct {
	code       string
	_type      types.DocumentType
	components []domain.Node
	prop       props.Barcode
}

func New(code string, barcodeProps ...props.Barcode) domain.Component {
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
	x := fpdf.GetX() - ctx.Margins.Left - ctx.Dimensions.Width
	y := fpdf.GetY() - ctx.Margins.Top
	fmt.Println(x, y)
	err := code.AddBar(b.code,
		internal.Cell{x, y, ctx.Dimensions.Width, ctx.Dimensions.Height},
		b.prop)

	if err != nil {
		fpdf.ClearError()
		txt := text.New("Failed to render code")
		txt.Render(fpdf, ctx)
	}
}

func (b *barcode) GetType() string {
	return b._type.String()
}

func (b *barcode) Add(component ...domain.Node) domain.Node {
	return b
}

func (b *barcode) GetStructure() *tree.Node[domain.Structure] {
	str := domain.Structure{
		Type:  string(b._type),
		Value: b.code,
	}

	return tree.NewNode(0, str)
}
