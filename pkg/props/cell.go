package props

import (
	"github.com/johnfercher/maroto/v2/pkg/consts/border"
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
)

type Cell struct {
	BackgroundColor *Color
	BorderColor     *Color
	BorderType      border.Type
	BorderThickness float64
	LineStyle       linestyle.Type
}

func (c *Cell) ToMap() map[string]interface{} {
	if c == nil {
		return nil
	}

	m := make(map[string]interface{})

	if c.BorderType != "" {
		m["cell_prop_border_type"] = c.BorderType
	}

	if c.BorderThickness != 0 {
		m["cell_prop_border_thickness"] = c.BorderThickness
	}

	if c.LineStyle != "" {
		m["cell_prop_border_line_style"] = c.LineStyle
	}

	if c.BackgroundColor != nil {
		m["cell_prop_backgrond_color"] = c.BackgroundColor.ToString()
	}

	if c.BorderColor != nil {
		m["cell_prop_border_color"] = c.BorderColor.ToString()
	}

	return m
}
