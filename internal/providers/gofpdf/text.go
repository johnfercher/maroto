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
)

type Text struct {
	pdf  gofpdfwrapper.Fpdf
	math core.Math
	font core.Font
}

type textLine struct {
	content string
	width   float64
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
	lines := s.getTextLines(text, textProp, width)

	accumulateOffsetY := 0.0

	for index, line := range lines {
		s.addLine(textProp, x, width, y+float64(index)*fontHeight+accumulateOffsetY, line.width, line.content)
		accumulateOffsetY += textProp.VerticalPadding
	}

	s.font.SetColor(originalColor)
}

// GetLinesQuantity retrieve the quantity of lines which a text will occupy to avoid that text to extrapolate a cell.
func (s *Text) GetLinesQuantity(text string, textProp *props.Text, colWidth float64) int {
	s.font.SetFont(textProp.Family, textProp.Style, textProp.Size)
	return len(s.getTextLines(text, textProp, colWidth))
}

func (s *Text) getLines(text string, strategy breakline.Strategy, colWidth float64) []string {
	switch strategy {
	case breakline.EmptySpaceStrategy:
		return s.getLinesBreakingLineFromSpace(strings.Split(text, " "), colWidth)
	case breakline.DashStrategy:
		return s.getLinesBreakingLineWithDash(text, colWidth)
	case breakline.CharacterStrategy:
		return s.getLinesBreakingLineByCharacter(text, colWidth)
	default:
		return s.getLinesBreakingLineFromSpace(strings.Split(text, " "), colWidth)
	}
}

func (s *Text) getLinesBreakingLineFromSpace(words []string, colWidth float64) []string {
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

		width := s.pdf.GetStringWidth(piece)
		if currentlySize+width <= colWidth {
			if len(lines) == 0 {
				lines = append(lines, "")
			}
			lines[len(lines)-1] += separator + word
			currentlySize += width
		} else {
			lines = append(lines, word)
			currentlySize = s.pdf.GetStringWidth(word)
		}
	}

	return lines
}

func (s *Text) getTextLines(text string, textProp *props.Text, colWidth float64) []textLine {
	normalizedText := normalizeLineBreaks(text, textProp.PreserveLineBreaks)
	unicodeText := s.textToUnicode(normalizedText, textProp)
	paragraphs := strings.Split(unicodeText, "\n")
	lines := make([]textLine, 0, len(paragraphs))

	for _, paragraph := range paragraphs {
		paragraphWidth := s.pdf.GetStringWidth(paragraph)
		if paragraphWidth <= colWidth {
			lines = append(lines, textLine{
				content: paragraph,
				width:   paragraphWidth,
			})
			continue
		}

		wrappedParagraph := s.getLines(paragraph, textProp.BreakLineStrategy, colWidth)

		for _, line := range wrappedParagraph {
			lines = append(lines, textLine{
				content: line,
				width:   s.pdf.GetStringWidth(line),
			})
		}
	}

	return lines
}

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

func (s *Text) getLinesBreakingLineByCharacter(words string, colWidth float64) []string {
	currentlySize := 0.0
	lines := []string{}
	var content string

	for _, letter := range words {
		letterString := fmt.Sprintf("%c", letter)
		width := s.pdf.GetStringWidth(letterString)

		if currentlySize+width > colWidth && content != "" {
			lines = append(lines, content)
			content = ""
			currentlySize = 0
		}

		// Skip spaces if they would be at the start of a new line.
		if letterString == " " && content == "" {
			continue
		}

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

func normalizeLineBreaks(text string, preserve bool) string {
	text = strings.ReplaceAll(text, "\r\n", "\n")
	text = strings.ReplaceAll(text, "\r", "\n")
	if preserve {
		return text
	}
	return strings.ReplaceAll(text, "\n", " ")
}
