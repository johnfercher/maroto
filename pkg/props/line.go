package props

import "github.com/johnfercher/maroto/v2/pkg/consts/linestyle"

// Line represents properties from a Line inside a cell.
type Line struct {
	// Color define the line color.
	Color Color
	// Style define the line style (solid, dashed or dotted).
	Style linestyle.Type
	// Width define the line width (it cannot be greater than line height)
	Width float64
}

// MakeValid from Line define default values for a Line.
func (s *Line) MakeValid(spaceHeight float64) {
	if s.Style == "" {
		s.Style = linestyle.Solid
	}

	if s.Width == 0 {
		s.Width = linestyle.DefaultLineWidth
	}

	if s.Width > spaceHeight {
		s.Width = spaceHeight
	}
}
