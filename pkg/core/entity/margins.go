package entity

// Margins is the representation of a margin.
type Margins struct {
	Left   float64
	Right  float64
	Top    float64
	Bottom float64
}

// AppendMap appends the margins to a map.
func (m *Margins) AppendMap(mp map[string]interface{}) map[string]interface{} {
	if m.Left != 0 {
		mp["config_margin_left"] = m.Left
	}

	if m.Top != 0 {
		mp["config_margin_top"] = m.Top
	}

	if m.Right != 0 {
		mp["config_margin_right"] = m.Right
	}

	if m.Bottom != 0 {
		mp["config_margin_bottom"] = m.Bottom
	}

	return mp
}
