package internal

import (
	"github.com/johnfercher/maroto/internal/fpdf"
	"github.com/johnfercher/maroto/pkg/props"
)

const (
	maxPercent = 100.0
)

// Math is the abstraction which deals with useful calc.
type Math interface {
	GetRectCenterColProperties(imageWidth float64, imageHeight float64, colWidth float64, colHeight float64,
		xColOffset float64, percent float64) (x float64, y float64, w float64, h float64)
	GetRectNonCenterColProperties(imageWidth float64, imageHeight float64, colWidth float64, colHeight float64,
		xColOffset float64, prop props.Rect) (x float64, y float64, w float64, h float64)
	GetCenterCorrection(outerSize, innerSize float64) float64
}

type math struct {
	pdf fpdf.Fpdf
}

// NewMath create a Math.
func NewMath(pdf fpdf.Fpdf) *math {
	return &math{
		pdf,
	}
}

// GetRectCenterColProperties define Width, Height, X Offset and Y Offset
// to and rectangle (QrCode, Barcode, Image) be centralized inside a cell.
func (s *math) GetRectCenterColProperties(imageWidth float64, imageHeight float64, colWidth float64, colHeight float64,
	xColOffset float64, percent float64,
) (x float64, y float64, w float64, h float64) {
	percent /= 100.0
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

	return x, y, w, h
}

// GetRectNonCenterColProperties define Width, Height to and rectangle (QrCode, Barcode, Image) inside a cell.
func (s *math) GetRectNonCenterColProperties(imageWidth float64, imageHeight float64, colWidth float64, colHeight float64,
	xColOffset float64, prop props.Rect,
) (x float64, y float64, w float64, h float64) {
	percent := prop.Percent / maxPercent
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
// centralize a line in relation with another line.
func (s *math) GetCenterCorrection(outerSize, innerSize float64) float64 {
	const divisorToGetHalf = 2.0
	return (outerSize - innerSize) / divisorToGetHalf
}
