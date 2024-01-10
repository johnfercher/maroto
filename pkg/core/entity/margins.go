// Package contains all core entities.
package entity

type Margins struct {
	Left   float64
	Right  float64
	Top    float64
	Bottom float64
}

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
