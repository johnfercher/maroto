package internal

import (
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/jung-kurt/gofpdf"
)

// Math is the abstraction which deals with useful calc
type Math interface {
	GetWidthPerCol(percent float64) float64
	GetRectCenterColProperties(imageWidth float64, imageHeight float64, colWidth float64, colHeight float64, xColOffset float64, percent float64) (x float64, y float64, w float64, h float64)
	GetRectNonCenterColProperties(imageWidth float64, imageHeight float64, colWidth float64, colHeight float64, xColOffset float64, prop props.Rect) (x float64, y float64, w float64, h float64)
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
func (s *math) GetWidthPerCol(percent float64) float64 {
	width, _ := s.pdf.GetPageSize()
	left, _, right, _ := s.pdf.GetMargins()
	return (width - right - left) * percent
}

// GetRectCenterColProperties define Width, Height, X Offset and Y Offset
// to and rectangle (QrCode, Barcode, Image) be centralized inside a cell
func (s *math) GetRectCenterColProperties(imageWidth float64, imageHeight float64, colWidth float64, colHeight float64, xColOffset float64, percent float64) (x float64, y float64, w float64, h float64) {
	percent = percent / 100.0
	left, top, _, _ := s.pdf.GetMargins()

	imageProportion := imageHeight / imageWidth
	celProportion := colHeight / colWidth

	if imageProportion > celProportion {
		newImageWidth := colHeight / imageProportion * percent
		newImageHeight := newImageWidth * imageProportion

		widthCorrection := s.GetCenterCorrection(colWidth, newImageWidth)
		heightCorrection := s.GetCenterCorrection(colHeight, newImageHeight)

		x = xColOffset + left + widthCorrection
		y = top + heightCorrection
		w = newImageWidth
		h = newImageHeight
	} else {
		newImageWidth := colWidth * percent
		newImageHeight := newImageWidth * imageProportion

		widthCorrection := s.GetCenterCorrection(colWidth, newImageWidth)
		heightCorrection := s.GetCenterCorrection(colHeight, newImageHeight)

		x = xColOffset + left + widthCorrection
		y = top + heightCorrection
		w = newImageWidth
		h = newImageHeight
	}

	return
}

// GetRectNonCenterColProperties define Width, Height to and rectangle (QrCode, Barcode, Image) inside a cell
func (s *math) GetRectNonCenterColProperties(imageWidth float64, imageHeight float64, colWidth float64, colHeight float64, xColOffset float64, prop props.Rect) (x float64, y float64, w float64, h float64) {
	percent := prop.Percent / 100.0
	left, top, _, _ := s.pdf.GetMargins()

	imageProportion := imageHeight / imageWidth
	celProportion := colHeight / colWidth

	if imageProportion > celProportion {
		newImageWidth := colHeight / imageProportion * percent
		newImageHeight := newImageWidth * imageProportion

		x = xColOffset + left + prop.Left
		y = top
		w = newImageWidth
		h = newImageHeight
	} else {
		newImageWidth := colWidth * percent
		newImageHeight := newImageWidth * imageProportion

		x = xColOffset + left + prop.Left
		y = top
		w = newImageWidth
		h = newImageHeight
	}

	return
}

// GetCenterCorrection return the correction of space in X or Y to
// centralize a line in relation with another line
func (s *math) GetCenterCorrection(outerSize, innerSize float64) float64 {
	return (outerSize - innerSize) / 2.0
}
