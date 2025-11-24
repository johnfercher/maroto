package props

import (
	"github.com/johnfercher/maroto/v2/pkg/consts/border"
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
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
	BorderType border.Type
	// BorderThickness defines the border thickness applied to a cell.
	// Default: 0.2
	BorderThickness float64
	// LineStyle defines which line style will be applied to a cell.
	// Default: Solid
	LineStyle linestyle.Type
}

// ToMap adds the Cell fields to the map.
func (c *Cell) ToMap() map[string]interface{} {
	if c == nil {
		return nil
	}

	m := make(map[string]interface{})

	if c.BorderType != border.None {
		m["prop_border_type"] = c.BorderType
	}

	if c.BorderThickness != 0 {
		m["prop_border_thickness"] = c.BorderThickness
	}

	if c.LineStyle != "" {
		m["prop_border_line_style"] = c.LineStyle
	}

	if c.BackgroundColor != nil {
		m["prop_background_color"] = c.BackgroundColor.ToString()
	}

	if c.BorderColor != nil {
		m["prop_border_color"] = c.BorderColor.ToString()
	}

	return m
}
