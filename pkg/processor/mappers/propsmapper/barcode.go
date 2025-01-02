package propsmapper

import "fmt"

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
	// Type represents the barcode type. Default: code128
	Type string
}

// NewBarcode is responsible for creating the barcode, if the font fields cannot be
// converted, an invalid value is set.
func NewBarcode(barcode interface{}) (*Barcode, error) {
	barcodeMap, ok := barcode.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("ensure barcode props can be converted to map[string] interface{}")
	}

	return &Barcode{
		Left:       *convertFields(barcodeMap["left"], -1.0),
		Top:        *convertFields(barcodeMap["top"], -1.0),
		Percent:    *convertFields(barcodeMap["percent"], -1.0),
		Proportion: NewProportion(barcodeMap["proportion"]),
		Center:     *convertFields(barcodeMap["center"], false),
		Type:       NewCodeType(*convertFields(barcodeMap["type"], "")),
	}, nil
}
