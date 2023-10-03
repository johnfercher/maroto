package props

// Rect represents properties from a rectangle (Image, QrCode or Barcode) inside a cell.
type Rect struct {
	// Left is the space between the left cell boundary to the rectangle, if center is false.
	Left float64
	// Top is space between the upper cell limit to the barcode, if center is false.
	Top float64
	// Percent is how much the rectangle will occupy the cell,
	// ex 100%: The rectangle will fulfill the entire cell
	// ex 50%: The greater side from the rectangle will have half the size of the cell.
	Percent float64
	// Center define that the barcode will be vertically and horizontally centralized.
	Center bool
}

func (r *Rect) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"prop_left":    r.Left,
		"prop_top":     r.Top,
		"prop_percent": r.Percent,
		"prop_center":  r.Center,
	}
}

// MakeValid from Rect will make the properties from a rectangle reliable to fit inside a cell
// and define default values for a rectangle.
func (r *Rect) MakeValid() {
	minPercentage := 0.0
	maxPercentage := 100.0
	minValue := 0.0

	if r.Percent <= minPercentage || r.Percent > maxPercentage {
		r.Percent = maxPercentage
	}

	if r.Center {
		r.Left = 0
		r.Top = 0
	}

	if r.Left < minValue {
		r.Left = minValue
	}

	if r.Top < minValue {
		r.Top = minValue
	}
}
