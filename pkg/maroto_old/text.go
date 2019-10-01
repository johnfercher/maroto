package maroto

import (
	"github.com/jung-kurt/gofpdf"
	"strings"
)

// Text is the abstraction which deals of how to add text inside PDF
type Text interface {
	Add(text string, fontFamily TextProp, marginTop float64, actualCol float64, qtdCols float64)
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
func (self *text) Add(text string, textProp TextProp, marginTop float64, actualCol float64, qtdCols float64) {
	actualWidthPerCol := self.math.GetWidthPerCol(qtdCols)

	translator := self.pdf.UnicodeTranslatorFromDescriptor("")
	self.font.SetFont(textProp.Family, textProp.Style, textProp.Size)

	textTranslated := translator(text)
	stringWidth := self.pdf.GetStringWidth(textTranslated)
	words := strings.Split(textTranslated, " ")

	if stringWidth < actualWidthPerCol || textProp.Extrapolate || len(words) == 1 {
		self.addLine(textProp, actualCol, actualWidthPerCol, marginTop, stringWidth, textTranslated)
	} else {
		currentlySize := 0.0
		actualLine := 0
		lines := []string{}
		lines = append(lines, "")

		for _, word := range words {
			if self.pdf.GetStringWidth(word+" ")+currentlySize < actualWidthPerCol {
				lines[actualLine] = lines[actualLine] + word + " "
				currentlySize += self.pdf.GetStringWidth(word + " ")
			} else {
				lines = append(lines, "")
				actualLine++
				lines[actualLine] = lines[actualLine] + word + " "
				currentlySize = self.pdf.GetStringWidth(word + " ")
			}
		}

		for index, line := range lines {
			lineWidth := self.pdf.GetStringWidth(line)
			self.addLine(textProp, actualCol, actualWidthPerCol, marginTop+float64(index)*textProp.Size/2.0, lineWidth, line)
		}
	}
}

func (self *text) addLine(textProp TextProp, actualCol, actualWidthPerCol, marginTop, stringWidth float64, textTranslated string) {
	left, top, _, _ := self.pdf.GetMargins()

	if textProp.Align == Left {
		self.pdf.Text(actualCol*actualWidthPerCol+left, marginTop+top, textTranslated)
		return
	}

	var modifier float64 = 2

	if textProp.Align == Right {
		modifier = 1
	}

	dx := (actualWidthPerCol - stringWidth) / modifier

	self.pdf.Text(dx+actualCol*actualWidthPerCol+left, marginTop+top, textTranslated)
}
