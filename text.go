package maroto

import (
	"github.com/jung-kurt/gofpdf"
)

type Text interface {
	Add(text string, fontFamily Family, fontStyle Style, fontSize float64, marginTop float64, align Align, actualCol float64, qtdCols float64)
}

type text struct {
	pdf  gofpdf.Pdf
	math Math
	font Font
}

func NewText(pdf gofpdf.Pdf, math Math, font Font) Text {
	return &text{
		pdf,
		math,
		font,
	}
}

func (m *text) Add(text string, fontFamily Family, fontStyle Style, fontSize float64, marginTop float64, align Align, actualCol float64, qtdCols float64) {
	actualWidthPerCol := m.math.GetWidthPerCol(qtdCols)

	m.font.SetFont(fontFamily, fontStyle, fontSize)

	left, top, _, _ := m.pdf.GetMargins()

	if align == Left {
		m.pdf.Text(actualCol*actualWidthPerCol+left, marginTop+top, text)
		return
	}

	stringWidth := m.pdf.GetStringWidth(text)
	dx := (actualWidthPerCol - stringWidth) / 2
	m.pdf.Text(dx+actualCol*actualWidthPerCol+left, marginTop+top, text)
}
