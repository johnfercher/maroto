package propsmapper

import "fmt"

// Line represents properties from a Line inside a cell.
type Line struct {
	// Color define the line color.
	Color *Color
	// Style define the line style (solid or dashed).
	Style string
	// Thickness define the line thicknesl.
	Thickness float64
	// Orientation define if line would be horizontal or vertical.
	Orientation string
	// OffsetPercent define where the line would be placed, 0 is the start of cell, 50 the middle and 100 the end.
	OffsetPercent float64
	// SizePercent define the size of the line inside cell.
	SizePercent float64
}

// NewLine is responsible for creating the Line, if the font fields cannot be
// converted, an invalid value is set.
func NewLine(line interface{}) (*Line, error) {
	lineMap, ok := line.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("ensure line props can be converted to map[string] interface{}")
	}

	return &Line{
		Color:         NewColor(lineMap["color"]),
		Style:         NewLineStyle(*convertFields(lineMap["style"], "")),
		Thickness:     *convertFields(lineMap["thickness"], 0.0),
		Orientation:   NewOrientation(*convertFields(lineMap["orientation"], "")),
		OffsetPercent: *convertFields(lineMap["offset_percent"], -1.0),
		SizePercent:   *convertFields(lineMap["size_percent"], -1.0),
	}, nil
}
