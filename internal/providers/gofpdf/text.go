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

	// Determine the lines up-front so multi-line rotation can pivot around the
	// whole block, not just the first line.
	var lines []string
	if stringWidth <= width {
		lines = []string{unicodeText}
	} else if textProp.BreakLineStrategy == breakline.EmptySpaceStrategy {
		lines = s.getLinesBreakingLineFromSpace(strings.Split(unicodeText, " "), width)
	} else {
		lines = s.getLinesBreakingLineWithDash(unicodeText, width)
	}

	// Rotation honours both axes of textProp.RotationPivot. The baseline of
	// the first line is shifted so the rotated bounding box of the whole
	// (multi-line) block sits inside the Text.GetHeight-expanded cell.
	if textProp.Rotation != 0 {
		marginLeft, marginTop, _, _ := s.pdf.GetMargins()
		n := float64(len(lines))
		textHeight := n*fontHeight + (n-1)*textProp.VerticalPadding
		blockWidth := stringWidth
		if blockWidth > width {
			blockWidth = width
		}

		var alignOffsetX float64
		switch textProp.Align {
		case align.Center:
			alignOffsetX = (width - blockWidth) / 2
		case align.Right:
			alignOffsetX = width - blockWidth
		}
		if alignOffsetX < 0 {
			alignOffsetX = 0
		}

		var pivotOffsetX float64
		switch textProp.RotationPivot.Horizontal {
		case rotationpivot.Start:
			pivotOffsetX = 0
		case rotationpivot.End:
			pivotOffsetX = blockWidth
		default: // Center
			pivotOffsetX = blockWidth / 2
		}
		var pivotOffsetY float64
		switch textProp.RotationPivot.Vertical {
		case rotationpivot.Top:
			pivotOffsetY = 0
		case rotationpivot.Bottom:
			pivotOffsetY = textHeight
		default: // Middle
			pivotOffsetY = textHeight / 2
		}

		rad := textProp.Rotation * math.Pi / 180
		sin, cos := math.Sin(rad), math.Cos(rad)
		// Distance any rotated corner rises above the pivot. Corners relative
		// to the pivot are TL=(-px,-py), TR=(W-px,-py), BR=(W-px,H-py), BL=(-px,H-py);
		// rotated y is -dx*sin + dy*cos, so -y = dx*sin - dy*cos. Take the max.
		px, py := pivotOffsetX, pivotOffsetY
		W, H := blockWidth, textHeight
		upExtent := math.Max(0, math.Max(py*cos-px*sin,
			math.Max((W-px)*sin+py*cos,
				math.Max((W-px)*sin-(H-py)*cos,
					-px*sin-(H-py)*cos))))

		contentHeight := cell.Height - textProp.Top - textProp.Bottom
		if contentHeight > textHeight {
			// place the rotated bbox top at the cell content top
			y = cell.Y + textProp.Top + upExtent + fontHeight - pivotOffsetY
		}
		pivotX := x + alignOffsetX + pivotOffsetX + marginLeft
		pivotY := y + (pivotOffsetY - fontHeight) + marginTop
		s.pdf.TransformBegin()
		s.pdf.TransformRotate(textProp.Rotation, pivotX, pivotY)
		defer s.pdf.TransformEnd()
	}

	if len(lines) == 1 {
		s.addLine(textProp, x, width, y, stringWidth, lines[0])
		s.font.SetColor(originalColor)
		return
	}

	accumulateOffsetY := 0.0

	for index, line := range lines {
		lineWidth := s.pdf.GetStringWidth(line)

		s.addLine(textProp, x, width, y+float64(index)*fontHeight+accumulateOffsetY, lineWidth, line)
		accumulateOffsetY += textProp.VerticalPadding
	}

	s.font.SetColor(originalColor)
}

// GetStringWidth returns the rendered width of the text after font selection
// and unicode translation.
func (s *Text) GetStringWidth(text string, textProp *props.Text) float64 {
	s.font.SetFont(textProp.Family, textProp.Style, textProp.Size)
	return s.pdf.GetStringWidth(s.textToUnicode(text, textProp))
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
