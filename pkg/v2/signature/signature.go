package signature

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/internal/fpdf"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/johnfercher/maroto/pkg/v2/context"
	"github.com/johnfercher/maroto/pkg/v2/domain"
	"github.com/johnfercher/maroto/pkg/v2/types"
)

type signature struct {
	value      string
	_type      types.DocumentType
	components []domain.Node
	prop       props.Text
}

func New(value string, textProps ...props.Text) domain.Component {
	prop := props.Text{}
	if len(textProps) > 0 {
		prop = textProps[0]
	}
	prop.MakeValid(consts.Arial)

	return &signature{
		_type: types.Signature,
		value: value,
		prop:  prop,
	}
}

func (s *signature) Render(fpdf fpdf.Fpdf, ctx context.Context) {
	font := internal.NewFont(fpdf, 2, consts.Arial, consts.Normal)
	math := internal.NewMath(fpdf)
	text := internal.NewText(fpdf, math, font)
	signature := internal.NewSignature(fpdf, math, text)
	x := fpdf.GetX() - ctx.Margins.Left - ctx.Dimensions.Width
	y := fpdf.GetY() - ctx.Margins.Top
	signature.AddSpaceFor(
		s.value,
		internal.Cell{x, y, ctx.Dimensions.Width, ctx.Dimensions.Height},
		s.prop)
}

func (s *signature) GetType() string {
	return s._type.String()
}

func (s *signature) Add(_ ...domain.Node) domain.Node {
	return s
}

func (s *signature) GetStructure() *tree.Node[domain.Structure] {
	str := domain.Structure{
		Type:  string(s._type),
		Value: s.value,
	}

	return tree.NewNode(0, str)
}
