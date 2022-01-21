package internal_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/internal/mocks"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewFont(t *testing.T) {
	font := internal.NewFont(&mocks.Fpdf{}, 10, consts.Arial, consts.Bold)

	assert.NotNil(t, font)
	assert.Equal(t, fmt.Sprintf("%T", font), "*internal.font")
	assert.Equal(t, font.GetFamily(), consts.Arial)
	assert.Equal(t, font.GetStyle(), consts.Bold)
	assert.Equal(t, font.GetSize(), 10.0)
	assert.Equal(t, font.GetColor(), color.Color{Red: 0, Green: 0, Blue: 0})
}

func TestFont_GetSetFamily(t *testing.T) {
	cases := []struct {
		name        string
		family      string
		fpdf        func() *mocks.Fpdf
		assertCalls func(t *testing.T, Fpdf *mocks.Fpdf)
		assertFont  func(t *testing.T, family string)
	}{
		{
			"PdfMaroto.Arial",
			consts.Arial,
			func() *mocks.Fpdf {
				fpdf := &mocks.Fpdf{}
				fpdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return fpdf
			},
			func(t *testing.T, fpdf *mocks.Fpdf) {
				fpdf.AssertNumberOfCalls(t, "SetFont", 1)
				fpdf.AssertCalled(t, "SetFont", "arial", "B", 10.0)
			},
			func(t *testing.T, family string) {
				assert.Equal(t, family, consts.Arial)
			},
		},
		{
			"PdfMaroto.Helvetica",
			consts.Helvetica,
			func() *mocks.Fpdf {
				fpdf := &mocks.Fpdf{}
				fpdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return fpdf
			},
			func(t *testing.T, fpdf *mocks.Fpdf) {
				fpdf.AssertNumberOfCalls(t, "SetFont", 1)
				fpdf.AssertCalled(t, "SetFont", "helvetica", "B", 10.0)
			},
			func(t *testing.T, family string) {
				assert.Equal(t, family, consts.Helvetica)
			},
		},
		{
			"PdfMaroto.Symbol",
			consts.Symbol,
			func() *mocks.Fpdf {
				fpdf := &mocks.Fpdf{}
				fpdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return fpdf
			},
			func(t *testing.T, fpdf *mocks.Fpdf) {
				fpdf.AssertNumberOfCalls(t, "SetFont", 1)
				fpdf.AssertCalled(t, "SetFont", "symbol", "B", 10.0)
			},
			func(t *testing.T, family string) {
				assert.Equal(t, family, consts.Symbol)
			},
		},
		{
			"PdfMaroto.ZapBats",
			consts.ZapBats,
			func() *mocks.Fpdf {
				fpdf := &mocks.Fpdf{}
				fpdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return fpdf
			},
			func(t *testing.T, fpdf *mocks.Fpdf) {
				fpdf.AssertNumberOfCalls(t, "SetFont", 1)
				fpdf.AssertCalled(t, "SetFont", "zapfdingbats", "B", 10.0)
			},
			func(t *testing.T, family string) {
				assert.Equal(t, family, consts.ZapBats)
			},
		},
		{
			"PdfMaroto.Courier",
			consts.Courier,
			func() *mocks.Fpdf {
				fpdf := &mocks.Fpdf{}
				fpdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return fpdf
			},
			func(t *testing.T, fpdf *mocks.Fpdf) {
				fpdf.AssertNumberOfCalls(t, "SetFont", 1)
				fpdf.AssertCalled(t, "SetFont", "courier", "B", 10.0)
			},
			func(t *testing.T, family string) {
				assert.Equal(t, family, consts.Courier)
			},
		},
	}

	for _, c := range cases {
		// Arrange
		fpdf := c.fpdf()
		font := internal.NewFont(fpdf, 10, consts.Arial, consts.Bold)

		// Act
		font.SetFamily(c.family)

		// Assert
		c.assertCalls(t, fpdf)
		c.assertFont(t, font.GetFamily())
	}
}

func TestFont_GetSetStyle(t *testing.T) {
	cases := []struct {
		name        string
		style       consts.Style
		fpdf        func() *mocks.Fpdf
		assertCalls func(t *testing.T, Fpdf *mocks.Fpdf)
		assertStyle func(t *testing.T, style consts.Style)
	}{
		{
			"PdfMaroto.Normal",
			consts.Normal,
			func() *mocks.Fpdf {
				fpdf := &mocks.Fpdf{}
				fpdf.On("SetFontStyle", mock.Anything)
				return fpdf
			},
			func(t *testing.T, fpdf *mocks.Fpdf) {
				fpdf.AssertNumberOfCalls(t, "SetFontStyle", 1)
				fpdf.AssertCalled(t, "SetFontStyle", "")
			},
			func(t *testing.T, style consts.Style) {
				assert.Equal(t, style, consts.Normal)
			},
		},
		{
			"PdfMaroto.Bold",
			consts.Bold,
			func() *mocks.Fpdf {
				fpdf := &mocks.Fpdf{}
				fpdf.On("SetFontStyle", mock.Anything)
				return fpdf
			},
			func(t *testing.T, fpdf *mocks.Fpdf) {
				fpdf.AssertNumberOfCalls(t, "SetFontStyle", 1)
				fpdf.AssertCalled(t, "SetFontStyle", "B")
			},
			func(t *testing.T, style consts.Style) {
				assert.Equal(t, style, consts.Bold)
			},
		},
		{
			"PdfMaroto.Italic",
			consts.Italic,
			func() *mocks.Fpdf {
				fpdf := &mocks.Fpdf{}
				fpdf.On("SetFontStyle", mock.Anything)
				return fpdf
			},
			func(t *testing.T, fpdf *mocks.Fpdf) {
				fpdf.AssertNumberOfCalls(t, "SetFontStyle", 1)
				fpdf.AssertCalled(t, "SetFontStyle", "I")
			},
			func(t *testing.T, style consts.Style) {
				assert.Equal(t, style, consts.Italic)
			},
		},
		{
			"PdfMaroto.BoldItalic",
			consts.BoldItalic,
			func() *mocks.Fpdf {
				fpdf := &mocks.Fpdf{}
				fpdf.On("SetFontStyle", mock.Anything)
				return fpdf
			},
			func(t *testing.T, fpdf *mocks.Fpdf) {
				fpdf.AssertNumberOfCalls(t, "SetFontStyle", 1)
				fpdf.AssertCalled(t, "SetFontStyle", "BI")
			},
			func(t *testing.T, style consts.Style) {
				assert.Equal(t, style, consts.BoldItalic)
			},
		},
	}

	for _, c := range cases {
		// Arrange
		fpdf := c.fpdf()
		font := internal.NewFont(fpdf, 10, consts.Arial, consts.Bold)

		// Act
		font.SetStyle(c.style)

		// Assert
		c.assertCalls(t, fpdf)
		c.assertStyle(t, font.GetStyle())
	}
}

func TestFont_GetSetSize(t *testing.T) {
	// Arrange
	fpdf := &mocks.Fpdf{}
	fpdf.On("SetFontSize", mock.Anything)
	font := internal.NewFont(fpdf, 10, consts.Arial, consts.Bold)

	// Act
	font.SetSize(16)

	// Assert
	fpdf.AssertNumberOfCalls(t, "SetFontSize", 1)
	fpdf.MethodCalled("SetFontSize", 16)
	assert.Equal(t, font.GetSize(), 16.0)
}

func TestFont_GetSetFont(t *testing.T) {
	cases := []struct {
		name        string
		family      string
		style       consts.Style
		size        float64
		fpdf        func() *mocks.Fpdf
		assertCalls func(t *testing.T, Fpdf *mocks.Fpdf)
		assertFont  func(t *testing.T, family string, style consts.Style, size float64)
	}{
		{
			"PdfMaroto.Arial, PdfMaroto.Normal, 16",
			consts.Arial,
			consts.Normal,
			16.0,
			func() *mocks.Fpdf {
				fpdf := &mocks.Fpdf{}
				fpdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return fpdf
			},
			func(t *testing.T, fpdf *mocks.Fpdf) {
				fpdf.AssertNumberOfCalls(t, "SetFont", 1)
				fpdf.AssertCalled(t, "SetFont", "arial", "", 16.0)
			},
			func(t *testing.T, family string, style consts.Style, size float64) {
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
			func() *mocks.Fpdf {
				fpdf := &mocks.Fpdf{}
				fpdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return fpdf
			},
			func(t *testing.T, fpdf *mocks.Fpdf) {
				fpdf.AssertNumberOfCalls(t, "SetFont", 1)
				fpdf.AssertCalled(t, "SetFont", "helvetica", "B", 13.0)
			},
			func(t *testing.T, family string, style consts.Style, size float64) {
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
			func() *mocks.Fpdf {
				fpdf := &mocks.Fpdf{}
				fpdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return fpdf
			},
			func(t *testing.T, fpdf *mocks.Fpdf) {
				fpdf.AssertNumberOfCalls(t, "SetFont", 1)
				fpdf.AssertCalled(t, "SetFont", "symbol", "I", 10.0)
			},
			func(t *testing.T, family string, style consts.Style, size float64) {
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
			func() *mocks.Fpdf {
				fpdf := &mocks.Fpdf{}
				fpdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return fpdf
			},
			func(t *testing.T, fpdf *mocks.Fpdf) {
				fpdf.AssertNumberOfCalls(t, "SetFont", 1)
				fpdf.AssertCalled(t, "SetFont", "zapfdingbats", "BI", 5.0)
			},
			func(t *testing.T, family string, style consts.Style, size float64) {
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
			func() *mocks.Fpdf {
				fpdf := &mocks.Fpdf{}
				fpdf.On("SetFont", mock.Anything, mock.Anything, mock.Anything)
				return fpdf
			},
			func(t *testing.T, fpdf *mocks.Fpdf) {
				fpdf.AssertNumberOfCalls(t, "SetFont", 1)
				fpdf.AssertCalled(t, "SetFont", "courier", "", 12.0)
			},
			func(t *testing.T, family string, style consts.Style, size float64) {
				assert.Equal(t, family, consts.Courier)
				assert.Equal(t, style, consts.Normal)
				assert.Equal(t, 12, int(size))
			},
		},
	}

	for _, c := range cases {
		// Arrange
		fpdf := c.fpdf()
		font := internal.NewFont(fpdf, 10, consts.Arial, consts.Bold)

		// Act
		font.SetFont(c.family, c.style, c.size)
		family, style, size := font.GetFont()

		// Assert
		c.assertCalls(t, fpdf)
		c.assertFont(t, family, style, size)
	}
}

func TestFont_GetScaleFactor(t *testing.T) {
	// Arrange
	fpdf := &mocks.Fpdf{}
	fpdf.On("GetFontSize").Return(1.0, 1.0)
	sut := internal.NewFont(fpdf, 0, consts.Arial, consts.Normal)

	// Act
	scalarFactor := sut.GetScaleFactor()

	// Assert
	assert.InDelta(t, scalarFactor, 2.83, 0.1)
}

func TestFont_GetSetColor(t *testing.T) {
	cases := []struct {
		name        string
		fontColor   color.Color
		Fpdf        func() *mocks.Fpdf
		assertCalls func(t *testing.T, Fpdf *mocks.Fpdf)
		assertFont  func(t *testing.T, fontColor color.Color)
	}{
		{
			"Without custom color",
			color.Color{Red: 0, Green: 0, Blue: 0},
			func() *mocks.Fpdf {
				Fpdf := &mocks.Fpdf{}
				Fpdf.On("SetTextColor", mock.Anything, mock.Anything, mock.Anything)
				return Fpdf
			},
			func(t *testing.T, Fpdf *mocks.Fpdf) {
				Fpdf.AssertNumberOfCalls(t, "SetTextColor", 1)
				Fpdf.AssertCalled(t, "SetTextColor", 0, 0, 0)
			},
			func(t *testing.T, fontColor color.Color) {
				assert.Equal(t, fontColor.Red, 0)
				assert.Equal(t, fontColor.Green, 0)
				assert.Equal(t, fontColor.Blue, 0)
			},
		},
		{
			"With custom color",
			color.Color{Red: 20, Green: 20, Blue: 20},
			func() *mocks.Fpdf {
				Fpdf := &mocks.Fpdf{}
				Fpdf.On("SetTextColor", mock.Anything, mock.Anything, mock.Anything)
				return Fpdf
			},
			func(t *testing.T, Fpdf *mocks.Fpdf) {
				Fpdf.AssertNumberOfCalls(t, "SetTextColor", 1)
				Fpdf.AssertCalled(t, "SetTextColor", 20, 20, 20)
			},
			func(t *testing.T, fontColor color.Color) {
				assert.Equal(t, fontColor.Red, 20)
				assert.Equal(t, fontColor.Green, 20)
				assert.Equal(t, fontColor.Blue, 20)
			},
		},
	}

	for _, c := range cases {
		// Arrange
		Fpdf := c.Fpdf()
		font := internal.NewFont(Fpdf, 10, consts.Arial, consts.Bold)

		// Act
		font.SetColor(c.fontColor)
		fontColor := font.GetColor()

		// Assert
		c.assertCalls(t, Fpdf)
		c.assertFont(t, fontColor)
	}
}
