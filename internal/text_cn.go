package internal

import (
	"github.com/johnfercher/maroto/internal/fpdf"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
)

type textCN struct {
	pdf  fpdf.Fpdf
	math Math
	font Font
}

// NewTextCN create a Text can handle Chinese content.
func NewTextCN(pdf fpdf.Fpdf, math Math, font Font) *textCN {
	return &textCN{
		pdf,
		math,
		font,
	}
}

// Add a text inside a cell.
func (s *textCN) Add(text string, cell Cell, textProp props.Text) {
	s.font.SetFont(textProp.Family, textProp.Style, textProp.Size)

	originalColor := s.font.GetColor()
	s.font.SetColor(textProp.Color)

	// duplicated
	_, _, fontSize := s.font.GetFont()
	fontHeight := fontSize / s.font.GetScaleFactor()

	cell.Y += fontHeight

	// Apply Unicode before calc spaces
	unicodeText := s.textToUnicode(text, textProp)
	stringWidth := s.pdf.GetStringWidth(unicodeText)
	accumulateOffsetY := 0.0

	// If should add one line
	if stringWidth < cell.Width || textProp.Extrapolate {
		s.addLine(textProp, cell.X, cell.Width, cell.Y, stringWidth, unicodeText)
	} else {
		// 在英语里，使用空格作分隔符是合理的；
		// 但在中文里，字与字之间没有空格
		// 根据宽度计算wc
		wc := getWorldCountByWidth(stringWidth, cell.Width)
		words := splitByWorldCount(unicodeText, wc)
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

func getWorldCountByWidth(stringWidth, cellWidth float64) int {
	return int(stringWidth / cellWidth)
}

// splitByWorldCount split the string by world count
func splitByWorldCount(in string, wc int) []string {
	if wc == 0 {
		return []string{in}
	}

	runeLen := len([]rune(in))
	r := make([]string, 0, runeLen)

	var runes = make([]rune, 0, runeLen)
	for _, re := range in {
		runes = append(runes, re)

		if len(runes) == wc {
			r = append(r, string(runes))
			runes = []rune{}
		}
	}
	if len(runes) != 0 {
		r = append(r, string(runes))
	}

	return r
}

const (
	space = ""
)

// GetLinesQuantity retrieve the quantity of lines which a text will occupy to avoid that text to extrapolate a cell.
func (s *textCN) GetLinesQuantity(text string, textProp props.Text, colWidth float64) int {
	translator := s.pdf.UnicodeTranslatorFromDescriptor("")
	s.font.SetFont(textProp.Family, textProp.Style, textProp.Size)

	// Apply Unicode.
	textTranslated := translator(text)

	stringWidth := s.pdf.GetStringWidth(textTranslated)
	wc := getWorldCountByWidth(stringWidth, colWidth)
	words := splitByWorldCount(textTranslated, wc)

	// If should add one line.
	if stringWidth < colWidth || textProp.Extrapolate || len(words) == 1 {
		return 1
	}

	lines := s.getLines(words, colWidth)
	return len(lines)
}

func (s *textCN) getLines(words []string, colWidth float64) []string {
	currentlySize := 0.0
	actualLine := 0

	lines := []string{}
	lines = append(lines, "")

	for _, word := range words {
		if s.pdf.GetStringWidth(word+space)+currentlySize < colWidth {
			lines[actualLine] = lines[actualLine] + word + space
			currentlySize += s.pdf.GetStringWidth(word + space)
		} else {
			lines = append(lines, "")
			actualLine++
			lines[actualLine] = lines[actualLine] + word + space
			currentlySize = s.pdf.GetStringWidth(word + space)
		}
	}

	return lines
}

func (s *textCN) addLine(textProp props.Text, xColOffset, colWidth, yColOffset, textWidth float64, text string) {
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

func (s *textCN) textToUnicode(txt string, props props.Text) string {
	if props.Family == consts.Arial ||
		props.Family == consts.Helvetica ||
		props.Family == consts.Symbol ||
		props.Family == consts.ZapBats ||
		props.Family == consts.Courier {
		translator := s.pdf.UnicodeTranslatorFromDescriptor("")
		return translator(txt)
	}

	return txt
}
