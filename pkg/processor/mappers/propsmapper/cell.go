package propsmapper

import (
	"fmt"
)

// Cell is the representation of a cell in the grid system.
// This can be applied to Col or Row
type Cell struct {
	// BackgroundColor defines which color will be applied to a cell.
	// Default: nil
	BackgroundColor *Color
	// BorderColor defines which color will be applied to a border cell
	// Default: nil
	BorderColor *Color
	// BorderType defines which kind of border will be applied to a cell.
	// Default: border.None
	BorderType string
	// BorderThickness defines the border thickness applied to a cell.
	// Default: 0.2
	BorderThickness float64
	// LineStyle defines which line style will be applied to a cell.
	// Default: Solid
	LineStyle string
}

// NewCell is responsible for creating the cell
func NewCell(cell interface{}) (*Cell, error) {
	cellMap, ok := cell.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("ensure cell props can be converted to map[string] interface{}")
	}

	return &Cell{
		BackgroundColor: NewColor(cellMap["background_color"]),
		BorderColor:     NewColor(cellMap["border_color"]),
		BorderType:      NewBorder(*convertFields(cellMap["border_type"], "")),
		BorderThickness: *convertFields(cellMap["border_thickness"], 0.0),
		LineStyle:       NewLineStyle(*convertFields(cellMap["line_style"], "")),
	}, nil
}
