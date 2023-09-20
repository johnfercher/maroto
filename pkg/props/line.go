package props

import "github.com/johnfercher/maroto/v2/pkg/consts/line"

// Line represents properties from a Line inside a cell.
type Line struct {
	// Color define the line color.
	Color Color
	// Style define the line style (solid, dashed or dotted).
	Style line.Style
	// Width define the line width (it cannot be greater than line height)
	Width float64
}

// MakeValid from Line define default values for a Line.
func (s *Line) MakeValid(spaceHeight float64) {
	if s.Style == "" {
		s.Style = line.Solid
	}

	if s.Width == 0 {
		s.Width = line.DefaultLineWidth
	}

	if s.Width > spaceHeight {
		s.Width = spaceHeight
	}
}
