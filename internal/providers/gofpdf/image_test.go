package gofpdf_test

import (
	"bytes"
	"fmt"
	"testing"

	gofpdf2 "github.com/johnfercher/maroto/v2/internal/providers/gofpdf"

	"github.com/johnfercher/maroto/v2/internal/fixture"
	"github.com/johnfercher/maroto/v2/internal/math"
	"github.com/jung-kurt/gofpdf"
	"github.com/stretchr/testify/mock"

	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNewImage(t *testing.T) {
	image := gofpdf2.NewImage(mocks.NewFpdf(t), mocks.NewMath(t))

	assert.NotNil(t, image)
	assert.Equal(t, fmt.Sprintf("%T", image), "*gofpdf.image")
}

func TestImage_Add(t *testing.T) {
	t.Run("when RegisterImageOptionsReader return nil, should return error", func(t *testing.T) {
		// Arrange
		cell := fixture.CellEntity()
		margins := fixture.MarginsEntity()
		rect := fixture.RectProp()
		img := fixture.ImageEntity()
		options := gofpdf.ImageOptions{
			ReadDpi:   false,
			ImageType: string(img.Extension),
		}

		pdf := mocks.NewFpdf(t)
		pdf.EXPECT().RegisterImageOptionsReader(mock.Anything, options, bytes.NewReader(img.Bytes)).Return(nil)

		image := gofpdf2.NewImage(pdf, mocks.NewMath(t))

		// Act
		err := image.Add(&img, &cell, &margins, &rect, img.Extension, true)

		// Assert
		assert.NotNil(t, err)
	})
	t.Run("when prop is not center, should work properly", func(t *testing.T) {
		// Arrange
		cell := fixture.CellEntity()
		margins := fixture.MarginsEntity()
		rect := fixture.RectProp()
		img := fixture.ImageEntity()
		options := gofpdf.ImageOptions{
			ReadDpi:   false,
			ImageType: string(img.Extension),
		}

		pdf := mocks.NewFpdf(t)
		pdf.EXPECT().RegisterImageOptionsReader(mock.Anything, options, bytes.NewReader(img.Bytes)).Return(&gofpdf.ImageInfoType{})
		pdf.EXPECT().Image(mock.Anything, 30.0, 35.0, 98.0, mock.Anything, true, "", 0, "")

		m := math.New()

		image := gofpdf2.NewImage(pdf, m)

		// Act
		err := image.Add(&img, &cell, &margins, &rect, img.Extension, true)

		// Assert
		assert.Nil(t, err)
	})
	t.Run("when prop is center, should work properly", func(t *testing.T) {
		// Arrange
		cell := fixture.CellEntity()
		margins := fixture.MarginsEntity()
		rect := fixture.RectProp()
		rect.Center = true
		img := fixture.ImageEntity()
		options := gofpdf.ImageOptions{
			ReadDpi:   false,
			ImageType: string(img.Extension),
		}

		pdf := mocks.NewFpdf(t)
		pdf.EXPECT().RegisterImageOptionsReader(mock.Anything, options, bytes.NewReader(img.Bytes)).Return(&gofpdf.ImageInfoType{})
		pdf.EXPECT().Image(mock.Anything, 21.0, mock.Anything, 98.0, mock.Anything, true, "", 0, "")

		m := math.New()

		image := gofpdf2.NewImage(pdf, m)

		// Act
		err := image.Add(&img, &cell, &margins, &rect, img.Extension, true)

		// Assert
		assert.Nil(t, err)
	})
}

func TestImage_GetImageInfo(t *testing.T) {
	t.Run("when RegisterImageOptionsReader return nil, should return nil", func(t *testing.T) {
		// Arrange
		img := fixture.ImageEntity()
		options := gofpdf.ImageOptions{
			ReadDpi:   false,
			ImageType: string(img.Extension),
		}

		pdf := mocks.NewFpdf(t)
		pdf.EXPECT().RegisterImageOptionsReader(mock.Anything, options, bytes.NewReader(img.Bytes)).Return(nil)

		image := gofpdf2.NewImage(pdf, mocks.NewMath(t))

		// Act
		info, _ := image.GetImageInfo(&img, img.Extension)

		// Assert
		assert.Nil(t, info)
	})

	t.Run("when RegisterImageOptionsReader return info, should return info", func(t *testing.T) {
		// Arrange
		img := fixture.ImageEntity()
		options := gofpdf.ImageOptions{
			ReadDpi:   false,
			ImageType: string(img.Extension),
		}

		pdf := mocks.NewFpdf(t)
		pdf.EXPECT().RegisterImageOptionsReader(mock.Anything, options, bytes.NewReader(img.Bytes)).Return(&gofpdf.ImageInfoType{})

		image := gofpdf2.NewImage(pdf, mocks.NewMath(t))

		// Act
		info, _ := image.GetImageInfo(&img, img.Extension)

		// Assert
		assert.NotNil(t, info)
	})
}
