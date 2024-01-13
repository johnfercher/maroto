// Package signature implements creation of signatures.
package signature

import (
	"github.com/johnfercher/go-tree/node"

	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type Signature struct {
	value  string
	prop   props.Signature
	config *entity.Config
}

// New is responsible to create an instance of a Signature.
func New(value string, ps ...props.Signature) core.Component {
	prop := props.Signature{}
	if len(ps) > 0 {
		prop = ps[0]
	}
	prop.MakeValid(fontfamily.Arial)

	return &Signature{
		value: value,
		prop:  prop,
	}
}

// NewCol is responsible to create an instance of a Signature wrapped in a Col.
func NewCol(size int, value string, ps ...props.Signature) core.Col {
	signature := New(value, ps...)
	return col.New(size).Add(signature)
}

// NewRow is responsible to create an instance of a Signature wrapped in a Row.
func NewRow(height float64, value string, ps ...props.Signature) core.Row {
	signature := New(value, ps...)
	c := col.New().Add(signature)
	return row.New(height).Add(c)
}

// Render renders a Signature into a PDF context.
func (s *Signature) Render(provider core.Provider, cell *entity.Cell) {
	fontProp := s.prop.ToFontProp()
	safePadding := 1.5
	fontSize := provider.GetTextHeight(fontProp) * safePadding

	textProp := s.prop.ToTextProp(align.Center, cell.Height-fontSize, 0)

	offsetPercent := (cell.Height - fontSize) / cell.Height * 100.0

	provider.AddText(s.value, cell, textProp)
	provider.AddLine(cell, s.prop.ToLineProp(offsetPercent))
}

// GetStructure returns the Structure of a Signature.
func (s *Signature) GetStructure() *node.Node[core.Structure] {
	str := core.Structure{
		Type:    "signature",
		Value:   s.value,
		Details: s.prop.ToMap(),
	}

	return node.New(str)
}

// SetConfig sets the config.
func (s *Signature) SetConfig(config *entity.Config) {
	s.config = config
}
