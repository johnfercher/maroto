package props

import (
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
)

// Font represents properties from a text.
type Font struct {
	// Family of the text, ex: consts.Arial, helvetica and etc.
	Family string
	// Style of the text, ex: consts.Normal, bold and etc.
	Style fontstyle.Type
	// Size of the text.
	Size float64
	// Color define the fontstyle color.
	Color *Color
}

// MakeValid from Font define default values for a Signature.
func (s *Font) MakeValid(defaultFamily string) {
	undefinedValue := 0.0

	if s.Family == "" {
		s.Family = defaultFamily
	}

	if s.Style == "" {
		s.Style = fontstyle.Bold
	}

	if s.Size == undefinedValue {
		s.Size = 8.0
	}
}

// ToTextProp from Font return a Text based on Font.
func (s *Font) ToTextProp(align align.Type, top float64, verticalPadding float64) *Text {
	textProp := &Text{
		Family:          s.Family,
		Style:           s.Style,
		Size:            s.Size,
		Align:           align,
		Top:             top,
		VerticalPadding: verticalPadding,
		Color:           s.Color,
	}

	textProp.MakeValid(s)

	return textProp
}
