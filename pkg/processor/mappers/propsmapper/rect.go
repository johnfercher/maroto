package propsmapper

import "fmt"

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
	// indicate whether only the width should be used as a reference to calculate the component size, disregarding the height
	// ex true: The component will be scaled only based on the available width, disregarding the available height
	JustReferenceWidth bool
	// Center define that the barcode will be vertically and horizontally centralized.
	Center bool
}

// NewRect is responsible for creating the Rect, if the font fields cannot be
// converted, an invalid value is set.
func NewRect(rect interface{}) (*Rect, error) {
	rectMap, ok := rect.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("ensure matrix code props can be converted to map[string] interface{}")
	}

	return &Rect{
		Left:               *convertFields(rectMap["left"], -1.0),
		Top:                *convertFields(rectMap["top"], -1.0),
		Percent:            *convertFields(rectMap["percent"], -1.0),
		JustReferenceWidth: *convertFields(rectMap["just_reference_width"], false),
		Center:             *convertFields(rectMap["center"], false),
	}, nil
}
