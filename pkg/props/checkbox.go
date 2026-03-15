package props

// Checkbox represents properties from a Checkbox inside a cell.
type Checkbox struct {
	// Checked defines whether the checkbox is marked.
	Checked bool
	// Top is the amount of space between the upper cell limit and the checkbox.
	Top float64
	// Left is the amount of space between the left cell boundary and the checkbox.
	Left float64
	// Size is the size of the checkbox square in mm.
	Size float64
}

// ToMap converts a Checkbox to a map.
func (c *Checkbox) ToMap() map[string]any {
	m := make(map[string]any)

	if c.Checked {
		m["prop_checked"] = c.Checked
	}

	if c.Top != 0 {
		m["prop_top"] = c.Top
	}

	if c.Left != 0 {
		m["prop_left"] = c.Left
	}

	if c.Size != 0 {
		m["prop_size"] = c.Size
	}

	return m
}

// MakeValid from Checkbox define default values for a Checkbox.
func (c *Checkbox) MakeValid() {
	if c.Size <= 0 {
		c.Size = 5.0
	}

	if c.Top < 0 {
		c.Top = 0
	}

	if c.Left < 0 {
		c.Left = 0
	}
}
