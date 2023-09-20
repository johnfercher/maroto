package props

// Barcode represents properties from a barcode inside a cell.
type Barcode struct {
	// Left is the space between the left cell boundary to the barcode, if center is false.
	Left float64
	// Top is space between the upper cell limit to the barcode, if center is false.
	Top float64
	// Percent is how much the barcode will occupy the cell,
	// ex 100%: The barcode will fulfill the entire cell
	// ex 50%: The greater side from the barcode will have half the size of the cell.
	Percent float64
	// Proportion is the proportion between size of the barcode.
	// Ex: 16x9, 4x3...
	Proportion Proportion
	// Center define that the barcode will be vertically and horizontally centralized.
	Center bool
}

// MakeValid from Barcode will make the properties from a barcode reliable to fit inside a cell
// and define default values for a barcode.
func (s *Barcode) MakeValid() {
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

	if s.Proportion.Width <= 0 {
		s.Proportion.Width = 1
	}

	if s.Proportion.Height <= 0 {
		s.Proportion.Height = 1
	}

	maxHeightProportionBasedOnWidth := 0.20
	minHeightProportionBasedOnWidth := 0.10

	if s.Proportion.Height > s.Proportion.Width*maxHeightProportionBasedOnWidth {
		s.Proportion.Height = s.Proportion.Width * maxHeightProportionBasedOnWidth
	} else if s.Proportion.Height < s.Proportion.Width*minHeightProportionBasedOnWidth {
		s.Proportion.Height = s.Proportion.Width * minHeightProportionBasedOnWidth
	}
}
