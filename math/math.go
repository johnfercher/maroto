package math

import "github.com/jung-kurt/gofpdf"

type Math interface {
	GetWidthPerCol(qtdCols float64) float64
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
