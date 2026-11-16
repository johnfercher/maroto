package gofpdf_test

import (
	"bytes"
	"fmt"
	stdimage "image"
	"image/png"
	"testing"

	"github.com/stretchr/testify/require"

	gofpdf2 "github.com/johnfercher/maroto/v2/internal/providers/gofpdf"

	"github.com/johnfercher/maroto/v2/internal/fixture"
	"github.com/johnfercher/maroto/v2/internal/math"
	"github.com/phpdave11/gofpdf"
	"github.com/stretchr/testify/mock"

	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNewImage(t *testing.T) {
	t.Parallel()
	image := gofpdf2.NewImage(mocks.NewFpdf(t), mocks.NewMath(t))

	assert.NotNil(t, image)
	assert.Equal(t, "*gofpdf.Image", fmt.Sprintf("%T", image))
}

func TestImage_Add(t *testing.T) {
	t.Parallel()
	t.Run("when RegisterImageOptionsReader return nil, should return error", func(t *testing.T) {
		t.Parallel()
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
		t.Parallel()
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
		t.Parallel()
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

func TestImage_Add_Rotation(t *testing.T) {
	t.Parallel()
	// Cell 100x150 at (10, 15), margins 10, rect Left/Top 10 and Percent 98, image 200x100.
	cases := []struct {
		name          string
		angle         int
		center        bool
		expectedAngle float64
		pivotX        float64
		pivotY        float64
		x, y, w, h    float64
	}{
		{
			name: "when rotation angle is 180, should rotate around the image center without swapping sides",
			// footprint 98x49 at (30, 35)
			angle: 180, expectedAngle: 180, pivotX: 79, pivotY: 59.5, x: 30, y: 35, w: 98, h: 49,
		},
		{
			name: "when rotation angle is 90 and not centered, should place the rotated footprint at the rect origin",
			// footprint 73.5x147 at (30, 35)
			angle: 90, expectedAngle: 90, pivotX: 66.75, pivotY: 108.5, x: -6.75, y: 71.75, w: 147, h: 73.5,
		},
		{
			name: "when rotation angle is 270 and centered, should center the rotated footprint in the cell",
			// footprint 73.5x147 at (33.25, 26.5)
			angle: 270, center: true, expectedAngle: 270, pivotX: 70, pivotY: 100, x: -3.5, y: 63.25, w: 147, h: 73.5,
		},
		{
			name: "when rotation angle is -90, should be handled as 270",
			// footprint 73.5x147 at (30, 35)
			angle: -90, expectedAngle: 270, pivotX: 66.75, pivotY: 108.5, x: -6.75, y: 71.75, w: 147, h: 73.5,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			// Arrange
			cell := fixture.CellEntity()
			margins := fixture.MarginsEntity()
			rect := fixture.RectProp()
			rect.RotationAngle = c.angle
			rect.Center = c.center
			rect.MakeValid()
			img := fixture.ImageEntity()
			options := gofpdf.ImageOptions{
				ReadDpi:   false,
				ImageType: string(img.Extension),
			}

			var pivotX, pivotY, x, y, w, h float64

			pdf := mocks.NewFpdf(t)
			pdf.EXPECT().RegisterImageOptionsReader(mock.Anything, options, bytes.NewReader(img.Bytes)).
				Return(imageInfo(t, 200, 100))
			pdf.EXPECT().TransformBegin()
			pdf.EXPECT().TransformRotate(c.expectedAngle, mock.Anything, mock.Anything).
				Run(func(_, px, py float64) { pivotX, pivotY = px, py })
			pdf.EXPECT().ImageOptions(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, false,
				gofpdf.ImageOptions{AllowNegativePosition: true}, 0, "").
				Run(func(_ string, ix, iy, iw, ih float64, _ bool, _ gofpdf.ImageOptions, _ int, _ string) {
					x, y, w, h = ix, iy, iw, ih
				})
			pdf.EXPECT().TransformEnd()

			image := gofpdf2.NewImage(pdf, math.New())

			// Act
			err := image.Add(&img, &cell, &margins, &rect, img.Extension, true)

			// Assert
			assert.Nil(t, err)
			const delta = 1e-9
			assert.InDelta(t, c.pivotX, pivotX, delta)
			assert.InDelta(t, c.pivotY, pivotY, delta)
			assert.InDelta(t, c.x, x, delta)
			assert.InDelta(t, c.y, y, delta)
			assert.InDelta(t, c.w, w, delta)
			assert.InDelta(t, c.h, h, delta)
		})
	}
}

func imageInfo(t *testing.T, width, height int) *gofpdf.ImageInfoType {
	t.Helper()

	var buf bytes.Buffer
	require.NoError(t, png.Encode(&buf, stdimage.NewGray(stdimage.Rect(0, 0, width, height))))

	pdf := gofpdf.New("P", "mm", "A4", "")
	info := pdf.RegisterImageOptionsReader("probe", gofpdf.ImageOptions{ImageType: "png"}, &buf)
	require.NoError(t, pdf.Error())

	return info
}

func TestImage_GetImageInfo(t *testing.T) {
	t.Parallel()
	t.Run("when RegisterImageOptionsReader return nil, should return nil", func(t *testing.T) {
		t.Parallel()
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
		t.Parallel()
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
