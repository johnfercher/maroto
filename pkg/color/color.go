package color

const (
	minRGB8Bytes = 0
	maxRGB8Bytes = 255
)

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

// IsWhite from Color will return true if all components of color.
// are in the maximum value.
func (s *Color) IsWhite() bool {
	return s.Red == maxRGB8Bytes && s.Green == maxRGB8Bytes && s.Blue == maxRGB8Bytes
}

// NewWhite return a Color with all components (red, green and blue) as 255.
func NewWhite() Color {
	return Color{
		Red:   maxRGB8Bytes,
		Green: maxRGB8Bytes,
		Blue:  maxRGB8Bytes,
	}
}

// NewBlack return a Color with all components (red, green and blue) as 0.
func NewBlack() Color {
	return Color{
		Red:   minRGB8Bytes,
		Green: minRGB8Bytes,
		Blue:  minRGB8Bytes,
	}
}
