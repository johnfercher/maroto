package maroto

import (
	"github.com/jung-kurt/gofpdf"
)

// Abstraction of useful calcs used in Maroto
type Math interface {
	GetWidthPerCol(qtdCols float64) float64
	GetRectCenterColProperties(imageWidth float64, imageHeight float64, qtdCols float64, colHeight float64, indexCol float64, percent float64) (x float64, y float64, w float64, h float64)
	GetCenterCorrection(outerSize, innerSize float64) float64
}

type math struct {
	pdf gofpdf.Pdf
}

// Create a Math instance with useful calcs used in Maroto
func NewMath(pdf gofpdf.Pdf) Math {
	return &math{
		pdf,
	}
}

func (self *math) GetWidthPerCol(qtdCols float64) float64 {
	width, _ := self.pdf.GetPageSize()
	left, _, right, _ := self.pdf.GetMargins()
	return (width - right - left) / qtdCols
}

func (self *math) GetRectCenterColProperties(imageWidth float64, imageHeight float64, qtdCols float64, colHeight float64, indexCol float64, percent float64) (x float64, y float64, w float64, h float64) {
	percent = percent / 100.0
	width, _ := self.pdf.GetPageSize()
	left, top, right, _ := self.pdf.GetMargins()
	widthPerCol := ((width - right - left) / qtdCols)

	proportion := imageHeight / imageWidth

	newImageHeight := widthPerCol * proportion * percent
	newImageWidth := widthPerCol * percent

	if newImageHeight > colHeight {
		newImageWidth := colHeight / proportion * percent
		newImageHeight := newImageWidth * proportion

		widthCorrection := self.GetCenterCorrection(widthPerCol, newImageWidth)
		heightCorrection := self.GetCenterCorrection(colHeight, newImageHeight)

		x = (widthPerCol * indexCol) + left + widthCorrection
		y = top + heightCorrection
		w = newImageWidth
		h = newImageHeight
	} else {
		widthCorrection := self.GetCenterCorrection(widthPerCol, newImageWidth)
		heightCorrection := self.GetCenterCorrection(colHeight, newImageHeight)

		x = (widthPerCol * indexCol) + left + widthCorrection
		y = top + heightCorrection
		w = newImageWidth
		h = newImageHeight
	}

	return
}

func (self *math) GetCenterCorrection(outerSize, innerSize float64) float64 {
	return (outerSize - innerSize) / 2.0
}
