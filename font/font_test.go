package font_test

import (
	"fmt"
	"github.com/johnfercher/maroto/font"
	"github.com/johnfercher/maroto/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewFont(t *testing.T) {
	_font := font.NewFont(&mocks.Pdf{}, 10, font.Arial, font.Bold)

	assert.NotNil(t, _font)
	assert.Equal(t, fmt.Sprintf("%T", _font), "*font.font")
	assert.Equal(t, _font.GetFamily(), font.Arial)
	assert.Equal(t, _font.GetStyle(), font.Bold)
	assert.Equal(t, _font.GetSize(), 10.0)
}

func TestFont_GetSetFamily(t *testing.T) {
	cases := []struct {
		name        string
		family      font.Family
		pdf         func() *mocks.Pdf
		assertCalls func(t *testing.T, pdf *mocks.Pdf)
		assertFont  func(t *testing.T, family font.Family)
	}{
		{
			"font.Arial",
			font.Arial,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "SetFont", 1)
				pdf.AssertCalled(t, "SetFont", "arial", "B", 10.0)
			},
			func(t *testing.T, family font.Family) {
				assert.Equal(t, family, font.Arial)
			},
		},
		{
			"font.Helvetica",
			font.Helvetica,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "SetFont", 1)
				pdf.AssertCalled(t, "SetFont", "helvetica", "B", 10.0)
			},
			func(t *testing.T, family font.Family) {
				assert.Equal(t, family, font.Helvetica)
			},
		},
		{
			"font.Symbol",
			font.Symbol,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "SetFont", 1)
				pdf.AssertCalled(t, "SetFont", "symbol", "B", 10.0)
			},
			func(t *testing.T, family font.Family) {
				assert.Equal(t, family, font.Symbol)
			},
		},
		{
			"font.ZapBats",
			font.ZapBats,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "SetFont", 1)
				pdf.AssertCalled(t, "SetFont", "zapfdingbats", "B", 10.0)
			},
			func(t *testing.T, family font.Family) {
				assert.Equal(t, family, font.ZapBats)
			},
		},
		{
			"font.Courier",
			font.Courier,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "SetFont", 1)
				pdf.AssertCalled(t, "SetFont", "courier", "B", 10.0)
			},
			func(t *testing.T, family font.Family) {
				assert.Equal(t, family, font.Courier)
			},
		},
	}

	for _, c := range cases {
		// Arrange
		pdf := c.pdf()
		font := font.NewFont(pdf, 10, font.Arial, font.Bold)

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
		style       font.Style
		pdf         func() *mocks.Pdf
		assertCalls func(t *testing.T, pdf *mocks.Pdf)
		assertStyle func(t *testing.T, style font.Style)
	}{
		{
			"font.Normal",
			font.Normal,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("SetFontStyle", mock.Anything)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "SetFontStyle", 1)
				pdf.AssertCalled(t, "SetFontStyle", "")
			},
			func(t *testing.T, style font.Style) {
				assert.Equal(t, style, font.Normal)
			},
		},
		{
			"font.Bold",
			font.Bold,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("SetFontStyle", mock.Anything)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "SetFontStyle", 1)
				pdf.AssertCalled(t, "SetFontStyle", "B")
			},
			func(t *testing.T, style font.Style) {
				assert.Equal(t, style, font.Bold)
			},
		},
		{
			"font.Italic",
			font.Italic,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("SetFontStyle", mock.Anything)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "SetFontStyle", 1)
				pdf.AssertCalled(t, "SetFontStyle", "I")
			},
			func(t *testing.T, style font.Style) {
				assert.Equal(t, style, font.Italic)
			},
		},
		{
			"font.BoldItalic",
			font.BoldItalic,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("SetFontStyle", mock.Anything)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "SetFontStyle", 1)
				pdf.AssertCalled(t, "SetFontStyle", "BI")
			},
			func(t *testing.T, style font.Style) {
				assert.Equal(t, style, font.BoldItalic)
			},
		},
	}

	for _, c := range cases {
		// Arrange
		pdf := c.pdf()
		font := font.NewFont(pdf, 10, font.Arial, font.Bold)

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
	font := font.NewFont(pdf, 10, font.Arial, font.Bold)

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
		family      font.Family
		style       font.Style
		size        float64
		pdf         func() *mocks.Pdf
		assertCalls func(t *testing.T, pdf *mocks.Pdf)
		assertFont  func(t *testing.T, family font.Family, style font.Style, size float64)
	}{
		{
			"font.Arial, font.Normal, 16",
			font.Arial,
			font.Normal,
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
			func(t *testing.T, family font.Family, style font.Style, size float64) {
				assert.Equal(t, family, font.Arial)
				assert.Equal(t, style, font.Normal)
				assert.Equal(t, 16, int(size))
			},
		},
		{
			"font.Helvetica, font.Bold, 13",
			font.Helvetica,
			font.Bold,
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
			func(t *testing.T, family font.Family, style font.Style, size float64) {
				assert.Equal(t, family, font.Helvetica)
				assert.Equal(t, style, font.Bold)
				assert.Equal(t, 13, int(size))
			},
		},
		{
			"font.Symbol, font.Italic, 10",
			font.Symbol,
			font.Italic,
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
			func(t *testing.T, family font.Family, style font.Style, size float64) {
				assert.Equal(t, family, font.Symbol)
				assert.Equal(t, style, font.Italic)
				assert.Equal(t, 10, int(size))
			},
		},
		{
			"font.ZapBats, font.BoldItalic, 5",
			font.ZapBats,
			font.BoldItalic,
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
			func(t *testing.T, family font.Family, style font.Style, size float64) {
				assert.Equal(t, family, font.ZapBats)
				assert.Equal(t, style, font.BoldItalic)
				assert.Equal(t, 5, int(size))
			},
		},
		{
			"font.Courier, font.Normal, 12",
			font.Courier,
			font.Normal,
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
			func(t *testing.T, family font.Family, style font.Style, size float64) {
				assert.Equal(t, family, font.Courier)
				assert.Equal(t, style, font.Normal)
				assert.Equal(t, 12, int(size))
			},
		},
	}

	for _, c := range cases {
		// Arrange
		pdf := c.pdf()
		font := font.NewFont(pdf, 10, font.Arial, font.Bold)

		// Act
		font.SetFont(c.family, c.style, c.size)
		family, style, size := font.GetFont()

		// Assert
		c.assertCalls(t, pdf)
		c.assertFont(t, family, style, size)
	}
}

func TestFont_GetStyleString(t *testing.T) {
	cases := []struct {
		name        string
		style       font.Style
		styleString string
	}{
		{
			"font.Normal",
			font.Normal,
			"",
		},
		{
			"font.Bold",
			font.Bold,
			"B",
		},
		{
			"font.Italic",
			font.Italic,
			"I",
		},
		{
			"font.BoldItalic",
			font.BoldItalic,
			"BI",
		},
	}

	for _, c := range cases {
		// Arrange
		_font := font.NewFont(&mocks.Pdf{}, 16.0, font.Arial, font.Bold)

		// Act
		styleString := _font.GetStyleString(c.style)

		// Assert
		assert.Equal(t, styleString, c.styleString)
	}
}

func TestFont_GetFamilyString(t *testing.T) {
	cases := []struct {
		name         string
		family       font.Family
		familyString string
	}{
		{
			"font.Arial",
			font.Arial,
			"arial",
		},
		{
			"font.Helvetica",
			font.Helvetica,
			"helvetica",
		},
		{
			"font.Symbol",
			font.Symbol,
			"symbol",
		},
		{
			"font.ZapBats",
			font.ZapBats,
			"zapfdingbats",
		},
		{
			"font.Courier",
			font.Courier,
			"courier",
		},
	}

	for _, c := range cases {
		// Arrange
		_font := font.NewFont(&mocks.Pdf{}, 16.0, font.Arial, font.Bold)

		// Act
		familyString := _font.GetFamilyString(c.family)

		// Assert
		assert.Equal(t, familyString, c.familyString)
	}
}
