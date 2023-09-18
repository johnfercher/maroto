package context

import "github.com/johnfercher/maroto/v2/internal"

func NewRootContext(pageWidth, pageHeight float64, margins Margins) internal.Cell {
	return internal.Cell{
		X:      0,
		Y:      0,
		Width:  pageWidth - margins.Left - margins.Right,
		Height: pageHeight - margins.Top - margins.Bottom,
	}
}
