package maroto

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
)

type Math interface {
	GetWidthPerCol(qtdCols float64) float64
	GetRectCenterColProperties(imageWidth float64, imageHeight float64, qtdCols float64, colHeight float64, indexCol float64) (x float64, y float64, w float64, h float64)
}

type math struct {
	pdf gofpdf.Pdf
}

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

func (m *math) GetRectCenterColProperties(imageWidth float64, imageHeight float64, qtdCols float64, colHeight float64, indexCol float64) (x float64, y float64, w float64, h float64) {
	width, _ := m.pdf.GetPageSize()
	left, top, right, _ := m.pdf.GetMargins()
	widthPerCol := (width - right - left) / qtdCols

	proportion := imageHeight / imageWidth

	heightForWidth := widthPerCol * proportion

	fmt.Println(heightForWidth)
	if heightForWidth > colHeight {
		widthForColHeight := colHeight / proportion
		widthCorrection := (widthPerCol - widthForColHeight) / 2.0
		x = widthPerCol*indexCol + left + widthCorrection
		y = top
		w = widthForColHeight
		h = 0
	} else {
		x = widthPerCol*indexCol + left
		y = top
		w = widthPerCol
		h = 0
	}

	return
}
