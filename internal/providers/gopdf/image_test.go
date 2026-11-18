package gopdf_test

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/png"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/johnfercher/maroto/v2/internal/fixture"
	"github.com/johnfercher/maroto/v2/internal/math"
	"github.com/johnfercher/maroto/v2/internal/providers/gopdf"
	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
)

// pngEntity creates a valid PNG of 72x144 pixels, which is 25.4x50.8 millimeters at 72 DPI.
func pngEntity(t *testing.T) entity.Image {
	t.Helper()

	var buf bytes.Buffer
	err := png.Encode(&buf, image.NewRGBA(image.Rect(0, 0, 72, 144)))
	assert.NoError(t, err)

	return entity.Image{Bytes: buf.Bytes(), Extension: extension.Png}
}

func TestNewImage(t *testing.T) {
	t.Parallel()
	image := gopdf.NewImage(mocks.NewGopdf(t), mocks.NewMath(t))

	assert.NotNil(t, image)
	assert.Equal(t, "*gopdf.Image", fmt.Sprintf("%T", image))
}

func TestImage_Add(t *testing.T) {
	t.Parallel()
	t.Run("when image cannot be decoded, should return error", func(t *testing.T) {
		t.Parallel()
		// Arrange
		cell := fixture.CellEntity()
		margins := fixture.MarginsEntity()
		rect := fixture.RectProp()
		img := fixture.ImageEntity()

		sut := gopdf.NewImage(mocks.NewGopdf(t), mocks.NewMath(t))

		// Act
		err := sut.Add(&img, &cell, &margins, &rect, img.Extension, true)

		// Assert
		assert.NotNil(t, err)
	})
	t.Run("when pdf cannot draw the image, should return error", func(t *testing.T) {
		t.Parallel()
		// Arrange
		cell := fixture.CellEntity()
		margins := fixture.MarginsEntity()
		rect := fixture.RectProp()
		img := pngEntity(t)
		drawErr := errors.New("draw error")

		pdf := mocks.NewGopdf(t)
		pdf.EXPECT().Image(img.Bytes, 30.0, 35.0, mock.Anything, mock.Anything, false).Return(drawErr)

		sut := gopdf.NewImage(pdf, math.New())

		// Act
		err := sut.Add(&img, &cell, &margins, &rect, img.Extension, false)

		// Assert
		assert.ErrorIs(t, err, drawErr)
	})
	t.Run("when prop is not center, should work properly", func(t *testing.T) {
		t.Parallel()
		// Arrange
		cell := fixture.CellEntity()
		margins := fixture.MarginsEntity()
		rect := fixture.RectProp()
		img := pngEntity(t)

		pdf := mocks.NewGopdf(t)
		pdf.EXPECT().Image(img.Bytes, 30.0, 35.0, mock.Anything, mock.Anything, true).Return(nil)

		sut := gopdf.NewImage(pdf, math.New())

		// Act
		err := sut.Add(&img, &cell, &margins, &rect, img.Extension, true)

		// Assert
		assert.Nil(t, err)
	})
	t.Run("when prop is center, should work properly", func(t *testing.T) {
		t.Parallel()
		// Arrange
		cell := fixture.CellEntity()
		margins := fixture.MarginsEntity()
		rect := fixture.RectProp()
		rect.Center = true
		img := pngEntity(t)

		pdf := mocks.NewGopdf(t)
		pdf.EXPECT().Image(img.Bytes, mock.Anything, mock.Anything, mock.Anything, mock.Anything, true).Return(nil)

		sut := gopdf.NewImage(pdf, math.New())

		// Act
		err := sut.Add(&img, &cell, &margins, &rect, img.Extension, true)

		// Assert
		assert.Nil(t, err)
	})
}

func TestImage_GetImageInfo(t *testing.T) {
	t.Parallel()
	t.Run("when image cannot be decoded, should return error", func(t *testing.T) {
		t.Parallel()
		// Arrange
		img := fixture.ImageEntity()
		sut := gopdf.NewImage(mocks.NewGopdf(t), mocks.NewMath(t))

		// Act
		info, err := sut.GetImageInfo(&img, img.Extension)

		// Assert
		assert.Nil(t, info)
		assert.NotNil(t, err)
	})
	t.Run("when image is valid, should return its dimensions at 72 DPI", func(t *testing.T) {
		t.Parallel()
		// Arrange
		img := pngEntity(t)
		sut := gopdf.NewImage(mocks.NewGopdf(t), mocks.NewMath(t))

		// Act
		info, err := sut.GetImageInfo(&img, img.Extension)

		// Assert
		assert.Nil(t, err)
		assert.InDelta(t, 25.4, info.Width, 1e-9)
		assert.InDelta(t, 50.8, info.Height, 1e-9)
	})
}
