package internal_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/internal/mocks"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewCode(t *testing.T) {
	code := internal.NewCode(&mocks.Fpdf{}, &mocks.Math{})

	assert.NotNil(t, code)
	assert.Equal(t, fmt.Sprintf("%T", code), "*internal.code")
}

func TestCode_AddBar(t *testing.T) {
	cases := []struct {
		name        string
		code        string
		fpdf        func() *mocks.Fpdf
		math        func() *mocks.Math
		assertFpdf  func(t *testing.T, fpdf *mocks.Fpdf)
		assertMath  func(t *testing.T, math *mocks.Math)
		assertError func(t *testing.T, err error)
		prop        props.Barcode
	}{
		{
			"When everything works",
			"AnyCode",
			func() *mocks.Fpdf {
				fpdf := &mocks.Fpdf{}
				fpdf.On("GetImageInfo", mock.Anything).Return(widthGreaterThanHeightImageInfo())
				fpdf.On("Image", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return fpdf
			},
			func() *mocks.Math {
				math := &mocks.Math{}
				math.On("GetRectNonCenterColProperties", mock.Anything, mock.Anything, mock.Anything, mock.Anything,
					mock.Anything, mock.Anything).Return(100.0, 20.0, 33.0, 0.0)
				return math
			},
			func(t *testing.T, fpdf *mocks.Fpdf) {
				fpdf.AssertNumberOfCalls(t, "GetImageInfo", 1)
				fpdf.AssertCalled(t, "GetImageInfo", "barcode-Code 128AnyCode-1E+022.2E+01")

				fpdf.AssertNumberOfCalls(t, "Image", 1)
				fpdf.AssertCalled(t, "Image", "", 100, 22, 33, 0)
			},
			func(t *testing.T, math *mocks.Math) {
				math.AssertNumberOfCalls(t, "GetRectNonCenterColProperties", 1)
				math.AssertCalled(t, "GetRectNonCenterColProperties", 5, 2, 5, 40, 10, props.Rect{Center: false, Left: 10, Top: 10})
			},
			func(t *testing.T, err error) {
				assert.Nil(t, err)
			},
			props.Barcode{Center: false, Left: 10, Top: 10, Proportion: props.Proportion{Width: 16, Height: 9}},
		},
		{
			"When everything works and code centered",
			"AnyCode",
			func() *mocks.Fpdf {
				fpdf := &mocks.Fpdf{}
				fpdf.On("GetImageInfo", mock.Anything).Return(widthGreaterThanHeightImageInfo())
				fpdf.On("Image", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return fpdf
			},
			func() *mocks.Math {
				math := &mocks.Math{}
				math.On("GetRectCenterColProperties", mock.Anything, mock.Anything, mock.Anything, mock.Anything,
					mock.Anything, mock.Anything).Return(100.0, 20.0, 33.0, 0.0)
				return math
			},
			func(t *testing.T, fpdf *mocks.Fpdf) {
				fpdf.AssertNumberOfCalls(t, "GetImageInfo", 1)
				fpdf.AssertCalled(t, "GetImageInfo", "barcode-Code 128AnyCode-1E+022.2E+01")

				fpdf.AssertNumberOfCalls(t, "Image", 1)
				fpdf.AssertCalled(t, "Image", "", 100, 22, 33, 0)
			},
			func(t *testing.T, math *mocks.Math) {
				math.AssertNumberOfCalls(t, "GetRectCenterColProperties", 1)
				math.AssertCalled(t, "GetRectCenterColProperties", 5, 5, 5, 40, 10, 100)
			},
			func(t *testing.T, err error) {
				assert.Nil(t, err)
			},
			props.Barcode{Center: true, Percent: 100, Proportion: props.Proportion{Width: 1, Height: 1}},
		},
		{
			"When cannot generate QrCode",
			"",
			func() *mocks.Fpdf {
				fpdf := &mocks.Fpdf{}
				fpdf.On("GetImageInfo", mock.Anything).Return(widthGreaterThanHeightImageInfo())
				fpdf.On("Image", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return fpdf
			},
			func() *mocks.Math {
				math := &mocks.Math{}
				math.On("GetRectCenterColProperties", mock.Anything, mock.Anything, mock.Anything, mock.Anything,
					mock.Anything, mock.Anything).Return(100.0, 20.0, 33.0, 0.0)
				return math
			},
			func(t *testing.T, fpdf *mocks.Fpdf) {
				fpdf.AssertNotCalled(t, "GetImageInfo")
				fpdf.AssertNotCalled(t, "Image")
			},
			func(t *testing.T, math *mocks.Math) {
				math.AssertNotCalled(t, "GetRectCenterColProperties")
			},
			func(t *testing.T, err error) {
				assert.NotNil(t, err)
			},
			props.Barcode{Center: true, Percent: 100},
		},
	}

	for _, c := range cases {
		// Arrange
		fpdf := c.fpdf()
		math := c.math()

		code := internal.NewCode(fpdf, math)
		cell := internal.Cell{
			X:      10,
			Y:      2,
			Width:  5,
			Height: 40,
		}

		// Act
		err := code.AddBar(c.code, cell, c.prop)

		// Assert
		c.assertFpdf(t, fpdf)
		c.assertMath(t, math)
		c.assertError(t, err)
	}
}

// nolint:dupl // better this way
func TestCode_AddQr(t *testing.T) {
	cases := []struct {
		name       string
		code       string
		fpdf       func() *mocks.Fpdf
		math       func() *mocks.Math
		assertFpdf func(t *testing.T, fpdf *mocks.Fpdf)
		assertMath func(t *testing.T, math *mocks.Math)
		prop       props.Rect
	}{
		{
			"When everything works",
			"AnyCode",
			func() *mocks.Fpdf {
				fpdf := &mocks.Fpdf{}
				fpdf.On("GetImageInfo", mock.Anything).Return(widthGreaterThanHeightImageInfo())
				fpdf.On("Image", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return fpdf
			},
			func() *mocks.Math {
				math := &mocks.Math{}
				math.On("GetRectCenterColProperties", mock.Anything, mock.Anything, mock.Anything, mock.Anything,
					mock.Anything, mock.Anything).Return(100.0, 20.0, 33.0, 0.0)
				return math
			},
			func(t *testing.T, fpdf *mocks.Fpdf) {
				fpdf.AssertNumberOfCalls(t, "GetImageInfo", 1)
				fpdf.AssertCalled(t, "GetImageInfo", "barcode-QR CodeAnyCode-1E+023E+01")

				fpdf.AssertNumberOfCalls(t, "Image", 1)
				fpdf.AssertCalled(t, "Image", "", 100, 30, 33, 0)
			},
			func(t *testing.T, math *mocks.Math) {
				math.AssertNumberOfCalls(t, "GetRectCenterColProperties", 1)
				math.AssertCalled(t, "GetRectCenterColProperties", 5, 5, 5, 40, 2, 100)
			},
			props.Rect{Center: true, Percent: 100},
		},
		{
			"When everything works not-centered",
			"AnyCode",
			func() *mocks.Fpdf {
				fpdf := &mocks.Fpdf{}
				fpdf.On("GetImageInfo", mock.Anything).Return(widthGreaterThanHeightImageInfo())
				fpdf.On("Image", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return fpdf
			},
			func() *mocks.Math {
				math := &mocks.Math{}
				math.On("GetRectNonCenterColProperties", mock.Anything, mock.Anything, mock.Anything, mock.Anything,
					mock.Anything, mock.Anything).Return(100.0, 20.0, 33.0, 0.0)
				return math
			},
			func(t *testing.T, fpdf *mocks.Fpdf) {
				fpdf.AssertNumberOfCalls(t, "GetImageInfo", 1)
				fpdf.AssertCalled(t, "GetImageInfo", "barcode-QR CodeAnyCode-1E+023E+01")

				fpdf.AssertNumberOfCalls(t, "Image", 1)
				fpdf.AssertCalled(t, "Image", "", 100, 30, 33, 0)
			},
			func(t *testing.T, math *mocks.Math) {
				math.AssertNumberOfCalls(t, "GetRectNonCenterColProperties", 1)
				math.AssertCalled(t, "GetRectNonCenterColProperties", 5, 5, 5, 40, 2, props.Rect{Center: false, Left: 10, Top: 10})
			},
			props.Rect{Center: false, Left: 10, Top: 10},
		},
	}

	for _, c := range cases {
		// Arrange
		fpdf := c.fpdf()
		math := c.math()

		code := internal.NewCode(fpdf, math)
		cell := internal.Cell{
			X:      2,
			Y:      10,
			Width:  5,
			Height: 40,
		}

		// Act
		code.AddQr(c.code, cell, c.prop)

		// Assert
		c.assertFpdf(t, fpdf)
		c.assertMath(t, math)
	}
}

// nolint:dupl // better this way
func TestCode_AddDataMatrix(t *testing.T) {
	cases := []struct {
		name       string
		code       string
		fpdf       func() *mocks.Fpdf
		math       func() *mocks.Math
		assertFpdf func(t *testing.T, fpdf *mocks.Fpdf)
		assertMath func(t *testing.T, math *mocks.Math)
		prop       props.Rect
	}{
		{
			"When everything works",
			"AnyCode",
			func() *mocks.Fpdf {
				fpdf := &mocks.Fpdf{}
				fpdf.On("GetImageInfo", mock.Anything).Return(widthGreaterThanHeightImageInfo())
				fpdf.On("Image", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return fpdf
			},
			func() *mocks.Math {
				math := &mocks.Math{}
				math.On("GetRectCenterColProperties", mock.Anything, mock.Anything, mock.Anything, mock.Anything,
					mock.Anything, mock.Anything).Return(100.0, 20.0, 33.0, 0.0)
				return math
			},
			func(t *testing.T, fpdf *mocks.Fpdf) {
				fpdf.AssertNumberOfCalls(t, "GetImageInfo", 1)
				fpdf.AssertCalled(t, "GetImageInfo", "barcode-DataMatrixAnyCode-1E+023E+01")

				fpdf.AssertNumberOfCalls(t, "Image", 1)
				fpdf.AssertCalled(t, "Image", "", 100, 30, 33, 0)
			},
			func(t *testing.T, math *mocks.Math) {
				math.AssertNumberOfCalls(t, "GetRectCenterColProperties", 1)
				math.AssertCalled(t, "GetRectCenterColProperties", 5, 5, 5, 40, 2, 100)
			},
			props.Rect{Center: true, Percent: 100},
		},
		{
			"When everything works not-centered",
			"AnyCode",
			func() *mocks.Fpdf {
				fpdf := &mocks.Fpdf{}
				fpdf.On("GetImageInfo", mock.Anything).Return(widthGreaterThanHeightImageInfo())
				fpdf.On("Image", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return fpdf
			},
			func() *mocks.Math {
				math := &mocks.Math{}
				math.On("GetRectNonCenterColProperties", mock.Anything, mock.Anything, mock.Anything, mock.Anything,
					mock.Anything, mock.Anything).Return(100.0, 20.0, 33.0, 0.0)
				return math
			},
			func(t *testing.T, fpdf *mocks.Fpdf) {
				fpdf.AssertNumberOfCalls(t, "GetImageInfo", 1)
				fpdf.AssertCalled(t, "GetImageInfo", "barcode-DataMatrixAnyCode-1E+023E+01")

				fpdf.AssertNumberOfCalls(t, "Image", 1)
				fpdf.AssertCalled(t, "Image", "", 100, 30, 33, 0)
			},
			func(t *testing.T, math *mocks.Math) {
				math.AssertNumberOfCalls(t, "GetRectNonCenterColProperties", 1)
				math.AssertCalled(t, "GetRectNonCenterColProperties", 5, 5, 5, 40, 2, props.Rect{Center: false, Left: 10, Top: 10})
			},
			props.Rect{Center: false, Left: 10, Top: 10},
		},
	}

	for _, c := range cases {
		// Arrange
		fpdf := c.fpdf()
		math := c.math()

		code := internal.NewCode(fpdf, math)
		cell := internal.Cell{
			X:      2,
			Y:      10,
			Width:  5,
			Height: 40,
		}

		// Act
		code.AddDataMatrix(c.code, cell, c.prop)

		// Assert
		c.assertFpdf(t, fpdf)
		c.assertMath(t, math)
	}
}
