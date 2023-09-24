package internal

import (
	"github.com/johnfercher/maroto/v2/internal/fpdf"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

const (
	maxPercent = 100.0
)

// Math is the abstraction which deals with useful calc.
type Math interface {
	GetRectCenterColProperties(dimensions *config.Dimensions, cell *core.Cell, percent float64) (x float64, y float64, w float64, h float64)
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

func (s *math) GetRectCenterColProperties(dimensions *config.Dimensions, cell *core.Cell, percent float64,
) (x float64, y float64, w float64, h float64) {
	percent /= 100.0
	left, top, _, _ := s.pdf.GetMargins()

	imageProportion := dimensions.Height / dimensions.Width
	celProportion := cell.Height / cell.Width

	if imageProportion > celProportion {
		newImageWidth := cell.Height / imageProportion * percent
		newImageHeight := newImageWidth * imageProportion

		widthCorrection := s.GetCenterCorrection(cell.Width, newImageWidth)
		heightCorrection := s.GetCenterCorrection(cell.Height, newImageHeight)

		x = cell.X + left + widthCorrection
		y = top + heightCorrection
		w = newImageWidth
		h = newImageHeight
	} else {
		newImageWidth := cell.Width * percent
		newImageHeight := newImageWidth * imageProportion

		widthCorrection := s.GetCenterCorrection(cell.Width, newImageWidth)
		heightCorrection := s.GetCenterCorrection(cell.Height, newImageHeight)

		x = cell.X + left + widthCorrection
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
