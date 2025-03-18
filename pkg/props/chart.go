package props

import (
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
)

type Chart struct {
	XLabels []float64
	YLabels []float64
	Font    Font
}

func (c *Chart) MakeValid() {
	if c.Font.Family == "" {
		c.Font.Family = fontfamily.Arial
	}

	if c.Font.Style == "" {
		c.Font.Style = fontstyle.Normal
	}

	if c.Font.Size == 0.0 {
		c.Font.Size = 8.0
	}
}
