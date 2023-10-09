package props

import "fmt"

var (
	WhiteColor = Color{Red: 255, Green: 255, Blue: 255}
	BlackColor = Color{Red: 0, Green: 0, Blue: 0}
	RedColor   = Color{Red: 255, Green: 0, Blue: 0}
	GreenColor = Color{Red: 0, Green: 255, Blue: 0}
	BlueColor  = Color{Red: 0, Green: 0, Blue: 255}
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

func (c *Color) ToString() string {
	if c == nil {
		return ""
	}

	return fmt.Sprintf("RGB(%d, %d, %d)", c.Red, c.Green, c.Blue)
}

// IsWhite from Color will return true if all components of color.
// are in the maximum value.
func (c *Color) IsWhite() bool {
	return c.Red == 255 && c.Green == 255 && c.Blue == 255
}
