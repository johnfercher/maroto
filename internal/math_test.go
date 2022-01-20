package internal_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/internal/mocks"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/stretchr/testify/assert"
)

func TestNewMath(t *testing.T) {
	math := internal.NewMath(&mocks.Fpdf{})

	assert.NotNil(t, math)
	assert.Equal(t, fmt.Sprintf("%T", math), "*internal.math")
}

func TestMath_GetRectCenterColProperties(t *testing.T) {
	cases := []struct {
		name           string
		width          float64
		height         float64
		percent        float64
		pdf            func() *mocks.Fpdf
		assertPdfCalls func(t *testing.T, pdf *mocks.Fpdf)
		assertResult   func(t *testing.T, x, y, w, h float64)
	}{
		{
			"When cel proportion is greater than col",
			20,
			26,
			100.0,
			func() *mocks.Fpdf {
				pdf := &mocks.Fpdf{}
				pdf.On("GetMargins").Return(12.0, 11.0, 13.0, 15.0)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Fpdf) {
				pdf.AssertNumberOfCalls(t, "GetMargins", 1)
			},
			func(t *testing.T, x, y, w, h float64) {
				assert.InDelta(t, x, 14.3, 0.1)
				assert.InDelta(t, y, 11.0, 0.1)
				assert.InDelta(t, w, 19.2, 0.1)
				assert.InDelta(t, h, 25.0, 0.1)
			},
		},
		{
			"When cel proportion is greater than col, 45 percent",
			20,
			26,
			45.0,
			func() *mocks.Fpdf {
				pdf := &mocks.Fpdf{}
				pdf.On("GetMargins").Return(12.0, 11.0, 13.0, 15.0)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Fpdf) {
				pdf.AssertNumberOfCalls(t, "GetMargins", 1)
			},
			func(t *testing.T, x, y, w, h float64) {
				assert.InDelta(t, x, 19.6, 0.1)
				assert.InDelta(t, y, 17.8, 0.1)
				assert.InDelta(t, w, 8.6, 0.1)
				assert.InDelta(t, h, 11.2, 0.1)
			},
		},
		{
			"When cen proportion is less than col",
			26,
			20,
			100.0,
			func() *mocks.Fpdf {
				pdf := &mocks.Fpdf{}
				pdf.On("GetMargins").Return(15.0, 12.0, 17.0, 10.0)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Fpdf) {
				pdf.AssertNumberOfCalls(t, "GetMargins", 1)
			},
			func(t *testing.T, x, y, w, h float64) {
				assert.InDelta(t, x, 17.0, 0.1)
				assert.InDelta(t, y, 16.8, 0.1)
				assert.InDelta(t, w, 20.0, 0.1)
				assert.InDelta(t, h, 15.3, 0.1)
			},
		},
		{
			"When cen proportion is less than col, 45 percent",
			26,
			20,
			45.0,
			func() *mocks.Fpdf {
				pdf := &mocks.Fpdf{}
				pdf.On("GetMargins").Return(15.0, 12.0, 17.0, 10.0)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Fpdf) {
				pdf.AssertNumberOfCalls(t, "GetMargins", 1)
			},
			func(t *testing.T, x, y, w, h float64) {
				assert.InDelta(t, x, 22.5, 0.1)
				assert.InDelta(t, y, 21.0, 0.1)
				assert.InDelta(t, w, 9.0, 0.1)
				assert.InDelta(t, h, 6.9, 0.1)
			},
		},
	}

	for _, c := range cases {
		// Arrange
		pdf := c.pdf()

		math := internal.NewMath(pdf)

		// Act
		x, y, w, h := math.GetRectCenterColProperties(c.width, c.height, 20, 25.0, 2, c.percent)

		// Assert
		c.assertPdfCalls(t, pdf)
		c.assertResult(t, x, y, w, h)
	}
}

func TestMath_GetRectNonCenterColProperties(t *testing.T) {
	cases := []struct {
		name           string
		width          float64
		height         float64
		prop           props.Rect
		pdf            func() *mocks.Fpdf
		assertPdfCalls func(t *testing.T, pdf *mocks.Fpdf)
		assertResult   func(t *testing.T, x, y, w, h float64)
	}{
		{
			"When cel proportion is greater than rectangle",
			20,
			26,
			props.Rect{
				Percent: 100,
				Center:  false,
				Left:    0,
				Top:     0,
			},
			func() *mocks.Fpdf {
				pdf := &mocks.Fpdf{}
				pdf.On("GetMargins").Return(12.0, 11.0, 13.0, 15.0)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Fpdf) {
				pdf.AssertNumberOfCalls(t, "GetMargins", 1)
			},
			func(t *testing.T, x, y, w, h float64) {
				assert.InDelta(t, x, 14.0, 0.1)
				assert.InDelta(t, y, 11, 0.1)
				assert.InDelta(t, w, 19.2, 0.1)
				assert.InDelta(t, h, 25.0, 0.1)
			},
		},
		{
			"When cel proportion is greater than rectangle, 45 percent",
			20,
			26,
			props.Rect{
				Percent: 45,
				Center:  false,
				Left:    0,
				Top:     0,
			},
			func() *mocks.Fpdf {
				pdf := &mocks.Fpdf{}
				pdf.On("GetMargins").Return(12.0, 11.0, 13.0, 15.0)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Fpdf) {
				pdf.AssertNumberOfCalls(t, "GetMargins", 1)
			},
			func(t *testing.T, x, y, w, h float64) {
				assert.InDelta(t, x, 14.0, 0.1)
				assert.InDelta(t, y, 11, 0.1)
				assert.InDelta(t, w, 8.6, 0.1)
				assert.InDelta(t, h, 11.25, 0.1)
			},
		},
		{
			"When cel proportion is less than rectangle",
			26,
			20,
			props.Rect{
				Percent: 100,
				Center:  false,
				Left:    0,
				Top:     0,
			},
			func() *mocks.Fpdf {
				pdf := &mocks.Fpdf{}
				pdf.On("GetMargins").Return(15.0, 12.0, 17.0, 10.0)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Fpdf) {
				pdf.AssertNumberOfCalls(t, "GetMargins", 1)
			},
			func(t *testing.T, x, y, w, h float64) {
				assert.InDelta(t, x, 17.0, 0.1)
				assert.InDelta(t, y, 12.0, 0.1)
				assert.InDelta(t, w, 20.0, 0.1)
				assert.InDelta(t, h, 15.3, 0.1)
			},
		},
		{
			"When cel proportion is less than rectangle, 45 percent",
			26,
			20,
			props.Rect{
				Percent: 45,
				Center:  false,
				Left:    0,
				Top:     0,
			},
			func() *mocks.Fpdf {
				pdf := &mocks.Fpdf{}
				pdf.On("GetMargins").Return(15.0, 12.0, 17.0, 10.0)
				return pdf
			},
			func(t *testing.T, pdf *mocks.Fpdf) {
				pdf.AssertNumberOfCalls(t, "GetMargins", 1)
			},
			func(t *testing.T, x, y, w, h float64) {
				assert.InDelta(t, x, 17.0, 0.1)
				assert.InDelta(t, y, 12.0, 0.1)
				assert.InDelta(t, w, 9.0, 0.1)
				assert.InDelta(t, h, 6.9, 0.1)
			},
		},
	}

	for _, c := range cases {
		// Arrange
		pdf := c.pdf()

		math := internal.NewMath(pdf)

		// Act
		x, y, w, h := math.GetRectNonCenterColProperties(c.width, c.height, 20.0, 25.0, 2, c.prop)

		// Assert
		c.assertPdfCalls(t, pdf)
		c.assertResult(t, x, y, w, h)
	}
}

func TestMath_GetCenterCorrection(t *testing.T) {
	// Arrange
	pdf := &mocks.Fpdf{}
	math := internal.NewMath(pdf)

	// Act
	correction := math.GetCenterCorrection(10, 5)

	// Assert
	assert.Equal(t, correction, 2.5)
}
