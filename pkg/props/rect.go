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

// MakeValid from Rect will make the properties from a rectangle reliable to fit inside a cell
// and define default values for a rectangle.
func (s *Rect) MakeValid() {
	minPercentage := 0.0
	maxPercentage := 100.0
	minValue := 0.0

	if s.Percent <= minPercentage || s.Percent > maxPercentage {
		s.Percent = maxPercentage
	}

	if s.Center {
		s.Left = 0
		s.Top = 0
	}

	if s.Left < minValue {
		s.Left = minValue
	}

	if s.Top < minValue {
		s.Top = minValue
	}
}
