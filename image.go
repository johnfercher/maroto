package maroto

import (
	"bytes"
	"encoding/base64"
	"github.com/google/uuid"
	"github.com/jung-kurt/gofpdf"
)

type Image interface {
	AddFromFile(path string, marginTop float64, indexCol float64, qtdCols float64, colHeight float64)
	AddFromBase64(b64 string, marginTop float64, indexCol float64, qtdCols float64, colHeight float64, extension Extension)
}

type image struct {
	pdf  gofpdf.Pdf
	math Math
}

func NewImage(pdf gofpdf.Pdf, math Math) Image {
	return &image{
		pdf,
		math,
	}
}

func (i *image) AddFromFile(path string, marginTop float64, indexCol float64, qtdCols float64, colHeight float64) {
	info := i.pdf.RegisterImageOptions(path, gofpdf.ImageOptions{
		ReadDpi:   false,
		ImageType: "",
	})

	x, y, w, h := i.math.GetRectCenterColProperties(info.Width(), info.Height(), qtdCols, colHeight, indexCol)

	i.pdf.ImageOptions(path, x, y+marginTop, w, h, false, gofpdf.ImageOptions{}, 0, "")
}

func (i *image) AddFromBase64(b64 string, marginTop float64, indexCol float64, qtdCols float64, colHeight float64, extension Extension) {
	imageId, _ := uuid.NewRandom()

	ss, _ := base64.StdEncoding.DecodeString(b64)

	info := i.pdf.RegisterImageOptionsReader(
		imageId.String(),
		gofpdf.ImageOptions{
			ReadDpi:   false,
			ImageType: GetExtensionString(extension),
		},
		bytes.NewReader(ss),
	)

	x, y, w, h := i.math.GetRectCenterColProperties(info.Width(), info.Height(), qtdCols, colHeight, indexCol)
	i.pdf.Image(imageId.String(), x, y+marginTop, w, h, false, "", 0, "")
}
