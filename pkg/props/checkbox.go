package props

// Checkbox represents properties from a checkbox inside a cell.
type Checkbox struct {
	Font
	// Checked indicates whether the checkbox is checked.
	Checked bool
	// BoxSize is the size of the checkbox box in PDF units.
	BoxSize float64
	// Left is the space between the left cell boundary to the checkbox.
	Left float64
	// Top is the space between the upper cell limit to the checkbox.
	Top float64
	// Right is the space between the checkbox to the right cell boundary.
	Right float64
	// Bottom is the space between the checkbox to the bottom cell boundary.
	Bottom float64
}

// ToMap returns a map representation of Checkbox properties.
func (c *Checkbox) ToMap() map[string]interface{} {
	m := make(map[string]interface{})

	if c.Checked {
		m["prop_checked"] = c.Checked
	}

	if c.BoxSize != 0 {
		m["prop_box_size"] = c.BoxSize
	}

	if c.Left != 0 {
		m["prop_left"] = c.Left
	}

	if c.Top != 0 {
		m["prop_top"] = c.Top
	}

	if c.Right != 0 {
		m["prop_right"] = c.Right
	}

	if c.Bottom != 0 {
		m["prop_bottom"] = c.Bottom
	}

	m = c.Font.AppendMap(m)

	return m
}

// MakeValid ensures Checkbox properties are valid and sets defaults.
func (c *Checkbox) MakeValid(defaultFont *Font) {
	const (
		defaultBoxSize = 4.0
		minBoxSize     = 1.0
		maxBoxSize     = 20.0
		minValue       = 0.0
	)

	if c.BoxSize < minBoxSize || c.BoxSize > maxBoxSize {
		c.BoxSize = defaultBoxSize
	}

	if c.Left < minValue {
		c.Left = minValue
	}

	if c.Top < minValue {
		c.Top = minValue
	}

	if c.Right < minValue {
		c.Right = minValue
	}

	if c.Bottom < minValue {
		c.Bottom = minValue
	}

	if defaultFont != nil {
		c.Font.MakeValid(defaultFont.Family)
	} else {
		c.Font.MakeValid("")
	}
}
