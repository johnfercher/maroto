package maroto_test

import (
	"fmt"
	"github.com/johnfercher/maroto"
	"github.com/johnfercher/maroto/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewCode(t *testing.T) {
	code := maroto.NewCode(&mocks.Pdf{}, &mocks.Math{})

	assert.NotNil(t, code)
	assert.Equal(t, fmt.Sprintf("%T", code), "*maroto.code")
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
		},
	}

	for _, c := range cases {
		// Arrange
		pdf := c.pdf()
		math := c.math()

		code := maroto.NewCode(pdf, math)

		// Act
		err := code.AddBar(c.code, 10, 2, 5, 40, 100, 0)

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
		},
	}

	for _, c := range cases {
		// Arrange
		pdf := c.pdf()
		math := c.math()

		code := maroto.NewCode(pdf, math)

		// Act
		code.AddQr(c.code, 10, 2, 5, 40, 100)

		// Assert
		c.assertPdf(t, pdf)
		c.assertMath(t, math)
	}
}
