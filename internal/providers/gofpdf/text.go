package gofpdf

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/breakline"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
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

// validates whether the text position respects the cell limits
func validTextPosition(textProp *props.Text, cell *entity.Cell) {
	if textProp.Top > cell.Height {
		textProp.Top = cell.Height
	}

	if textProp.Left > cell.Width {
		textProp.Left = cell.Width
	}

	if textProp.Right > cell.Width {
		textProp.Right = cell.Width
	}
}

// Add a text inside a cell.
// When usePageMargin is false, text can be positioned at any position in the cell
func (t *text) Add(text string, cell *entity.Cell, textProp *props.Text, usePageMargin ...bool) {
	if len(usePageMargin) > 0 && !usePageMargin[0] {
		left, top, right, _ := t.pdf.GetMargins()
		defer t.pdf.SetMargins(left, top, right)
		t.pdf.SetMargins(0.0, 0.0, 0.0)
	}
	t.font.SetFont(textProp.Family, textProp.Style, textProp.Size)
	fontHeight := t.font.GetHeight(textProp.Family, textProp.Style, textProp.Size)

	validTextPosition(textProp, cell)

	width := cell.Width - textProp.Left - textProp.Right
	if width < 0 {
		width = 0
	}

	x := cell.X + textProp.Left
	y := cell.Y + textProp.Top

	originalColor := t.font.GetColor()
	if textProp.Color != nil {
		t.font.SetColor(textProp.Color)
	}

	// override style if hyperlink is set
	if textProp.Hyperlink != nil {
		t.font.SetColor(&props.BlueColor)
	}

	y += fontHeight

	// Apply Unicode before calc spaces
	unicodeText := t.textToUnicode(text, textProp)
	stringWidth := t.pdf.GetStringWidth(unicodeText)

	// If should add one line
	if stringWidth < width {
		t.addLine(textProp, x, width, y, stringWidth, unicodeText)
		if textProp.Color != nil {
			t.font.SetColor(originalColor)
		}
		return
	}

	var lines []string

	if textProp.BreakLineStrategy == breakline.EmptySpaceStrategy {
		words := strings.Split(unicodeText, " ")
		lines = t.getLinesBreakingLineFromSpace(words, width)
	} else {
		lines = t.getLinesBreakingLineWithDash(unicodeText, width)
	}

	accumulateOffsetY := 0.0

	for index, line := range lines {
		lineWidth := t.pdf.GetStringWidth(line)

		t.addLine(textProp, x, width, y+float64(index)*fontHeight+accumulateOffsetY, lineWidth, line)
		accumulateOffsetY += textProp.VerticalPadding
	}

	if textProp.Color != nil {
		t.font.SetColor(originalColor)
	}
}

// GetLinesQuantity retrieve the quantity of lines which a text will occupy to avoid that text to extrapolate a cell.
func (t *text) GetLinesQuantity(text string, textProp *props.Text, colWidth float64) int {
	t.font.SetFont(textProp.Family, textProp.Style, textProp.Size)

	textTranslated := t.textToUnicode(text, textProp)

	if textProp.BreakLineStrategy == breakline.DashStrategy {
		return len(t.getLinesBreakingLineWithDash(text, colWidth))
	} else {
		return len(t.getLinesBreakingLineFromSpace(strings.Split(textTranslated, " "), colWidth))
	}
}

func (t *text) getLinesBreakingLineFromSpace(words []string, colWidth float64) []string {
	currentlySize := 0.0
	actualLine := 0

	lines := []string{}
	lines = append(lines, "")

	for _, word := range words {
		if t.pdf.GetStringWidth(word+" ")+currentlySize < colWidth {
			lines[actualLine] = lines[actualLine] + word + " "
			currentlySize += t.pdf.GetStringWidth(word + " ")
		} else {
			lines = append(lines, "")
			actualLine++
			lines[actualLine] = lines[actualLine] + word + " "
			currentlySize = t.pdf.GetStringWidth(word + " ")
		}
	}

	return lines
}

func (t *text) getLinesBreakingLineWithDash(words string, colWidth float64) []string {
	currentlySize := 0.0

	lines := []string{}

	dashSize := t.pdf.GetStringWidth(" - ")

	var content string
	for _, letter := range words {
		if currentlySize+dashSize > colWidth-dashSize {
			content += "-"
			lines = append(lines, content)
			content = ""
			currentlySize = 0
		}

		letterString := fmt.Sprintf("%c", letter)
		width := t.pdf.GetStringWidth(letterString)
		content += letterString
		currentlySize += width
	}

	if content != "" {
		lines = append(lines, content)
	}

	return lines
}

func (t *text) addLine(textProp *props.Text, xColOffset, colWidth, yColOffset, textWidth float64, text string) {
	left, top, _, _ := t.pdf.GetMargins()

	fontHeight := t.font.GetHeight(textProp.Family, textProp.Style, textProp.Size)

	if textProp.Align == align.Left {
		t.pdf.Text(xColOffset+left, yColOffset+top, text)

		if textProp.Hyperlink != nil {
			t.pdf.LinkString(xColOffset+left, yColOffset+top-fontHeight, textWidth, fontHeight, *textProp.Hyperlink)
		}

		return
	}

	if textProp.Align == align.Justify {
		const spaceString = " "
		const emptyString = ""

		text = strings.TrimRight(text, spaceString)
		textNotSpaces := strings.ReplaceAll(text, spaceString, emptyString)
		textWidth = t.pdf.GetStringWidth(textNotSpaces)
		defaultSpaceWidth := t.pdf.GetStringWidth(spaceString)
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
			t.pdf.Text(x, yColOffset+top, word)
			finishX = x + t.pdf.GetStringWidth(word)
			x = finishX + spaceWidth
		}

		if textProp.Hyperlink != nil {
			t.pdf.LinkString(initX, yColOffset+top-fontHeight, finishX-initX, fontHeight, *textProp.Hyperlink)
		}

		return
	}

	var modifier float64 = 2

	if textProp.Align == align.Right {
		modifier = 1
	}

	dx := (colWidth - textWidth) / modifier

	if textProp.Hyperlink != nil {
		t.pdf.LinkString(dx+xColOffset+left, yColOffset+top-fontHeight, textWidth, fontHeight, *textProp.Hyperlink)
	}

	t.pdf.Text(dx+xColOffset+left, yColOffset+top, text)
}

func (t *text) textToUnicode(txt string, props *props.Text) string {
	if props.Family == fontfamily.Arial ||
		props.Family == fontfamily.Helvetica ||
		props.Family == fontfamily.Symbol ||
		props.Family == fontfamily.ZapBats ||
		props.Family == fontfamily.Courier {
		translator := t.pdf.UnicodeTranslatorFromDescriptor("")
		return translator(txt)
	}

	return txt
}

func isIncorrectSpaceWidth(textWidth, spaceWidth, defaultSpaceWidth float64, text string) bool {
	if textWidth <= 0 || spaceWidth <= defaultSpaceWidth*10 {
		return false
	}

	lastChar := rune(text[len(text)-1])
	return !unicode.IsLetter(lastChar) && !unicode.IsNumber(lastChar)
}
