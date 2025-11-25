package props

import (
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
)

type Chart struct {
	Scale ChartScale
	Title ChartTitle
}

func (c *Chart) MakeValid() {
	c.Scale.MakeValid()
	c.Title.MakeValid()
}

func (c *Chart) HasScale() bool {
	return c.Scale.X != nil || c.Scale.Y != nil
}

type ChartScale struct {
	X    []float64
	Y    []float64
	Font Font
}

func (c *ChartScale) MakeValid() {
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

type ChartTitle struct {
	Text string
	Font Font
}

func (c *ChartTitle) MakeValid() {
	if c.Text == "" {
		return
	}

	if c.Font.Family == "" {
		c.Font.Family = fontfamily.Arial
	}

	if c.Font.Style == "" {
		c.Font.Style = fontstyle.Normal
	}

	if c.Font.Size == 0.0 {
		c.Font.Size = 10.0
	}
}
