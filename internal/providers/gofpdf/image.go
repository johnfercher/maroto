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
	// footprint is the area the image occupies on the page, so a quarter turn swaps its sides.
	footprint := &entity.Dimensions{Width: info.Width(), Height: info.Height()}
	if prop.IsQuarterTurn() {
		footprint = &entity.Dimensions{Width: info.Height(), Height: info.Width()}
	}

	dimensions := s.math.Resize(footprint, cell.GetDimensions(), prop.Percent, prop.JustReferenceWidth)

	rectCell := &entity.Cell{X: prop.Left, Y: prop.Top, Width: dimensions.Width, Height: dimensions.Height}

	if prop.Center {
		rectCell = s.math.GetInnerCenterCell(dimensions, cell.GetDimensions())
	}

	x := cell.X + rectCell.X + margins.Left
	y := cell.Y + rectCell.Y + margins.Top
	w := rectCell.Width
	h := rectCell.Height

	if prop.RotationAngle == 0 {
		s.pdf.Image(imageLabel, x, y, w, h, flow, "", 0, "")
		return
	}

	imageWidth, imageHeight := w, h
	if prop.IsQuarterTurn() {
		imageWidth, imageHeight = h, w
	}

	s.addRotatedImageToPdf(imageLabel, float64(prop.RotationAngle), x+w/2, y+h/2, imageWidth, imageHeight)
}

// addRotatedImageToPdf draws an unrotated image of the given size centered on (centerX, centerY)
// and rotates it around that point using PDF matrix transformations.
func (s *Image) addRotatedImageToPdf(imageLabel string, angle, centerX, centerY, width, height float64) {
	s.pdf.TransformBegin()
	s.pdf.TransformRotate(angle, centerX, centerY)
	s.pdf.ImageOptions(imageLabel, centerX-width/2, centerY-height/2, width, height, false, gofpdf.ImageOptions{
		AllowNegativePosition: true,
	}, 0, "")
	s.pdf.TransformEnd()
}
