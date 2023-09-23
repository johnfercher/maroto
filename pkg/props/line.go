package props

import (
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"
)

// Line represents properties from a Line inside a cell.
type Line struct {
	// Color define the linestyle color.
	Color Color
	// Style define the linestyle style (solid, dashed or dotted).
	Style         linestyle.Type
	Thickness     float64
	Orientation   orientation.Type
	OffsetPercent float64
	SizePercent   float64
}

// MakeValid from Line define default values for a Line.
func (s *Line) MakeValid() {
	if s.Style == "" {
		s.Style = linestyle.Solid
	}

	if s.Thickness == 0 {
		s.Thickness = linestyle.DefaultLineThickness
	}

	if s.Orientation == "" {
		s.Orientation = orientation.Horizontal
	}

	if s.OffsetPercent < 5 {
		s.OffsetPercent = 5
	}

	if s.OffsetPercent > 95 {
		s.OffsetPercent = 95
	}

	if s.SizePercent <= 0 {
		s.SizePercent = 90
	}

	if s.SizePercent > 100 {
		s.SizePercent = 100
	}
}
