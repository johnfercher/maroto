package context

import "fmt"

type Dimensions struct {
	Width  float64
	Height float64
}

func (d *Dimensions) Print() {
	if d == nil {
		return
	}

	fmt.Printf("Width: %1.f, Height: %1.f\n", d.Width, d.Height)
}
