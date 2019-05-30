package maroto_test

import (
	"fmt"
	"github.com/johnfercher/maroto"
	"github.com/johnfercher/maroto/mocks"
	"github.com/jung-kurt/gofpdf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewImage(t *testing.T) {
	image := maroto.NewImage(&mocks.Pdf{}, &mocks.Math{})

	assert.NotNil(t, image)
	assert.Equal(t, fmt.Sprintf("%T", image), "*maroto.image")
}

func TestImage_AddFromPath(t *testing.T) {
	cases := []struct {
		name            string
		pdf             func() *mocks.Pdf
		math            func() *mocks.Math
		assertPdfCalls  func(t *testing.T, pdf *mocks.Pdf)
		assertMathCalls func(t *testing.T, pdf *mocks.Math)
	}{
		{
			"When image has width greater than height",
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("GetMargins").Return(10.0, 10.0, 10.0, 10.0)
				pdf.On("RegisterImageOptions", mock.Anything, mock.Anything).Return(widthGreaterThanHeightImageInfo())
				pdf.On("ImageOptions", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return pdf
			},
			func() *mocks.Math {
				math := &mocks.Math{}
				math.On("GetWidthPerCol", mock.Anything).Return(100.0)
				return math
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "GetMargins", 1)

				pdf.AssertNumberOfCalls(t, "ImageOptions", 1)
				pdf.AssertCalled(t, "ImageOptions", "AnyPath", 155, 20, 9, 0)

				pdf.AssertNumberOfCalls(t, "RegisterImageOptions", 1)
				pdf.AssertCalled(t, "RegisterImageOptions", "AnyPath", gofpdf.ImageOptions{
					ReadDpi:   false,
					ImageType: "",
				})
			},
			func(t *testing.T, math *mocks.Math) {
				math.AssertNumberOfCalls(t, "GetWidthPerCol", 1)
				math.AssertCalled(t, "GetWidthPerCol", 4.0)
			},
		},
		{
			"When image has height greater than width",
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("GetMargins").Return(10.0, 10.0, 10.0, 10.0)
				pdf.On("RegisterImageOptions", mock.Anything, mock.Anything).Return(heightGreaterThanWidthImageInfo())
				pdf.On("ImageOptions", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return pdf
			},
			func() *mocks.Math {
				math := &mocks.Math{}
				math.On("GetWidthPerCol", mock.Anything).Return(100.0)
				return math
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "GetMargins", 1)

				pdf.AssertNumberOfCalls(t, "ImageOptions", 1)
				pdf.AssertCalled(t, "ImageOptions", "AnyPath", 158, 20, 3, 0)

				pdf.AssertNumberOfCalls(t, "RegisterImageOptions", 1)
				pdf.AssertCalled(t, "RegisterImageOptions", "AnyPath", gofpdf.ImageOptions{
					ReadDpi:   false,
					ImageType: "",
				})
			},
			func(t *testing.T, math *mocks.Math) {
				math.AssertNumberOfCalls(t, "GetWidthPerCol", 1)
				math.AssertCalled(t, "GetWidthPerCol", 4.0)
			},
		},
	}

	for _, c := range cases {
		// Arrange
		pdf := c.pdf()
		math := c.math()

		image := maroto.NewImage(pdf, math)

		// Act
		image.AddFromPath("AnyPath", 10.0, 1.0, 4.0, 5.0)

		// Assert
		c.assertPdfCalls(t, pdf)
		c.assertMathCalls(t, math)
	}
}

func heightGreaterThanWidthImageInfo() *gofpdf.ImageInfoType {
	truePdf := gofpdf.New("P", "mm", "A4", "")

	info := truePdf.RegisterImageOptions("assets/images/shape.jpg", gofpdf.ImageOptions{
		ReadDpi:   false,
		ImageType: "",
	})

	return info
}

func widthGreaterThanHeightImageInfo() *gofpdf.ImageInfoType {
	truePdf := gofpdf.New("P", "mm", "A4", "")

	info := truePdf.RegisterImageOptions("assets/images/marvel.jpg", gofpdf.ImageOptions{
		ReadDpi:   false,
		ImageType: "",
	})

	return info
}
