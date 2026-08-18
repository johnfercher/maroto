package gofpdf_test

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"testing"

	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFromBytes(t *testing.T) {
	t.Parallel()
	t.Run("when extension is not valid, should return error", func(t *testing.T) {
		t.Parallel()
		// Act
		img, err := gofpdf.FromBytes([]byte{1, 2, 3}, "invalid")

		// Assert
		assert.Nil(t, img)
		assert.NotNil(t, err)
	})
	t.Run("when extension is valid, should return image", func(t *testing.T) {
		t.Parallel()
		// Act
		img, err := gofpdf.FromBytes([]byte{1, 2, 3}, extension.Jpg)

		// Assert
		assert.NotNil(t, img)
		assert.Nil(t, err)
	})
	t.Run("when image is 8 bit, should not change anything", func(t *testing.T) {
		t.Parallel()
		// Arrange
		m := image.NewNRGBA(image.Rect(0, 0, 2, 2))
		m.SetNRGBA(0, 0, color.NRGBA{R: 255, G: 128, B: 64, A: 255})

		var input bytes.Buffer
		require.NoError(t, png.Encode(&input, m))
		original := append([]byte(nil), input.Bytes()...)

		// Act
		img, err := gofpdf.FromBytes(input.Bytes(), extension.Png)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, original, img.Bytes)
	})
	t.Run("when image is NRGBA64, should standardize with 8 bit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		m := image.NewNRGBA64(image.Rect(0, 0, 1, 1))
		m.SetNRGBA64(0, 0, color.NRGBA64{R: 65535, G: 32768, B: 16384, A: 65535})

		var input bytes.Buffer
		_ = png.Encode(&input, m)

		// Act
		img, err := gofpdf.FromBytes(input.Bytes(), extension.Png)

		// Assert
		decoded, err := png.Decode(bytes.NewReader(img.Bytes))
		assert.Equal(t, image.Rect(0, 0, 1, 1), decoded.Bounds())
		c := color.NRGBAModel.Convert(decoded.At(0, 0)).(color.NRGBA)
		assert.Equal(t, color.NRGBA{R: 255, G: 128, B: 64, A: 255}, c)
		require.NoError(t, err)
	})
	t.Run("when image is RGBA64, should standardize with 8 bit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		m := image.NewRGBA64(image.Rect(0, 0, 1, 1))
		m.SetRGBA64(0, 0, color.RGBA64{R: 65535, G: 32768, B: 16384, A: 65535})

		var input bytes.Buffer
		_ = png.Encode(&input, m)

		// Act
		img, err := gofpdf.FromBytes(input.Bytes(), extension.Png)

		// Assert
		decoded, err := png.Decode(bytes.NewReader(img.Bytes))
		assert.Equal(t, image.Rect(0, 0, 1, 1), decoded.Bounds())
		c := color.NRGBAModel.Convert(decoded.At(0, 0)).(color.NRGBA)
		assert.Equal(t, color.NRGBA{R: 255, G: 128, B: 64, A: 255}, c)
		require.NoError(t, err)
	})
	t.Run("when image is Gray16, should standardize with 8 bit", func(t *testing.T) {
		t.Parallel()
		// Arrange
		m := image.NewGray16(image.Rect(0, 0, 1, 1))
		m.SetGray16(0, 0, color.Gray16{Y: 32768})

		var input bytes.Buffer
		_ = png.Encode(&input, m)

		// Act
		img, err := gofpdf.FromBytes(input.Bytes(), extension.Png)

		// Assert
		decoded, err := png.Decode(bytes.NewReader(img.Bytes))
		assert.Equal(t, image.Rect(0, 0, 1, 1), decoded.Bounds())
		c := color.NRGBAModel.Convert(decoded.At(0, 0)).(color.NRGBA)
		assert.Equal(t, color.NRGBA{R: 128, G: 128, B: 128, A: 255}, c)
		require.NoError(t, err)
	})
}
