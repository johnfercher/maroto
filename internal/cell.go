package internal

// Cell represents a cell inside the PDF.
type Cell struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

func (c Cell) Copy() Cell {
	return Cell{
		X:      c.X,
		Y:      c.Y,
		Width:  c.Width,
		Height: c.Height,
	}
}
