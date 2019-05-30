package maroto

import (
	"github.com/jung-kurt/gofpdf"
)

type Image interface {
	AddFromPath(path string, marginTop float64, indexCol float64, qtdCol float64, colHeight float64)
	AddFromBase64(base64 string, x, y, width, height float64)
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

func (i *image) AddFromPath(path string, marginTop float64, indexCol float64, qtdCol float64, colHeight float64) {
	widthPerCol := i.math.GetWidthPerCol(qtdCol)

	left, top, _, _ := i.pdf.GetMargins()

	i.pdf.RegisterImage(path, "")
	info := i.pdf.GetImageInfo(path)

	height := info.Height()
	width := info.Width()

	proportion := height / width

	heightForWidth := widthPerCol * proportion

	if heightForWidth > colHeight {
		widthForColHeight := colHeight / proportion
		widthCorrection := (widthPerCol - widthForColHeight) / 2.0
		i.pdf.ImageOptions(path, widthPerCol*indexCol+left+widthCorrection, marginTop+top, widthForColHeight, 0, false, gofpdf.ImageOptions{}, 0, "")
	} else {
		i.pdf.ImageOptions(path, widthPerCol*indexCol+left, marginTop+top, widthPerCol, 0, false, gofpdf.ImageOptions{}, 0, "")
	}
}

func (image) AddFromBase64(base64 string, x, y, width, height float64) {
	panic("implement me")
}
