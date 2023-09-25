package internal

import (
	"strings"

	textconsts "github.com/johnfercher/maroto/v2/pkg/consts/text"
	"github.com/johnfercher/maroto/v2/pkg/core"

	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"

	"github.com/johnfercher/maroto/v2/internal/fpdf"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

// Text is the abstraction which deals of how to add text inside PDF.
type Text interface {
	Add(text string, cell core.Cell, textProp props.Text)
	GetComputedHeight(text string, fontFamily props.Text, colWidth float64) float64
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

	accumulateOffsetY := 0.0
	lines := s.getLines(text, textProp, cell.Width)

	for index, line := range lines {
		lineWidth := s.pdf.GetStringWidth(line)
		_, _, fontSize := s.font.GetFont()
		textHeight := fontSize / s.font.GetScaleFactor()

		s.addLine(textProp, cell.X, cell.Width, cell.Y+float64(index)*textHeight+accumulateOffsetY, lineWidth, line)
		accumulateOffsetY += textProp.VerticalPadding
	}

	s.font.SetColor(originalColor)
}

// GetComputedHeight retrieve the quantity of lines which a text will occupy to avoid that text to extrapolate a cell.
func (s *text) GetComputedHeight(text string, textProp props.Text, colWidth float64) float64 {
	s.font.SetFont(textProp.Family, textProp.Style, textProp.Size)

	lines := s.getLines(text, textProp, colWidth-textProp.Right-textProp.Left)

	_, _, fontSize := s.font.GetFont()
	textHeight := fontSize / s.font.GetScaleFactor()

	return textProp.Top + textProp.Bottom + float64(len(lines))*textHeight + float64(len(lines))*textProp.VerticalPadding
}

func (s *text) getLines(text string, textProp props.Text, colWidth float64) []string {
	// Apply Unicode.
	unicodeText := s.textToUnicode(text, textProp)
	words := strings.Split(unicodeText, " ")

	if len(words) == 0 {
		return []string{}
	}

	switch textProp.ExtrapolateStrategy {
	case textconsts.ExtrapolateStrategyWords:
		return s.extrapolateWords(words, colWidth)
	case textconsts.ExtrapolateStrategySymbols:
		return s.extrapolateSymbols(words, colWidth)
	default:
		return []string{unicodeText}
	}
}

type splitter struct {
	lines          []string
	actualLine     int
	colWidth       float64
	getStringWidth func(string) float64
}

func (s *splitter) addWord(word string) {
	if s.lines == nil {
		s.lines = []string{""}
	}

	if s.getStringWidth(s.lines[s.actualLine]+word+" ") < s.colWidth {
		s.lines[s.actualLine] += word + " "
		return
	}

	for sIndex, symbol := range word {
		if s.getStringWidth(s.lines[s.actualLine]+string(symbol)) < s.colWidth {
			s.lines[s.actualLine] += string(symbol)
		} else {
			s.lines = append(s.lines, "")
			s.actualLine++
			s.addWord(word[sIndex:])
			break
		}
	}
}

func (s *text) extrapolateSymbols(words []string, colWidth float64) []string {
	wordSplitter := splitter{
		colWidth:       colWidth,
		getStringWidth: s.pdf.GetStringWidth,
	}
	for _, word := range words {
		wordSplitter.addWord(word)
	}

	return wordSplitter.lines
}

func (s *text) extrapolateWords(words []string, colWidth float64) []string {
	currentlySize := 0.0
	actualLine := 0

	var lines []string

	if len(words) == 1 {
		return append(lines, words[0])
	}

	lines = append(lines, "")

	for _, word := range words {
		if s.pdf.GetStringWidth(word+" ")+currentlySize < colWidth {
			lines[actualLine] = lines[actualLine] + word + " "
			currentlySize += s.pdf.GetStringWidth(word + " ")
			continue
		}

		if lines[actualLine] != "" {
			lines = append(lines, "")
			actualLine++
		}
		lines[actualLine] = lines[actualLine] + word + " "
		currentlySize = s.pdf.GetStringWidth(word + " ")
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
