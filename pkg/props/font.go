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

// MakeValid from Font define default values for a Signature.
func (f *Font) MakeValid(defaultFamily string) {
	if f.Family == "" {
		f.Family = defaultFamily
	}

	if f.Style == "" {
		f.Style = fontstyle.Bold
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
