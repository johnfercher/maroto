package props

import "fmt"

var (
	// WhiteColor is a Color with all values in 255.
	WhiteColor = Color{Red: 255, Green: 255, Blue: 255}
	// BlackColor is a Color with all values in 0.
	BlackColor = Color{Red: 0, Green: 0, Blue: 0}
	// RedColor is a Color with only Red in 255.
	RedColor = Color{Red: 255, Green: 0, Blue: 0}
	// GreenColor is a Color with only Green in 255.
	GreenColor = Color{Red: 0, Green: 255, Blue: 0}
	// BlueColor is a Color with only Blue in 255.
	BlueColor = Color{Red: 0, Green: 0, Blue: 255}
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

// ToString returns a string representation of the Color.
func (c *Color) ToString() string {
	if c == nil {
		return ""
	}

	return fmt.Sprintf("RGB(%d, %d, %d)", c.Red, c.Green, c.Blue)
}
