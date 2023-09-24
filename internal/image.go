package internal

import (
	"bytes"
	"encoding/base64"
	"errors"

	"github.com/johnfercher/maroto/v2/internal/math"
	"github.com/johnfercher/maroto/v2/pkg/config"

	"github.com/johnfercher/maroto/v2/pkg/core"

	"github.com/google/uuid"
	"github.com/johnfercher/maroto/v2/internal/fpdf"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/jung-kurt/gofpdf"
)

// Image is the abstraction which deals of how to add images in a PDF.
type Image interface {
	AddFromFile(path string, cell *core.Cell, margins *config.Margins, prop *props.Rect) (err error)
	AddFromBase64(stringBase64 string, cell *core.Cell, margins *config.Margins, prop *props.Rect, extension extension.Type) (err error)
}

type image struct {
	pdf  fpdf.Fpdf
	math math.Math
}

// NewImage create an Image.
func NewImage(pdf fpdf.Fpdf, math math.Math) *image {
	return &image{
		pdf,
		math,
	}
}

// AddFromFile open an image from disk and add to PDF.
func (s *image) AddFromFile(path string, cell *core.Cell, margins *config.Margins, prop *props.Rect) error {
	info := s.pdf.RegisterImageOptions(path, gofpdf.ImageOptions{
		ReadDpi:   false,
		ImageType: "",
	})

	if info == nil {
		return errors.New("could not register image options, maybe path/name is wrong")
	}

	s.addImageToPdf(path, info, cell, margins, prop)
	return nil
}

// AddFromBase64 use a base64 string to add to PDF.
func (s *image) AddFromBase64(stringBase64 string, cell *core.Cell, margins *config.Margins,
	prop *props.Rect, extension extension.Type,
) error {
	imageID, _ := uuid.NewRandom()

	ss, _ := base64.StdEncoding.DecodeString(stringBase64)

	info := s.pdf.RegisterImageOptionsReader(
		imageID.String(),
		gofpdf.ImageOptions{
			ReadDpi:   false,
			ImageType: string(extension),
		},
		bytes.NewReader(ss),
	)

	if info == nil {
		return errors.New("could not register image options, maybe path/name is wrong")
	}

	s.addImageToPdf(imageID.String(), info, cell, margins, prop)
	return nil
}

func (s *image) addImageToPdf(imageLabel string, info *gofpdf.ImageInfoType, cell *core.Cell, margins *config.Margins, prop *props.Rect) {
	rectCell := &core.Cell{}
	dimensions := &config.Dimensions{Width: info.Width(), Height: info.Height()}

	if prop.Center {
		rectCell = s.math.GetInnerCenterCell(dimensions, cell.GetDimensions(), prop.Percent)
	} else {
		rectCell = s.math.GetInnerNonCenterCell(dimensions, cell.GetDimensions(), prop)
	}
	s.pdf.Image(imageLabel, cell.X+rectCell.X+margins.Left, cell.Y+rectCell.Y+prop.Top+margins.Top,
		rectCell.Width, rectCell.Height, false, "", 0, "")
}
