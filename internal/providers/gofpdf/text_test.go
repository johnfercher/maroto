package gofpdf_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/breakline"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"

	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewText(t *testing.T) {
	t.Parallel()
	text := gofpdf.NewText(mocks.NewFpdf(t), mocks.NewMath(t), mocks.NewFont(t))

	assert.NotNil(t, text)
	assert.Equal(t, "*gofpdf.Text", fmt.Sprintf("%T", text))
}

func TestGetLinesHeight(t *testing.T) {
	t.Parallel()
	t.Run("when a text that occupies two lines is sent with EmptySpaceStrategy, should two is returned", func(t *testing.T) {
		t.Parallel()
		textProp := &props.Text{}
		textProp.MakeValid(&props.Font{Family: fontfamily.Arial, Size: 10, Style: fontstyle.Normal})

		font := mocks.NewFont(t)
		font.EXPECT().SetFont(textProp.Family, textProp.Style, textProp.Size)

		pdf := mocks.NewFpdf(t)
		pdf.EXPECT().UnicodeTranslatorFromDescriptor("").Return(func(s string) string { return s })
		pdf.EXPECT().GetStringWidth("text text text text").Return(22.0)
		pdf.EXPECT().GetStringWidth("text").Return(5.0)  // First token just returns text
		pdf.EXPECT().GetStringWidth(" text").Return(6.0) // subsequent tokens return leading space
		pdf.EXPECT().GetStringWidth("text text").Return(11.0)

		text := gofpdf.NewText(pdf, mocks.NewMath(t), font)

		height := text.GetLinesQuantity("text text text text", textProp, 11)

		assert.Equal(t, 2, height)
	})

	t.Run("When a text that occupies two lines is sent with EmptySpaceStrategy, should two is returned", func(t *testing.T) {
		t.Parallel()
		textProp := &props.Text{BreakLineStrategy: breakline.DashStrategy}
		textProp.MakeValid(&props.Font{Family: fontfamily.Arial, Size: 10, Style: fontstyle.Normal})

		font := mocks.NewFont(t)
		font.EXPECT().SetFont(textProp.Family, textProp.Style, textProp.Size)

		pdf := mocks.NewFpdf(t)
		pdf.EXPECT().GetStringWidth("tttt tttt tttt tttt").Return(12)
		pdf.EXPECT().GetStringWidth("t").Return(1)
		pdf.EXPECT().GetStringWidth(" ").Return(1)
		pdf.EXPECT().GetStringWidth(" - ").Return(1)
		pdf.EXPECT().GetStringWidth("tttt tttt -").Return(10)
		pdf.EXPECT().GetStringWidth("tttt tttt").Return(9)
		pdf.EXPECT().UnicodeTranslatorFromDescriptor("").Return(func(s string) string { return s })

		text := gofpdf.NewText(pdf, mocks.NewMath(t), font)

		height := text.GetLinesQuantity("tttt tttt tttt tttt", textProp, 11)

		assert.Equal(t, 2, height)
	})

	t.Run("when translated text occupies two lines with CharacterStrategy, should return two", func(t *testing.T) {
		t.Parallel()
		textProp := &props.Text{BreakLineStrategy: breakline.CharacterStrategy}
		textProp.MakeValid(&props.Font{Family: fontfamily.Arial, Size: 10, Style: fontstyle.Normal})

		font := mocks.NewFont(t)
		font.EXPECT().SetFont(textProp.Family, textProp.Style, textProp.Size)

		pdf := mocks.NewFpdf(t)
		pdf.EXPECT().UnicodeTranslatorFromDescriptor("").Return(func(string) string { return "aaaa" })
		pdf.EXPECT().GetStringWidth("aaaa").Return(12.0)
		pdf.EXPECT().GetStringWidth("a").Return(3.0).Times(4)
		pdf.EXPECT().GetStringWidth("aa").Return(6.0).Twice()

		text := gofpdf.NewText(pdf, mocks.NewMath(t), font)

		height := text.GetLinesQuantity("ääää", textProp, 6)

		assert.Equal(t, 2, height)
	})

	t.Run("when cp1252 translated text occupies two lines with CharacterStrategy, should measure each byte", func(t *testing.T) {
		t.Parallel()
		textProp := &props.Text{BreakLineStrategy: breakline.CharacterStrategy}
		textProp.MakeValid(&props.Font{Family: fontfamily.Arial, Size: 10, Style: fontstyle.Normal})

		font := mocks.NewFont(t)
		font.EXPECT().SetFont(textProp.Family, textProp.Style, textProp.Size)

		pdf := mocks.NewFpdf(t)
		pdf.EXPECT().UnicodeTranslatorFromDescriptor("").Return(func(string) string { return "\xe4\xe4\xe4\xe4" })
		pdf.EXPECT().GetStringWidth("\xe4\xe4\xe4\xe4").Return(12.0)
		pdf.EXPECT().GetStringWidth("\xe4").Return(3.0).Times(4)
		pdf.EXPECT().GetStringWidth("\xe4\xe4").Return(6.0).Twice()

		text := gofpdf.NewText(pdf, mocks.NewMath(t), font)

		height := text.GetLinesQuantity("ääää", textProp, 6)

		assert.Equal(t, 2, height)
	})

	t.Run("when utf-8 text occupies two lines with CharacterStrategy, should measure each rune", func(t *testing.T) {
		t.Parallel()
		textProp := &props.Text{BreakLineStrategy: breakline.CharacterStrategy}
		textProp.MakeValid(&props.Font{Family: "custom", Size: 10, Style: fontstyle.Normal})

		font := mocks.NewFont(t)
		font.EXPECT().SetFont(textProp.Family, textProp.Style, textProp.Size)

		pdf := mocks.NewFpdf(t)
		pdf.EXPECT().GetStringWidth("їїїї").Return(12.0)
		pdf.EXPECT().GetStringWidth("ї").Return(3.0).Times(4)
		pdf.EXPECT().GetStringWidth("її").Return(6.0).Twice()

		text := gofpdf.NewText(pdf, mocks.NewMath(t), font)

		height := text.GetLinesQuantity("їїїї", textProp, 6)

		assert.Equal(t, 2, height)
	})

	t.Run("when text is blank with CharacterStrategy, should return one line as rendered", func(t *testing.T) {
		t.Parallel()
		for _, blank := range []string{"", " "} {
			textProp := &props.Text{BreakLineStrategy: breakline.CharacterStrategy, VerticalPadding: 10}
			textProp.MakeValid(&props.Font{Family: "custom", Size: 10, Style: fontstyle.Normal})

			font := mocks.NewFont(t)
			font.EXPECT().SetFont(textProp.Family, textProp.Style, textProp.Size)

			pdf := mocks.NewFpdf(t)
			pdf.EXPECT().GetStringWidth("").Return(0.0).Maybe()
			pdf.EXPECT().GetStringWidth(" ").Return(1.0).Maybe()

			text := gofpdf.NewText(pdf, mocks.NewMath(t), font)

			height := text.GetLinesQuantity(blank, textProp, 6)

			assert.Equal(t, 1, height, "text %q", blank)
		}
	})

	t.Run("when a space-only paragraph does not fit a zero width, should keep it as an empty line", func(t *testing.T) {
		t.Parallel()
		for _, strategy := range []breakline.Strategy{breakline.EmptySpaceStrategy, breakline.CharacterStrategy} {
			textProp := &props.Text{PreserveLineBreaks: true, BreakLineStrategy: strategy}
			textProp.MakeValid(&props.Font{Family: "custom", Size: 10, Style: fontstyle.Normal})

			font := mocks.NewFont(t)
			font.EXPECT().SetFont(textProp.Family, textProp.Style, textProp.Size)

			pdf := mocks.NewFpdf(t)
			pdf.EXPECT().GetStringWidth("A").Return(5.0)
			pdf.EXPECT().GetStringWidth(" ").Return(1.0)
			pdf.EXPECT().GetStringWidth("B").Return(5.0)

			text := gofpdf.NewText(pdf, mocks.NewMath(t), font)

			height := text.GetLinesQuantity("A\n \nB", textProp, 0)

			assert.Equal(t, 3, height, "strategy %s", strategy)
		}
	})

	t.Run("when LF, CRLF or CR line breaks are counted, should depend on PreserveLineBreaks", func(t *testing.T) {
		t.Parallel()
		for _, value := range []string{"A\nB", "A\r\nB", "A\rB"} {
			for preserve, expected := range map[bool]int{true: 2, false: 1} {
				textProp := &props.Text{PreserveLineBreaks: preserve}
				textProp.MakeValid(&props.Font{Family: "custom", Size: 10, Style: fontstyle.Normal})

				font := mocks.NewFont(t)
				font.EXPECT().SetFont(textProp.Family, textProp.Style, textProp.Size)

				pdf := mocks.NewFpdf(t)
				pdf.EXPECT().GetStringWidth("A").Return(5.0).Maybe()
				pdf.EXPECT().GetStringWidth("B").Return(5.0).Maybe()
				pdf.EXPECT().GetStringWidth("A B").Return(12.0).Maybe()

				text := gofpdf.NewText(pdf, mocks.NewMath(t), font)

				height := text.GetLinesQuantity(value, textProp, 100)

				assert.Equal(t, expected, height, "text %q, preserve %t", value, preserve)
			}
		}
	})

	t.Run("when a text contains explicit empty lines, should count each paragraph break", func(t *testing.T) {
		t.Parallel()
		textProp := &props.Text{PreserveLineBreaks: true}
		textProp.MakeValid(&props.Font{Family: fontfamily.Arial, Size: 10, Style: fontstyle.Normal})

		font := mocks.NewFont(t)
		font.EXPECT().SetFont(textProp.Family, textProp.Style, textProp.Size)

		pdf := mocks.NewFpdf(t)
		pdf.EXPECT().UnicodeTranslatorFromDescriptor("").Return(func(s string) string { return s })
		pdf.EXPECT().GetStringWidth("line one").Return(20.0)
		pdf.EXPECT().GetStringWidth("").Return(0.0)
		pdf.EXPECT().GetStringWidth("line two").Return(20.0)

		text := gofpdf.NewText(pdf, mocks.NewMath(t), font)

		height := text.GetLinesQuantity("line one\n\nline two", textProp, 100)

		assert.Equal(t, 3, height)
	})

	t.Run("when margins exceed the column width, should count lines as rendered with zero width", func(t *testing.T) {
		t.Parallel()
		textProp := &props.Text{PreserveLineBreaks: true}
		textProp.MakeValid(&props.Font{Family: fontfamily.Arial, Size: 10, Style: fontstyle.Normal})

		font := mocks.NewFont(t)
		font.EXPECT().SetFont(textProp.Family, textProp.Style, textProp.Size)

		pdf := mocks.NewFpdf(t)
		pdf.EXPECT().UnicodeTranslatorFromDescriptor("").Return(func(s string) string { return s })
		pdf.EXPECT().GetStringWidth("A").Return(5.0)
		pdf.EXPECT().GetStringWidth("").Return(0.0)
		pdf.EXPECT().GetStringWidth("B").Return(5.0)

		text := gofpdf.NewText(pdf, mocks.NewMath(t), font)

		height := text.GetLinesQuantity("A\n\nB", textProp, -2)

		assert.Equal(t, 3, height)
	})
}

func TestText_Add(t *testing.T) {
	t.Parallel()
	t.Run("when single line with left align and no color and no hyperlink, should render text once", func(t *testing.T) {
		t.Parallel()
		// Arrange
		cell := &entity.Cell{X: 0, Y: 0, Width: 100, Height: 50}
		originalColor := &props.Color{Red: 0, Green: 0, Blue: 0}
		textProp := &props.Text{
			Family: fontfamily.Arial,
			Style:  fontstyle.Normal,
			Size:   10,
			Align:  align.Left,
			Top:    0,
			Left:   0,
			Right:  0,
		}

		font := mocks.NewFont(t)
		font.EXPECT().SetFont(fontfamily.Arial, fontstyle.Normal, 10.0)
		font.EXPECT().GetHeight(fontfamily.Arial, fontstyle.Normal, 10.0).Return(5.0)
		font.EXPECT().GetColor().Return(originalColor)
		font.EXPECT().SetColor(originalColor)

		pdf := mocks.NewFpdf(t)
		pdf.EXPECT().UnicodeTranslatorFromDescriptor("").Return(func(s string) string { return s })
		pdf.EXPECT().GetStringWidth("hello").Return(20.0)
		pdf.EXPECT().GetMargins().Return(0.0, 0.0, 0.0, 0.0)
		// x = cell.X + Left = 0, y = cell.Y + Top + fontHeight = 5
		// Text(x + left_margin, y + top_margin, text) = Text(0, 5, "hello")
		pdf.EXPECT().Text(0.0, 5.0, "hello")

		sut := gofpdf.NewText(pdf, mocks.NewMath(t), font)

		// Act
		sut.Add("hello", cell, textProp)
	})
	t.Run("when single line with left align and color is set, should apply color and restore original", func(t *testing.T) {
		t.Parallel()
		// Arrange
		cell := &entity.Cell{X: 0, Y: 0, Width: 100, Height: 50}
		originalColor := &props.Color{Red: 0, Green: 0, Blue: 0}
		customColor := &props.Color{Red: 200, Green: 100, Blue: 50}
		textProp := &props.Text{
			Family: fontfamily.Arial,
			Style:  fontstyle.Normal,
			Size:   10,
			Align:  align.Left,
			Color:  customColor,
		}

		font := mocks.NewFont(t)
		font.EXPECT().SetFont(fontfamily.Arial, fontstyle.Normal, 10.0)
		font.EXPECT().GetHeight(fontfamily.Arial, fontstyle.Normal, 10.0).Return(5.0)
		font.EXPECT().GetColor().Return(originalColor)
		font.EXPECT().SetColor(customColor)
		font.EXPECT().SetColor(originalColor)

		pdf := mocks.NewFpdf(t)
		pdf.EXPECT().UnicodeTranslatorFromDescriptor("").Return(func(s string) string { return s })
		pdf.EXPECT().GetStringWidth("hello").Return(20.0)
		pdf.EXPECT().GetMargins().Return(0.0, 0.0, 0.0, 0.0)
		pdf.EXPECT().Text(0.0, 5.0, "hello")

		sut := gofpdf.NewText(pdf, mocks.NewMath(t), font)

		// Act
		sut.Add("hello", cell, textProp)
	})
	t.Run("when single line with left align and hyperlink is set, should apply blue color and render link", func(t *testing.T) {
		t.Parallel()
		// Arrange
		cell := &entity.Cell{X: 0, Y: 0, Width: 100, Height: 50}
		originalColor := &props.Color{Red: 0, Green: 0, Blue: 0}
		url := "https://example.com"
		textProp := &props.Text{
			Family:    fontfamily.Arial,
			Style:     fontstyle.Normal,
			Size:      10,
			Align:     align.Left,
			Hyperlink: &url,
		}

		font := mocks.NewFont(t)
		font.EXPECT().SetFont(fontfamily.Arial, fontstyle.Normal, 10.0)
		font.EXPECT().GetHeight(fontfamily.Arial, fontstyle.Normal, 10.0).Return(5.0)
		font.EXPECT().GetColor().Return(originalColor)
		font.EXPECT().SetColor(&props.BlueColor)
		font.EXPECT().SetColor(originalColor)

		pdf := mocks.NewFpdf(t)
		pdf.EXPECT().UnicodeTranslatorFromDescriptor("").Return(func(s string) string { return s })
		pdf.EXPECT().GetStringWidth("hello").Return(20.0)
		pdf.EXPECT().GetMargins().Return(0.0, 0.0, 0.0, 0.0)
		pdf.EXPECT().Text(0.0, 5.0, "hello")
		// LinkString(x+left, y+top-fontHeight, textWidth, fontHeight, url) = LinkString(0, 5-5, 20, 5, url)
		pdf.EXPECT().LinkString(0.0, 0.0, 20.0, 5.0, url)

		sut := gofpdf.NewText(pdf, mocks.NewMath(t), font)

		// Act
		sut.Add("hello", cell, textProp)
	})
	t.Run("when single line with right align, should offset text by full remaining width", func(t *testing.T) {
		t.Parallel()
		// Arrange
		cell := &entity.Cell{X: 0, Y: 0, Width: 100, Height: 50}
		originalColor := &props.Color{Red: 0, Green: 0, Blue: 0}
		textProp := &props.Text{
			Family: fontfamily.Arial,
			Style:  fontstyle.Normal,
			Size:   10,
			Align:  align.Right,
		}

		font := mocks.NewFont(t)
		font.EXPECT().SetFont(fontfamily.Arial, fontstyle.Normal, 10.0)
		font.EXPECT().GetHeight(fontfamily.Arial, fontstyle.Normal, 10.0).Return(5.0)
		font.EXPECT().GetColor().Return(originalColor)
		font.EXPECT().SetColor(originalColor)

		pdf := mocks.NewFpdf(t)
		pdf.EXPECT().UnicodeTranslatorFromDescriptor("").Return(func(s string) string { return s })
		pdf.EXPECT().GetStringWidth("hello").Return(20.0)
		pdf.EXPECT().GetMargins().Return(0.0, 0.0, 0.0, 0.0)
		// dx = (colWidth - textWidth) / 1 = (100 - 20) / 1 = 80
		// Text(dx + xColOffset + left, yColOffset + top, text) = Text(80, 5, "hello")
		pdf.EXPECT().Text(80.0, 5.0, "hello")

		sut := gofpdf.NewText(pdf, mocks.NewMath(t), font)

		// Act
		sut.Add("hello", cell, textProp)
	})
	t.Run("when single line with center align, should offset text by half remaining width", func(t *testing.T) {
		t.Parallel()
		// Arrange
		cell := &entity.Cell{X: 0, Y: 0, Width: 100, Height: 50}
		originalColor := &props.Color{Red: 0, Green: 0, Blue: 0}
		textProp := &props.Text{
			Family: fontfamily.Arial,
			Style:  fontstyle.Normal,
			Size:   10,
			Align:  align.Center,
		}

		font := mocks.NewFont(t)
		font.EXPECT().SetFont(fontfamily.Arial, fontstyle.Normal, 10.0)
		font.EXPECT().GetHeight(fontfamily.Arial, fontstyle.Normal, 10.0).Return(5.0)
		font.EXPECT().GetColor().Return(originalColor)
		font.EXPECT().SetColor(originalColor)

		pdf := mocks.NewFpdf(t)
		pdf.EXPECT().UnicodeTranslatorFromDescriptor("").Return(func(s string) string { return s })
		pdf.EXPECT().GetStringWidth("hello").Return(20.0)
		pdf.EXPECT().GetMargins().Return(0.0, 0.0, 0.0, 0.0)
		// dx = (colWidth - textWidth) / 2 = (100 - 20) / 2 = 40
		// Text(dx + xColOffset + left, yColOffset + top, text) = Text(40, 5, "hello")
		pdf.EXPECT().Text(40.0, 5.0, "hello")

		sut := gofpdf.NewText(pdf, mocks.NewMath(t), font)

		// Act
		sut.Add("hello", cell, textProp)
	})
	t.Run("when single line with justify align, should render each word at calculated position", func(t *testing.T) {
		t.Parallel()
		// Arrange
		cell := &entity.Cell{X: 0, Y: 0, Width: 100, Height: 50}
		originalColor := &props.Color{Red: 0, Green: 0, Blue: 0}
		textProp := &props.Text{
			Family: fontfamily.Arial,
			Style:  fontstyle.Normal,
			Size:   10,
			Align:  align.Justify,
		}

		font := mocks.NewFont(t)
		font.EXPECT().SetFont(fontfamily.Arial, fontstyle.Normal, 10.0)
		font.EXPECT().GetHeight(fontfamily.Arial, fontstyle.Normal, 10.0).Return(5.0)
		font.EXPECT().GetColor().Return(originalColor)
		font.EXPECT().SetColor(originalColor)

		pdf := mocks.NewFpdf(t)
		pdf.EXPECT().UnicodeTranslatorFromDescriptor("").Return(func(s string) string { return s })
		// initial width check in Add: fits in 100
		pdf.EXPECT().GetStringWidth("hello world").Return(30.0)
		pdf.EXPECT().GetMargins().Return(0.0, 0.0, 0.0, 0.0)
		// addLine justify: textNotSpaces="helloworld", GetStringWidth("helloworld")=25
		// defaultSpaceWidth=GetStringWidth(" ")=3
		// spaceWidth=(100-25)/1=75
		// word "hello": Text(0, 5, "hello"), finishX=0+10=10, x=10+75=85
		// word "world": Text(85, 5, "world")
		pdf.EXPECT().GetStringWidth("helloworld").Return(25.0)
		pdf.EXPECT().GetStringWidth(" ").Return(3.0)
		pdf.EXPECT().GetStringWidth("hello").Return(10.0)
		pdf.EXPECT().GetStringWidth("world").Return(15.0)
		pdf.EXPECT().Text(0.0, 5.0, "hello")
		pdf.EXPECT().Text(85.0, 5.0, "world")

		sut := gofpdf.NewText(pdf, mocks.NewMath(t), font)

		// Act
		sut.Add("hello world", cell, textProp)
	})
	t.Run("when a justified hyperlink text has an empty line, should not link the empty line", func(t *testing.T) {
		t.Parallel()
		// Arrange
		url := "https://example.com"
		cell := &entity.Cell{X: 40, Y: 0, Width: 60, Height: 50}
		textProp := &props.Text{
			Family:             "custom",
			Style:              fontstyle.Normal,
			Size:               10,
			Align:              align.Justify,
			PreserveLineBreaks: true,
			Hyperlink:          &url,
		}

		font := mocks.NewFont(t)
		font.EXPECT().SetFont("custom", fontstyle.Normal, 10.0)
		font.EXPECT().GetHeight("custom", fontstyle.Normal, 10.0).Return(5.0)
		font.EXPECT().GetColor().Return(&props.BlackColor)
		font.EXPECT().SetColor(&props.BlueColor)
		font.EXPECT().SetColor(&props.BlackColor)

		pdf := mocks.NewFpdf(t)
		pdf.EXPECT().GetStringWidth("A").Return(5.0)
		pdf.EXPECT().GetStringWidth("").Return(0.0)
		pdf.EXPECT().GetStringWidth("B").Return(5.0)
		pdf.EXPECT().GetStringWidth(" ").Return(2.0)
		pdf.EXPECT().GetMargins().Return(0.0, 0.0, 0.0, 0.0)
		pdf.EXPECT().Text(40.0, 5.0, "A").Once()
		pdf.EXPECT().Text(40.0, 15.0, "B").Once()
		linkWidths := []float64{}
		pdf.EXPECT().LinkString(40.0, mock.Anything, mock.Anything, 5.0, url).
			Run(func(_, _, w, _ float64, _ string) { linkWidths = append(linkWidths, w) })

		sut := gofpdf.NewText(pdf, mocks.NewMath(t), font)

		// Act
		sut.Add("A\n\nB", cell, textProp)

		// Assert
		assert.Equal(t, []float64{5, 5}, linkWidths)
	})
	t.Run("when text exceeds cell width with empty space strategy, should split into multiple lines", func(t *testing.T) {
		t.Parallel()
		// Arrange
		cell := &entity.Cell{X: 0, Y: 0, Width: 40, Height: 100}
		originalColor := &props.Color{Red: 0, Green: 0, Blue: 0}
		textProp := &props.Text{
			Family:            fontfamily.Arial,
			Style:             fontstyle.Normal,
			Size:              10,
			Align:             align.Left,
			BreakLineStrategy: breakline.EmptySpaceStrategy,
		}

		font := mocks.NewFont(t)
		font.EXPECT().SetFont(fontfamily.Arial, fontstyle.Normal, 10.0)
		font.EXPECT().GetHeight(fontfamily.Arial, fontstyle.Normal, 10.0).Return(5.0)
		font.EXPECT().GetColor().Return(originalColor)
		font.EXPECT().SetColor(originalColor)

		pdf := mocks.NewFpdf(t)
		pdf.EXPECT().UnicodeTranslatorFromDescriptor("").Return(func(s string) string { return s })
		// full text does not fit in width=40
		pdf.EXPECT().GetStringWidth("word1 word2").Return(60.0)
		// getLinesBreakingLineFromSpace: "word1" fits, " word2" doesn't → new line
		pdf.EXPECT().GetStringWidth("word1").Return(25.0)
		pdf.EXPECT().GetStringWidth(" word2").Return(30.0)
		pdf.EXPECT().GetStringWidth("word2").Return(25.0)
		// addLine margins and height (called once per line)
		pdf.EXPECT().GetMargins().Return(0.0, 0.0, 0.0, 0.0)
		// line 0: y = 0 + fontHeight(5) + index(0)*5 = 5
		pdf.EXPECT().Text(0.0, 5.0, "word1")
		// line 1: y = 5 + 1*5 = 10
		pdf.EXPECT().Text(0.0, 10.0, "word2")

		sut := gofpdf.NewText(pdf, mocks.NewMath(t), font)

		// Act
		sut.Add("word1 word2", cell, textProp)
	})
	t.Run("when text exceeds cell width with dash strategy, should split into lines with dashes", func(t *testing.T) {
		t.Parallel()
		// Arrange
		cell := &entity.Cell{X: 0, Y: 0, Width: 8, Height: 100}
		originalColor := &props.Color{Red: 0, Green: 0, Blue: 0}
		textProp := &props.Text{
			Family:            fontfamily.Arial,
			Style:             fontstyle.Normal,
			Size:              10,
			Align:             align.Left,
			BreakLineStrategy: breakline.DashStrategy,
		}

		font := mocks.NewFont(t)
		font.EXPECT().SetFont(fontfamily.Arial, fontstyle.Normal, 10.0)
		font.EXPECT().GetHeight(fontfamily.Arial, fontstyle.Normal, 10.0).Return(5.0)
		font.EXPECT().GetColor().Return(originalColor)
		font.EXPECT().SetColor(originalColor)

		pdf := mocks.NewFpdf(t)
		pdf.EXPECT().UnicodeTranslatorFromDescriptor("").Return(func(s string) string { return s })
		// full text "ab" does not fit in width=8
		pdf.EXPECT().GetStringWidth("ab").Return(12.0)
		// getLinesBreakingLineWithDash: dashSize=2
		// 'a': 0+2 > 8-2=6? No. width("a")=5. content="a", size=5
		// 'b': 5+2=7 > 6? Yes. content="a-", lines=["a-"]. width("b")=5. content="b", size=5
		// end: lines=["a-","b"]
		pdf.EXPECT().GetStringWidth(" - ").Return(2.0)
		pdf.EXPECT().GetStringWidth("a").Return(5.0)
		pdf.EXPECT().GetStringWidth("b").Return(5.0)
		// addLine per line: lineWidth from loop
		pdf.EXPECT().GetStringWidth("a-").Return(6.0)
		pdf.EXPECT().GetMargins().Return(0.0, 0.0, 0.0, 0.0)
		// line 0: y = 0 + 5 + 0*5 = 5
		pdf.EXPECT().Text(0.0, 5.0, "a-")
		// line 1: y = 5 + 1*5 = 10
		pdf.EXPECT().Text(0.0, 10.0, "b")

		sut := gofpdf.NewText(pdf, mocks.NewMath(t), font)

		// Act
		sut.Add("ab", cell, textProp)
	})
	t.Run("when character strategy text is wider than an empty column, should not render an empty line first", func(t *testing.T) {
		t.Parallel()
		cell := &entity.Cell{X: 0, Y: 0, Width: 0, Height: 100}
		originalColor := &props.Color{Red: 0, Green: 0, Blue: 0}
		textProp := &props.Text{
			Family:            fontfamily.Arial,
			Style:             fontstyle.Normal,
			Size:              10,
			Align:             align.Left,
			BreakLineStrategy: breakline.CharacterStrategy,
		}

		font := mocks.NewFont(t)
		font.EXPECT().SetFont(fontfamily.Arial, fontstyle.Normal, 10.0)
		font.EXPECT().GetHeight(fontfamily.Arial, fontstyle.Normal, 10.0).Return(5.0)
		font.EXPECT().GetColor().Return(originalColor)
		font.EXPECT().SetColor(originalColor)

		pdf := mocks.NewFpdf(t)
		pdf.EXPECT().UnicodeTranslatorFromDescriptor("").Return(func(s string) string { return s })
		pdf.EXPECT().GetStringWidth("a").Return(5.0).Times(3)
		pdf.EXPECT().GetMargins().Return(0.0, 0.0, 0.0, 0.0)
		pdf.EXPECT().Text(0.0, 5.0, "a")

		sut := gofpdf.NewText(pdf, mocks.NewMath(t), font)

		sut.Add("a", cell, textProp)
	})
	t.Run("when character strategy wraps cp1252 translated text, should keep the translated bytes", func(t *testing.T) {
		t.Parallel()
		cell := &entity.Cell{X: 0, Y: 0, Width: 6, Height: 100}
		originalColor := &props.Color{Red: 0, Green: 0, Blue: 0}
		textProp := &props.Text{
			Family:            fontfamily.Arial,
			Style:             fontstyle.Normal,
			Size:              10,
			Align:             align.Left,
			BreakLineStrategy: breakline.CharacterStrategy,
		}

		font := mocks.NewFont(t)
		font.EXPECT().SetFont(fontfamily.Arial, fontstyle.Normal, 10.0)
		font.EXPECT().GetHeight(fontfamily.Arial, fontstyle.Normal, 10.0).Return(5.0)
		font.EXPECT().GetColor().Return(originalColor)
		font.EXPECT().SetColor(originalColor)

		pdf := mocks.NewFpdf(t)
		pdf.EXPECT().UnicodeTranslatorFromDescriptor("").Return(func(string) string { return "\xe9\xe9" })
		pdf.EXPECT().GetStringWidth("\xe9\xe9").Return(8.0)
		pdf.EXPECT().GetStringWidth("\xe9").Return(4.0).Times(4)
		pdf.EXPECT().GetMargins().Return(0.0, 0.0, 0.0, 0.0)
		pdf.EXPECT().Text(0.0, 5.0, "\xe9")
		pdf.EXPECT().Text(0.0, 10.0, "\xe9")

		sut := gofpdf.NewText(pdf, mocks.NewMath(t), font)

		sut.Add("éé", cell, textProp)
	})
	t.Run("when text contains explicit empty lines, should preserve paragraph gaps", func(t *testing.T) {
		t.Parallel()
		// Arrange
		cell := &entity.Cell{X: 0, Y: 0, Width: 100, Height: 100}
		originalColor := &props.Color{Red: 0, Green: 0, Blue: 0}
		textProp := &props.Text{
			Family:             fontfamily.Arial,
			Style:              fontstyle.Normal,
			Size:               10,
			Align:              align.Left,
			PreserveLineBreaks: true,
		}

		font := mocks.NewFont(t)
		font.EXPECT().SetFont(fontfamily.Arial, fontstyle.Normal, 10.0)
		font.EXPECT().GetHeight(fontfamily.Arial, fontstyle.Normal, 10.0).Return(5.0)
		font.EXPECT().GetColor().Return(originalColor)
		font.EXPECT().SetColor(originalColor)

		pdf := mocks.NewFpdf(t)
		pdf.EXPECT().UnicodeTranslatorFromDescriptor("").Return(func(s string) string { return s })
		pdf.EXPECT().GetStringWidth("first").Return(20.0)
		pdf.EXPECT().GetStringWidth("").Return(0.0)
		pdf.EXPECT().GetStringWidth("third").Return(20.0)
		pdf.EXPECT().GetMargins().Return(0.0, 0.0, 0.0, 0.0)
		pdf.EXPECT().Text(0.0, 5.0, "first")
		pdf.EXPECT().GetMargins().Return(0.0, 0.0, 0.0, 0.0)
		pdf.EXPECT().Text(0.0, 10.0, "")
		pdf.EXPECT().GetMargins().Return(0.0, 0.0, 0.0, 0.0)
		pdf.EXPECT().Text(0.0, 15.0, "third")

		sut := gofpdf.NewText(pdf, mocks.NewMath(t), font)

		// Act
		sut.Add("first\n\nthird", cell, textProp)
	})
	t.Run("when line breaks are preserved, should render LF, CRLF and CR as one break each", func(t *testing.T) {
		t.Parallel()
		for _, value := range []string{"A\nB", "A\r\nB", "A\rB"} {
			// Arrange
			cell := &entity.Cell{X: 0, Y: 0, Width: 100, Height: 100}
			textProp := &props.Text{Family: "custom", Style: fontstyle.Normal, Size: 10, Align: align.Left, PreserveLineBreaks: true}

			font := mocks.NewFont(t)
			font.EXPECT().SetFont("custom", fontstyle.Normal, 10.0)
			font.EXPECT().GetHeight("custom", fontstyle.Normal, 10.0).Return(5.0)
			font.EXPECT().GetColor().Return(&props.BlackColor)
			font.EXPECT().SetColor(&props.BlackColor)

			pdf := mocks.NewFpdf(t)
			pdf.EXPECT().GetStringWidth("A").Return(5.0)
			pdf.EXPECT().GetStringWidth("B").Return(5.0)
			pdf.EXPECT().GetMargins().Return(0.0, 0.0, 0.0, 0.0)
			pdf.EXPECT().Text(0.0, 5.0, "A").Once()
			pdf.EXPECT().Text(0.0, 10.0, "B").Once()

			sut := gofpdf.NewText(pdf, mocks.NewMath(t), font)

			// Act
			sut.Add(value, cell, textProp)
		}
	})
	t.Run("when line breaks are not preserved, should render LF, CRLF and CR as a space", func(t *testing.T) {
		t.Parallel()
		for _, value := range []string{"A\nB", "A\r\nB", "A\rB"} {
			// Arrange
			cell := &entity.Cell{X: 0, Y: 0, Width: 100, Height: 100}
			textProp := &props.Text{Family: "custom", Style: fontstyle.Normal, Size: 10, Align: align.Left}

			font := mocks.NewFont(t)
			font.EXPECT().SetFont("custom", fontstyle.Normal, 10.0)
			font.EXPECT().GetHeight("custom", fontstyle.Normal, 10.0).Return(5.0)
			font.EXPECT().GetColor().Return(&props.BlackColor)
			font.EXPECT().SetColor(&props.BlackColor)

			pdf := mocks.NewFpdf(t)
			pdf.EXPECT().GetStringWidth("A B").Return(12.0)
			pdf.EXPECT().GetMargins().Return(0.0, 0.0, 0.0, 0.0)
			pdf.EXPECT().Text(0.0, 5.0, "A B").Once()

			sut := gofpdf.NewText(pdf, mocks.NewMath(t), font)

			// Act
			sut.Add(value, cell, textProp)
		}
	})
	t.Run("when top exceeds cell height, should clamp top to cell height", func(t *testing.T) {
		t.Parallel()
		// Arrange
		cell := &entity.Cell{X: 0, Y: 0, Width: 100, Height: 50}
		originalColor := &props.Color{Red: 0, Green: 0, Blue: 0}
		textProp := &props.Text{
			Family: fontfamily.Arial,
			Style:  fontstyle.Normal,
			Size:   10,
			Align:  align.Left,
			Top:    100, // exceeds cell.Height=50, gets clamped to 50
		}

		font := mocks.NewFont(t)
		font.EXPECT().SetFont(fontfamily.Arial, fontstyle.Normal, 10.0)
		font.EXPECT().GetHeight(fontfamily.Arial, fontstyle.Normal, 10.0).Return(5.0)
		font.EXPECT().GetColor().Return(originalColor)
		font.EXPECT().SetColor(originalColor)

		pdf := mocks.NewFpdf(t)
		pdf.EXPECT().UnicodeTranslatorFromDescriptor("").Return(func(s string) string { return s })
		pdf.EXPECT().GetStringWidth("hi").Return(5.0)
		pdf.EXPECT().GetMargins().Return(0.0, 0.0, 0.0, 0.0)
		// top clamped to 50: y = cell.Y + 50 + fontHeight = 0 + 50 + 5 = 55
		pdf.EXPECT().Text(0.0, 55.0, "hi")

		sut := gofpdf.NewText(pdf, mocks.NewMath(t), font)

		// Act
		sut.Add("hi", cell, textProp)
	})
	t.Run("when left exceeds cell width, should clamp left to cell width", func(t *testing.T) {
		t.Parallel()
		// Arrange
		cell := &entity.Cell{X: 0, Y: 0, Width: 100, Height: 50}
		originalColor := &props.Color{Red: 0, Green: 0, Blue: 0}
		textProp := &props.Text{
			Family: fontfamily.Arial,
			Style:  fontstyle.Normal,
			Size:   10,
			Align:  align.Left,
			Left:   150, // exceeds cell.Width=100, gets clamped to 100
		}

		font := mocks.NewFont(t)
		font.EXPECT().SetFont(fontfamily.Arial, fontstyle.Normal, 10.0)
		font.EXPECT().GetHeight(fontfamily.Arial, fontstyle.Normal, 10.0).Return(5.0)
		font.EXPECT().GetColor().Return(originalColor)
		font.EXPECT().SetColor(originalColor)

		pdf := mocks.NewFpdf(t)
		pdf.EXPECT().UnicodeTranslatorFromDescriptor("").Return(func(s string) string { return s })
		// left clamped to 100, right=0 → width = 100-100-0 = 0
		// x = 0 + 100 = 100
		// GetStringWidth("") = 0 ≤ 0 → single line
		pdf.EXPECT().GetStringWidth("").Return(0.0)
		pdf.EXPECT().GetMargins().Return(0.0, 0.0, 0.0, 0.0)
		// Text(x + left_margin, y + top_margin) = Text(100, 5, "")
		pdf.EXPECT().Text(100.0, 5.0, "")

		sut := gofpdf.NewText(pdf, mocks.NewMath(t), font)

		// Act
		sut.Add("", cell, textProp)
	})
	t.Run("when right exceeds cell width, should clamp right to cell width", func(t *testing.T) {
		t.Parallel()
		// Arrange
		cell := &entity.Cell{X: 0, Y: 0, Width: 100, Height: 50}
		originalColor := &props.Color{Red: 0, Green: 0, Blue: 0}
		textProp := &props.Text{
			Family: fontfamily.Arial,
			Style:  fontstyle.Normal,
			Size:   10,
			Align:  align.Left,
			Right:  150, // exceeds cell.Width=100, gets clamped to 100
		}

		font := mocks.NewFont(t)
		font.EXPECT().SetFont(fontfamily.Arial, fontstyle.Normal, 10.0)
		font.EXPECT().GetHeight(fontfamily.Arial, fontstyle.Normal, 10.0).Return(5.0)
		font.EXPECT().GetColor().Return(originalColor)
		font.EXPECT().SetColor(originalColor)

		pdf := mocks.NewFpdf(t)
		pdf.EXPECT().UnicodeTranslatorFromDescriptor("").Return(func(s string) string { return s })
		// right clamped to 100, left=0 → width = 100-0-100 = 0
		// x = 0 + 0 = 0
		// GetStringWidth("") = 0 ≤ 0 → single line
		pdf.EXPECT().GetStringWidth("").Return(0.0)
		pdf.EXPECT().GetMargins().Return(0.0, 0.0, 0.0, 0.0)
		pdf.EXPECT().Text(0.0, 5.0, "")

		sut := gofpdf.NewText(pdf, mocks.NewMath(t), font)

		// Act
		sut.Add("", cell, textProp)
	})
}
