package maroto

import (
	"github.com/jung-kurt/gofpdf"
)

type Sign interface {
	Sign(label string, fontFamily Family, fontStyle Style, fontSize float64, qtdCols float64, marginTop float64, actualCol float64)
}

type sign struct {
	pdf  gofpdf.Pdf
	math Math
	text Text
}

func NewSign(pdf gofpdf.Pdf, math Math, text Text) Sign {
	return &sign{
		pdf,
		math,
		text,
	}
}

func (s *sign) Sign(label string, fontFamily Family, fontStyle Style, fontSize float64, qtdCols float64, marginTop float64, actualCol float64) {
	widthPerCol := s.math.GetWidthPerCol(qtdCols)
	left, _, right, _ := s.pdf.GetMargins()

	s.pdf.Line((widthPerCol*actualCol)+left, marginTop+5.0, widthPerCol*(actualCol+1)+right, marginTop+5.0)
	s.text.Add(label, fontFamily, fontStyle, fontSize, marginTop, CenterH, actualCol, qtdCols)
}
