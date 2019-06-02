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

func (s *signature) AddSpaceFor(label string, fontFamily Family, fontStyle Style, fontSize float64, qtdCols float64, marginTop float64, actualCol float64) {
	widthPerCol := s.math.GetWidthPerCol(qtdCols)
	left, _, right, _ := s.pdf.GetMargins()
	space := 4.0

	s.pdf.Line((widthPerCol*actualCol)+left+space, marginTop+5.0, widthPerCol*(actualCol+1)+right-space, marginTop+5.0)
	s.text.Add(label, fontFamily, fontStyle, fontSize, marginTop, Center, actualCol, qtdCols)
}
