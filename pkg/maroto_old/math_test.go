package maroto_test

import (
	"fmt"
	"github.com/johnfercher/maroto"
	"github.com/johnfercher/maroto/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMath(t *testing.T) {
	math := maroto.NewMath(&mocks.Pdf{})

	assert.NotNil(t, math)
	assert.Equal(t, fmt.Sprintf("%T", math), "*maroto.math")
}

func TestMath_GetWidthPerCol(t *testing.T) {
	cases := []struct {
		name        string
		qtdCols     float64
		pdf         func() *mocks.Pdf
		assertCalls func(t *testing.T, pdf *mocks.Pdf)
		assertWidth func(t *testing.T, width float64)
	}{
		{
			"1 col, margins 10 10",
			1,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("GetPageSize").Return(210.0, 0.0)
				pdf.On("GetMargins").Return(10.0, 10.0, 10.0, 10.0)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "GetPageSize", 1)
				pdf.AssertNumberOfCalls(t, "GetMargins", 1)
			},
			func(t *testing.T, width float64) {
				assert.Equal(t, int(width), 190)
			},
		},
		{
			"2 col, margins 10 10",
			2,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("GetPageSize").Return(210.0, 0.0)
				pdf.On("GetMargins").Return(10.0, 10.0, 10.0, 10.0)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "GetPageSize", 1)
				pdf.AssertNumberOfCalls(t, "GetMargins", 1)
			},
			func(t *testing.T, width float64) {
				assert.Equal(t, int(width), 95)
			},
		},
		{
			"4 col, margins 10 10",
			4,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("GetPageSize").Return(210.0, 0.0)
				pdf.On("GetMargins").Return(10.0, 10.0, 10.0, 10.0)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "GetPageSize", 1)
				pdf.AssertNumberOfCalls(t, "GetMargins", 1)
			},
			func(t *testing.T, width float64) {
				assert.Equal(t, int(width), 47)
			},
		},
		{
			"1 col, margins 20 20",
			1,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("GetPageSize").Return(210.0, 0.0)
				pdf.On("GetMargins").Return(20.0, 20.0, 20.0, 20.0)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "GetPageSize", 1)
				pdf.AssertNumberOfCalls(t, "GetMargins", 1)
			},
			func(t *testing.T, width float64) {
				assert.Equal(t, int(width), 170)
			},
		},
		{
			"2 col, margins 20 20",
			2,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("GetPageSize").Return(210.0, 0.0)
				pdf.On("GetMargins").Return(20.0, 20.0, 20.0, 20.0)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "GetPageSize", 1)
				pdf.AssertNumberOfCalls(t, "GetMargins", 1)
			},
			func(t *testing.T, width float64) {
				assert.Equal(t, int(width), 85)
			},
		},
		{
			"4 col, margins 20 20",
			4,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("GetPageSize").Return(210.0, 0.0)
				pdf.On("GetMargins").Return(20.0, 20.0, 20.0, 20.0)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "GetPageSize", 1)
				pdf.AssertNumberOfCalls(t, "GetMargins", 1)
			},
			func(t *testing.T, width float64) {
				assert.Equal(t, int(width), 42)
			},
		},
	}

	for _, c := range cases {
		// Arrange
		pdf := c.pdf()
		math := maroto.NewMath(pdf)

		// Act
		width := math.GetWidthPerCol(c.qtdCols)

		// Assert
		c.assertWidth(t, width)
		c.assertCalls(t, pdf)
	}
}

func TestMath_GetRectCenterColProperties(t *testing.T) {
	cases := []struct {
		name           string
		width          float64
		height         float64
		percent        float64
		pdf            func() *mocks.Pdf
		assertPdfCalls func(t *testing.T, pdf *mocks.Pdf)
		assertResult   func(t *testing.T, x, y, w, h float64)
	}{
		{
			"When Image has width greater than height",
			300,
			200,
			100.0,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("GetMargins").Return(12.0, 11.0, 13.0, 15.0)
				pdf.On("GetPageSize").Return(201.0, 301.0)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "GetMargins", 1)
				pdf.AssertNumberOfCalls(t, "GetPageSize", 1)
			},
			func(t *testing.T, x, y, w, h float64) {
				assert.InDelta(t, x, 82.4, 0.1)
				assert.InDelta(t, y, 11.7, 0.1)
				assert.InDelta(t, w, 35.2, 0.1)
				assert.InDelta(t, h, 23.4, 0.1)
			},
		},
		{
			"When Image has height greater than width",
			200,
			300,
			100.0,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("GetMargins").Return(15.0, 12.0, 17.0, 10.0)
				pdf.On("GetPageSize").Return(211.0, 233.0)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "GetMargins", 1)
				pdf.AssertNumberOfCalls(t, "GetPageSize", 1)
			},
			func(t *testing.T, x, y, w, h float64) {
				assert.InDelta(t, x, 96.1, 0.1)
				assert.InDelta(t, y, 12.0, 0.1)
				assert.InDelta(t, w, 16.6, 0.1)
				assert.InDelta(t, h, 25.0, 0.1)
			},
		},
		{
			"When Image has height greater than width, percent 45",
			200,
			300,
			45.0,
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("GetMargins").Return(15.0, 12.0, 17.0, 10.0)
				pdf.On("GetPageSize").Return(211.0, 233.0)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "GetMargins", 1)
				pdf.AssertNumberOfCalls(t, "GetPageSize", 1)
			},
			func(t *testing.T, x, y, w, h float64) {
				assert.InDelta(t, x, 96.4, 0.1)
				assert.InDelta(t, y, 12.4, 0.1)
				assert.InDelta(t, w, 16.1, 0.1)
				assert.InDelta(t, h, 24.1, 0.1)
			},
		},
	}

	for _, c := range cases {
		// Arrange
		pdf := c.pdf()

		math := maroto.NewMath(pdf)

		// Act
		x, y, w, h := math.GetRectCenterColProperties(c.width, c.height, 5, 25.0, 2, c.percent)

		// Assert
		c.assertPdfCalls(t, pdf)
		c.assertResult(t, x, y, w, h)
	}
}

func TestMath_GetCenterCorrection(t *testing.T) {
	// Arrange
	pdf := &mocks.Pdf{}
	math := maroto.NewMath(pdf)

	// Act
	correction := math.GetCenterCorrection(10, 5)

	// Assert
	assert.Equal(t, correction, 2.5)
}
