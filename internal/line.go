package internal

import (
	"github.com/johnfercher/maroto/internal/fpdf"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
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
	s.pdf.SetLineWidth(lineProp.Width)
	s.drawStylizedLine(cell, lineProp)
	s.pdf.SetDrawColor(s.defaultLineColor.Red, s.defaultLineColor.Green, s.defaultLineColor.Blue)
	s.pdf.SetLineWidth(consts.DefaultLineWidth)
}

func (s *line) drawStylizedLine(cell Cell, prop props.Line) {
	if prop.Style == consts.Solid {
		s.pdf.Line(cell.X, cell.Y, cell.Width, cell.Height)
		return
	}

	if prop.Style == consts.Dashed {
		s.drawDashedLine(cell)
		return
	}

	s.drawDottedLine(cell, prop.Width)
}

func (s *line) drawDashedLine(cell Cell) {
	xStep := 5.0
	halfDivisor := 2.0
	xHalfStep := xStep / halfDivisor
	for x := cell.X; x < cell.Width; x += xStep {
		s.pdf.Line(x, cell.Y, x+xHalfStep, cell.Height)
	}
}

func (s *line) drawDottedLine(cell Cell, width float64) {
	xStep := 3.0
	for x := cell.X; x < cell.Width; x += xStep {
		s.pdf.Line(x, cell.Y, x+width, cell.Height)
	}
}
