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
			maroto.Left,
			func() *mocks.Pdf {
				_pdf := &mocks.Pdf{}
				_pdf.On("GetStringWidth", mock.Anything).Return(12.0)
				_pdf.On("GetMargins").Return(10.0, 10.0, 10.0, 10.0)
				_pdf.On("Text", mock.Anything, mock.Anything, mock.Anything)
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
				_pdf.AssertCalled(t, "Text", 133.0, 15.0, "Text")
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
			maroto.Center,
			func() *mocks.Pdf {
				_pdf := &mocks.Pdf{}
				_pdf.On("GetStringWidth", mock.Anything).Return(12.0)
				_pdf.On("GetMargins").Return(10.0, 10.0, 10.0, 10.0)
				_pdf.On("Text", mock.Anything, mock.Anything, mock.Anything)
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
				_pdf.AssertCalled(t, "GetStringWidth", "Text")

				_pdf.AssertNumberOfCalls(t, "GetMargins", 1)

				_pdf.AssertNumberOfCalls(t, "Text", 1)
				_pdf.AssertCalled(t, "Text", 188.5, 15.0, "Text")
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
		text.Add("Text", maroto.Arial, maroto.BoldItalic, 16.0, 5.0, c.align, 1, 15.0)

		// Assert
		c.assertPdf(t, _pdf)
		c.assertMath(t, _math)
		c.assertFont(t, _font)
	}
}
