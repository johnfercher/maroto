package propsmapper

// Color represents a color in the RGB (Red, Green, Blue) space,
// is possible mix values, when all values are 0 the result color is black
// when all values are 255 the result color is white.
type Color struct {
	// Red is the amount of red
	Red int
	// Green is the amount of red
	Green int
	// Blue is the amount of red
	Blue int
}

// NewColor is responsible for creating the color, if the font fields cannot be
// converted, an invalid value is set.
func NewColor(color interface{}) *Color {
	colorMap, ok := color.(map[string]interface{})
	if !ok {
		return nil
	}

	return &Color{
		Red:   int(*convertFields(colorMap["red"], 0.0)),
		Green: int(*convertFields(colorMap["green"], 0.0)),
		Blue:  int(*convertFields(colorMap["blue"], 0.0)),
	}
}
