package internal

import (
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/jung-kurt/gofpdf"
	"strings"
)

// Text is the abstraction which deals of how to add text inside PDF
type Text interface {
	Add(text string, textProp props.Text, yColOffset float64, xColOffset float64, colWidth float64)
	GetLinesQuantity(text string, fontFamily props.Text, colWidth float64) int
}

type text struct {
	pdf  gofpdf.Pdf
	math Math
	font Font
}

// NewText create a Text
func NewText(pdf gofpdf.Pdf, math Math, font Font) *text {
	return &text{
		pdf,
		math,
		font,
	}
}

// Add a text inside a cell.
func (s *text) Add(text string, textProp props.Text, yColOffset float64, xColOffset float64, colWidth float64) {
	translator := s.pdf.UnicodeTranslatorFromDescriptor("")
	s.font.SetFont(textProp.Family, textProp.Style, textProp.Size)

	// duplicated
	_, _, fontSize := s.font.GetFont()
	fontHeight := fontSize / s.font.GetScaleFactor()

	yColOffset += fontHeight

	// Apply Unicode before calc spaces
	unicodeText := translator(text)

	stringWidth := s.pdf.GetStringWidth(unicodeText)
	words := strings.Split(unicodeText, " ")
	accumulateOffsetY := 0.0

	// If should add one line
	if stringWidth < colWidth || textProp.Extrapolate || len(words) == 1 {
		s.addLine(textProp, xColOffset, colWidth, yColOffset, stringWidth, unicodeText)
	} else {
		lines := s.getLines(words, colWidth)

		for index, line := range lines {
			lineWidth := s.pdf.GetStringWidth(line)
			_, _, fontSize := s.font.GetFont()
			textHeight := fontSize / s.font.GetScaleFactor()

			s.addLine(textProp, xColOffset, colWidth, yColOffset+float64(index)*textHeight+accumulateOffsetY, lineWidth, line)
			accumulateOffsetY += textProp.VerticalPadding
		}
	}
}

// GetLinesQuantity retrieve the quantity of lines which a text will occupy to avoid that text to extrapolate a cell
func (s *text) GetLinesQuantity(text string, textProp props.Text, colWidth float64) int {
	translator := s.pdf.UnicodeTranslatorFromDescriptor("")
	s.font.SetFont(textProp.Family, textProp.Style, textProp.Size)

	// Apply Unicode
	textTranslated := translator(text)

	stringWidth := s.pdf.GetStringWidth(textTranslated)
	words := strings.Split(textTranslated, " ")

	// If should add one line
	if stringWidth < colWidth || textProp.Extrapolate || len(words) == 1 {
		return 1
	}

	lines := s.getLines(words, colWidth)
	return len(lines)
}

func (s *text) getLines(words []string, colWidth float64) []string {
	currentlySize := 0.0
	actualLine := 0

	lines := []string{}
	lines = append(lines, "")

	for _, word := range words {
		if s.pdf.GetStringWidth(word+" ")+currentlySize < colWidth {
			lines[actualLine] = lines[actualLine] + word + " "
			currentlySize += s.pdf.GetStringWidth(word + " ")
		} else {
			lines = append(lines, "")
			actualLine++
			lines[actualLine] = lines[actualLine] + word + " "
			currentlySize = s.pdf.GetStringWidth(word + " ")
		}
	}

	return lines
}

func (s *text) addLine(textProp props.Text, xColOffset, colWidth, yColOffset, textWidth float64, text string) {
	left, top, _, _ := s.pdf.GetMargins()

	if textProp.Align == consts.Left {
		s.pdf.Text(xColOffset+left, yColOffset+top, text)
		return
	}

	var modifier float64 = 2

	if textProp.Align == consts.Right {
		modifier = 1
	}

	dx := (colWidth - textWidth) / modifier

	s.pdf.Text(dx+xColOffset+left, yColOffset+top, text)
}
