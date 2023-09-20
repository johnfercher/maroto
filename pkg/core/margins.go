package core

import "fmt"

type Margins struct {
	Left   float64
	Right  float64
	Top    float64
	Bottom float64
}

func (m *Margins) Print() {
	if m == nil {
		return
	}

	fmt.Printf("Left: %1.f, Right: %1.f, Top: %1.f, Bottom: %1.f\n", m.Left, m.Right, m.Left, m.Bottom)
}
