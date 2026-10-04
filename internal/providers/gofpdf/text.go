package gofpdf

import (
	"fmt"
	"math"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/breakline"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/rotationpivot"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
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
	stringWidth := s.pdf.GetStringWidth(unicodeText)

	// Split up-front so multi-line rotation can pivot around the whole block.
	lines := s.splitLines(unicodeText, textProp, width)

	lineProp := textProp
	if textProp.Rotation != 0 {
		y = s.rotate(lines, cell, textProp, x, width, fontHeight)
		defer s.pdf.TransformEnd()

		if textProp.Hyperlink != nil {
			unlinked := *textProp
			unlinked.Hyperlink = nil
			lineProp = &unlinked
		}
	}

	if len(lines) == 1 {
		s.addLine(lineProp, x, width, y, stringWidth, lines[0])
		s.font.SetColor(originalColor)
		return
	}

	accumulateOffsetY := 0.0

	for index, line := range lines {
		lineWidth := s.pdf.GetStringWidth(line)

		s.addLine(lineProp, x, width, y+float64(index)*fontHeight+accumulateOffsetY, lineWidth, line)
		accumulateOffsetY += textProp.VerticalPadding
	}

	s.font.SetColor(originalColor)
}

// GetLinesWidth returns the width of the widest line the text is drawn with in a column of colWidth.
// It can exceed colWidth: a word too long to wrap, or a narrow DashStrategy line, is drawn wider.
func (s *Text) GetLinesWidth(text string, textProp *props.Text, colWidth float64) float64 {
	s.font.SetFont(textProp.Family, textProp.Style, textProp.Size)

	var width float64
	for _, line := range s.splitLines(s.textToUnicode(text, textProp), textProp, colWidth) {
		width = max(width, s.pdf.GetStringWidth(line))
	}
	return width
}

// GetLinesQuantity retrieve the quantity of lines which a text will occupy to avoid that text to extrapolate a cell.
func (s *Text) GetLinesQuantity(text string, textProp *props.Text, colWidth float64) int {
	s.font.SetFont(textProp.Family, textProp.Style, textProp.Size)

	textTranslated := s.textToUnicode(text, textProp)

	if textProp.BreakLineStrategy == breakline.DashStrategy {
		return len(s.getLinesBreakingLineWithDash(text, colWidth))
	}

	return len(s.getLinesBreakingLineFromSpace(strings.Split(textTranslated, " "), colWidth))
}

// rotate starts the rotation transform of a text block and returns the baseline of its first line,
// moved so the rotated block starts at the top of the cell. Both axes of RotationPivot are honored,
// and for multi-line text the whole block rotates as one.
func (s *Text) rotate(lines []string, cell *entity.Cell, textProp *props.Text, x, width, fontHeight float64) float64 {
	marginLeft, marginTop, _, _ := s.pdf.GetMargins()
	n := float64(len(lines))
	textHeight := n*fontHeight + (n-1)*textProp.VerticalPadding

	// The widest drawn line, which can exceed the column when a word is too long to wrap.
	var blockWidth float64
	for _, line := range lines {
		blockWidth = max(blockWidth, s.pdf.GetStringWidth(line))
	}

	var alignOffsetX float64
	switch textProp.Align {
	case align.Center:
		alignOffsetX = (width - blockWidth) / 2
	case align.Right:
		alignOffsetX = width - blockWidth
	case align.Left, align.Top, align.Bottom, align.Middle:
		alignOffsetX = 0
	}
	alignOffsetX = max(alignOffsetX, 0)

	var pivotOffsetX float64
	switch textProp.RotationPivot.Horizontal {
	case rotationpivot.Start:
		pivotOffsetX = 0
	case rotationpivot.End:
		pivotOffsetX = blockWidth
	case rotationpivot.Center:
		pivotOffsetX = blockWidth / 2
	default:
		pivotOffsetX = blockWidth / 2
	}
	var pivotOffsetY float64
	switch textProp.RotationPivot.Vertical {
	case rotationpivot.Top:
		pivotOffsetY = 0
	case rotationpivot.Bottom:
		pivotOffsetY = textHeight
	case rotationpivot.Middle:
		pivotOffsetY = textHeight / 2
	default:
		pivotOffsetY = textHeight / 2
	}

	rad := textProp.Rotation * math.Pi / 180
	sin, cos := math.Sin(rad), math.Cos(rad)
	// Corners relative to the pivot, rotated counter-clockwise on a y-down page.
	px, py := pivotOffsetX, pivotOffsetY
	corners := [4][2]float64{{-px, -py}, {blockWidth - px, -py}, {blockWidth - px, textHeight - py}, {-px, textHeight - py}}
	minX, minY, maxX, maxY := math.Inf(1), math.Inf(1), math.Inf(-1), math.Inf(-1)
	for _, c := range corners {
		rx, ry := c[0]*cos+c[1]*sin, -c[0]*sin+c[1]*cos
		minX, maxX = min(minX, rx), max(maxX, rx)
		minY, maxY = min(minY, ry), max(maxY, ry)
	}

	y := cell.Y + textProp.Top - min(minY, 0) + fontHeight - pivotOffsetY
	pivotX := x + alignOffsetX + pivotOffsetX + marginLeft
	pivotY := y + (pivotOffsetY - fontHeight) + marginTop

	s.pdf.TransformBegin()
	s.pdf.TransformRotate(textProp.Rotation, pivotX, pivotY)

	// A link annotation ignores the transform, so it covers the rotated block instead of each line.
	if textProp.Hyperlink != nil {
		s.pdf.LinkString(pivotX+minX, pivotY+minY, maxX-minX, maxY-minY, *textProp.Hyperlink)
	}

	return y
}

// splitLines breaks the text into the lines Add draws it with.
func (s *Text) splitLines(text string, textProp *props.Text, colWidth float64) []string {
	switch {
	case s.pdf.GetStringWidth(text) <= colWidth:
		return []string{text}
	case textProp.BreakLineStrategy == breakline.EmptySpaceStrategy:
		return s.getLinesBreakingLineFromSpace(strings.Split(text, " "), colWidth)
	default:
		return s.getLinesBreakingLineWithDash(text, colWidth)
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
