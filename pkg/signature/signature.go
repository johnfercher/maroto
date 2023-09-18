package signature

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/v2/internal"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts"
	"github.com/johnfercher/maroto/v2/pkg/domain"
	"github.com/johnfercher/maroto/v2/pkg/grid/col"
	"github.com/johnfercher/maroto/v2/pkg/grid/row"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type signature struct {
	value  string
	prop   props.Font
	config *config.Maroto
}

func New(value string, ps ...props.Font) domain.Component {
	prop := props.Font{}
	if len(ps) > 0 {
		prop = ps[0]
	}
	prop.MakeValid(consts.Arial)

	return &signature{
		value: value,
		prop:  prop,
	}
}

func NewCol(size int, value string, ps ...props.Font) domain.Col {
	signature := New(value, ps...)
	return col.New(size).Add(signature)
}

func NewRow(height float64, value string, ps ...props.Font) domain.Row {
	signature := New(value, ps...)
	c := col.New().Add(signature)
	return row.New(height).Add(c)
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

func (s *signature) SetConfig(config *config.Maroto) {
	s.config = config
}
