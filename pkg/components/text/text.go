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
	text   []*entity.SubText
	config *entity.Config
	props  props.Text
}

// New is responsible to create an instance of a Text.
func New(value string, ps ...props.Text) core.Component {
	textProp := props.Text{}
	if len(ps) > 0 {
		textProp = ps[0]
	}

	return &Text{
		text:  []*entity.SubText{entity.NewSubText(value, props.NewSubText(&textProp))},
		props: textProp,
	}
}

// NewCustomText is responsible for creating an instance of a Text based on SubTexts
func NewCustomText(subText []*entity.SubText, ps ...props.Text) core.Component {
	textProp := props.Text{}
	if len(ps) > 0 {
		textProp = ps[0]
	}
	if len(subText) == 0 {
		subText = append(subText, entity.NewSubText(""))
	}

	return &Text{
		text:  subText,
		props: textProp,
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
	node := node.New(
		core.Structure{
			Type:    "text",
			Details: t.props.ToMap(),
		})

	for _, sub := range t.text {
		node.AddNext(t.getStructSubText(sub))
	}

	return node
}

func (t *Text) getStructSubText(sub *entity.SubText) *node.Node[core.Structure] {
	str := core.Structure{
		Type:    "sub_text",
		Value:   sub.Value,
		Details: sub.Prop.ToMap(),
	}

	return node.New(str)
}

// SetConfig sets the config.
func (t *Text) SetConfig(config *entity.Config) {
	t.config = config
	t.props.MakeValid(t.config.DefaultFont)
	for _, sub := range t.text {
		sub.Prop.MakeValid(t.config.DefaultFont)
	}
}

// Render renders a Text into a PDF context.
func (t *Text) Render(provider core.Provider, cell *entity.Cell) {
	innerCell := cell.Copy()
	provider.AddCustomText(t.text, &innerCell, &t.props)
}
