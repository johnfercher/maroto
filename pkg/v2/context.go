package v2

import "fmt"

type Context struct {
	Coordinate *Coordinate
	Dimensions *Dimensions
	Margins    *Margins
}

func (c *Context) GetUsefulWidth() float64 {
	return c.Dimensions.Width - c.Margins.Left - c.Margins.Right
}

func (c *Context) Print(label interface{}) {
	fmt.Println(label)
	c.Margins.Print()
}

type Dimensions struct {
	Width  float64
	Height float64
}

func (c *Dimensions) Print() {

}

type Coordinate struct {
	X float64
	Y float64
}

func (c *Coordinate) Print() {

}

type Margins struct {
	Left   float64
	Right  float64
	Top    float64
	Bottom float64
}

func (c *Margins) Print() {
	if c == nil {
		return
	}

	fmt.Printf("Left: %1.f, Right: %1.f, Top: %1.f, Bottom: %1.f\n", c.Left, c.Right, c.Left, c.Bottom)
}
