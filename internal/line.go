package internal

import (
	"github.com/johnfercher/maroto/internal/fpdf"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/props"
)

// Line is the abstraction which deals with lines.
type Line interface {
	Draw(cell Cell, lineProp props.Line)
}

type line struct {
	pdf              fpdf.Fpdf
	defaultLineColor color.Color
}

// NewLine create a Line Helper.
func NewLine(pdf fpdf.Fpdf) *line {
	return &line{
		pdf:              pdf,
		defaultLineColor: color.NewBlack(),
	}
}

func (s *line) Draw(cell Cell, lineProp props.Line) {
	s.pdf.SetDrawColor(lineProp.Color.Red, lineProp.Color.Green, lineProp.Color.Blue)

	s.pdf.Line(cell.X, cell.Y, cell.Width, cell.Height)

	s.pdf.SetDrawColor(s.defaultLineColor.Red, s.defaultLineColor.Green, s.defaultLineColor.Blue)
}
