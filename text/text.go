package text

import (
	"github.com/johnfercher/maroto/enums"
	"github.com/johnfercher/maroto/font"
	"github.com/johnfercher/maroto/math"
	"github.com/jung-kurt/gofpdf"
)

type Text interface {
	Add(text string, fontFamily font.Family, fontStyle font.Style, fontSize float64, marginTop float64, align enums.HorizontalAlign, actualCol float64, qtdCols float64)
}

type text struct {
	pdf  gofpdf.Pdf
	math math.Math
	font font.Font
}

func NewText(pdf gofpdf.Pdf, math math.Math, font font.Font) Text {
	return &text{
		pdf,
		math,
		font,
	}
}

func (m *text) Add(text string, fontFamily font.Family, fontStyle font.Style, fontSize float64, marginTop float64, align enums.HorizontalAlign, actualCol float64, qtdCols float64) {
	actualWidthPerCol := m.math.GetWidthPerCol(qtdCols)

	m.font.SetFont(fontFamily, fontStyle, fontSize)

	left, top, _, _ := m.pdf.GetMargins()

	if align == enums.Left {
		m.pdf.Text(actualCol*actualWidthPerCol+left, marginTop+top, text)
		return
	}

	stringWidth := m.pdf.GetStringWidth(text)
	dx := (actualWidthPerCol - stringWidth) / 2
	m.pdf.Text(dx+actualCol*actualWidthPerCol+left, marginTop+top, text)
}
