package internal

import (
	"github.com/johnfercher/maroto/v2/internal/fpdf"
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
	"github.com/johnfercher/maroto/v2/pkg/core/context"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

// Line is the abstraction which deals with lines.
type Line interface {
	Draw(cell context.Cell, lineProp props.Line)
}

type line struct {
	pdf              fpdf.Fpdf
	defaultLineColor *props.Color
}

// NewLine create a Line Helper.
func NewLine(pdf fpdf.Fpdf) *line {
	return &line{
		pdf:              pdf,
		defaultLineColor: props.NewBlack(),
	}
}

func (s *line) Draw(cell context.Cell, lineProp props.Line) {
	s.pdf.SetDrawColor(lineProp.Color.Red, lineProp.Color.Green, lineProp.Color.Blue)
	s.pdf.SetLineWidth(lineProp.Width)
	s.drawStylizedLine(cell, lineProp)
	s.pdf.SetDrawColor(s.defaultLineColor.Red, s.defaultLineColor.Green, s.defaultLineColor.Blue)
	s.pdf.SetLineWidth(linestyle.DefaultLineWidth)
}

func (s *line) drawStylizedLine(cell context.Cell, prop props.Line) {
	if prop.Style == linestyle.Solid {
		s.pdf.Line(cell.X, cell.Y, cell.Width, cell.Height)
		return
	}

	if prop.Style == linestyle.Dashed {
		s.drawDashedLine(cell)
		return
	}

	s.drawDottedLine(cell, prop.Width)
}

func (s *line) drawDashedLine(cell context.Cell) {
	xStep := 5.0
	halfDivisor := 2.0
	xHalfStep := xStep / halfDivisor
	for x := cell.X; x < cell.Width; x += xStep {
		s.pdf.Line(x, cell.Y, x+xHalfStep, cell.Height)
	}
}

func (s *line) drawDottedLine(cell context.Cell, width float64) {
	xStep := 3.0
	for x := cell.X; x < cell.Width; x += xStep {
		s.pdf.Line(x, cell.Y, x+width, cell.Height)
	}
}
