package signature

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/johnfercher/maroto/pkg/v2/domain"
)

type signature struct {
	value string
	prop  props.Font
}

func New(value string, textProps ...props.Font) domain.Node {
	prop := props.Font{}
	if len(textProps) > 0 {
		prop = textProps[0]
	}
	prop.MakeValid(consts.Arial)

	return &signature{
		value: value,
		prop:  prop,
	}
}

func (s *signature) Render(provider domain.Provider, cell internal.Cell) {
	provider.AddSignature(s.value, cell, s.prop.ToTextProp(consts.Center, 0.0, false, 0))
}

func (s *signature) GetStructure() *tree.Node[domain.Structure] {
	str := domain.Structure{
		Type:  "signature",
		Value: s.value,
	}

	return tree.NewNode(str)
}
