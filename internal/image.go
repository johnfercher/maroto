package internal

import (
	"bytes"
	"encoding/base64"
	"errors"
	"github.com/google/uuid"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/jung-kurt/gofpdf"
)

// Image is the abstraction which deals of how to add images in a PDF
type Image interface {
	AddFromFile(path string, yColOffset float64, xColOffset float64, colWidth float64, colHeight float64, prop props.Rect) (err error)
	AddFromBase64(b64 string, yColOffset float64, xColOffset float64, colWidth float64, colHeight float64, prop props.Rect, extension consts.Extension) (err error)
}

type image struct {
	pdf  gofpdf.Pdf
	math Math
}

// NewImage create an Image
func NewImage(pdf gofpdf.Pdf, math Math) *image {
	return &image{
		pdf,
		math,
	}
}

// AddFromFile open an image from disk and add to PDF
func (s *image) AddFromFile(path string, yColOffset float64, xColOffset float64, colWidth float64, colHeight float64, prop props.Rect) error {
	info := s.pdf.RegisterImageOptions(path, gofpdf.ImageOptions{
		ReadDpi:   false,
		ImageType: "",
	})

	if info == nil {
		return errors.New("Could not register image options, maybe path/name is wrong")
	}

	s.addImageToPdf(path, info, yColOffset, colWidth, colHeight, xColOffset, prop)
	return nil
}

// AddFromBase64 use a base64 string to add to PDF
func (s *image) AddFromBase64(b64 string, yColOffset float64, xColOffset float64, colWidth float64, colHeight float64, prop props.Rect, extension consts.Extension) error {
	imageId, _ := uuid.NewRandom()

	ss, _ := base64.StdEncoding.DecodeString(b64)

	info := s.pdf.RegisterImageOptionsReader(
		imageId.String(),
		gofpdf.ImageOptions{
			ReadDpi:   false,
			ImageType: string(extension),
		},
		bytes.NewReader(ss),
	)

	if info == nil {
		return errors.New("Could not register image options, maybe path/name is wrong")
	}

	s.addImageToPdf(imageId.String(), info, yColOffset, colWidth, colHeight, xColOffset, prop)
	return nil
}

func (s *image) addImageToPdf(imageLabel string, info *gofpdf.ImageInfoType, marginTop, colWidth, colHeight, xColOffset float64, prop props.Rect) {
	var x, y, w, h float64
	if prop.Center {
		x, y, w, h = s.math.GetRectCenterColProperties(info.Width(), info.Height(), colWidth, colHeight, xColOffset, prop.Percent)
	} else {
		x, y, w, h = s.math.GetRectNonCenterColProperties(info.Width(), info.Height(), colWidth, colHeight, xColOffset, prop)
	}
	s.pdf.Image(imageLabel, x, y+marginTop, w, h, false, "", 0, "")
}
