package v2

import "fmt"

const LineHeight = 20

// region Types
type Context struct {
	Coordinate     *Coordinate
	Dimensions     *Dimensions
	Margins        *Margins
	pageDimensions *Dimensions
}

type Dimensions struct {
	Width  float64
	Height float64
}

type Coordinate struct {
	X float64
	Y float64
}

type Margins struct {
	Left   float64
	Right  float64
	Top    float64
	Bottom float64
}

// endregion

func NewRootContext(pageWidth, pageHeight float64, margins *Margins) Context {
	return Context{
		Coordinate:     &Coordinate{0, 0},
		Dimensions:     &Dimensions{0, 0},
		pageDimensions: &Dimensions{pageWidth, pageHeight},
		Margins:        margins,
	}
}

func (c *Context) MaxHeight() float64 {
	return c.pageDimensions.Height - c.Margins.Top - c.Margins.Bottom
}

func (c *Context) MaxWidth() float64 {
	return c.pageDimensions.Width - c.Margins.Left - c.Margins.Right
}

func (c *Context) WithDimension(width, height float64) Context {
	ctx := *c
	ctx.Dimensions.Width = width
	ctx.Dimensions.Height = height
	return ctx
}

func (c *Context) WithCoordinateOffset(x, y float64) Context {
	ctx := *c
	ctx.Coordinate.X = x
	ctx.Coordinate.Y = y
	return ctx
}

func (c *Context) GetX() float64 {
	baseX := c.Coordinate.X + c.Dimensions.Width
	if baseX > c.MaxWidth() {
		c.Coordinate.Y = c.Coordinate.Y + LineHeight
		return 0
	}
	return baseX
}

func (c *Context) GetY() float64 {
	baseY := c.Coordinate.Y + c.Dimensions.Height
	if baseY > c.MaxHeight() {
		//Verificar como indicar quebra de pagina
		return 0
	}
	return baseY
}

func (c *Context) Print(label interface{}) {
	fmt.Println(label)
	c.Margins.Print()
}

func (c *Margins) Print() {
	if c == nil {
		return
	}

	fmt.Printf("Left: %1.f, Right: %1.f, Top: %1.f, Bottom: %1.f\n", c.Left, c.Right, c.Left, c.Bottom)
}
