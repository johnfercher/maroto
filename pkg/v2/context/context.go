package context

import (
	"fmt"
	"math"
)

const LineHeight = 20

// region Types
type Context struct {
	Coordinate     *Coordinate
	Dimensions     *Dimensions
	Margins        *Margins
	CurrentPage    int
	pageDimensions *Dimensions
}

// endregion

func NewRootContext(pageWidth, pageHeight float64, margins *Margins) Context {
	return Context{
		Coordinate:     &Coordinate{0, 0},
		Dimensions:     &Dimensions{0, 0},
		pageDimensions: &Dimensions{pageWidth, pageHeight},
		Margins:        margins,
		CurrentPage:    1,
	}
}

func (c Context) MaxHeight() float64 {
	return c.pageDimensions.Height - c.Margins.Top - c.Margins.Bottom
}

func (c Context) MaxWidth() float64 {
	return c.pageDimensions.Width - c.Margins.Left - c.Margins.Right
}

func (c Context) WithDimension(width, height float64) Context {
	ctx := c.copy()
	ctx.Dimensions.Width = width
	ctx.Dimensions.Height = height
	return ctx
}

func (c Context) WithCoordinates(x, y float64) Context {
	ctx := c.copy()
	ctx.Coordinate.X = x
	ctx.Coordinate.Y = y
	return ctx
}

func (c Context) GetXOffset() float64 {
	baseX := c.Coordinate.X + c.Dimensions.Width
	if baseX > c.MaxWidth() {
		c.Coordinate.Y = c.Coordinate.Y + LineHeight
		return c.Coordinate.X
	}
	return baseX
}

func (c Context) GetYOffset() float64 {
	baseY := c.Coordinate.Y + c.Dimensions.Height
	if baseY > c.MaxHeight() {
		c.CurrentPage = c.CurrentPage + 1
		return 0
	}
	return baseY
}

func (c Context) NewPage(pageCount int) Context {
	ctx := c.copy()
	ctx.CurrentPage = int(math.Max(float64(pageCount), float64(c.CurrentPage+1)))
	ctx.Coordinate.X = 0
	ctx.Coordinate.Y = 0
	return ctx
}

func (c Context) GetCurrentPage() int {
	return c.CurrentPage
}
func (c Context) Print(label interface{}) {
	fmt.Println(label)

	c.Coordinate.Print()
	c.Dimensions.Print()
	c.pageDimensions.Print()
	c.Margins.Print()
}

func (c Context) copy() Context {
	return Context{
		Coordinate:     &Coordinate{c.Coordinate.X, c.Coordinate.Y},
		Dimensions:     &Dimensions{c.Dimensions.Width, c.Dimensions.Height},
		Margins:        c.Margins,
		CurrentPage:    c.CurrentPage,
		pageDimensions: c.pageDimensions,
	}
}
