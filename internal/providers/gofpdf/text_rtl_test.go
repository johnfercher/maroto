package gofpdf_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf"
	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/breakline"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/stretchr/testify/assert"
)

// The strings these tests expect the writer to receive hold Arabic
// Presentation Forms-B characters in visual order, so they read as garbled
// Arabic in an editor. That is the point: it is what has to reach Fpdf.Text.
//
// Every width the pipeline asks for has to be registered on the mock, so a
// call made with an unregistered string fails the test. That is what proves
// the widths are measured on the shaped text and not on the logical one.
const (
	// "مرحبا", shaped and reversed.
	marhabaShaped = "ﺎﺒﺣﺮﻣ"
	// "جدا", shaped and reversed.
	jiddanShaped = "ﺍﺪﺟ"
	// "لا", contracted into the single lam-alef ligature glyph.
	laShaped = "ﻻ"
)

func TestText_Add_RTL(t *testing.T) {
	t.Parallel()

	t.Run("when rtl is enabled, should shape and reorder the text before writing it", func(t *testing.T) {
		t.Parallel()
		// Arrange
		cell := &entity.Cell{X: 0, Y: 0, Width: 100, Height: 50}
		originalColor := &props.Color{Red: 0, Green: 0, Blue: 0}
		textProp := &props.Text{
			Family: fontfamily.Arial,
			Style:  fontstyle.Normal,
			Size:   10,
			Align:  align.Left,
			RTL:    true,
		}

		font := mocks.NewFont(t)
		font.EXPECT().SetFont(fontfamily.Arial, fontstyle.Normal, 10.0)
		font.EXPECT().GetHeight(fontfamily.Arial, fontstyle.Normal, 10.0).Return(5.0)
		font.EXPECT().GetColor().Return(originalColor)
		font.EXPECT().SetColor(originalColor)

		pdf := mocks.NewFpdf(t)
		pdf.EXPECT().UnicodeTranslatorFromDescriptor("").Return(func(s string) string { return s })
		// The width is asked for on the shaped form, never on the logical one.
		pdf.EXPECT().GetStringWidth(marhabaShaped).Return(20.0)
		pdf.EXPECT().GetMargins().Return(0.0, 0.0, 0.0, 0.0)
		pdf.EXPECT().Text(0.0, 5.0, marhabaShaped)

		sut := gofpdf.NewText(pdf, mocks.NewMath(t), font)

		// Act
		sut.Add("مرحبا", cell, textProp)
	})

	t.Run("when rtl is disabled, should write the arabic text untouched", func(t *testing.T) {
		t.Parallel()
		// Arrange
		cell := &entity.Cell{X: 0, Y: 0, Width: 100, Height: 50}
		originalColor := &props.Color{Red: 0, Green: 0, Blue: 0}
		textProp := &props.Text{
			Family: fontfamily.Arial,
			Style:  fontstyle.Normal,
			Size:   10,
			Align:  align.Left,
		}

		font := mocks.NewFont(t)
		font.EXPECT().SetFont(fontfamily.Arial, fontstyle.Normal, 10.0)
		font.EXPECT().GetHeight(fontfamily.Arial, fontstyle.Normal, 10.0).Return(5.0)
		font.EXPECT().GetColor().Return(originalColor)
		font.EXPECT().SetColor(originalColor)

		pdf := mocks.NewFpdf(t)
		pdf.EXPECT().UnicodeTranslatorFromDescriptor("").Return(func(s string) string { return s })
		// The flag is opt-in, so the behaviour existing users get is unchanged.
		pdf.EXPECT().GetStringWidth("مرحبا").Return(20.0)
		pdf.EXPECT().GetMargins().Return(0.0, 0.0, 0.0, 0.0)
		pdf.EXPECT().Text(0.0, 5.0, "مرحبا")

		sut := gofpdf.NewText(pdf, mocks.NewMath(t), font)

		// Act
		sut.Add("مرحبا", cell, textProp)
	})

	t.Run("when rtl is enabled and the text has no arabic, should write it untouched", func(t *testing.T) {
		t.Parallel()
		// Arrange
		cell := &entity.Cell{X: 0, Y: 0, Width: 100, Height: 50}
		originalColor := &props.Color{Red: 0, Green: 0, Blue: 0}
		textProp := &props.Text{
			Family: fontfamily.Arial,
			Style:  fontstyle.Normal,
			Size:   10,
			Align:  align.Left,
			RTL:    true,
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
		pdf.EXPECT().Text(0.0, 5.0, "hello")

		sut := gofpdf.NewText(pdf, mocks.NewMath(t), font)

		// Act
		sut.Add("hello", cell, textProp)
	})

	t.Run("when a ligature makes the text narrower, should fit the line using the shaped width", func(t *testing.T) {
		t.Parallel()
		// Arrange
		// "لا لا" is five characters logically but only three once each lam-alef
		// pair contracts into its ligature glyph. Measuring the logical text
		// would overestimate the line and wrap it; measuring the shaped text
		// keeps it on one line.
		cell := &entity.Cell{X: 0, Y: 0, Width: 10, Height: 50}
		originalColor := &props.Color{Red: 0, Green: 0, Blue: 0}
		textProp := &props.Text{
			Family: fontfamily.Arial,
			Style:  fontstyle.Normal,
			Size:   10,
			Align:  align.Left,
			RTL:    true,
		}

		font := mocks.NewFont(t)
		font.EXPECT().SetFont(fontfamily.Arial, fontstyle.Normal, 10.0)
		font.EXPECT().GetHeight(fontfamily.Arial, fontstyle.Normal, 10.0).Return(5.0)
		font.EXPECT().GetColor().Return(originalColor)
		font.EXPECT().SetColor(originalColor)

		pdf := mocks.NewFpdf(t)
		pdf.EXPECT().UnicodeTranslatorFromDescriptor("").Return(func(s string) string { return s })
		// Only the shaped form is registered: measuring "لا لا" would fail here.
		pdf.EXPECT().GetStringWidth(laShaped + " " + laShaped).Return(8.0)
		pdf.EXPECT().GetMargins().Return(0.0, 0.0, 0.0, 0.0)
		pdf.EXPECT().Text(0.0, 5.0, laShaped+" "+laShaped)

		sut := gofpdf.NewText(pdf, mocks.NewMath(t), font)

		// Act
		sut.Add("لا لا", cell, textProp)
	})

	t.Run("when the text wraps, should break on the logical text and give each line its own base direction", func(t *testing.T) {
		t.Parallel()
		// Arrange
		// The whole text starts with Arabic, but the second line starts with a
		// Latin word. Processing the paragraph as a unit would give both lines
		// the same base direction and reverse their order; processing each
		// emitted line on its own gives line one a right to left base and line
		// two a left to right one.
		cell := &entity.Cell{X: 0, Y: 0, Width: 40, Height: 100}
		originalColor := &props.Color{Red: 0, Green: 0, Blue: 0}
		textProp := &props.Text{
			Family:            fontfamily.Arial,
			Style:             fontstyle.Normal,
			Size:              10,
			Align:             align.Left,
			BreakLineStrategy: breakline.EmptySpaceStrategy,
			RTL:               true,
		}

		font := mocks.NewFont(t)
		font.EXPECT().SetFont(fontfamily.Arial, fontstyle.Normal, 10.0)
		font.EXPECT().GetHeight(fontfamily.Arial, fontstyle.Normal, 10.0).Return(5.0)
		font.EXPECT().GetColor().Return(originalColor)
		font.EXPECT().SetColor(originalColor)

		pdf := mocks.NewFpdf(t)
		pdf.EXPECT().UnicodeTranslatorFromDescriptor("").Return(func(s string) string { return s })
		// The whole text does not fit in width 40.
		pdf.EXPECT().GetStringWidth(jiddanShaped + " Hello Wide " + marhabaShaped).Return(70.0)
		// Line breaking, word by word, each candidate measured shaped:
		// "مرحبا" 20 fits, " Hello" 15 fits (35), " Wide" 15 does not (50 > 40).
		pdf.EXPECT().GetStringWidth(marhabaShaped).Return(20.0)
		pdf.EXPECT().GetStringWidth(" Hello").Return(15.0)
		pdf.EXPECT().GetStringWidth(" Wide").Return(15.0)
		// Second line restarts from "Wide" 12, then " جدا" 12 fits (24).
		pdf.EXPECT().GetStringWidth("Wide").Return(12.0)
		pdf.EXPECT().GetStringWidth(jiddanShaped + " ").Return(12.0)
		// Each emitted line is processed on its own before being measured again.
		pdf.EXPECT().GetStringWidth("Hello " + marhabaShaped).Return(35.0)
		pdf.EXPECT().GetStringWidth("Wide " + jiddanShaped).Return(24.0)
		pdf.EXPECT().GetMargins().Return(0.0, 0.0, 0.0, 0.0)
		// Line 0 has a right to left base: the Latin run goes to the left.
		pdf.EXPECT().Text(0.0, 5.0, "Hello "+marhabaShaped)
		// Line 1 has a left to right base: the run order is kept.
		pdf.EXPECT().Text(0.0, 10.0, "Wide "+jiddanShaped)

		sut := gofpdf.NewText(pdf, mocks.NewMath(t), font)

		// Act
		sut.Add("مرحبا Hello Wide جدا", cell, textProp)
	})

	t.Run("when the text is justified, should write the words in visual order at processed widths", func(t *testing.T) {
		t.Parallel()
		// Arrange
		// "مرحبا جدا" is logically مرحبا then جدا, so visually جدا comes first.
		// The justified output has to draw جدا at the smallest x.
		cell := &entity.Cell{X: 0, Y: 0, Width: 100, Height: 50}
		originalColor := &props.Color{Red: 0, Green: 0, Blue: 0}
		textProp := &props.Text{
			Family: fontfamily.Arial,
			Style:  fontstyle.Normal,
			Size:   10,
			Align:  align.Justify,
			RTL:    true,
		}

		font := mocks.NewFont(t)
		font.EXPECT().SetFont(fontfamily.Arial, fontstyle.Normal, 10.0)
		font.EXPECT().GetHeight(fontfamily.Arial, fontstyle.Normal, 10.0).Return(5.0)
		font.EXPECT().GetColor().Return(originalColor)
		font.EXPECT().SetColor(originalColor)

		pdf := mocks.NewFpdf(t)
		pdf.EXPECT().UnicodeTranslatorFromDescriptor("").Return(func(s string) string { return s })
		pdf.EXPECT().GetStringWidth(jiddanShaped + " " + marhabaShaped).Return(30.0)
		pdf.EXPECT().GetMargins().Return(0.0, 0.0, 0.0, 0.0)
		// Justify measures the processed line without its spaces, then spreads
		// the leftover: spaceWidth = (100 - 25) / 1 = 75.
		pdf.EXPECT().GetStringWidth(jiddanShaped + marhabaShaped).Return(25.0)
		pdf.EXPECT().GetStringWidth(" ").Return(3.0)
		pdf.EXPECT().GetStringWidth(jiddanShaped).Return(10.0)
		pdf.EXPECT().GetStringWidth(marhabaShaped).Return(15.0)
		// The logically last word is drawn first, at x = 0.
		pdf.EXPECT().Text(0.0, 5.0, jiddanShaped)
		// x = 0 + 10 + 75 = 85
		pdf.EXPECT().Text(85.0, 5.0, marhabaShaped)

		sut := gofpdf.NewText(pdf, mocks.NewMath(t), font)

		// Act
		sut.Add("مرحبا جدا", cell, textProp)
	})

	t.Run("when the text is right aligned, should offset each line by its own processed width", func(t *testing.T) {
		t.Parallel()
		// Arrange
		cell := &entity.Cell{X: 0, Y: 0, Width: 40, Height: 100}
		originalColor := &props.Color{Red: 0, Green: 0, Blue: 0}
		textProp := &props.Text{
			Family:            fontfamily.Arial,
			Style:             fontstyle.Normal,
			Size:              10,
			Align:             align.Right,
			BreakLineStrategy: breakline.EmptySpaceStrategy,
			RTL:               true,
		}

		font := mocks.NewFont(t)
		font.EXPECT().SetFont(fontfamily.Arial, fontstyle.Normal, 10.0)
		font.EXPECT().GetHeight(fontfamily.Arial, fontstyle.Normal, 10.0).Return(5.0)
		font.EXPECT().GetColor().Return(originalColor)
		font.EXPECT().SetColor(originalColor)

		pdf := mocks.NewFpdf(t)
		pdf.EXPECT().UnicodeTranslatorFromDescriptor("").Return(func(s string) string { return s })
		pdf.EXPECT().GetStringWidth(jiddanShaped + " Hello Wide " + marhabaShaped).Return(70.0)
		pdf.EXPECT().GetStringWidth(marhabaShaped).Return(20.0)
		pdf.EXPECT().GetStringWidth(" Hello").Return(15.0)
		pdf.EXPECT().GetStringWidth(" Wide").Return(15.0)
		pdf.EXPECT().GetStringWidth("Wide").Return(12.0)
		pdf.EXPECT().GetStringWidth(jiddanShaped + " ").Return(12.0)
		pdf.EXPECT().GetStringWidth("Hello " + marhabaShaped).Return(35.0)
		pdf.EXPECT().GetStringWidth("Wide " + jiddanShaped).Return(24.0)
		pdf.EXPECT().GetMargins().Return(0.0, 0.0, 0.0, 0.0)
		// dx comes from the processed width of each line: 40 - 35 = 5
		pdf.EXPECT().Text(5.0, 5.0, "Hello "+marhabaShaped)
		// 40 - 24 = 16
		pdf.EXPECT().Text(16.0, 10.0, "Wide "+jiddanShaped)

		sut := gofpdf.NewText(pdf, mocks.NewMath(t), font)

		// Act
		sut.Add("مرحبا Hello Wide جدا", cell, textProp)
	})
}

func TestText_GetLinesQuantity_RTL(t *testing.T) {
	t.Parallel()

	t.Run("when rtl is enabled, should count the lines using the shaped widths", func(t *testing.T) {
		t.Parallel()
		// Arrange
		// Measured logically "لا لا" would need more room than the column has;
		// measured shaped both ligatures fit on a single line.
		textProp := &props.Text{RTL: true}
		textProp.MakeValid(&props.Font{Family: fontfamily.Arial, Size: 10, Style: fontstyle.Normal})

		font := mocks.NewFont(t)
		font.EXPECT().SetFont(textProp.Family, textProp.Style, textProp.Size)

		pdf := mocks.NewFpdf(t)
		pdf.EXPECT().UnicodeTranslatorFromDescriptor("").Return(func(s string) string { return s })
		pdf.EXPECT().GetStringWidth(laShaped).Return(4.0)
		pdf.EXPECT().GetStringWidth(laShaped + " ").Return(5.0)

		sut := gofpdf.NewText(pdf, mocks.NewMath(t), font)

		// Act
		quantity := sut.GetLinesQuantity("لا لا", textProp, 10)

		// Assert
		assert.Equal(t, 1, quantity)
	})
}
