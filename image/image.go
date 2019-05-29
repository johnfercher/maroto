package image

import (
	"github.com/johnfercher/maroto/math"
	"github.com/jung-kurt/gofpdf"
)

type Image interface {
	AddFromPath(path string, marginTop float64, indexCol float64, qtdCol float64)
	AddFromBase64(base64 string, x, y, width, height float64)
}

type image struct {
	pdf  gofpdf.Pdf
	math math.Math
}

func NewImage(pdf gofpdf.Pdf, math math.Math) Image {
	return &image{
		pdf,
		math,
	}
}

func (i *image) AddFromPath(path string, marginTop float64, indexCol float64, qtdCol float64) {
	var opt gofpdf.ImageOptions
	widthPerCol := i.math.GetWidthPerCol(qtdCol)

	left, top, _, _ := i.pdf.GetMargins()

	i.pdf.ImageOptions(path, widthPerCol*indexCol+left, marginTop+top, widthPerCol, 0, false, opt, 0, "")

}

func (image) AddFromBase64(base64 string, x, y, width, height float64) {
	panic("implement me")
}
