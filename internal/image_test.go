package internal_test

import (
	"encoding/base64"
	"fmt"
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/internal/mocks"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/jung-kurt/gofpdf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"testing"
)

func TestNewImage(t *testing.T) {
	image := internal.NewImage(&mocks.Pdf{}, &mocks.Math{})

	assert.NotNil(t, image)
	assert.Equal(t, fmt.Sprintf("%T", image), "*internal.image")
}

func TestImage_AddFromFile(t *testing.T) {
	cases := []struct {
		name            string
		pdf             func() *mocks.Pdf
		math            func() *mocks.Math
		assertPdfCalls  func(t *testing.T, pdf *mocks.Pdf)
		assertMathCalls func(t *testing.T, pdf *mocks.Math)
	}{
		/*{
			"When Image has width greater than height",
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("RegisterImageOptions", mock.Anything, mock.Anything).Return(widthGreaterThanHeightImageInfo())
				pdf.On("Image", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return pdf
			},
			func() *mocks.Math {
				math := &mocks.Math{}
				math.On("GetRectCenterColProperties", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(100.0, 20.0, 33.0, 0.0)
				return math
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "Image", 1)
				pdf.AssertCalled(t, "Image", "", 100, 30, 33, 0)

				pdf.AssertNumberOfCalls(t, "RegisterImageOptions", 1)
				pdf.AssertCalled(t, "RegisterImageOptions", "AnyPath", gofpdf.ImageOptions{
					ReadDpi:   false,
					ImageType: "",
				})
			},
			func(t *testing.T, math *mocks.Math) {
				math.AssertNumberOfCalls(t, "GetRectCenterColProperties", 1)
				math.AssertCalled(t, "GetRectCenterColProperties", 98, 63, 4, 5, 1, 100)
			},
		},*/
		{
			"When Image has height greater than width",
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("RegisterImageOptions", mock.Anything, mock.Anything).Return(heightGreaterThanWidthImageInfo())
				pdf.On("Image", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return pdf
			},
			func() *mocks.Math {
				math := &mocks.Math{}
				math.On("GetRectCenterColProperties", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(100.0, 20.0, 33.0, 0.0)
				return math
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "Image", 1)
				pdf.AssertCalled(t, "Image", "", 100, 30, 33, 0)

				pdf.AssertNumberOfCalls(t, "RegisterImageOptions", 1)
				pdf.AssertCalled(t, "RegisterImageOptions", "AnyPath", gofpdf.ImageOptions{
					ReadDpi:   false,
					ImageType: "",
				})
			},
			func(t *testing.T, math *mocks.Math) {
				math.AssertNumberOfCalls(t, "GetRectCenterColProperties", 1)
				math.AssertCalled(t, "GetRectCenterColProperties", 661, 521, 4, 5, 1, 100)
			},
		},
	}

	for _, c := range cases {
		// Arrange
		pdf := c.pdf()
		math := c.math()

		image := internal.NewImage(pdf, math)

		// Act
		image.AddFromFile("AnyPath", 10.0, 1.0, 4.0, 5.0, props.Rect{Center: true, Percent: 100})

		// Assert
		c.assertPdfCalls(t, pdf)
		c.assertMathCalls(t, math)
	}
}

func TestImage_AddFromBase64(t *testing.T) {
	cases := []struct {
		name            string
		pdf             func() *mocks.Pdf
		math            func() *mocks.Math
		assertPdfCalls  func(t *testing.T, pdf *mocks.Pdf)
		assertMathCalls func(t *testing.T, pdf *mocks.Math)
	}{
		{
			"When ImageHelper has width greater than height",
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("RegisterImageOptionsReader", mock.Anything, mock.Anything, mock.Anything).Return(widthGreaterThanHeightImageInfo())
				pdf.On("Image", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return pdf
			},
			func() *mocks.Math {
				math := &mocks.Math{}
				math.On("GetRectCenterColProperties", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(100.0, 20.0, 33.0, 0.0)
				return math
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "Image", 1)
				pdf.AssertCalled(t, "Image", "", 100, 30, 33, 0)

				pdf.AssertNumberOfCalls(t, "RegisterImageOptionsReader", 1)
				pdf.AssertCalled(t, "RegisterImageOptionsReader", "", gofpdf.ImageOptions{
					ReadDpi:   false,
					ImageType: string(consts.Jpg),
				},
					"")
			},
			func(t *testing.T, math *mocks.Math) {
				math.AssertNumberOfCalls(t, "GetRectCenterColProperties", 1)
				math.AssertCalled(t, "GetRectCenterColProperties", 88, 119, 4, 5, 1, 100)
			},
		},
		{
			"When ImageHelper has height greater than width",
			func() *mocks.Pdf {
				pdf := &mocks.Pdf{}
				pdf.On("RegisterImageOptionsReader", mock.Anything, mock.Anything, mock.Anything).Return(heightGreaterThanWidthImageInfo())
				pdf.On("Image", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				return pdf
			},
			func() *mocks.Math {
				math := &mocks.Math{}
				math.On("GetRectCenterColProperties", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(100.0, 20.0, 33.0, 0.0)
				return math
			},
			func(t *testing.T, pdf *mocks.Pdf) {
				pdf.AssertNumberOfCalls(t, "Image", 1)
				pdf.AssertCalled(t, "Image", "", 100, 30, 33, 0)

				pdf.AssertNumberOfCalls(t, "RegisterImageOptionsReader", 1)
				pdf.AssertCalled(t, "RegisterImageOptionsReader", "", gofpdf.ImageOptions{
					ReadDpi:   false,
					ImageType: string(consts.Jpg),
				},
					"")
			},
			func(t *testing.T, math *mocks.Math) {
				math.AssertNumberOfCalls(t, "GetRectCenterColProperties", 1)
				math.AssertCalled(t, "GetRectCenterColProperties", 661, 521, 4, 5, 1, 100)
			},
		},
	}

	for _, c := range cases {
		// Arrange
		pdf := c.pdf()
		math := c.math()

		image := internal.NewImage(pdf, math)
		base64 := getBase64String()

		// Act
		image.AddFromBase64(base64, 10.0, 1.0, 4.0, 5.0, props.Rect{Center: true, Percent: 100}, consts.Jpg)

		// Assert
		c.assertPdfCalls(t, pdf)
		c.assertMathCalls(t, math)
	}
}

func heightGreaterThanWidthImageInfo() *gofpdf.ImageInfoType {
	truePdf := gofpdf.New("P", "mm", "A4", "")

	info := truePdf.RegisterImageOptions("assets/images/biplane.jpg", gofpdf.ImageOptions{
		ReadDpi:   false,
		ImageType: "",
	})

	return info
}

func widthGreaterThanHeightImageInfo() *gofpdf.ImageInfoType {
	truePdf := gofpdf.New("P", "mm", "A4", "")

	info := truePdf.RegisterImageOptions("assets/images/frontpage.png", gofpdf.ImageOptions{
		ReadDpi:   false,
		ImageType: "",
	})

	return info
}

func getBase64String() string {
	byteSlices, _ := ioutil.ReadFile("assets/images/frontpage.png")
	return base64.StdEncoding.EncodeToString(byteSlices)
}
