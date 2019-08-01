package maroto

import (
	"github.com/jung-kurt/gofpdf"
)

// Text is the abstraction which deals of how to add text inside PDF
type Text interface {
	Add(text string, fontFamily Family, fontStyle Style, fontSize float64, marginTop float64, align Align, actualCol float64, qtdCols float64)
}

type text struct {
	pdf  gofpdf.Pdf
	math Math
	font Font
}

// NewText create a Text
func NewText(pdf gofpdf.Pdf, math Math, font Font) Text {
	return &text{
		pdf,
		math,
		font,
	}
}

// Add a text inside a cell.
func (self *text) Add(text string, fontFamily Family, fontStyle Style, fontSize float64, marginTop float64, align Align, actualCol float64, qtdCols float64) {
	actualWidthPerCol := self.math.GetWidthPerCol(qtdCols)

	self.font.SetFont(fontFamily, fontStyle, fontSize)

	left, top, _, _ := self.pdf.GetMargins()

	if align == Left {
		self.pdf.Text(actualCol*actualWidthPerCol+left, marginTop+top, text)
		return
	}

	var modifier float64 = 2

	if align == Right {
		modifier = 1
	}

	translator := self.pdf.UnicodeTranslatorFromDescriptor("")

	stringWidth := self.pdf.GetStringWidth(translator(text))
	dx := (actualWidthPerCol - stringWidth) / modifier

	self.pdf.Text(dx+actualCol*actualWidthPerCol+left, marginTop+top, translator(text))
}
