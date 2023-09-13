package signature

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/internal/fpdf"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
	v2 "github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/context"
)

type signature struct {
	value      string
	_type      v2.DocumentType
	components []v2.Component
	prop       props.Text
}

func New(value string, textProps ...props.Text) *signature {
	prop := props.Text{}
	if len(textProps) > 0 {
		prop = textProps[0]
	}
	prop.MakeValid(consts.Arial)

	return &signature{
		_type: v2.Signature,
		value: value,
		prop:  prop,
	}
}

func (s *signature) Render(fpdf fpdf.Fpdf, ctx context.Context) {
	font := internal.NewFont(fpdf, 2, consts.Arial, consts.Normal)
	math := internal.NewMath(fpdf)
	text := internal.NewText(fpdf, math, font)
	signature := internal.NewSignature(fpdf, math, text)

	signature.AddSpaceFor(
		s.value,
		internal.Cell{fpdf.GetX() - ctx.Margins.Left,
			fpdf.GetY() - ctx.Margins.Top,
			ctx.Dimensions.Width,
			ctx.Dimensions.Height},
		s.prop)
}

func (s *signature) GetType() string {
	return s._type.String()
}

func (s *signature) Add(_ ...v2.Component) v2.Component {
	return s
}

func (s *signature) GetStructure() *tree.Node[v2.Structure] {
	str := v2.Structure{
		Type:  string(s._type),
		Value: s.value,
	}

	return tree.NewNode(0, str)
}
