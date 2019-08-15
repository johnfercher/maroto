package maroto_test

import (
	"fmt"
	"github.com/johnfercher/maroto"
	"github.com/johnfercher/maroto/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewText(t *testing.T) {
	text := maroto.NewText(&mocks.Pdf{}, &mocks.Math{}, &mocks.Font{})

	assert.NotNil(t, text)
	assert.Equal(t, fmt.Sprintf("%T", text), "*maroto.text")
}

func TestText_Add(t *testing.T) {
	cases := []struct {
		name       string
		text       string
		align      maroto.Align
		pdf        func() *mocks.Pdf
		math       func() *mocks.Math
		font       func() *mocks.Font
		assertPdf  func(t *testing.T, pdf *mocks.Pdf)
		assertMath func(t *testing.T, math *mocks.Math)
		assertFont func(t *testing.T, font *mocks.Font)
	}{
		{
			"Left Align",
			"TextHelper1",
			maroto.Left,
			func() *mocks.Pdf {
				_pdf := &mocks.Pdf{}
				_pdf.On("GetStringWidth", mock.Anything).Return(12.0)
				_pdf.On("GetMargins").Return(10.0, 10.0, 10.0, 10.0)
				_pdf.On("Text", mock.Anything, mock.Anything, mock.Anything)
				_pdf.On("UnicodeTranslatorFromDescriptor", mock.Anything).Return(func(value string) string { return value })
				return _pdf
			},
			func() *mocks.Math {
				_math := &mocks.Math{}
				_math.On("GetWidthPerCol", mock.Anything).Return(123.0)
				return _math
			},
			func() *mocks.Font {
				_font := &mocks.Font{}
				_font.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return _font
			},
			func(t *testing.T, _pdf *mocks.Pdf) {
				_pdf.AssertNotCalled(t, "GetStringWidth")

				_pdf.AssertNumberOfCalls(t, "GetMargins", 1)

				_pdf.AssertNumberOfCalls(t, "Text", 1)
				_pdf.AssertCalled(t, "Text", 133.0, 15.0, "TextHelper1")
			},
			func(t *testing.T, _math *mocks.Math) {
				_math.AssertNumberOfCalls(t, "GetWidthPerCol", 1)
				_math.AssertCalled(t, "GetWidthPerCol", 15.0)
			},
			func(t *testing.T, _font *mocks.Font) {
				_font.AssertNumberOfCalls(t, "SetFont", 1)
				_font.AssertCalled(t, "SetFont", maroto.Arial, maroto.BoldItalic, 16.0)
			},
		},
		{
			"Center Align",
			"TextHelper2",
			maroto.Center,
			func() *mocks.Pdf {
				_pdf := &mocks.Pdf{}
				_pdf.On("GetStringWidth", mock.Anything).Return(12.0)
				_pdf.On("GetMargins").Return(10.0, 10.0, 10.0, 10.0)
				_pdf.On("Text", mock.Anything, mock.Anything, mock.Anything)
				_pdf.On("UnicodeTranslatorFromDescriptor", mock.Anything).Return(func(value string) string { return value })
				return _pdf
			},
			func() *mocks.Math {
				_math := &mocks.Math{}
				_math.On("GetWidthPerCol", mock.Anything).Return(123.0)
				return _math
			},
			func() *mocks.Font {
				_font := &mocks.Font{}
				_font.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return _font
			},
			func(t *testing.T, _pdf *mocks.Pdf) {
				_pdf.AssertNumberOfCalls(t, "GetStringWidth", 1)
				_pdf.AssertCalled(t, "GetStringWidth", "TextHelper2")

				_pdf.AssertNumberOfCalls(t, "GetMargins", 1)

				_pdf.AssertNumberOfCalls(t, "Text", 1)
				_pdf.AssertCalled(t, "Text", 188.5, 15.0, "TextHelper2")
			},
			func(t *testing.T, _math *mocks.Math) {
				_math.AssertNumberOfCalls(t, "GetWidthPerCol", 1)
				_math.AssertCalled(t, "GetWidthPerCol", 15.0)
			},
			func(t *testing.T, _font *mocks.Font) {
				_font.AssertNumberOfCalls(t, "SetFont", 1)
				_font.AssertCalled(t, "SetFont", maroto.Arial, maroto.BoldItalic, 16.0)
			},
		},
		{
			"Right Align",
			"TextHelper3",
			maroto.Right,
			func() *mocks.Pdf {
				_pdf := &mocks.Pdf{}
				_pdf.On("GetStringWidth", mock.Anything).Return(12.0)
				_pdf.On("GetMargins").Return(10.0, 10.0, 10.0, 10.0)
				_pdf.On("Text", mock.Anything, mock.Anything, mock.Anything)
				_pdf.On("UnicodeTranslatorFromDescriptor", mock.Anything).Return(func(value string) string { return value })
				return _pdf
			},
			func() *mocks.Math {
				_math := &mocks.Math{}
				_math.On("GetWidthPerCol", mock.Anything).Return(123.0)
				return _math
			},
			func() *mocks.Font {
				_font := &mocks.Font{}
				_font.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return _font
			},
			func(t *testing.T, _pdf *mocks.Pdf) {
				_pdf.AssertNumberOfCalls(t, "GetStringWidth", 1)
				_pdf.AssertCalled(t, "GetStringWidth", "TextHelper3")

				_pdf.AssertNumberOfCalls(t, "GetMargins", 1)

				_pdf.AssertNumberOfCalls(t, "Text", 1)
				_pdf.AssertCalled(t, "Text", 244.0, 15.0, "TextHelper3")
			},
			func(t *testing.T, _math *mocks.Math) {
				_math.AssertNumberOfCalls(t, "GetWidthPerCol", 1)
				_math.AssertCalled(t, "GetWidthPerCol", 15.0)
			},
			func(t *testing.T, _font *mocks.Font) {
				_font.AssertNumberOfCalls(t, "SetFont", 1)
				_font.AssertCalled(t, "SetFont", maroto.Arial, maroto.BoldItalic, 16.0)
			},
		},
		{
			"Right Align",
			"TextHelper4",
			maroto.Right,
			func() *mocks.Pdf {
				_pdf := &mocks.Pdf{}
				_pdf.On("GetStringWidth", mock.Anything).Return(12.0)
				_pdf.On("GetMargins").Return(10.0, 10.0, 10.0, 10.0)
				_pdf.On("Text", mock.Anything, mock.Anything, mock.Anything)
				_pdf.On("UnicodeTranslatorFromDescriptor", mock.Anything).Return(func(value string) string { return value })
				return _pdf
			},
			func() *mocks.Math {
				_math := &mocks.Math{}
				_math.On("GetWidthPerCol", mock.Anything).Return(123.0)
				return _math
			},
			func() *mocks.Font {
				_font := &mocks.Font{}
				_font.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return _font
			},
			func(t *testing.T, _pdf *mocks.Pdf) {
				_pdf.AssertNumberOfCalls(t, "GetStringWidth", 1)
				_pdf.AssertCalled(t, "GetStringWidth", "TextHelper4")

				_pdf.AssertNumberOfCalls(t, "GetMargins", 1)

				_pdf.AssertNumberOfCalls(t, "Text", 1)
				_pdf.AssertCalled(t, "Text", 244.0, 15.0, "TextHelper4")
			},
			func(t *testing.T, _math *mocks.Math) {
				_math.AssertNumberOfCalls(t, "GetWidthPerCol", 1)
				_math.AssertCalled(t, "GetWidthPerCol", 15.0)
			},
			func(t *testing.T, _font *mocks.Font) {
				_font.AssertNumberOfCalls(t, "SetFont", 1)
				_font.AssertCalled(t, "SetFont", maroto.Arial, maroto.BoldItalic, 16.0)
			},
		},
		{
			"Bigger than cell width",
			"Lorem Ipsum is simply dummy textá of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
			maroto.Center,
			func() *mocks.Pdf {
				_pdf := &mocks.Pdf{}
				_pdf.On("GetStringWidth", "Lorem Ipsum is simply dummy textá of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.").Return(900.0)
				_pdf.On("GetStringWidth", mock.Anything).Return(20.0)
				_pdf.On("GetMargins").Return(10.0, 10.0, 10.0, 10.0)
				_pdf.On("Text", mock.Anything, mock.Anything, mock.Anything)
				_pdf.On("UnicodeTranslatorFromDescriptor", mock.Anything).Return(func(value string) string { return value })
				return _pdf
			},
			func() *mocks.Math {
				_math := &mocks.Math{}
				_math.On("GetWidthPerCol", mock.Anything).Return(123.0)
				return _math
			},
			func() *mocks.Font {
				_font := &mocks.Font{}
				_font.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return _font
			},
			func(t *testing.T, _pdf *mocks.Pdf) {
				_pdf.AssertNumberOfCalls(t, "GetStringWidth", 199)
				_pdf.AssertCalled(t, "GetStringWidth", "Lorem Ipsum is simply dummy textá of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.")

				_pdf.AssertNumberOfCalls(t, "GetMargins", 16)

				_pdf.AssertNumberOfCalls(t, "Text", 16)
			},
			func(t *testing.T, _math *mocks.Math) {
				_math.AssertNumberOfCalls(t, "GetWidthPerCol", 1)
				_math.AssertCalled(t, "GetWidthPerCol", 15.0)
			},
			func(t *testing.T, _font *mocks.Font) {
				_font.AssertNumberOfCalls(t, "SetFont", 1)
				_font.AssertCalled(t, "SetFont", maroto.Arial, maroto.BoldItalic, 16.0)
			},
		},
	}

	for _, c := range cases {
		// Arrange
		_pdf := c.pdf()
		_math := c.math()
		_font := c.font()

		text := maroto.NewText(_pdf, _math, _font)

		// Act
		text.Add(c.text, maroto.TextProp{Family: maroto.Arial, Style: maroto.BoldItalic, Size: 16.0, Align: c.align}, 5.0, 1, 15.0)

		// Assert
		c.assertPdf(t, _pdf)
		c.assertMath(t, _math)
		c.assertFont(t, _font)
	}
}
