package propsmapper

// Dimensions is the representation of a width and height.
type Dimensions struct {
	Width  float64
	Height float64
}

// NewDimensions is responsible for creating the dimensions, if the font fields cannot be
// converted, an invalid value is set.
func NewDimensions(dimensions interface{}) *Dimensions {
	dimensionsMap, ok := dimensions.(map[string]interface{})
	if !ok {
		return nil
	}

	return &Dimensions{
		Width:  *convertFields(dimensionsMap["width"], -1.0),
		Height: *convertFields(dimensionsMap["height"], -1.0),
	}
}
