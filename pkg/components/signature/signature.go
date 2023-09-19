package signature

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type signature struct {
	value  string
	prop   props.Font
	config *config.Config
}

func New(value string, ps ...props.Font) core.Component {
	prop := props.Font{}
	if len(ps) > 0 {
		prop = ps[0]
	}
	prop.MakeValid(fontfamily.Arial)

	return &signature{
		value: value,
		prop:  prop,
	}
}

func NewCol(size int, value string, ps ...props.Font) core.Col {
	signature := New(value, ps...)
	return col.New(size).Add(signature)
}

func NewRow(height float64, value string, ps ...props.Font) core.Row {
	signature := New(value, ps...)
	c := col.New().Add(signature)
	return row.New(height).Add(c)
}

func (s *signature) Render(provider core.Provider, cell core.Cell) {
	provider.AddSignature(s.value, cell, s.prop.ToTextProp(align.Center, 0.0, false, 0))
}

func (s *signature) GetStructure() *tree.Node[core.Structure] {
	str := core.Structure{
		Type:  "signature",
		Value: s.value,
	}

	return tree.NewNode(str)
}

func (s *signature) SetConfig(config *config.Config) {
	s.config = config
}
