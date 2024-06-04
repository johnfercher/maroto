// Package props contains all props used to customize components.
package props

import (
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/breakline"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
)

// SubText represents properties from a SubText inside a cell.
type SubText struct {
	// Family of the text, ex: consts.Arial, helvetica and etc.
	Family string
	// Style of the text, ex: consts.Normal, bold and etc.
	Style fontstyle.Type
	// Size of the text.
	Size float64
	// Color define the font style color.
	Color *Color
	// Hyperlink define a link to be opened when the text is clicked.
	Hyperlink *string
}

func (sub *SubText) MakeValid(font *Font) {
	undefinedValue := 0.0

	if sub.Family == "" {
		sub.Family = font.Family
	}

	if sub.Style == "" {
		sub.Style = font.Style
	}

	if sub.Size == undefinedValue {
		sub.Size = font.Size
	}

	if sub.Color == nil {
		sub.Color = font.Color
	}
}

func (t *SubText) ToMap() map[string]interface{} {
	m := make(map[string]interface{})
	if t.Family != "" {
		m["prop_font_family"] = t.Family
	}

	if t.Style != "" {
		m["prop_font_style"] = t.Style
	}

	if t.Size != 0 {
		m["prop_font_size"] = t.Size
	}

	if t.Color != nil {
		m["prop_color"] = t.Color.ToString()
	}

	if t.Hyperlink != nil {
		m["prop_hyperlink"] = *t.Hyperlink
	}

	return m
}

func NewSubText(t *Text) SubText {
	return SubText{
		Family:    t.Family,
		Style:     t.Style,
		Size:      t.Size,
		Color:     t.Color,
		Hyperlink: t.Hyperlink,
	}
}

// Text represents properties from a Text inside a cell.
type Text struct {
	// Top is the amount of space between the upper cell limit and the text.
	Top float64
	// Left is the minimal amount of space between the left cell boundary and the text.
	Left float64
	// Right is the minimal amount of space between the right cell boundary and the text.
	Right float64
	// Family of the text, ex: consts.Arial, helvetica and etc.
	Family string
	// Style of the text, ex: consts.Normal, bold and etc.
	Style fontstyle.Type
	// Size of the text.
	Size float64
	// Align of the text.
	Align align.Type
	// BreakLineStrategy define the break line strategy.
	BreakLineStrategy breakline.Strategy
	// VerticalPadding define an additional space between linet.
	VerticalPadding float64
	// Color define the font style color.
	Color *Color
	// Hyperlink define a link to be opened when the text is clicked.
	Hyperlink *string
}

// ToMap converts a Text to a map.
func (t *Text) ToMap() map[string]interface{} {
	m := make(map[string]interface{})
	if t.Top != 0 {
		m["prop_top"] = t.Top
	}

	if t.Left != 0 {
		m["prop_left"] = t.Left
	}

	if t.Right != 0 {
		m["prop_right"] = t.Right
	}

	if t.Family != "" {
		m["prop_font_family"] = t.Family
	}

	if t.Style != "" {
		m["prop_font_style"] = t.Style
	}

	if t.Size != 0 {
		m["prop_font_size"] = t.Size
	}

	if t.Align != "" {
		m["prop_align"] = t.Align
	}

	if t.BreakLineStrategy != "" {
		m["prop_breakline_strategy"] = t.BreakLineStrategy
	}

	if t.VerticalPadding != 0 {
		m["prop_vertical_padding"] = t.VerticalPadding
	}

	if t.Color != nil {
		m["prop_color"] = t.Color.ToString()
	}

	if t.Hyperlink != nil {
		m["prop_hyperlink"] = *t.Hyperlink
	}

	return m
}

// MakeValid from Text define default values for a Text.
func (t *Text) MakeValid(font *Font) {
	minValue := 0.0
	undefinedValue := 0.0

	if t.Family == "" {
		t.Family = font.Family
	}

	if t.Style == "" {
		t.Style = font.Style
	}

	if t.Size == undefinedValue {
		t.Size = font.Size
	}

	if t.Color == nil {
		t.Color = font.Color
	}

	if t.Align == "" {
		t.Align = align.Left
	}

	if t.Top < minValue {
		t.Top = minValue
	}

	if t.Left < minValue {
		t.Left = minValue
	}

	if t.Right < minValue {
		t.Right = minValue
	}

	if t.VerticalPadding < 0 {
		t.VerticalPadding = 0
	}

	if t.BreakLineStrategy == "" {
		t.BreakLineStrategy = breakline.EmptySpaceStrategy
	}
}
