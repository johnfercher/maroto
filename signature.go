package maroto

import (
	"github.com/jung-kurt/gofpdf"
)

// Signature is the abstraction which deals of how to add a signature space inside PDF
type Signature interface {
	AddSpaceFor(label string, textProp TextProp, qtdCols float64, marginTop float64, actualCol float64)
}

type signature struct {
	pdf  gofpdf.Pdf
	math Math
	text Text
}

// NewSignature create a Signature
func NewSignature(pdf gofpdf.Pdf, math Math, text Text) Signature {
	return &signature{
		pdf,
		math,
		text,
	}
}

// AddSpaceFor create a space for a signature inside a cell
func (self *signature) AddSpaceFor(label string, textProp TextProp, qtdCols float64, marginTop float64, actualCol float64) {
	widthPerCol := self.math.GetWidthPerCol(qtdCols)
	left, _, right, _ := self.pdf.GetMargins()
	space := 4.0

	self.pdf.Line((widthPerCol*actualCol)+left+space, marginTop+5.0, widthPerCol*(actualCol+1)+right-space, marginTop+5.0)
	self.text.Add(label, textProp, marginTop, actualCol, qtdCols)
}
