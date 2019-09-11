package internal

import (
	"github.com/jung-kurt/gofpdf"
)

// Math is the abstraction which deals with useful calc
type Math interface {
	GetWidthPerCol(qtdCols float64) float64
	GetRectCenterColProperties(imageWidth float64, imageHeight float64, qtdCols float64, colHeight float64, indexCol float64, percent float64) (x float64, y float64, w float64, h float64)
	GetCenterCorrection(outerSize, innerSize float64) float64
}

type math struct {
	pdf gofpdf.Pdf
}

// NewMath create a Math
func NewMath(pdf gofpdf.Pdf) *math {
	return &math{
		pdf,
	}
}

// GetWidthPerCol return a width which a col will have
// using margins and page size information
func (self *math) GetWidthPerCol(qtdCols float64) float64 {
	width, _ := self.pdf.GetPageSize()
	left, _, right, _ := self.pdf.GetMargins()
	return (width - right - left) / qtdCols
}

// GetRectCenterColProperties define Width, Height, X Offset and Y Offset
// to and rectangle (QrCode, Barcode, Image) be centralized inside a cell
func (self *math) GetRectCenterColProperties(imageWidth float64, imageHeight float64, qtdCols float64, colHeight float64, indexCol float64, percent float64) (x float64, y float64, w float64, h float64) {
	percent = percent / 100.0
	width, _ := self.pdf.GetPageSize()
	left, top, right, _ := self.pdf.GetMargins()
	widthPerCol := (width - right - left) / qtdCols

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

// GetCenterCorrection return the correction of space in X or Y to
// centralize a line in relation with another line
func (self *math) GetCenterCorrection(outerSize, innerSize float64) float64 {
	return (outerSize - innerSize) / 2.0
}
