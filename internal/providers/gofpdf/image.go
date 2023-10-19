package gofpdf

import (
	"bytes"
	"errors"

	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/jung-kurt/gofpdf"

	"github.com/johnfercher/maroto/v2/pkg/core"

	"github.com/google/uuid"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type image struct {
	pdf  Fpdf
	math core.Math
}

// NewImage create an Image.
func NewImage(pdf Fpdf, math core.Math) *image {
	return &image{
		pdf,
		math,
	}
}

// Add use a byte array to add image to PDF.
func (s *image) Add(img *entity.Image, cell *entity.Cell, margins *entity.Margins,
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
		return errors.New("could not register image options, maybe path/name is wrong")
	}

	s.addImageToPdf(imageID.String(), info, cell, margins, prop, flow)
	return nil
}

func (s *image) addImageToPdf(imageLabel string, info *gofpdf.ImageInfoType, cell *entity.Cell, margins *entity.Margins,
	prop *props.Rect, flow bool,
) {
	rectCell := &entity.Cell{}
	dimensions := &entity.Dimensions{Width: info.Width(), Height: info.Height()}

	if prop.Center {
		rectCell = s.math.GetInnerCenterCell(dimensions, cell.GetDimensions(), prop.Percent)
	} else {
		rectCell = s.math.GetInnerNonCenterCell(dimensions, cell.GetDimensions(), prop)
	}
	s.pdf.Image(imageLabel, cell.X+rectCell.X+margins.Left, cell.Y+rectCell.Y+margins.Top,
		rectCell.Width, rectCell.Height, flow, "", 0, "")
}
