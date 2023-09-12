package context

import "fmt"

type Coordinate struct {
	X float64
	Y float64
}

func (c *Coordinate) Print() {
	if c == nil {
		return
	}

	fmt.Printf("(%1.f, %1.f)\n", c.X, c.Y)
}
