package internal

import (
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/jung-kurt/gofpdf"
)

// Signature is the abstraction which deals of how to add a signature space inside PDF
type Signature interface {
	AddSpaceFor(label string, cell Cell, textProp props.Text)
}

type signature struct {
	pdf  gofpdf.Pdf
	math Math
	text Text
}

// NewSignature create a Signature
func NewSignature(pdf gofpdf.Pdf, math Math, text Text) *signature {
	return &signature{
		pdf,
		math,
		text,
	}
}

// AddSpaceFor create a space for a signature inside a cell
func (s *signature) AddSpaceFor(label string, cell Cell, textProp props.Text) {
	left, top, _, _ := s.pdf.GetMargins()
	space := 4.0

	lineCenterY := cell.Height / 1.33

	s.pdf.Line(cell.X+left+space, cell.Y+top+lineCenterY, cell.X+cell.Width+left-space, cell.Y+top+lineCenterY)
	s.text.Add(label, textProp, cell.Y+lineCenterY+2.0, cell.X, cell.Width)
}
