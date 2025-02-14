package props

import (
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"
)

// Line represents properties from a Line inside a cell.
type Line struct {
	// Color define the line color.
	Color *Color
	// Style define the line style (solid or dashed).
	Style linestyle.Type
	// Thickness define the line thickness.
	Thickness float64
	// Orientation define if line would be horizontal or vertical.
	Orientation orientation.Type
	// OffsetPercent define where the line would be placed, 0 is the start of cell, 50 the middle and 100 the end.
	OffsetPercent float64
	// SizePercent define the size of the line inside cell.
	SizePercent float64
}

// ToMap returns a map with the Line fields.
func (l *Line) ToMap() map[string]interface{} {
	if l == nil {
		return nil
	}

	m := make(map[string]interface{})

	if l.Color != nil {
		m["prop_color"] = l.Color.ToString()
	}

	if l.Style != "" {
		m["prop_style"] = l.Style
	}

	if l.Thickness != 0 {
		m["prop_thickness"] = l.Thickness
	}

	if l.Orientation != "" {
		m["prop_orientation"] = l.Orientation
	}

	if l.OffsetPercent != 0 {
		m["prop_offset_percent"] = l.OffsetPercent
	}

	if l.SizePercent != 0 {
		m["prop_size_percent"] = l.SizePercent
	}

	return m
}

// MakeValid from Line define default values for a Line.
func (l *Line) MakeValid() {
	if l.Style == "" {
		l.Style = linestyle.Solid
	}

	if l.Thickness == 0 {
		l.Thickness = linestyle.DefaultLineThickness
	}

	if l.Orientation == "" {
		l.Orientation = orientation.Horizontal
	}

	if l.OffsetPercent < 5 {
		l.OffsetPercent = 5
	}

	if l.OffsetPercent > 95 {
		l.OffsetPercent = 95
	}

	if l.SizePercent <= 0 {
		l.SizePercent = 90
	}

	if l.SizePercent > 100 {
		l.SizePercent = 100
	}
}
