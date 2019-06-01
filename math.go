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

// Create a math instance with useful calcs used in Maroto
func NewMath(pdf gofpdf.Pdf) Math {
	return &math{
		pdf,
	}
}

func (m *math) GetWidthPerCol(qtdCols float64) float64 {
	width, _ := m.pdf.GetPageSize()
	left, _, right, _ := m.pdf.GetMargins()
	return (width - right - left) / qtdCols
}

func (m *math) GetRectCenterColProperties(imageWidth float64, imageHeight float64, qtdCols float64, colHeight float64, indexCol float64, percent float64) (x float64, y float64, w float64, h float64) {
	percent = percent / 100.0
	width, _ := m.pdf.GetPageSize()
	left, top, right, _ := m.pdf.GetMargins()
	widthPerCol := ((width - right - left) / qtdCols)

	proportion := imageHeight / imageWidth

	newImageHeight := widthPerCol * proportion * percent
	newImageWidth := widthPerCol * percent

	if newImageHeight > colHeight {
		newImageWidth := colHeight / proportion * percent
		newImageHeight := newImageWidth * proportion

		widthCorrection := m.GetCenterCorrection(widthPerCol, newImageWidth)
		heightCorrection := m.GetCenterCorrection(colHeight, newImageHeight)

		x = (widthPerCol * indexCol) + left + widthCorrection
		y = top + heightCorrection
		w = newImageWidth
		h = newImageHeight
	} else {
		widthCorrection := m.GetCenterCorrection(widthPerCol, newImageWidth)
		heightCorrection := m.GetCenterCorrection(colHeight, newImageHeight)

		x = (widthPerCol * indexCol) + left + widthCorrection
		y = top + heightCorrection
		w = newImageWidth
		h = newImageHeight
	}

	return
}

func (m *math) GetCenterCorrection(outerSize, innerSize float64) float64 {
	return (outerSize - innerSize) / 2.0
}
