package core

import "github.com/johnfercher/maroto/v2/pkg/config"

// Cell represents a cell inside the PDF.
type Cell struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

func (c *Cell) GetDimensions() *config.Dimensions {
	return &config.Dimensions{Width: c.Width, Height: c.Height}
}

func (c Cell) Copy() Cell {
	return Cell{
		X:      c.X,
		Y:      c.Y,
		Width:  c.Width,
		Height: c.Height,
	}
}

func NewRootContext(pageWidth, pageHeight float64, margins config.Margins) Cell {
	return Cell{
		X:      0,
		Y:      0,
		Width:  pageWidth - margins.Left - margins.Right,
		Height: pageHeight - margins.Top - margins.Bottom,
	}
}
