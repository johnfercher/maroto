package props

import (
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/text"
)

// Text represents properties from a Text inside a cell.
type Text struct {
	// Top is the amount of space between the upper cell limit and the text.
	Top float64
	// Bottom is the amount of space from bottom line of cell. Work only with adaptive rows
	Bottom float64
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
	// Extrapolate define if the text will automatically add a new line when.
	// text reach the right cell boundary.
	Extrapolate bool
	// VerticalPadding define an additional space between lines.
	VerticalPadding float64
	// Color define the fontstyle color.
	Color *Color
	// ExtrapolateStrategy strategy of text extrapolation
	ExtrapolateStrategy text.ExtrapolateStrategy
}

// MakeValid from Text define default values for a Text.
func (s *Text) MakeValid(font *Font) {
	minValue := 0.0
	undefinedValue := 0.0

	if s.Family == "" {
		s.Family = font.Family
	}

	if s.Style == "" {
		s.Style = font.Style
	}

	if s.Size == undefinedValue {
		s.Size = font.Size
	}

	if s.Color == nil {
		s.Color = font.Color
	}

	if s.Align == "" {
		s.Align = align.Center
	}

	if s.Top < minValue {
		s.Top = minValue
	}

	if s.Left < minValue {
		s.Left = minValue
	}

	if s.Right < minValue {
		s.Right = minValue
	}

	if s.VerticalPadding < 0 {
		s.VerticalPadding = 0
	}

	if s.ExtrapolateStrategy == "" {
		s.ExtrapolateStrategy = text.ExtrapolateStrategyWords
	}
}
