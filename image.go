package maroto

import (
	"bytes"
	"encoding/base64"
	"github.com/google/uuid"
	"github.com/jung-kurt/gofpdf"
)

// Abstraction of Image adding features used in Maroto
type Image interface {
	AddFromFile(path string, marginTop float64, indexCol float64, qtdCols float64, colHeight float64, percent float64)
	AddFromBase64(b64 string, marginTop float64, indexCol float64, qtdCols float64, colHeight float64, percent float64, extension Extension)
}

type image struct {
	pdf  gofpdf.Pdf
	math Math
}

// Create an Image adding used in Maroto
func NewImage(pdf gofpdf.Pdf, math Math) Image {
	return &image{
		pdf,
		math,
	}
}

func (self *image) AddFromFile(path string, marginTop float64, indexCol float64, qtdCols float64, colHeight float64, percent float64) {
	info := self.pdf.RegisterImageOptions(path, gofpdf.ImageOptions{
		ReadDpi:   false,
		ImageType: "",
	})

	self.addImageToPdf(path, info, marginTop, qtdCols, colHeight, indexCol, percent)
}

func (self *image) AddFromBase64(b64 string, marginTop float64, indexCol float64, qtdCols float64, colHeight float64, percent float64, extension Extension) {
	imageId, _ := uuid.NewRandom()

	ss, _ := base64.StdEncoding.DecodeString(b64)

	info := self.pdf.RegisterImageOptionsReader(
		imageId.String(),
		gofpdf.ImageOptions{
			ReadDpi:   false,
			ImageType: string(extension),
		},
		bytes.NewReader(ss),
	)

	self.addImageToPdf(imageId.String(), info, marginTop, qtdCols, colHeight, indexCol, percent)
}

func (self *image) addImageToPdf(imageLabel string, info *gofpdf.ImageInfoType, marginTop, qtdCols, colHeight, indexCol, percent float64) {
	x, y, w, h := self.math.GetRectCenterColProperties(info.Width(), info.Height(), qtdCols, colHeight, indexCol, percent)
	self.pdf.Image(imageLabel, x, y+marginTop, w, h, false, "", 0, "")
}
