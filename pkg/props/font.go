// Package contains all props used to customize components.
package props

import (
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
)

// Font represents properties from a text.
type Font struct {
	// Family of the text, ex: constf.Arial, helvetica and etc.
	Family string
	// Style of the text, ex: constf.Normal, bold and etc.
	Style fontstyle.Type
	// Size of the text.
	Size float64
	// Color define the font color.
	Color *Color
}

func (f *Font) AppendMap(m map[string]interface{}) map[string]interface{} {
	if f.Family != "" {
		m["prop_font_family"] = f.Family
	}

	if f.Style != "" {
		m["prop_font_style"] = f.Style
	}

	if f.Size != 0 {
		m["prop_font_size"] = f.Size
	}

	if f.Color != nil {
		m["prop_font_color"] = f.Color.ToString()
	}

	return m
}

// MakeValid from Font define default values for a Signature.
func (f *Font) MakeValid(defaultFamily string) {
	if f.Family == "" {
		f.Family = defaultFamily
	}

	if f.Style == "" {
		f.Style = fontstyle.Normal
	}

	if f.Size == 0.0 {
		f.Size = 8.0
	}
}

// ToTextProp from Font return a Text based on Font.
func (f *Font) ToTextProp(align align.Type, top float64, verticalPadding float64) *Text {
	textProp := &Text{
		Family:          f.Family,
		Style:           f.Style,
		Size:            f.Size,
		Align:           align,
		Top:             top,
		VerticalPadding: verticalPadding,
		Color:           f.Color,
	}

	textProp.MakeValid(f)

	return textProp
}
