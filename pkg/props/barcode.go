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

func (b *Barcode) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"prop_left":              b.Left,
		"prop_top":               b.Top,
		"prop_percent":           b.Percent,
		"prop_proportion_width":  b.Proportion.Width,
		"prop_proportion_height": b.Proportion.Height,
		"prop_center":            b.Center,
	}
}

func (b *Barcode) ToRectProp() *Rect {
	return &Rect{
		Left:    b.Left,
		Top:     b.Top,
		Percent: b.Percent,
		Center:  b.Center,
	}
}

// MakeValid from Barcode will make the properties from a barcode reliable to fit inside a cell
// and define default values for a barcode.
func (b *Barcode) MakeValid() {
	minPercentage := 0.0
	maxPercentage := 100.0
	minValue := 0.0

	if b.Percent <= minPercentage || b.Percent > maxPercentage {
		b.Percent = maxPercentage
	}

	if b.Center {
		b.Left = 0
		b.Top = 0
	}

	if b.Left < minValue {
		b.Left = minValue
	}

	if b.Top < minValue {
		b.Top = minValue
	}

	if b.Proportion.Width <= 0 {
		b.Proportion.Width = 1
	}

	if b.Proportion.Height <= 0 {
		b.Proportion.Height = 1
	}

	maxHeightProportionBasedOnWidth := 0.20
	minHeightProportionBasedOnWidth := 0.10

	if b.Proportion.Height > b.Proportion.Width*maxHeightProportionBasedOnWidth {
		b.Proportion.Height = b.Proportion.Width * maxHeightProportionBasedOnWidth
	} else if b.Proportion.Height < b.Proportion.Width*minHeightProportionBasedOnWidth {
		b.Proportion.Height = b.Proportion.Width * minHeightProportionBasedOnWidth
	}
}
