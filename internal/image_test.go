package internal_test

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/internal/fpdf"
	"github.com/johnfercher/maroto/internal/mocks"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/jung-kurt/gofpdf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewImage(t *testing.T) {
	image := internal.NewImage(&mocks.Fpdf{}, &mocks.Math{})

	assert.NotNil(t, image)
	assert.Equal(t, fmt.Sprintf("%T", image), "*internal.image")
}

func TestImage_AddFromFile(t *testing.T) {
	cases := []struct {
		name            string
		Fpdf            func() *mocks.Fpdf
		math            func() *mocks.Math
		assertFpdfCalls func(t *testing.T, Fpdf *mocks.Fpdf)
		assertMathCalls func(t *testing.T, Fpdf *mocks.Math)
		assertErr       func(t *testing.T, err error)
		props           props.Rect
	}{
		{
			"When cannot load image",
			func() *mocks.Fpdf {
				Fpdf := &mocks.Fpdf{}
				Fpdf.On("RegisterImageOptions", mock.Anything, mock.Anything).Return(nil)
				Fpdf.On("Image", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything,
					mock.Anything, mock.Anything)
				return Fpdf
			},
			func() *mocks.Math {
				math := &mocks.Math{}
				math.On("GetRectCenterColProperties", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(100.0, 20.0, 33.0, 0.0)
				return math
			},
			func(t *testing.T, Fpdf *mocks.Fpdf) {
				Fpdf.AssertNumberOfCalls(t, "Image", 0)

				Fpdf.AssertNumberOfCalls(t, "RegisterImageOptions", 1)
				Fpdf.AssertCalled(t, "RegisterImageOptions", "AnyPath", gofpdf.ImageOptions{
					ReadDpi:   false,
					ImageType: "",
				})
			},
			func(t *testing.T, math *mocks.Math) {
				math.AssertNumberOfCalls(t, "GetRectCenterColProperties", 0)
			},
			func(t *testing.T, err error) {
				assert.NotNil(t, err)
			},
			props.Rect{Center: true, Percent: 100},
		},
		// nolint:dupl // better this way
		{
			"When Image has width greater than height",
			func() *mocks.Fpdf {
				Fpdf := &mocks.Fpdf{}
				Fpdf.On("RegisterImageOptions", mock.Anything, mock.Anything).Return(widthGreaterThanHeightImageInfo())
				Fpdf.On("Image", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything,
					mock.Anything, mock.Anything)
				return Fpdf
			},
			func() *mocks.Math {
				math := &mocks.Math{}
				math.On("GetRectCenterColProperties", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(100.0, 20.0, 33.0, 0.0)
				return math
			},
			func(t *testing.T, Fpdf *mocks.Fpdf) {
				Fpdf.AssertNumberOfCalls(t, "Image", 1)
				Fpdf.AssertCalled(t, "Image", "", 100, 30, 33, 0)

				Fpdf.AssertNumberOfCalls(t, "RegisterImageOptions", 1)
				Fpdf.AssertCalled(t, "RegisterImageOptions", "AnyPath", gofpdf.ImageOptions{
					ReadDpi:   false,
					ImageType: "",
				})
			},
			func(t *testing.T, math *mocks.Math) {
				math.AssertNumberOfCalls(t, "GetRectCenterColProperties", 1)
				math.AssertCalled(t, "GetRectCenterColProperties", 88, 119, 4, 5, 1, 100)
			},
			func(t *testing.T, err error) {
				assert.Nil(t, err)
			},
			props.Rect{Center: true, Percent: 100},
		},
		// nolint:dupl // better this way
		{
			"When Image has height greater than width",
			func() *mocks.Fpdf {
				Fpdf := &mocks.Fpdf{}
				Fpdf.On("RegisterImageOptions", mock.Anything, mock.Anything).Return(heightGreaterThanWidthImageInfo())
				Fpdf.On("Image", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything,
					mock.Anything, mock.Anything)
				return Fpdf
			},
			func() *mocks.Math {
				math := &mocks.Math{}
				math.On("GetRectCenterColProperties", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(100.0, 20.0, 33.0, 0.0)
				return math
			},
			func(t *testing.T, Fpdf *mocks.Fpdf) {
				Fpdf.AssertNumberOfCalls(t, "Image", 1)
				Fpdf.AssertCalled(t, "Image", "", 100, 30, 33, 0)

				Fpdf.AssertNumberOfCalls(t, "RegisterImageOptions", 1)
				Fpdf.AssertCalled(t, "RegisterImageOptions", "AnyPath", gofpdf.ImageOptions{
					ReadDpi:   false,
					ImageType: "",
				})
			},
			func(t *testing.T, math *mocks.Math) {
				math.AssertNumberOfCalls(t, "GetRectCenterColProperties", 1)
				math.AssertCalled(t, "GetRectCenterColProperties", 661, 521, 4, 5, 1, 100)
			},
			func(t *testing.T, err error) {
				assert.Nil(t, err)
			},
			props.Rect{Center: true, Percent: 100},
		},
		{
			"When Image must not be centered",
			func() *mocks.Fpdf {
				Fpdf := &mocks.Fpdf{}
				Fpdf.On("RegisterImageOptions", mock.Anything, mock.Anything).Return(nonCenteredImageInfo())
				Fpdf.On("Image", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything,
					mock.Anything, mock.Anything, mock.Anything)
				return Fpdf
			},
			func() *mocks.Math {
				math := &mocks.Math{}
				math.On("GetRectNonCenterColProperties", mock.Anything, mock.Anything, mock.Anything, mock.Anything,
					mock.Anything, mock.Anything).Return(100.0, 20.0, 33.0, 0.0)
				return math
			},
			func(t *testing.T, Fpdf *mocks.Fpdf) {
				Fpdf.AssertNumberOfCalls(t, "Image", 1)
				Fpdf.AssertCalled(t, "Image", "", 100, 30, 33, 0)

				Fpdf.AssertNumberOfCalls(t, "RegisterImageOptions", 1)
				Fpdf.AssertCalled(t, "RegisterImageOptions", "AnyPath", gofpdf.ImageOptions{
					ReadDpi:   false,
					ImageType: "",
				})
			},
			func(t *testing.T, math *mocks.Math) {
				math.AssertNumberOfCalls(t, "GetRectNonCenterColProperties", 1)
				math.AssertCalled(t, "GetRectNonCenterColProperties", 661, 521, 4, 5, 1, props.Rect{Center: false, Percent: 100})
			},
			func(t *testing.T, err error) {
				assert.Nil(t, err)
			},
			props.Rect{Center: false, Percent: 100},
		},
	}

	for _, c := range cases {
		// Arrange
		Fpdf := c.Fpdf()
		math := c.math()

		image := internal.NewImage(Fpdf, math)
		cell := internal.Cell{
			X:      1.0,
			Y:      10.0,
			Width:  4.0,
			Height: 5.0,
		}

		// Act
		err := image.AddFromFile("AnyPath", cell, c.props)

		// Assert
		c.assertFpdfCalls(t, Fpdf)
		c.assertMathCalls(t, math)
		c.assertErr(t, err)
	}
}

func TestImage_AddFromBase64(t *testing.T) {
	cases := []struct {
		name            string
		Fpdf            func() *mocks.Fpdf
		math            func() *mocks.Math
		assertFpdfCalls func(t *testing.T, Fpdf *mocks.Fpdf)
		assertMathCalls func(t *testing.T, Fpdf *mocks.Math)
		assertErr       func(t *testing.T, err error)
	}{
		{
			"When cannot RegisterImageOptionsReader",
			func() *mocks.Fpdf {
				Fpdf := &mocks.Fpdf{}
				Fpdf.On("RegisterImageOptionsReader", mock.Anything, mock.Anything, mock.Anything).Return(nil)
				Fpdf.On("Image", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything,
					mock.Anything, mock.Anything, mock.Anything)
				return Fpdf
			},
			func() *mocks.Math {
				math := &mocks.Math{}
				math.On("GetRectCenterColProperties", mock.Anything, mock.Anything, mock.Anything, mock.Anything,
					mock.Anything, mock.Anything).Return(100.0, 20.0, 33.0, 0.0)
				return math
			},
			func(t *testing.T, Fpdf *mocks.Fpdf) {
				Fpdf.AssertNumberOfCalls(t, "Image", 0)
				Fpdf.AssertNumberOfCalls(t, "RegisterImageOptionsReader", 1)
				Fpdf.AssertCalled(t, "RegisterImageOptionsReader", "", gofpdf.ImageOptions{
					ReadDpi:   false,
					ImageType: string(consts.Jpg),
				},
					"")
			},
			func(t *testing.T, math *mocks.Math) {
				math.AssertNumberOfCalls(t, "GetRectCenterColProperties", 0)
			},
			func(t *testing.T, err error) {
				assert.NotNil(t, err)
			},
		},
		// nolint:dupl // better this way
		{
			"When ImageHelper has width greater than height",
			func() *mocks.Fpdf {
				Fpdf := &mocks.Fpdf{}
				Fpdf.On("RegisterImageOptionsReader", mock.Anything, mock.Anything, mock.Anything).Return(widthGreaterThanHeightImageInfo())
				Fpdf.On("Image", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything,
					mock.Anything, mock.Anything)
				return Fpdf
			},
			func() *mocks.Math {
				math := &mocks.Math{}
				math.On("GetRectCenterColProperties", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(100.0, 20.0, 33.0, 0.0)
				return math
			},
			func(t *testing.T, Fpdf *mocks.Fpdf) {
				Fpdf.AssertNumberOfCalls(t, "Image", 1)
				Fpdf.AssertCalled(t, "Image", "", 100, 30, 33, 0)

				Fpdf.AssertNumberOfCalls(t, "RegisterImageOptionsReader", 1)
				Fpdf.AssertCalled(t, "RegisterImageOptionsReader", "", gofpdf.ImageOptions{
					ReadDpi:   false,
					ImageType: string(consts.Jpg),
				},
					"")
			},
			func(t *testing.T, math *mocks.Math) {
				math.AssertNumberOfCalls(t, "GetRectCenterColProperties", 1)
				math.AssertCalled(t, "GetRectCenterColProperties", 88, 119, 4, 5, 1, 100)
			},
			func(t *testing.T, err error) {
				assert.Nil(t, err)
			},
		},
		// nolint:dupl // better this way
		{
			"When ImageHelper has height greater than width",
			func() *mocks.Fpdf {
				Fpdf := &mocks.Fpdf{}
				Fpdf.On("RegisterImageOptionsReader", mock.Anything, mock.Anything, mock.Anything).Return(heightGreaterThanWidthImageInfo())
				Fpdf.On("Image", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything,
					mock.Anything, mock.Anything, mock.Anything)
				return Fpdf
			},
			func() *mocks.Math {
				math := &mocks.Math{}
				math.On("GetRectCenterColProperties", mock.Anything, mock.Anything, mock.Anything, mock.Anything,
					mock.Anything, mock.Anything).Return(100.0, 20.0, 33.0, 0.0)
				return math
			},
			func(t *testing.T, Fpdf *mocks.Fpdf) {
				Fpdf.AssertNumberOfCalls(t, "Image", 1)
				Fpdf.AssertCalled(t, "Image", "", 100, 30, 33, 0)

				Fpdf.AssertNumberOfCalls(t, "RegisterImageOptionsReader", 1)
				Fpdf.AssertCalled(t, "RegisterImageOptionsReader", "", gofpdf.ImageOptions{
					ReadDpi:   false,
					ImageType: string(consts.Jpg),
				},
					"")
			},
			func(t *testing.T, math *mocks.Math) {
				math.AssertNumberOfCalls(t, "GetRectCenterColProperties", 1)
				math.AssertCalled(t, "GetRectCenterColProperties", 661, 521, 4, 5, 1, 100)
			},
			func(t *testing.T, err error) {
				assert.Nil(t, err)
			},
		},
	}

	for _, c := range cases {
		// Arrange
		Fpdf := c.Fpdf()
		math := c.math()

		image := internal.NewImage(Fpdf, math)
		base64 := getBase64String()
		cell := internal.Cell{
			X:      1.0,
			Y:      10.0,
			Width:  4.0,
			Height: 5.0,
		}

		// Act
		err := image.AddFromBase64(base64, cell, props.Rect{Center: true, Percent: 100}, consts.Jpg)

		// Assert
		c.assertFpdfCalls(t, Fpdf)
		c.assertMathCalls(t, math)
		c.assertErr(t, err)
	}
}

func heightGreaterThanWidthImageInfo() *gofpdf.ImageInfoType {
	trueFpdf := gofpdf.New("P", "mm", "A4", "")

	info := trueFpdf.RegisterImageOptions("assets/images/biplane.jpg", gofpdf.ImageOptions{
		ReadDpi:   false,
		ImageType: "",
	})

	return info
}

func widthGreaterThanHeightImageInfo() *gofpdf.ImageInfoType {
	trueFpdf := gofpdf.New("P", "mm", "A4", "")
	wrapper := fpdf.NewWrapper(trueFpdf)

	info := wrapper.RegisterImageOptions("assets/images/frontpage.png", gofpdf.ImageOptions{
		ReadDpi:   false,
		ImageType: "",
	})

	return info
}

func nonCenteredImageInfo() *gofpdf.ImageInfoType {
	trueFpdf := gofpdf.New("P", "mm", "A4", "")

	info := trueFpdf.RegisterImageOptions("assets/images/biplane.jpg", gofpdf.ImageOptions{
		ReadDpi:   false,
		ImageType: "",
	})

	return info
}

func getBase64String() string {
	byteSlices, _ := ioutil.ReadFile("assets/images/frontpage.png")
	return base64.StdEncoding.EncodeToString(byteSlices)
}
