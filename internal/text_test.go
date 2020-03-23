package internal_test

import (
	"fmt"
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/internal/mocks"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewText(t *testing.T) {
	text := internal.NewText(&mocks.Pdf{}, &mocks.Math{}, &mocks.Font{})

	assert.NotNil(t, text)
	assert.Equal(t, fmt.Sprintf("%T", text), "*internal.text")
}

func TestText_GetLinesQuantity_WhenStringSmallerThanLimits(t *testing.T) {
	// Arrange
	pdf := &mocks.Pdf{}
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
	pdf := &mocks.Pdf{}
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
	pdf := &mocks.Pdf{}
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
	pdf := &mocks.Pdf{}
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
		pdf        func() *mocks.Pdf
		font       func() *mocks.Font
		assertPdf  func(t *testing.T, pdf *mocks.Pdf)
		assertFont func(t *testing.T, font *mocks.Font)
	}{
		{
			"Left Align",
			"TextHelper1",
			consts.Left,
			func() *mocks.Pdf {
				_pdf := &mocks.Pdf{}
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
				return _font
			},
			func(t *testing.T, _pdf *mocks.Pdf) {
				_pdf.AssertNotCalled(t, "GetStringWidth")

				_pdf.AssertNumberOfCalls(t, "GetMargins", 1)

				_pdf.AssertNumberOfCalls(t, "Text", 1)
				_pdf.AssertCalled(t, "Text", 11.0, 16.0, "TextHelper1")
			},
			func(t *testing.T, _font *mocks.Font) {
				_font.AssertNumberOfCalls(t, "SetFont", 1)
				_font.AssertCalled(t, "SetFont", consts.Arial, consts.BoldItalic, 16.0)
			},
		},
		{
			"Center Align",
			"TextHelper2",
			consts.Center,
			func() *mocks.Pdf {
				_pdf := &mocks.Pdf{}
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
				return _font
			},
			func(t *testing.T, _pdf *mocks.Pdf) {
				_pdf.AssertNumberOfCalls(t, "GetStringWidth", 1)
				_pdf.AssertCalled(t, "GetStringWidth", "TextHelper2")

				_pdf.AssertNumberOfCalls(t, "GetMargins", 1)

				_pdf.AssertNumberOfCalls(t, "Text", 1)
				_pdf.AssertCalled(t, "Text", 12.5, 16.0, "TextHelper2")
			},
			func(t *testing.T, _font *mocks.Font) {
				_font.AssertNumberOfCalls(t, "SetFont", 1)
				_font.AssertCalled(t, "SetFont", consts.Arial, consts.BoldItalic, 16.0)
			},
		},
		{
			"Right Align",
			"TextHelper3",
			consts.Right,
			func() *mocks.Pdf {
				_pdf := &mocks.Pdf{}
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
				return _font
			},
			func(t *testing.T, _pdf *mocks.Pdf) {
				_pdf.AssertNumberOfCalls(t, "GetStringWidth", 1)
				_pdf.AssertCalled(t, "GetStringWidth", "TextHelper3")

				_pdf.AssertNumberOfCalls(t, "GetMargins", 1)

				_pdf.AssertNumberOfCalls(t, "Text", 1)
				_pdf.AssertCalled(t, "Text", 14.0, 16.0, "TextHelper3")
			},
			func(t *testing.T, _font *mocks.Font) {
				_font.AssertNumberOfCalls(t, "SetFont", 1)
				_font.AssertCalled(t, "SetFont", consts.Arial, consts.BoldItalic, 16.0)
			},
		},
		{
			"Right Align",
			"TextHelper4",
			consts.Right,
			func() *mocks.Pdf {
				_pdf := &mocks.Pdf{}
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
				return _font
			},
			func(t *testing.T, _pdf *mocks.Pdf) {
				_pdf.AssertNumberOfCalls(t, "GetStringWidth", 1)
				_pdf.AssertCalled(t, "GetStringWidth", "TextHelper4")

				_pdf.AssertNumberOfCalls(t, "GetMargins", 1)

				_pdf.AssertNumberOfCalls(t, "Text", 1)
				_pdf.AssertCalled(t, "Text", 14.0, 16.0, "TextHelper4")
			},
			func(t *testing.T, _font *mocks.Font) {
				_font.AssertNumberOfCalls(t, "SetFont", 1)
				_font.AssertCalled(t, "SetFont", consts.Arial, consts.BoldItalic, 16.0)
			},
		},
		{
			"Bigger than cell width",
			"Lorem Ipsum is simply dummy textá of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
			consts.Center,
			func() *mocks.Pdf {
				_pdf := &mocks.Pdf{}
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
				return _font
			},
			func(t *testing.T, _pdf *mocks.Pdf) {
				_pdf.AssertNumberOfCalls(t, "GetStringWidth", 275)
				_pdf.AssertCalled(t, "GetStringWidth", "Lorem Ipsum is simply dummy textá of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.")

				_pdf.AssertNumberOfCalls(t, "GetMargins", 92)

				_pdf.AssertNumberOfCalls(t, "Text", 92)
			},
			func(t *testing.T, _font *mocks.Font) {
				_font.AssertNumberOfCalls(t, "SetFont", 1)
				_font.AssertCalled(t, "SetFont", consts.Arial, consts.BoldItalic, 16.0)
			},
		},
	}

	for _, c := range cases {
		// Arrange
		_pdf := c.pdf()
		_font := c.font()

		text := internal.NewText(_pdf, nil, _font)

		// Act
		text.Add(c.text, props.Text{Family: consts.Arial, Style: consts.BoldItalic, Size: 16.0, Align: c.align}, 5.0, 1, 15.0)

		// Assert
		c.assertPdf(t, _pdf)
		c.assertFont(t, _font)
	}
}
