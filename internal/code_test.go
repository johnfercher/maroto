package internal_test

import (
	"fmt"
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/internal/mocks"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewCode(t *testing.T) {
	code := internal.NewCode(&mocks.Pdf{}, &mocks.Math{})

	assert.NotNil(t, code)
	assert.Equal(t, fmt.Sprintf("%T", code), "*internal.code")
}

func TestCode_AddBar(t *testing.T) {
	cases := []struct {
		name        string
		code        string
		pdf         func() *mocks.Pdf
		math        func() *mocks.Math
		assertPdf   func(t *testing.T, pdf *mocks.Pdf)
		assertMath  func(t *testing.T, math *mocks.Math)
		assertError func(t *testing.T, err error)
		prop        props.Barcode
	}{
		{
			"When everything works",
			"AnyCode",
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("GetImageInfo", mock.Anything).Return(widthGreaterThanHeightImageInfo())
				pdf.On("Image", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return pdf
			},
			func() *mocks.Math {
				math := &mocks.Math{}
				math.On("GetWidthPerCol", mock.Anything).Return(50.0)
				math.On("GetRectNonCenterColProperties", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(100.0, 20.0, 33.0, 0.0)
				return math
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "GetImageInfo", 1)
				pdf.AssertCalled(t, "GetImageInfo", "barcode-Code 128AnyCode-1E+023E+01")

				pdf.AssertNumberOfCalls(t, "Image", 1)
				pdf.AssertCalled(t, "Image", "", 100, 30, 33, 0)
			},
			func(t *testing.T, math *mocks.Math) {
				math.AssertNumberOfCalls(t, "GetWidthPerCol", 1)
				math.AssertCalled(t, "GetWidthPerCol", 5.0)

				math.AssertNumberOfCalls(t, "GetRectNonCenterColProperties", 1)
				math.AssertCalled(t, "GetRectNonCenterColProperties", 50, 0, 5, 40, 2, props.Rect{Center: false, Left: 10, Top: 10})
			},
			func(t *testing.T, err error) {
				assert.Nil(t, err)
			},
			props.Barcode{Center: false, Left: 10, Top: 10, Proportion: props.Proportion{Width: 16, Height: 9}},
		},
		{
			"When everything works and code centered",
			"AnyCode",
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("GetImageInfo", mock.Anything).Return(widthGreaterThanHeightImageInfo())
				pdf.On("Image", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return pdf
			},
			func() *mocks.Math {
				math := &mocks.Math{}
				math.On("GetWidthPerCol", mock.Anything).Return(50.0)
				math.On("GetRectCenterColProperties", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(100.0, 20.0, 33.0, 0.0)
				return math
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "GetImageInfo", 1)
				pdf.AssertCalled(t, "GetImageInfo", "barcode-Code 128AnyCode-1E+023E+01")

				pdf.AssertNumberOfCalls(t, "Image", 1)
				pdf.AssertCalled(t, "Image", "", 100, 30, 33, 0)
			},
			func(t *testing.T, math *mocks.Math) {
				math.AssertNumberOfCalls(t, "GetWidthPerCol", 1)
				math.AssertCalled(t, "GetWidthPerCol", 5.0)

				math.AssertNumberOfCalls(t, "GetRectCenterColProperties", 1)
				math.AssertCalled(t, "GetRectCenterColProperties", 50, 0, 5, 40, 2, 100)
			},
			func(t *testing.T, err error) {
				assert.Nil(t, err)
			},
			props.Barcode{Center: true, Percent: 100},
		},
		{
			"When cannot generate QrCode",
			"",
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("GetImageInfo", mock.Anything).Return(widthGreaterThanHeightImageInfo())
				pdf.On("Image", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return pdf
			},
			func() *mocks.Math {
				math := &mocks.Math{}
				math.On("GetWidthPerCol", mock.Anything).Return(50.0)
				math.On("GetRectCenterColProperties", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(100.0, 20.0, 33.0, 0.0)
				return math
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNotCalled(t, "GetImageInfo")
				pdf.AssertNotCalled(t, "Image")
			},
			func(t *testing.T, math *mocks.Math) {
				math.AssertNotCalled(t, "GetWidthPerCol")
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
		pdf := c.pdf()
		math := c.math()

		code := internal.NewCode(pdf, math)

		// Act
		err := code.AddBar(c.code, 10, 2, 5, 40, c.prop, 0)

		// Assert
		c.assertPdf(t, pdf)
		c.assertMath(t, math)
		c.assertError(t, err)
	}
}

func TestCode_AddQr(t *testing.T) {
	cases := []struct {
		name       string
		code       string
		pdf        func() *mocks.Pdf
		math       func() *mocks.Math
		assertPdf  func(t *testing.T, pdf *mocks.Pdf)
		assertMath func(t *testing.T, math *mocks.Math)
		prop       props.Rect
	}{
		{
			"When everything works",
			"AnyCode",
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("GetImageInfo", mock.Anything).Return(widthGreaterThanHeightImageInfo())
				pdf.On("Image", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return pdf
			},
			func() *mocks.Math {
				math := &mocks.Math{}
				math.On("GetWidthPerCol", mock.Anything).Return(50.0)
				math.On("GetRectCenterColProperties", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(100.0, 20.0, 33.0, 0.0)
				return math
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "GetImageInfo", 1)
				pdf.AssertCalled(t, "GetImageInfo", "barcode-QR CodeAnyCode-1E+023E+01")

				pdf.AssertNumberOfCalls(t, "Image", 1)
				pdf.AssertCalled(t, "Image", "", 100, 30, 33, 0)
			},
			func(t *testing.T, math *mocks.Math) {
				math.AssertNumberOfCalls(t, "GetWidthPerCol", 1)
				math.AssertCalled(t, "GetWidthPerCol", 5.0)

				math.AssertNumberOfCalls(t, "GetRectCenterColProperties", 1)
				math.AssertCalled(t, "GetRectCenterColProperties", 50, 50, 5, 40, 2, 100)
			},
			props.Rect{Center: true, Percent: 100},
		},
		{
			"When everything works not-centered",
			"AnyCode",
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("GetImageInfo", mock.Anything).Return(widthGreaterThanHeightImageInfo())
				pdf.On("Image", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return pdf
			},
			func() *mocks.Math {
				math := &mocks.Math{}
				math.On("GetWidthPerCol", mock.Anything).Return(50.0)
				math.On("GetRectNonCenterColProperties", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(100.0, 20.0, 33.0, 0.0)
				return math
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "GetImageInfo", 1)
				pdf.AssertCalled(t, "GetImageInfo", "barcode-QR CodeAnyCode-1E+023E+01")

				pdf.AssertNumberOfCalls(t, "Image", 1)
				pdf.AssertCalled(t, "Image", "", 100, 30, 33, 0)
			},
			func(t *testing.T, math *mocks.Math) {
				math.AssertNumberOfCalls(t, "GetWidthPerCol", 1)
				math.AssertCalled(t, "GetWidthPerCol", 5.0)

				math.AssertNumberOfCalls(t, "GetRectNonCenterColProperties", 1)
				math.AssertCalled(t, "GetRectNonCenterColProperties", 50, 50, 5, 40, 2, props.Rect{Center: false, Left: 10, Top: 10})
			},
			props.Rect{Center: false, Left: 10, Top: 10},
		},
	}

	for _, c := range cases {
		// Arrange
		pdf := c.pdf()
		math := c.math()

		code := internal.NewCode(pdf, math)

		// Act
		code.AddQr(c.code, 10, 2, 5, 40, c.prop)

		// Assert
		c.assertPdf(t, pdf)
		c.assertMath(t, math)
	}
}
