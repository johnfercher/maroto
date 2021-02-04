package internal_test

import (
	"fmt"
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/internal/mocks"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewText(t *testing.T) {
	text := internal.NewText(&mocks.Fpdf{}, &mocks.Math{}, &mocks.Font{})

	assert.NotNil(t, text)
	assert.Equal(t, fmt.Sprintf("%T", text), "*internal.text")
}

func TestText_GetLinesQuantity_WhenStringSmallerThanLimits(t *testing.T) {
	// Arrange
	pdf := &mocks.Fpdf{}
	pdf.On("UnicodeTranslatorFromDescriptor", mock.Anything).Return(func(text string) string {
		return text
	})
	pdf.On("GetStringWidth", mock.Anything).Return(8.0)

	font := &mocks.Font{}
	font.On("SetFont", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	sut := internal.NewText(pdf, nil, font)

	// Act
	lines := sut.GetLinesQuantity("AnyText With Spaces", props.Text{}, 2)

	// Assert
	assert.Equal(t, lines, 4)
}

func TestText_GetLinesQuantity_WhenHasOneWord(t *testing.T) {
	// Arrange
	pdf := &mocks.Fpdf{}
	pdf.On("UnicodeTranslatorFromDescriptor", mock.Anything).Return(func(text string) string {
		return text
	})
	pdf.On("GetStringWidth", mock.Anything).Return(15.0)

	font := &mocks.Font{}
	font.On("SetFont", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	sut := internal.NewText(pdf, nil, font)

	// Act
	lines := sut.GetLinesQuantity("OneWord", props.Text{}, 2)

	// Assert
	assert.Equal(t, lines, 1)
}

func TestText_GetLinesQuantity_WhenExtrapolate(t *testing.T) {
	// Arrange
	pdf := &mocks.Fpdf{}
	pdf.On("UnicodeTranslatorFromDescriptor", mock.Anything).Return(func(text string) string {
		return text
	})
	pdf.On("GetStringWidth", mock.Anything).Return(15.0)

	font := &mocks.Font{}
	font.On("SetFont", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	sut := internal.NewText(pdf, nil, font)

	// Act
	lines := sut.GetLinesQuantity("Many words", props.Text{Extrapolate: true}, 2)

	// Assert
	assert.Equal(t, lines, 1)
}

func TestText_GetLinesQuantity_WhenHasToBreakLines(t *testing.T) {
	// Arrange
	pdf := &mocks.Fpdf{}
	pdf.On("UnicodeTranslatorFromDescriptor", mock.Anything).Return(func(text string) string {
		return text
	})
	pdf.On("GetStringWidth", mock.Anything).Return(15.0)

	math := &mocks.Math{}
	math.On("GetWidthPerCol", mock.Anything).Return(10.0)

	font := &mocks.Font{}
	font.On("SetFont", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	sut := internal.NewText(pdf, math, font)

	// Act
	lines := sut.GetLinesQuantity("Many words", props.Text{}, 2)

	// Assert
	assert.Equal(t, lines, 3)
}

func TestText_Add(t *testing.T) {
	cases := []struct {
		name       string
		text       string
		align      consts.Align
		family     string
		color      color.Color
		pdf        func() *mocks.Fpdf
		font       func() *mocks.Font
		cell       func() *internal.Cell
		assertPdf  func(t *testing.T, pdf *mocks.Fpdf)
		assertFont func(t *testing.T, font *mocks.Font)
	}{
		{
			"Left Align",
			"TextHelper1",
			consts.Left,
			consts.Arial,
			color.Color{Red: 0, Green: 0, Blue: 0},
			func() *mocks.Fpdf {
				_pdf := &mocks.Fpdf{}
				_pdf.On("GetStringWidth", mock.Anything).Return(12.0)
				_pdf.On("GetMargins").Return(10.0, 10.0, 10.0, 10.0)
				_pdf.On("Text", mock.Anything, mock.Anything, mock.Anything)
				_pdf.On("UnicodeTranslatorFromDescriptor", mock.Anything).Return(func(value string) string { return value })
				return _pdf
			},
			func() *mocks.Font {
				_font := &mocks.Font{}
				_font.On("GetScaleFactor").Return(1.0)
				_font.On("GetFont").Return(consts.Arial, consts.Bold, 1.0)
				_font.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				_font.On("GetColor").Return(color.Color{Red: 0, Green: 0, Blue: 0})
				_font.On("SetColor", mock.Anything)
				return _font
			},
			func() *internal.Cell {
				return nil
			},
			func(t *testing.T, _pdf *mocks.Fpdf) {
				_pdf.AssertNotCalled(t, "GetStringWidth")

				_pdf.AssertNumberOfCalls(t, "GetMargins", 1)

				_pdf.AssertNumberOfCalls(t, "Text", 1)
				_pdf.AssertCalled(t, "Text", 11.0, 16.0, "TextHelper1")
			},
			func(t *testing.T, _font *mocks.Font) {
				_font.AssertNumberOfCalls(t, "SetFont", 1)
				_font.AssertCalled(t, "SetFont", consts.Arial, consts.BoldItalic, 16.0)
				_font.AssertNumberOfCalls(t, "GetColor", 1)
				_font.AssertNumberOfCalls(t, "SetColor", 2)
				_font.AssertCalled(t, "SetColor", color.Color{Red: 0, Green: 0, Blue: 0})
			},
		},
		{
			"Custom Font",
			"TextHelper1",
			consts.Left,
			"CustomFont",
			color.Color{Red: 0, Green: 0, Blue: 0},
			func() *mocks.Fpdf {
				_pdf := &mocks.Fpdf{}
				_pdf.On("GetStringWidth", mock.Anything).Return(12.0)
				_pdf.On("GetMargins").Return(10.0, 10.0, 10.0, 10.0)
				_pdf.On("Text", mock.Anything, mock.Anything, mock.Anything)
				_pdf.On("UnicodeTranslatorFromDescriptor", mock.Anything).Return(func(value string) string { return value })
				return _pdf
			},
			func() *mocks.Font {
				_font := &mocks.Font{}
				_font.On("GetScaleFactor").Return(1.0)
				_font.On("GetFont").Return(consts.Arial, consts.Bold, 1.0)
				_font.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				_font.On("GetColor").Return(color.Color{Red: 0, Green: 0, Blue: 0})
				_font.On("SetColor", mock.Anything)
				return _font
			},
			func() *internal.Cell {
				return nil
			},
			func(t *testing.T, _pdf *mocks.Fpdf) {
				_pdf.AssertNotCalled(t, "GetStringWidth")

				_pdf.AssertNumberOfCalls(t, "GetMargins", 1)

				_pdf.AssertNumberOfCalls(t, "Text", 1)
				_pdf.AssertCalled(t, "Text", 11.0, 16.0, "TextHelper1")
			},
			func(t *testing.T, _font *mocks.Font) {
				_font.AssertNumberOfCalls(t, "SetFont", 1)
				_font.AssertCalled(t, "SetFont", "CustomFont", consts.BoldItalic, 16.0)
				_font.AssertNumberOfCalls(t, "GetColor", 1)
				_font.AssertNumberOfCalls(t, "SetColor", 2)
				_font.AssertCalled(t, "SetColor", color.Color{Red: 0, Green: 0, Blue: 0})
			},
		},
		{
			"Center Align",
			"TextHelper2",
			consts.Center,
			consts.Arial,
			color.Color{Red: 0, Green: 0, Blue: 0},
			func() *mocks.Fpdf {
				_pdf := &mocks.Fpdf{}
				_pdf.On("GetStringWidth", mock.Anything).Return(12.0)
				_pdf.On("GetMargins").Return(10.0, 10.0, 10.0, 10.0)
				_pdf.On("Text", mock.Anything, mock.Anything, mock.Anything)
				_pdf.On("UnicodeTranslatorFromDescriptor", mock.Anything).Return(func(value string) string { return value })
				return _pdf
			},
			func() *mocks.Font {
				_font := &mocks.Font{}
				_font.On("GetScaleFactor").Return(1.0)
				_font.On("GetFont").Return(consts.Arial, consts.Bold, 1.0)
				_font.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				_font.On("GetColor").Return(color.Color{Red: 0, Green: 0, Blue: 0})
				_font.On("SetColor", mock.Anything)
				return _font
			},
			func() *internal.Cell {
				return nil
			},
			func(t *testing.T, _pdf *mocks.Fpdf) {
				_pdf.AssertNumberOfCalls(t, "GetStringWidth", 1)
				_pdf.AssertCalled(t, "GetStringWidth", "TextHelper2")

				_pdf.AssertNumberOfCalls(t, "GetMargins", 1)

				_pdf.AssertNumberOfCalls(t, "Text", 1)
				_pdf.AssertCalled(t, "Text", 12.5, 16.0, "TextHelper2")
			},
			func(t *testing.T, _font *mocks.Font) {
				_font.AssertNumberOfCalls(t, "SetFont", 1)
				_font.AssertCalled(t, "SetFont", consts.Arial, consts.BoldItalic, 16.0)
				_font.AssertNumberOfCalls(t, "GetColor", 1)
				_font.AssertNumberOfCalls(t, "SetColor", 2)
				_font.AssertCalled(t, "SetColor", color.Color{Red: 0, Green: 0, Blue: 0})
			},
		},
		{
			"Right Align",
			"TextHelper3",
			consts.Right,
			consts.Arial,
			color.Color{Red: 0, Green: 0, Blue: 0},
			func() *mocks.Fpdf {
				_pdf := &mocks.Fpdf{}
				_pdf.On("GetStringWidth", mock.Anything).Return(12.0)
				_pdf.On("GetMargins").Return(10.0, 10.0, 10.0, 10.0)
				_pdf.On("Text", mock.Anything, mock.Anything, mock.Anything)
				_pdf.On("UnicodeTranslatorFromDescriptor", mock.Anything).Return(func(value string) string { return value })
				return _pdf
			},
			func() *mocks.Font {
				_font := &mocks.Font{}
				_font.On("GetScaleFactor").Return(1.0)
				_font.On("GetFont").Return(consts.Arial, consts.Bold, 1.0)
				_font.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				_font.On("GetColor").Return(color.Color{Red: 0, Green: 0, Blue: 0})
				_font.On("SetColor", mock.Anything)
				return _font
			}, func() *internal.Cell {
				return nil
			},
			func(t *testing.T, _pdf *mocks.Fpdf) {
				_pdf.AssertNumberOfCalls(t, "GetStringWidth", 1)
				_pdf.AssertCalled(t, "GetStringWidth", "TextHelper3")

				_pdf.AssertNumberOfCalls(t, "GetMargins", 1)

				_pdf.AssertNumberOfCalls(t, "Text", 1)
				_pdf.AssertCalled(t, "Text", 14.0, 16.0, "TextHelper3")
			},
			func(t *testing.T, _font *mocks.Font) {
				_font.AssertNumberOfCalls(t, "SetFont", 1)
				_font.AssertCalled(t, "SetFont", consts.Arial, consts.BoldItalic, 16.0)
				_font.AssertNumberOfCalls(t, "GetColor", 1)
				_font.AssertNumberOfCalls(t, "SetColor", 2)
				_font.AssertCalled(t, "SetColor", color.Color{Red: 0, Green: 0, Blue: 0})
			},
		},
		{
			"Right Align",
			"TextHelper4",
			consts.Right,
			consts.Arial,
			color.Color{Red: 0, Green: 0, Blue: 0},
			func() *mocks.Fpdf {
				_pdf := &mocks.Fpdf{}
				_pdf.On("GetStringWidth", mock.Anything).Return(12.0)
				_pdf.On("GetMargins").Return(10.0, 10.0, 10.0, 10.0)
				_pdf.On("Text", mock.Anything, mock.Anything, mock.Anything)
				_pdf.On("UnicodeTranslatorFromDescriptor", mock.Anything).Return(func(value string) string { return value })
				return _pdf
			},
			func() *mocks.Font {
				_font := &mocks.Font{}
				_font.On("GetScaleFactor").Return(1.0)
				_font.On("GetFont").Return(consts.Arial, consts.Bold, 1.0)
				_font.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				_font.On("GetColor").Return(color.Color{Red: 0, Green: 0, Blue: 0})
				_font.On("SetColor", mock.Anything)
				return _font
			}, func() *internal.Cell {
				return nil
			},
			func(t *testing.T, _pdf *mocks.Fpdf) {
				_pdf.AssertNumberOfCalls(t, "GetStringWidth", 1)
				_pdf.AssertCalled(t, "GetStringWidth", "TextHelper4")

				_pdf.AssertNumberOfCalls(t, "GetMargins", 1)

				_pdf.AssertNumberOfCalls(t, "Text", 1)
				_pdf.AssertCalled(t, "Text", 14.0, 16.0, "TextHelper4")
			},
			func(t *testing.T, _font *mocks.Font) {
				_font.AssertNumberOfCalls(t, "SetFont", 1)
				_font.AssertCalled(t, "SetFont", consts.Arial, consts.BoldItalic, 16.0)
				_font.AssertNumberOfCalls(t, "GetColor", 1)
				_font.AssertNumberOfCalls(t, "SetColor", 2)
				_font.AssertCalled(t, "SetColor", color.Color{Red: 0, Green: 0, Blue: 0})
			},
		},
		{
			"Bigger than cell width",
			"Lorem Ipsum is simply dummy textá of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
			consts.Center,
			consts.Arial,
			color.Color{Red: 0, Green: 0, Blue: 0},
			func() *mocks.Fpdf {
				_pdf := &mocks.Fpdf{}
				_pdf.On("GetStringWidth", "Lorem Ipsum is simply dummy textá of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.").Return(900.0)
				_pdf.On("GetStringWidth", mock.Anything).Return(20.0)
				_pdf.On("GetMargins").Return(10.0, 10.0, 10.0, 10.0)
				_pdf.On("Text", mock.Anything, mock.Anything, mock.Anything)
				_pdf.On("UnicodeTranslatorFromDescriptor", mock.Anything).Return(func(value string) string { return value })
				return _pdf
			},
			func() *mocks.Font {
				_font := &mocks.Font{}
				_font.On("GetScaleFactor").Return(1.0)
				_font.On("GetFont").Return(consts.Arial, consts.Bold, 1.0)
				_font.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				_font.On("GetColor").Return(color.Color{Red: 0, Green: 0, Blue: 0})
				_font.On("SetColor", mock.Anything)
				return _font
			}, func() *internal.Cell {
				return nil
			},
			func(t *testing.T, _pdf *mocks.Fpdf) {
				_pdf.AssertNumberOfCalls(t, "GetStringWidth", 275)
				_pdf.AssertCalled(t, "GetStringWidth", "Lorem Ipsum is simply dummy textá of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.")

				_pdf.AssertNumberOfCalls(t, "GetMargins", 92)

				_pdf.AssertNumberOfCalls(t, "Text", 92)
			},
			func(t *testing.T, _font *mocks.Font) {
				_font.AssertNumberOfCalls(t, "SetFont", 1)
				_font.AssertCalled(t, "SetFont", consts.Arial, consts.BoldItalic, 16.0)
				_font.AssertNumberOfCalls(t, "GetColor", 1)
				_font.AssertNumberOfCalls(t, "SetColor", 2)
				_font.AssertCalled(t, "SetColor", color.Color{Red: 0, Green: 0, Blue: 0})
			},
		}, {
			"Customizable col width",
			"Lorem Ipsum is simply dummy textá of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
			consts.Center,
			consts.Arial,
			color.Color{Red: 0, Green: 0, Blue: 0},
			func() *mocks.Fpdf {
				_pdf := &mocks.Fpdf{}
				_pdf.On("GetStringWidth", "Lorem Ipsum is simply dummy textá of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.").Return(900.0)
				_pdf.On("GetStringWidth", mock.Anything).Return(20.0)
				_pdf.On("GetMargins").Return(10.0, 10.0, 10.0, 10.0)
				_pdf.On("Text", mock.Anything, mock.Anything, mock.Anything)
				_pdf.On("UnicodeTranslatorFromDescriptor", mock.Anything).Return(func(value string) string { return value })
				return _pdf
			},
			func() *mocks.Font {
				_font := &mocks.Font{}
				_font.On("GetScaleFactor").Return(1.0)
				_font.On("GetFont").Return(consts.Arial, consts.Bold, 1.0)
				_font.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				_font.On("GetColor").Return(color.Color{Red: 0, Green: 0, Blue: 0})
				_font.On("SetColor", mock.Anything)
				return _font
			},
			func() *internal.Cell {
				return &internal.Cell{
					X:     1.0,
					Y:     5.0,
					Width: 25.0,
				}
			},
			func(t *testing.T, _pdf *mocks.Fpdf) {
				_pdf.AssertNumberOfCalls(t, "GetStringWidth", 274)
				_pdf.AssertCalled(t, "GetStringWidth", "Lorem Ipsum is simply dummy textá of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.")

				_pdf.AssertNumberOfCalls(t, "GetMargins", 91)

				_pdf.AssertNumberOfCalls(t, "Text", 91)
			},
			func(t *testing.T, _font *mocks.Font) {
				_font.AssertNumberOfCalls(t, "SetFont", 1)
				_font.AssertCalled(t, "SetFont", consts.Arial, consts.BoldItalic, 16.0)
				_font.AssertNumberOfCalls(t, "GetColor", 1)
				_font.AssertNumberOfCalls(t, "SetColor", 2)
				_font.AssertCalled(t, "SetColor", color.Color{Red: 0, Green: 0, Blue: 0})
			},
		},
		{
			"Custom Font Color",
			"CustomFontColor",
			consts.Left,
			consts.Arial,
			color.Color{Red: 20, Green: 20, Blue: 20},
			func() *mocks.Fpdf {
				_pdf := &mocks.Fpdf{}
				_pdf.On("GetStringWidth", mock.Anything).Return(12.0)
				_pdf.On("GetMargins").Return(10.0, 10.0, 10.0, 10.0)
				_pdf.On("Text", mock.Anything, mock.Anything, mock.Anything)
				_pdf.On("UnicodeTranslatorFromDescriptor", mock.Anything).Return(func(value string) string { return value })
				return _pdf
			},
			func() *mocks.Font {
				_font := &mocks.Font{}
				_font.On("GetScaleFactor").Return(1.0)
				_font.On("GetFont").Return(consts.Arial, consts.Bold, 1.0)
				_font.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				_font.On("GetColor").Return(color.Color{Red: 0, Green: 0, Blue: 0})
				_font.On("SetColor", mock.Anything)
				return _font
			},
			func() *internal.Cell {
				return nil
			},
			func(t *testing.T, _pdf *mocks.Fpdf) {
				_pdf.AssertNotCalled(t, "GetStringWidth")

				_pdf.AssertNumberOfCalls(t, "GetMargins", 1)

				_pdf.AssertNumberOfCalls(t, "Text", 1)
				_pdf.AssertCalled(t, "Text", 11.0, 16.0, "CustomFontColor")
			},
			func(t *testing.T, _font *mocks.Font) {
				_font.AssertNumberOfCalls(t, "SetFont", 1)
				_font.AssertCalled(t, "SetFont", consts.Arial, consts.BoldItalic, 16.0)
				_font.AssertNumberOfCalls(t, "GetColor", 1)
				_font.AssertNumberOfCalls(t, "SetColor", 2)
				_font.AssertCalled(t, "SetColor", color.Color{Red: 0, Green: 0, Blue: 0})
				_font.AssertCalled(t, "SetColor", color.Color{Red: 20, Green: 20, Blue: 20})
			},
		},
	}

	for _, c := range cases {
		// Arrange
		_pdf := c.pdf()
		_font := c.font()

		text := internal.NewText(_pdf, nil, _font)

		var cell internal.Cell
		if c.cell() == nil {
			cell = internal.Cell{
				X:     1.0,
				Y:     5.0,
				Width: 15.0,
			}
		} else {
			cell = *c.cell()
		}

		// Act
		text.Add(c.text, cell, props.Text{Family: c.family, Style: consts.BoldItalic, Size: 16.0, Align: c.align, Color: c.color})

		// Assert
		c.assertPdf(t, _pdf)
		c.assertFont(t, _font)
	}
}
