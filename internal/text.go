package internal

import (
	"strings"

	"github.com/johnfercher/maroto/v2/pkg/core"

	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"

	"github.com/johnfercher/maroto/v2/internal/fpdf"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

// Text is the abstraction which deals of how to add text inside PDF.
type Text interface {
	Add(text string, cell core.Cell, textProp props.Text)
	GetLinesQuantity(text string, fontFamily props.Text, colWidth float64) int
}

type text struct {
	pdf  fpdf.Fpdf
	math Math
	font Font
}

// NewText create a Text.
func NewText(pdf fpdf.Fpdf, math Math, font Font) *text {
	return &text{
		pdf,
		math,
		font,
	}
}

// Add a text inside a cell.
func (s *text) Add(text string, cell core.Cell, textProp props.Text) {
	s.font.SetFont(textProp.Family, textProp.Style, textProp.Size)

	if textProp.Top > cell.Height {
		textProp.Top = cell.Height
	}

	if textProp.Left > cell.Width {
		textProp.Left = cell.Width
	}

	if textProp.Right > cell.Width {
		textProp.Right = cell.Width
	}

	cell.Width = cell.Width - textProp.Left - textProp.Right
	if cell.Width < 0 {
		cell.Width = 0
	}

	cell.X += textProp.Left
	cell.Y += textProp.Top

	originalColor := s.font.GetColor()
	s.font.SetColor(textProp.Color)

	// duplicated
	_, _, fontSize := s.font.GetFont()
	fontHeight := fontSize / s.font.GetScaleFactor()

	cell.Y += fontHeight

	// Apply Unicode before calc spaces
	unicodeText := s.textToUnicode(text, textProp)
	stringWidth := s.pdf.GetStringWidth(unicodeText)
	words := strings.Split(unicodeText, " ")
	accumulateOffsetY := 0.0

	// If should add one line
	if stringWidth < cell.Width || textProp.Extrapolate || len(words) == 1 {
		s.addLine(textProp, cell.X, cell.Width, cell.Y, stringWidth, unicodeText)
	} else {
		lines := s.getLines(words, cell.Width)

		for index, line := range lines {
			lineWidth := s.pdf.GetStringWidth(line)
			_, _, fontSize := s.font.GetFont()
			textHeight := fontSize / s.font.GetScaleFactor()

			s.addLine(textProp, cell.X, cell.Width, cell.Y+float64(index)*textHeight+accumulateOffsetY, lineWidth, line)
			accumulateOffsetY += textProp.VerticalPadding
		}
	}

	s.font.SetColor(originalColor)
}

// GetLinesQuantity retrieve the quantity of lines which a text will occupy to avoid that text to extrapolate a cell.
func (s *text) GetLinesQuantity(text string, textProp props.Text, colWidth float64) int {
	translator := s.pdf.UnicodeTranslatorFromDescriptor("")
	s.font.SetFont(textProp.Family, textProp.Style, textProp.Size)

	// Apply Unicode.
	textTranslated := translator(text)

	stringWidth := s.pdf.GetStringWidth(textTranslated)
	words := strings.Split(textTranslated, " ")

	// If should add one line.
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

	if textProp.Align == align.Left {
		s.pdf.Text(xColOffset+left, yColOffset+top, text)
		return
	}

	var modifier float64 = 2

	if textProp.Align == align.Right {
		modifier = 1
	}

	dx := (colWidth - textWidth) / modifier

	s.pdf.Text(dx+xColOffset+left, yColOffset+top, text)
}

func (s *text) textToUnicode(txt string, props props.Text) string {
	if props.Family == fontfamily.Arial ||
		props.Family == fontfamily.Helvetica ||
		props.Family == fontfamily.Symbol ||
		props.Family == fontfamily.ZapBats ||
		props.Family == fontfamily.Courier {
		translator := s.pdf.UnicodeTranslatorFromDescriptor("")
		return translator(txt)
	}

	return txt
}
