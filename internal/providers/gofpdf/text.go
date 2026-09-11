package gofpdf

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/breakline"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/johnfercher/maroto/v2/pkg/rtl"
)

type Text struct {
	pdf  gofpdfwrapper.Fpdf
	math core.Math
	font core.Font
}

// NewText create a Text.
func NewText(pdf gofpdfwrapper.Fpdf, math core.Math, font core.Font) *Text {
	return &Text{
		pdf,
		math,
		font,
	}
}

// Add a text inside a cell.
func (s *Text) Add(text string, cell *entity.Cell, textProp *props.Text) {
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
	stringWidth := s.getStringWidth(unicodeText, textProp)

	// If should add one line
	if stringWidth <= width {
		s.addLine(textProp, x, width, y, stringWidth, s.processLine(unicodeText, textProp))
		s.font.SetColor(originalColor)
		return
	}

	var lines []string

	if textProp.BreakLineStrategy == breakline.EmptySpaceStrategy {
		words := strings.Split(unicodeText, " ")
		lines = s.getLinesBreakingLineFromSpace(words, width, textProp)
	} else {
		lines = s.getLinesBreakingLineWithDash(unicodeText, width)
	}

	accumulateOffsetY := 0.0

	// Lines are broken on the logical text, then each one is processed on its
	// own so that it is reordered with its own base direction.
	for index, line := range lines {
		processedLine := s.processLine(line, textProp)
		lineWidth := s.pdf.GetStringWidth(processedLine)

		s.addLine(textProp, x, width, y+float64(index)*fontHeight+accumulateOffsetY, lineWidth, processedLine)
		accumulateOffsetY += textProp.VerticalPadding
	}

	s.font.SetColor(originalColor)
}

// GetLinesQuantity retrieve the quantity of lines which a text will occupy to avoid that text to extrapolate a cell.
func (s *Text) GetLinesQuantity(text string, textProp *props.Text, colWidth float64) int {
	s.font.SetFont(textProp.Family, textProp.Style, textProp.Size)

	textTranslated := s.textToUnicode(text, textProp)

	if textProp.BreakLineStrategy == breakline.DashStrategy {
		return len(s.getLinesBreakingLineWithDash(text, colWidth))
	}

	return len(s.getLinesBreakingLineFromSpace(strings.Split(textTranslated, " "), colWidth, textProp))
}

// processLine returns the line in the form the writer has to draw it. For a
// right-to-left text that means the Arabic letters replaced by their
// contextual forms and the bidirectional runs laid out visually.
//
// It must only be called on a text that already is a single line: processing a
// whole paragraph before breaking it would reverse the order of its lines.
func (s *Text) processLine(text string, textProp *props.Text) string {
	if !textProp.RTL {
		return text
	}

	return rtl.Process(text)
}

// getStringWidth measures a text in the form it is going to be drawn. The
// presentation forms do not share the metrics of the letters they replace, and
// the mandatory ligatures contract two letters into a single glyph, so
// measuring the logical text would break the line wrapping of a right-to-left
// text.
func (s *Text) getStringWidth(text string, textProp *props.Text) float64 {
	return s.pdf.GetStringWidth(s.processLine(text, textProp))
}

func (s *Text) getLinesBreakingLineFromSpace(words []string, colWidth float64, textProp *props.Text) []string {
	currentlySize := 0.0
	lines := []string{}

	for _, word := range words {
		if word == "" {
			continue
		}
		var piece, separator string
		if len(lines) == 0 || lines[len(lines)-1] == "" {
			piece = word
			separator = ""
		} else {
			piece = " " + word
			separator = " "
		}

		width := s.getStringWidth(piece, textProp)
		if currentlySize+width <= colWidth {
			if len(lines) == 0 {
				lines = append(lines, "")
			}
			lines[len(lines)-1] += separator + word
			currentlySize += width
		} else {
			lines = append(lines, word)
			currentlySize = s.getStringWidth(word, textProp)
		}
	}

	return lines
}

// getLinesBreakingLineWithDash breaks a text into lines by hyphenating it at
// the character that no longer fits.
//
// Arabic does not hyphenate, and the width of a letter here is measured on the
// logical character rather than on the presentation form it will be drawn
// with, so the break positions of a right-to-left text are only approximate.
// The lines it returns are still shaped and reordered before being drawn. Use
// breakline.EmptySpaceStrategy, the default, for right-to-left text.
func (s *Text) getLinesBreakingLineWithDash(words string, colWidth float64) []string {
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

func (s *Text) addLine(textProp *props.Text, xColOffset, colWidth, yColOffset, textWidth float64, text string) {
	left, top, _, _ := s.pdf.GetMargins()

	fontHeight := s.font.GetHeight(textProp.Family, textProp.Style, textProp.Size)

	if textProp.Align == align.Left {
		s.pdf.Text(xColOffset+left, yColOffset+top, text)

		if textProp.Hyperlink != nil {
			s.pdf.LinkString(xColOffset+left, yColOffset+top-fontHeight, textWidth, fontHeight, *textProp.Hyperlink)
		}

		return
	}

	if textProp.Align == align.Justify {
		const spaceString = " "
		const emptyString = ""

		text = strings.TrimRight(text, spaceString)
		textNotSpaces := strings.ReplaceAll(text, spaceString, emptyString)
		textWidth = s.pdf.GetStringWidth(textNotSpaces)
		defaultSpaceWidth := s.pdf.GetStringWidth(spaceString)
		words := strings.Fields(text)

		numSpaces := max(len(words)-1, 1)
		spaceWidth := (colWidth - textWidth) / float64(numSpaces)
		x := xColOffset + left

		if isIncorrectSpaceWidth(textWidth, spaceWidth, defaultSpaceWidth, textNotSpaces) {
			spaceWidth = defaultSpaceWidth
		}
		initX := x
		var finishX float64
		for _, word := range words {
			s.pdf.Text(x, yColOffset+top, word)
			finishX = x + s.pdf.GetStringWidth(word)
			x = finishX + spaceWidth
		}

		if textProp.Hyperlink != nil {
			s.pdf.LinkString(initX, yColOffset+top-fontHeight, finishX-initX, fontHeight, *textProp.Hyperlink)
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

func (s *Text) textToUnicode(txt string, props *props.Text) string {
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

func isIncorrectSpaceWidth(textWidth, spaceWidth, defaultSpaceWidth float64, text string) bool {
	if textWidth <= 0 || spaceWidth <= defaultSpaceWidth*10 {
		return false
	}

	r, _ := utf8.DecodeLastRuneInString(text)
	lastChar := r
	return !unicode.IsLetter(lastChar) && !unicode.IsNumber(lastChar)
}
