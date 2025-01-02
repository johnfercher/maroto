package propsmapper

// Proportion represents a proportion from a rectangle, example: 16x9, 4x3...
type Proportion struct {
	// Width from the rectangle: Barcode, image and etc.
	Width float64
	// Height from the rectangle: Barcode, image and etc.
	Height float64
}

// NewProportion is responsible for creating the proportion, if the font fields cannot be
// converted, an invalid value is set.
func NewProportion(barcode interface{}) Proportion {
	barcodeMap, _ := barcode.(map[string]interface{})

	return Proportion{
		Width:  *convertFields(barcodeMap["width"], 0.0),
		Height: *convertFields(barcodeMap["height"], 0.0),
	}
}
