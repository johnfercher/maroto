package sign

import (
	"github.com/johnfercher/maroto/enums"
	"github.com/johnfercher/maroto/font"
	"github.com/johnfercher/maroto/math"
	"github.com/johnfercher/maroto/text"
	"github.com/jung-kurt/gofpdf"
)

type Sign interface {
	Sign(label string, fontFamily font.Family, fontStyle font.Style, fontSize float64, qtdCols float64, marginTop float64, actualCol float64)
}

type sign struct {
	pdf  gofpdf.Pdf
	math math.Math
	text text.Text
}

func NewSign(pdf gofpdf.Pdf, math math.Math, text text.Text) Sign {
	return &sign{
		pdf,
		math,
		text,
	}
}

func (s *sign) Sign(label string, fontFamily font.Family, fontStyle font.Style, fontSize float64, qtdCols float64, marginTop float64, actualCol float64) {
	widthPerCol := s.math.GetWidthPerCol(qtdCols)
	left, _, right, _ := s.pdf.GetMargins()

	s.pdf.Line((widthPerCol*actualCol)+left, marginTop+5.0, widthPerCol*(actualCol+1)+right, marginTop+5.0)
	s.text.Add(label, fontFamily, fontStyle, fontSize, marginTop, enums.CenterH, actualCol, qtdCols)
}
