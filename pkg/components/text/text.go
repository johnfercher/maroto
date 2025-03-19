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

// New is responsible for creating an instance of a Text. It will create a subtext
// with value and use props to derive the subtext props
func New(value string, ps ...props.Text) *Text {
	textProp := props.Text{}
	subText := []*entity.SubText{}

	if len(ps) > 0 {
		textProp = ps[0]
	}
	if len(value) > 0 {
		subText = append(subText, &entity.SubText{Value: value, Props: props.NewSubText(&textProp)})
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

// NewAutoRow is responsible for creating an instance of Text grouped in a Line with automatic height.
func NewAutoRow(value string, ps ...props.Text) core.Row {
	r := New(value, ps...)
	c := col.New().Add(r)
	return row.New().Add(c)
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

// AddSubText will be used to add a subText to the current text.
// This method allows you to use different styles on the same text
func (t *Text) AddSubText(text string, ps ...props.SubText) *Text {
	subTextPs := props.SubText{}
	if len(ps) > 0 {
		subTextPs = ps[0]
	}

	t.text = append(t.text, &entity.SubText{Value: text, Props: subTextPs})
	return t
}

func (t *Text) getStructSubText(sub *entity.SubText) *node.Node[core.Structure] {
	str := core.Structure{
		Type:    "sub_text",
		Value:   sub.Value,
		Details: sub.Props.ToMap(),
	}

	return node.New(str)
}

// GetHeight returns the height that the text will have in the PDF
func (t *Text) GetHeight(provider core.Provider, cell *entity.Cell) float64 {
	amountLines := provider.GetTextHeight(t.text, &t.props, cell.Width-t.props.Left-t.props.Right)
	return amountLines + t.props.Top + t.props.Bottom
}

// SetConfig sets the config.
func (t *Text) SetConfig(config *entity.Config) {
	t.config = config
	t.props.MakeValid(t.config.DefaultFont)
	for _, sub := range t.text {
		sub.Props.MakeValid(t.config.DefaultFont)
	}
}

// Render renders a Text into a PDF context.
func (t *Text) Render(provider core.Provider, cell *entity.Cell) {
	innerCell := cell.Copy()
	provider.AddCustomText(t.text, &innerCell, &t.props)
}
