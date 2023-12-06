package gofpdf

import (
	"fmt"
	"strings"

	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"

	"github.com/johnfercher/maroto/v2/pkg/core/entity"

	"github.com/johnfercher/maroto/v2/pkg/consts/breakline"

	"github.com/johnfercher/maroto/v2/pkg/core"

	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"

	"github.com/johnfercher/maroto/v2/pkg/props"
)

type text struct {
	pdf  gofpdfwrapper.Fpdf
	math core.Math
	font core.Font
}

// NewText create a Text.
func NewText(pdf gofpdfwrapper.Fpdf, math core.Math, font core.Font) *text {
	return &text{
		pdf,
		math,
		font,
	}
}

// Add a text inside a cell.
func (s *text) Add(text string, cell *entity.Cell, textProp *props.Text) {
	s.font.SetFont(textProp.Family, textProp.Style, textProp.Size)
	fontHeight := s.font.GetHeight(textProp.Family, textProp.Style, textProp.Size)

	if textProp.Top > cell.Height {
		textProp.Top = cell.Height
	}

	if textProp.Left > cell.Width {
		textProp.Left = cell.Width
	}

	if textProp.Right > cell.Width {
		textProp.Right = cell.Width
	}

	width := cell.Width - textProp.Left - textProp.Right
	if width < 0 {
		width = 0
	}

	x := cell.X + textProp.Left
	y := cell.Y + textProp.Top

	originalColor := s.font.GetColor()
	if textProp.Color != nil {
		s.font.SetColor(textProp.Color)
	}

	// override style if hyperlink is set
	if textProp.Hyperlink != nil {
		s.font.SetColor(&props.BlueColor)
	}

	y += fontHeight

	// Apply Unicode before calc spaces
	unicodeText := s.textToUnicode(text, textProp)
	stringWidth := s.pdf.GetStringWidth(unicodeText)

	// If should add one line
	if stringWidth < width {
		s.addLine(textProp, x, width, y, stringWidth, unicodeText)
		if textProp.Color != nil {
			s.font.SetColor(originalColor)
		}
		return
	}

	var lines []string

	if textProp.BreakLineStrategy == breakline.EmptyLineStrategy {
		words := strings.Split(unicodeText, " ")
		lines = s.getLinesBreakingLineFromSpace(words, width)
	} else {
		lines = s.getLinesBreakingLineWithDash(unicodeText, width)
	}

	accumulateOffsetY := 0.0

	for index, line := range lines {
		lineWidth := s.pdf.GetStringWidth(line)

		s.addLine(textProp, x, width, y+float64(index)*fontHeight+accumulateOffsetY, lineWidth, line)
		accumulateOffsetY += textProp.VerticalPadding
	}

	if textProp.Color != nil {
		s.font.SetColor(originalColor)
	}
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
	if stringWidth < colWidth || len(words) == 1 {
		return 1
	}

	lines := s.getLinesBreakingLineFromSpace(words, colWidth)
	return len(lines)
}

func (s *text) getLinesBreakingLineFromSpace(words []string, colWidth float64) []string {
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

func (s *text) getLinesBreakingLineWithDash(words string, colWidth float64) []string {
	currentlySize := 0.0

	lines := []string{}

	dashSize := s.pdf.GetStringWidth(" - ")

	var content string
	for _, letter := range words {
		if currentlySize+dashSize > colWidth-dashSize {
			content += "-"
			lines = append(lines, content)
			content = ""
			currentlySize = 0
		}

		letterString := fmt.Sprintf("%c", letter)
		width := s.pdf.GetStringWidth(letterString)
		content += letterString
		currentlySize += width
	}

	if content != "" {
		lines = append(lines, content)
	}

	return lines
}

func (s *text) addLine(textProp *props.Text, xColOffset, colWidth, yColOffset, textWidth float64, text string) {
	left, top, _, _ := s.pdf.GetMargins()

	fontHeight := s.font.GetHeight(textProp.Family, textProp.Style, textProp.Size)

	if textProp.Align == align.Left {
		s.pdf.Text(xColOffset+left, yColOffset+top, text)

		if textProp.Hyperlink != nil {
			s.pdf.LinkString(xColOffset+left, yColOffset+top-fontHeight, textWidth, fontHeight, *textProp.Hyperlink)
		}

		return
	}

	var modifier float64 = 2

	if textProp.Align == align.Right {
		modifier = 1
	}

	dx := (colWidth - textWidth) / modifier

	if textProp.Hyperlink != nil {
		s.pdf.LinkString(dx+xColOffset+left, yColOffset+top-fontHeight, textWidth, fontHeight, *textProp.Hyperlink)
	}

	s.pdf.Text(dx+xColOffset+left, yColOffset+top, text)
}

func (s *text) textToUnicode(txt string, props *props.Text) string {
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
