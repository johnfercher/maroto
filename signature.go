package maroto

import (
	"github.com/jung-kurt/gofpdf"
)

type Signature interface {
	AddSpaceFor(label string, fontFamily Family, fontStyle Style, fontSize float64, qtdCols float64, marginTop float64, actualCol float64)
}

type signature struct {
	pdf  gofpdf.Pdf
	math Math
	text Text
}

func NewSignature(pdf gofpdf.Pdf, math Math, text Text) Signature {
	return &signature{
		pdf,
		math,
		text,
	}
}

func (self *signature) AddSpaceFor(label string, fontFamily Family, fontStyle Style, fontSize float64, qtdCols float64, marginTop float64, actualCol float64) {
	widthPerCol := self.math.GetWidthPerCol(qtdCols)
	left, _, right, _ := self.pdf.GetMargins()
	space := 4.0

	self.pdf.Line((widthPerCol*actualCol)+left+space, marginTop+5.0, widthPerCol*(actualCol+1)+right-space, marginTop+5.0)
	self.text.Add(label, fontFamily, fontStyle, fontSize, marginTop, Center, actualCol, qtdCols)
}
