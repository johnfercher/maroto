package internal

import (
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/jung-kurt/gofpdf"
)

// Signature is the abstraction which deals of how to add a signature space inside PDF
type Signature interface {
	AddSpaceFor(label string, textProp props.Text, colWidth float64, yColOffset float64, xColOffset float64, colHeight float64)
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
func (s *signature) AddSpaceFor(label string, textProp props.Text, colWidth float64, yColOffset float64, xColOffset float64, colHeight float64) {
	left, top, _, _ := s.pdf.GetMargins()
	space := 4.0

	lineCenterY := colHeight / 1.33

	s.pdf.Line(xColOffset+left+space, yColOffset+top+lineCenterY, xColOffset+colWidth+left-space, yColOffset+top+lineCenterY)
	s.text.Add(label, textProp, yColOffset+lineCenterY+2.0, xColOffset, colWidth)
}
