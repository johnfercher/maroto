// Package text implements creation of texts.
package text

import (
	"github.com/johnfercher/go-tree/node"

	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type Text struct {
	value  string
	prop   props.Text
	config *entity.Config
}

// New is responsible to create an instance of a Text.
func New(value string, ps ...props.Text) core.Component {
	textProp := props.Text{}
	if len(ps) > 0 {
		textProp = ps[0]
	}

	return &Text{
		value: value,
		prop:  textProp,
	}
}

// NewCol is responsible to create an instance of a Text wrapped in a Col.
func NewCol(size int, value string, ps ...props.Text) core.Col {
	text := New(value, ps...)
	return col.New(size).Add(text)
}

// NewRow is responsible to create an instance of a Text wrapped in a Row.
func NewRow(height float64, value string, ps ...props.Text) core.Row {
	r := New(value, ps...)
	c := col.New().Add(r)
	return row.New(height).Add(c)
}

// GetStructure returns the Structure of a Text.
func (t *Text) GetStructure() *node.Node[core.Structure] {
	str := core.Structure{
		Type:    "text",
		Value:   t.value,
		Details: t.prop.ToMap(),
	}

	return node.New(str)
}

// SetConfig sets the config.
func (t *Text) SetConfig(config *entity.Config) {
	t.config = config
	t.prop.MakeValid(t.config.DefaultFont)
}

// Render renders a Text into a PDF context.
func (t *Text) Render(provider core.Provider, cell *entity.Cell) {
	provider.AddText(t.value, cell, &t.prop)
}
