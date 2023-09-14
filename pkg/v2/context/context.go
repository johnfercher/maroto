package context

import (
	"fmt"
)

const LineHeight = 20

// region Types
type Context struct {
	Coordinate Coordinate
	Dimensions Dimensions
}

// endregion

func NewRootContext(pageWidth, pageHeight float64, margins Margins) Context {
	return Context{
		Coordinate: Coordinate{0, 0},
		Dimensions: Dimensions{pageWidth - margins.Left - margins.Right, pageHeight - margins.Top - margins.Bottom},
	}
}

func (c Context) WithDimension(width, height float64) Context {
	ctx := c.Copy()
	ctx.Dimensions.Width = width
	ctx.Dimensions.Height = height
	return ctx
}

func (c Context) WithCoordinates(x, y float64) Context {
	ctx := c.Copy()
	ctx.Coordinate.X = x
	ctx.Coordinate.Y = y
	return ctx
}

func (c Context) Print(label interface{}) {
	fmt.Println(label)

	c.Coordinate.Print()
	c.Dimensions.Print()
}

func (c Context) Copy() Context {
	return Context{
		Coordinate: Coordinate{c.Coordinate.X, c.Coordinate.Y},
		Dimensions: Dimensions{c.Dimensions.Width, c.Dimensions.Height},
	}
}
