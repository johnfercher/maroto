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

// This method is responsible for allowing the union of a set of subtexts in the same text
func (s *text) AddCustomText(cell *entity.Cell, textPs *props.Text, subs ...*entity.SubText) {
	originalColor := s.font.GetColor()
	defer s.font.SetColor(originalColor)

	availableWidth := s.validTextProps(textPs, cell, cell.Width-textPs.Left-textPs.Right)
	lines := s.splitTextByLine(availableWidth, textPs.BreakLineStrategy, subs...)

	yOfLine := cell.Y + textPs.Top
	for _, line := range lines {
		x := cell.X + textPs.Left
		heightLargestFont := s.findLargestFontHeight(line...)
		yOfLine += heightLargestFont + textPs.VerticalPadding

		for _, subtext := range line {
			s.addLine(&subtext.Props, x, availableWidth, yOfLine, subtext.Value, textPs.Align)
			x += s.pdf.GetStringWidth(subtext.Value)
		}
	}
}

// Add a text inside a cell.
func (s *text) Add(text string, cell *entity.Cell, textProp *props.Text) {
	s.AddCustomText(cell, textProp, &entity.SubText{Value: text, Props: props.NewSubText(textProp)})
}

// GetTextHeight returns the height occupied by the text in the document
func (s *text) GetTextHeight(textProp *props.Text, availableWidth float64, text ...*entity.SubText) float64 {
	lines := s.splitTextByLine(availableWidth, textProp.BreakLineStrategy, text...)
	if len(text) < 2 {
		lineHeight := s.font.GetHeight(text[0].Props.Family, text[0].Props.Style, text[0].Props.Size) + textProp.VerticalPadding
		return lineHeight * float64(len(lines))
	}

	textHeight := 0.0
	for _, line := range lines {
		textHeight += s.findLargestFontHeight(line...) + textProp.VerticalPadding
	}
	return textHeight
}

// getLinesBreakingLineFromSpace is responsible for sending text in lines, breaking lines in the spaces between words.
func (s *text) getLinesBreakingLineFromSpace(text string, colWidth, currentlySize float64) ([]string, float64) {
	actualLine := 0
	words := strings.Split(text, " ")

	lines := []string{}
	lines = append(lines, "")

	for _, word := range words {
		if s.fitsInTheCurrentLine(word+" ", currentlySize, colWidth) {
			lines[actualLine] = lines[actualLine] + word + " "
			currentlySize += s.pdf.GetStringWidth(word + " ")
		} else {
			lines = append(lines, "")
			actualLine++
			lines[actualLine] = lines[actualLine] + word + " "
			currentlySize = s.pdf.GetStringWidth(word + " ")
		}
	}

	return lines, currentlySize
}

func (s *text) getLinesBreakingLineWithDash(words string, colWidth, currentlySize float64) ([]string, float64) {
	lines := []string{}
	dashSize := s.pdf.GetStringWidth(" - ")

	var content string
	for _, letter := range words {
		if !s.fitsInTheCurrentLine(" - ", currentlySize, colWidth-dashSize) {
			content += "-"
			lines = append(lines, content)
			content = ""
			currentlySize = 0
		}

		letterString := fmt.Sprintf("%c", letter)
		content += letterString
		currentlySize += s.pdf.GetStringWidth(letterString)
	}

	if content != "" {
		lines = append(lines, content)
	}
	return lines, currentlySize
}

func (s *text) addLine(subProp *props.SubText, xColOffset, colWidth, yColOffset float64, text string, alignText align.Type) {
	left, top, _, _ := s.pdf.GetMargins()
	textWidth := s.pdf.GetStringWidth(text)
	s.font.SetColor(subProp.Color)

	fontHeight := s.font.GetHeight(subProp.Family, subProp.Style, subProp.Size)

	if alignText == align.Left {
		s.pdf.Text(xColOffset+left, yColOffset+top, text)
		if subProp.Hyperlink != nil {
			s.pdf.LinkString(xColOffset+left, yColOffset+top-fontHeight, textWidth, fontHeight, *subProp.Hyperlink)
		}
		return
	}

	if alignText == align.Justify {
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

		if subProp.Hyperlink != nil {
			s.pdf.LinkString(initX, yColOffset+top-fontHeight, finishX-initX, fontHeight, *subProp.Hyperlink)
		}

		return
	}

	var modifier float64 = 2

	if alignText == align.Right {
		modifier = 1
	}

	dx := (colWidth - textWidth) / modifier

	if subProp.Hyperlink != nil {
		s.pdf.LinkString(dx+xColOffset+left, yColOffset+top-fontHeight, textWidth, fontHeight, *subProp.Hyperlink)
	}

	s.pdf.Text(dx+xColOffset+left, yColOffset+top, text)
}

func (s *text) textToUnicode(txt string, props *props.SubText) string {
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

	lastChar := rune(text[len(text)-1])
	return !unicode.IsLetter(lastChar) && !unicode.IsNumber(lastChar)
}

// This method is responsible for validating text properties, ensuring that they are within the standard
func (s *text) validTextProps(ps *props.Text, cell *entity.Cell, width float64) float64 {
	if ps.Top > cell.Height {
		ps.Top = cell.Height
	}

	if ps.Left > cell.Width {
		ps.Left = cell.Width
	}

	if ps.Right > cell.Width {
		ps.Right = cell.Width
	}
	if width < 0 {
		width = 0
	}
	return width
}

// splitTextByLine is responsible for order subtext in lines.
//
// The organization of the lines will be done considering the available space and the chosen line
// break strategy. The returned array has all the subtexts of the line and all the lines of the text
func (s *text) splitTextByLine(widthAvailable float64, breakLineStrategy breakline.Strategy, subs ...*entity.SubText) [][]*entity.SubText {
	sizeLasLine := 0.0
	newText := [][]*entity.SubText{}

	for _, sub := range subs {
		sizeLasLine, newText = s.factoryLine(sub, widthAvailable, sizeLasLine, newText, breakLineStrategy)
	}
	return newText
}

// This method is responsible for making a new line with the subText sent.
func (s *text) factoryLine(sub *entity.SubText, width float64, sizeLastLine float64, newText [][]*entity.SubText,
	strategy breakline.Strategy,
) (float64, [][]*entity.SubText) {
	getLines := s.selectStrategyBreak(strategy)
	s.font.SetFont(sub.Props.Family, sub.Props.Style, sub.Props.Size)
	lineValues, currentSize := getLines(s.textToUnicode(sub.Value, &sub.Props), width, sizeLastLine)
	return currentSize, s.mergeSubtext(newText, lineValues, s.fitsInTheCurrentLine(lineValues[0], sizeLastLine, width), sub.Props)
}

// This method is responsible for checking if the word fits on the current line
func (s *text) fitsInTheCurrentLine(word string, currentSize, widthAvailable float64) bool {
	return s.pdf.GetStringWidth(word)+currentSize < widthAvailable
}

// This method is responsible for selecting the correct function to break the line according to the passed strategy
func (s *text) selectStrategyBreak(strategy breakline.Strategy) func(string, float64, float64) ([]string, float64) {
	if strategy == breakline.EmptySpaceStrategy {
		return s.getLinesBreakingLineFromSpace
	} else {
		return s.getLinesBreakingLineWithDash
	}
}

// This method is responsible for finding the font with the highest height in the set of subtexts.
func (s *text) findLargestFontHeight(subs ...*entity.SubText) float64 {
	fontSize := 0.0
	for _, sub := range subs {
		size := s.font.GetHeight(sub.Props.Family, sub.Props.Style, sub.Props.Size)
		if size > fontSize {
			fontSize = size
		}
	}
	return fontSize
}

// This function is responsible for adding the new line to the group of lines already created,
// merging the subtexts into the same line if necessary.
func (s *text) mergeSubtext(currentLines [][]*entity.SubText, newLines []string, merge bool, ps props.SubText) [][]*entity.SubText {
	startInsert := 0

	if merge && len(currentLines) != 0 {
		currentLines[len(currentLines)-1] = append(currentLines[len(currentLines)-1], &entity.SubText{Value: newLines[0], Props: ps})
		startInsert = 1
	}

	for i := startInsert; i < len(newLines); i++ {
		currentLines = append(currentLines, []*entity.SubText{{Value: newLines[i], Props: ps}})
	}
	return currentLines
}
