package internal_test

import (
	"fmt"
	"github.com/Vale-sail/maroto/internal"
	"github.com/Vale-sail/maroto/internal/mocks"
	"github.com/Vale-sail/maroto/pkg/consts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewFont(t *testing.T) {
	font := internal.NewFont(&mocks.Pdf{}, 10, consts.Arial, consts.Bold)

	assert.NotNil(t, font)
	assert.Equal(t, fmt.Sprintf("%T", font), "*internal.font")
	assert.Equal(t, font.GetFamily(), consts.Arial)
	assert.Equal(t, font.GetStyle(), consts.Bold)
	assert.Equal(t, font.GetSize(), 10.0)
}

func TestFont_GetSetFamily(t *testing.T) {
	cases := []struct {
		name        string
		family      consts.Family
		pdf         func() *mocks.Pdf
		assertCalls func(t *testing.T, pdf *mocks.Pdf)
		assertFont  func(t *testing.T, family consts.Family)
	}{
		{
			"PdfMaroto.Arial",
			consts.Arial,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "SetFont", 1)
				pdf.AssertCalled(t, "SetFont", "arial", "B", 10.0)
			},
			func(t *testing.T, family consts.Family) {
				assert.Equal(t, family, consts.Arial)
			},
		},
		{
			"PdfMaroto.Helvetica",
			consts.Helvetica,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "SetFont", 1)
				pdf.AssertCalled(t, "SetFont", "helvetica", "B", 10.0)
			},
			func(t *testing.T, family consts.Family) {
				assert.Equal(t, family, consts.Helvetica)
			},
		},
		{
			"PdfMaroto.Symbol",
			consts.Symbol,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "SetFont", 1)
				pdf.AssertCalled(t, "SetFont", "symbol", "B", 10.0)
			},
			func(t *testing.T, family consts.Family) {
				assert.Equal(t, family, consts.Symbol)
			},
		},
		{
			"PdfMaroto.ZapBats",
			consts.ZapBats,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "SetFont", 1)
				pdf.AssertCalled(t, "SetFont", "zapfdingbats", "B", 10.0)
			},
			func(t *testing.T, family consts.Family) {
				assert.Equal(t, family, consts.ZapBats)
			},
		},
		{
			"PdfMaroto.Courier",
			consts.Courier,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "SetFont", 1)
				pdf.AssertCalled(t, "SetFont", "courier", "B", 10.0)
			},
			func(t *testing.T, family consts.Family) {
				assert.Equal(t, family, consts.Courier)
			},
		},
	}

	for _, c := range cases {
		// Arrange
		pdf := c.pdf()
		font := internal.NewFont(pdf, 10, consts.Arial, consts.Bold)

		// Act
		font.SetFamily(c.family)

		// Assert
		c.assertCalls(t, pdf)
		c.assertFont(t, font.GetFamily())
	}
}

func TestFont_GetSetStyle(t *testing.T) {
	cases := []struct {
		name        string
		style       consts.Style
		pdf         func() *mocks.Pdf
		assertCalls func(t *testing.T, pdf *mocks.Pdf)
		assertStyle func(t *testing.T, style consts.Style)
	}{
		{
			"PdfMaroto.Normal",
			consts.Normal,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("SetFontStyle", mock.Anything)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "SetFontStyle", 1)
				pdf.AssertCalled(t, "SetFontStyle", "")
			},
			func(t *testing.T, style consts.Style) {
				assert.Equal(t, style, consts.Normal)
			},
		},
		{
			"PdfMaroto.Bold",
			consts.Bold,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("SetFontStyle", mock.Anything)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "SetFontStyle", 1)
				pdf.AssertCalled(t, "SetFontStyle", "B")
			},
			func(t *testing.T, style consts.Style) {
				assert.Equal(t, style, consts.Bold)
			},
		},
		{
			"PdfMaroto.Italic",
			consts.Italic,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("SetFontStyle", mock.Anything)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "SetFontStyle", 1)
				pdf.AssertCalled(t, "SetFontStyle", "I")
			},
			func(t *testing.T, style consts.Style) {
				assert.Equal(t, style, consts.Italic)
			},
		},
		{
			"PdfMaroto.BoldItalic",
			consts.BoldItalic,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("SetFontStyle", mock.Anything)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "SetFontStyle", 1)
				pdf.AssertCalled(t, "SetFontStyle", "BI")
			},
			func(t *testing.T, style consts.Style) {
				assert.Equal(t, style, consts.BoldItalic)
			},
		},
	}

	for _, c := range cases {
		// Arrange
		pdf := c.pdf()
		font := internal.NewFont(pdf, 10, consts.Arial, consts.Bold)

		// Act
		font.SetStyle(c.style)

		// Assert
		c.assertCalls(t, pdf)
		c.assertStyle(t, font.GetStyle())
	}
}

func TestFont_GetSetSize(t *testing.T) {
	// Arrange
	pdf := &mocks.Pdf{}
	pdf.On("SetFontSize", mock.Anything)
	font := internal.NewFont(pdf, 10, consts.Arial, consts.Bold)

	// Act
	font.SetSize(16)

	// Assert
	pdf.AssertNumberOfCalls(t, "SetFontSize", 1)
	pdf.MethodCalled("SetFontSize", 16)
	assert.Equal(t, font.GetSize(), 16.0)
}

func TestFont_GetSetFont(t *testing.T) {
	cases := []struct {
		name        string
		family      consts.Family
		style       consts.Style
		size        float64
		pdf         func() *mocks.Pdf
		assertCalls func(t *testing.T, pdf *mocks.Pdf)
		assertFont  func(t *testing.T, family consts.Family, style consts.Style, size float64)
	}{
		{
			"PdfMaroto.Arial, PdfMaroto.Normal, 16",
			consts.Arial,
			consts.Normal,
			16.0,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "SetFont", 1)
				pdf.AssertCalled(t, "SetFont", "arial", "", 16.0)
			},
			func(t *testing.T, family consts.Family, style consts.Style, size float64) {
				assert.Equal(t, family, consts.Arial)
				assert.Equal(t, style, consts.Normal)
				assert.Equal(t, 16, int(size))
			},
		},
		{
			"PdfMaroto.Helvetica, PdfMaroto.Bold, 13",
			consts.Helvetica,
			consts.Bold,
			13,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "SetFont", 1)
				pdf.AssertCalled(t, "SetFont", "helvetica", "B", 13.0)
			},
			func(t *testing.T, family consts.Family, style consts.Style, size float64) {
				assert.Equal(t, family, consts.Helvetica)
				assert.Equal(t, style, consts.Bold)
				assert.Equal(t, 13, int(size))
			},
		},
		{
			"PdfMaroto.Symbol, PdfMaroto.Italic, 10",
			consts.Symbol,
			consts.Italic,
			10,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "SetFont", 1)
				pdf.AssertCalled(t, "SetFont", "symbol", "I", 10.0)
			},
			func(t *testing.T, family consts.Family, style consts.Style, size float64) {
				assert.Equal(t, family, consts.Symbol)
				assert.Equal(t, style, consts.Italic)
				assert.Equal(t, 10, int(size))
			},
		},
		{
			"PdfMaroto.ZapBats, PdfMaroto.BoldItalic, 5",
			consts.ZapBats,
			consts.BoldItalic,
			5,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "SetFont", 1)
				pdf.AssertCalled(t, "SetFont", "zapfdingbats", "BI", 5.0)
			},
			func(t *testing.T, family consts.Family, style consts.Style, size float64) {
				assert.Equal(t, family, consts.ZapBats)
				assert.Equal(t, style, consts.BoldItalic)
				assert.Equal(t, 5, int(size))
			},
		},
		{
			"PdfMaroto.Courier, PdfMaroto.Normal, 12",
			consts.Courier,
			consts.Normal,
			12,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "SetFont", 1)
				pdf.AssertCalled(t, "SetFont", "courier", "", 12.0)
			},
			func(t *testing.T, family consts.Family, style consts.Style, size float64) {
				assert.Equal(t, family, consts.Courier)
				assert.Equal(t, style, consts.Normal)
				assert.Equal(t, 12, int(size))
			},
		},
	}

	for _, c := range cases {
		// Arrange
		pdf := c.pdf()
		font := internal.NewFont(pdf, 10, consts.Arial, consts.Bold)

		// Act
		font.SetFont(c.family, c.style, c.size)
		family, style, size := font.GetFont()

		// Assert
		c.assertCalls(t, pdf)
		c.assertFont(t, family, style, size)
	}
}

func TestFont_GetScaleFactor(t *testing.T) {
	// Arrange
	pdf := &mocks.Pdf{}
	pdf.On("GetFontSize").Return(1.0, 1.0)
	sut := internal.NewFont(pdf, 0, consts.Arial, consts.Normal)

	// Act
	scalarFactor := sut.GetScaleFactor()

	// Assert
	assert.InDelta(t, scalarFactor, 2.83, 0.1)
}
