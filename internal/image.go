package internal

import (
	"bytes"
	"encoding/base64"
	"github.com/google/uuid"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/jung-kurt/gofpdf"
)

// Image is the abstraction which deals of how to add images in a PDF
type Image interface {
	AddFromFile(path string, marginTop float64, indexCol float64, qtdCols float64, colHeight float64, prop props.Rect)
	AddFromBase64(b64 string, marginTop float64, indexCol float64, qtdCols float64, colHeight float64, prop props.Rect, extension consts.Extension)
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
func (s *image) AddFromFile(path string, marginTop float64, indexCol float64, qtdCols float64, colHeight float64, prop props.Rect) {
	info := s.pdf.RegisterImageOptions(path, gofpdf.ImageOptions{
		ReadDpi:   false,
		ImageType: "",
	})

	s.addImageToPdf(path, info, marginTop, qtdCols, colHeight, indexCol, prop)
}

// AddFromBase64 use a base64 string to add to PDF
func (s *image) AddFromBase64(b64 string, marginTop float64, indexCol float64, qtdCols float64, colHeight float64, prop props.Rect, extension consts.Extension) {
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

	s.addImageToPdf(imageId.String(), info, marginTop, qtdCols, colHeight, indexCol, prop)
}

func (s *image) addImageToPdf(imageLabel string, info *gofpdf.ImageInfoType, marginTop, qtdCols, colHeight, indexCol float64, prop props.Rect) {
	var x, y, w, h float64
	if prop.Center {
		x, y, w, h = s.math.GetRectCenterColProperties(info.Width(), info.Height(), qtdCols, colHeight, indexCol, prop.Percent)
	} else {
		x, y, w, h = s.math.GetRectNonCenterColProperties(info.Width(), info.Height(), qtdCols, colHeight, indexCol, prop)
	}
	s.pdf.Image(imageLabel, x, y+marginTop, w, h, false, "", 0, "")
}
