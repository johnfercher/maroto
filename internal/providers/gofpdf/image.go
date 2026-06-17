package gofpdf

import (
	"bytes"
	"errors"

	"github.com/google/uuid"
	"github.com/phpdave11/gofpdf"

	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

var ErrCouldNotRegisterImageOptions = errors.New("could not register image options, maybe path/name is wrong")

type Image struct {
	pdf  gofpdfwrapper.Fpdf
	math core.Math
}

// NewImage create an Image.
func NewImage(pdf gofpdfwrapper.Fpdf, math core.Math) *Image {
	return &Image{
		pdf,
		math,
	}
}

// GetImageInfo is responsible for loading the image in PDF and returning its information.
func (s *Image) GetImageInfo(img *entity.Image, extension extension.Type) (*gofpdf.ImageInfoType, uuid.UUID) {
	imageID, _ := uuid.NewRandom()

	info := s.pdf.RegisterImageOptionsReader(
		imageID.String(),
		gofpdf.ImageOptions{
			ReadDpi:   false,
			ImageType: string(extension),
		},
		bytes.NewReader(img.Bytes),
	)
	return info, imageID
}

// Add use a byte array to add image to PDF.
func (s *Image) Add(img *entity.Image, cell *entity.Cell, margins *entity.Margins,
	prop *props.Rect, extension extension.Type, flow bool,
) error {
	imageID, _ := uuid.NewRandom()

	info := s.pdf.RegisterImageOptionsReader(
		imageID.String(),
		gofpdf.ImageOptions{
			ReadDpi:   false,
			ImageType: string(extension),
		},
		bytes.NewReader(img.Bytes),
	)

	if info == nil {
		return ErrCouldNotRegisterImageOptions
	}

	s.addImageToPdf(imageID.String(), info, cell, margins, prop, flow)
	return nil
}

func (s *Image) addImageToPdf(imageLabel string, info *gofpdf.ImageInfoType, cell *entity.Cell, margins *entity.Margins,
	prop *props.Rect, flow bool,
) {
	dimensions := s.math.Resize(&entity.Dimensions{
		Width:  info.Width(),
		Height: info.Height(),
	}, cell.GetDimensions(), prop.Percent, prop.JustReferenceWidth)

	rectCell := &entity.Cell{X: prop.Left, Y: prop.Top, Width: dimensions.Width, Height: dimensions.Height}

	if prop.Center {
		rectCell = s.math.GetInnerCenterCell(dimensions, cell.GetDimensions())
	}

	x := cell.X + rectCell.X + margins.Left
	y := cell.Y + rectCell.Y + margins.Top
	w := rectCell.Width
	h := rectCell.Height

	if prop.RotationAngle != 0 {
		s.addRotatedImageToPdf(imageLabel, info, cell, prop, x, y, w, h)
		return
	}

	s.pdf.Image(imageLabel, x, y, w, h, flow, "", 0, "")
}

// addRotatedImageToPdf renders an image rotated by prop.RotationAngle degrees using
// PDF matrix transformations. For 90 and 270 degrees the bounding box is recomputed
// with swapped width/height so the rotated image keeps fitting inside the cell.
func (s *Image) addRotatedImageToPdf(imageLabel string, info *gofpdf.ImageInfoType, cell *entity.Cell,
	prop *props.Rect, x, y, w, h float64,
) {
	xDelta := 0.0
	yDelta := 0.0
	finalW := w
	finalH := h

	if prop.RotationAngle == 90 || prop.RotationAngle == 270 {
		rotated := s.math.Resize(&entity.Dimensions{
			Width:  info.Height(),
			Height: info.Width(),
		}, cell.GetDimensions(), prop.Percent, prop.JustReferenceWidth)

		xDelta = (w - rotated.Height) / 2
		yDelta = (h - rotated.Width) / 2
		finalW = rotated.Height
		finalH = rotated.Width
	}

	centerX := x + w/2
	centerY := y + h/2

	s.pdf.TransformBegin()
	s.pdf.TransformRotate(float64(prop.RotationAngle), centerX, centerY)
	s.pdf.ImageOptions(imageLabel, x+xDelta, y+yDelta, finalW, finalH, false, gofpdf.ImageOptions{
		AllowNegativePosition: true,
	}, 0, "")
	s.pdf.TransformEnd()
}
