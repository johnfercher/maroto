package gofpdf

import (
	"fmt"
	"strings"

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

// This method is responsible for adding new text to the file
func (s *text) Add(text string, cell *entity.Cell, textProp *props.Text) {
	entity.NewSubText(text, props.NewSubText(textProp))
	s.AddCustomText([]*entity.SubText{entity.NewSubText(text, props.NewSubText(textProp))}, cell, textProp)
}

// This method is responsible for allowing the union of a set of subtexts in the same text
func (s *text) AddCustomText(subs []*entity.SubText, cell *entity.Cell, textPs *props.Text) {
	originalColor := s.font.GetColor()
	defer s.font.SetColor(originalColor)

	width := s.validTextProps(textPs, cell, cell.Width-textPs.Left-textPs.Right)
	lines := s.orderSubTexts(subs, width, textPs.BreakLineStrategy)

	accumulateOffsetY := 0.0
	sumFonts := 0.0
	for index, line := range lines {
		x := cell.X + textPs.Left
		heightLargestFont := s.findLargestFontHeight(line)
		if index > 0 {
			sumFonts += heightLargestFont
		}

		for _, subText := range line {
			y := s.setTheLineProps(subText, heightLargestFont, cell.Y+textPs.Top)
			s.addLine(&subText.Prop, x, width, y+sumFonts+accumulateOffsetY, subText.Value, textPs.Align)
			x += s.pdf.GetStringWidth(subText.Value)
		}
		accumulateOffsetY += textPs.VerticalPadding
	}
}

// GetLinesQuantity retrieve the quantity of lines which a text will occupy to avoid that text to extrapolate a cell.
func (s *text) GetLinesQuantity(text string, textProp props.Text, colWidth float64) int {
	translator := s.pdf.UnicodeTranslatorFromDescriptor("")
	s.font.SetFont(textProp.Family, textProp.Style, textProp.Size)

	// Apply Unicode.
	textTranslated := translator(text)

	words := strings.Split(textTranslated, " ")

	// If should add one line.
	if s.pdf.GetStringWidth(textTranslated) < colWidth || len(words) == 1 {
		return 1
	}

	lines, _ := s.getLinesBreakingLineFromSpace(textTranslated, colWidth, 0)
	return len(lines)
}

// This method is responsible for checking the subtexts and adding them to the line according to the chosen strategy
func (s *text) orderSubTexts(subs []*entity.SubText, widthAvailable float64, breakLineStrategy breakline.Strategy) [][]*entity.SubText {
	sizeLasLine := 0.0
	newText := [][]*entity.SubText{}

	for _, sub := range subs {
		sizeLasLine, newText = s.factoryLine(sub, widthAvailable, sizeLasLine, newText, s.selectStrategyBreak(breakLineStrategy))
	}
	return newText
}

// This method is responsible for making a new line with the subText sent.
func (s *text) factoryLine(sub *entity.SubText, widthAvailable float64, sizeLasLine float64, newText [][]*entity.SubText, getLines func(string, float64, float64) ([]string, float64)) (float64, [][]*entity.SubText) {
	s.font.SetFont(sub.Prop.Family, sub.Prop.Style, sub.Prop.Size)
	lineValues, currentSize := getLines(s.textToUnicode(sub.Value, &sub.Prop), widthAvailable, sizeLasLine)
	return currentSize, s.mergeSubtext(newText, lineValues, s.fitsInTheCurrentLine(lineValues[0], sizeLasLine, widthAvailable), sub.Prop)
}

// This method is responsible for defining the line props, returning the position of the text on the y axis
func (s *text) setTheLineProps(sub *entity.SubText, heightLargestFont, y float64) float64 {
	s.font.SetFont(sub.Prop.Family, sub.Prop.Style, sub.Prop.Size)
	s.font.SetColor(sub.Prop.Color)

	fontHeight := s.font.GetHeight(sub.Prop.Family, sub.Prop.Style, sub.Prop.Size)
	if heightLargestFont > fontHeight {
		y -= (heightLargestFont - fontHeight) / 2
	}
	return y + heightLargestFont
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

// This method is responsible for finding the font with the highest height in the set of subtexts.
func (s *text) findLargestFontHeight(subs []*entity.SubText) float64 {
	fontSize := 0.0
	for _, sub := range subs {
		size := s.font.GetHeight(sub.Prop.Family, sub.Prop.Style, sub.Prop.Size)
		if size > fontSize {
			fontSize = size
		}
	}
	return fontSize
}

// This function is responsible for merging the subText, ensuring that when necessary they will occupy the same line
func (s *text) mergeSubtext(currentLines [][]*entity.SubText, newLines []string, joinOnTheSameLine bool, ps props.SubText) [][]*entity.SubText {
	startInsert := 0

	if joinOnTheSameLine && len(currentLines) != 0 {
		currentLines[len(currentLines)-1] = append(currentLines[len(currentLines)-1], entity.NewSubText(newLines[0], ps))
		startInsert = 1
	}

	for i := startInsert; i < len(newLines); i++ {
		currentLines = append(currentLines, []*entity.SubText{entity.NewSubText(newLines[i], ps)})
	}
	return currentLines
}

// This method is responsible for selecting the correct function to break the line according to the passed strategy
func (s *text) selectStrategyBreak(strategy breakline.Strategy) func(text string, colWidth float64, currentlySize float64) ([]string, float64) {
	if strategy == breakline.EmptySpaceStrategy {
		return s.getLinesBreakingLineFromSpace
	} else {
		return s.getLinesBreakingLineWithDash
	}
}

// This method is responsible for checking if the word fits on the current line
func (s *text) fitsInTheCurrentLine(word string, currentSize, widthAvailable float64) bool {
	return s.pdf.GetStringWidth(word)+currentSize < widthAvailable
}

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

	fontHeight := s.font.GetHeight(subProp.Family, subProp.Style, subProp.Size)

	if alignText == align.Left {
		s.pdf.Text(xColOffset+left, yColOffset+top, text)

		if subProp.Hyperlink != nil {
			s.pdf.LinkString(xColOffset+left, yColOffset+top-fontHeight, textWidth, fontHeight, *subProp.Hyperlink)
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
