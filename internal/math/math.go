package math

import (
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

const (
	maxPercent = 100.0
)

// Math is the abstraction which deals with useful calc.
type Math interface {
	GetRectCenterColProperties(rectDimensions *config.Dimensions, cell *core.Cell, margins *config.Margins, percent float64) *core.Cell
	GetRectNonCenterColProperties(rectDimensions *config.Dimensions, cell *core.Cell, margins *config.Margins, prop *props.Rect) *core.Cell
	GetCenterCorrection(outerSize, innerSize float64) float64
}

type math struct{}

// NewMath create a Math.
func NewMath() *math {
	return &math{}
}

func (s *math) GetRectCenterColProperties(rectDimensions *config.Dimensions, cell *core.Cell, margins *config.Margins, percent float64) *core.Cell {
	percent /= 100.0
	left, top := margins.Left, margins.Top

	imageProportion := rectDimensions.Height / rectDimensions.Width
	celProportion := cell.Height / cell.Width

	rectCell := &core.Cell{}
	if imageProportion > celProportion {
		newImageWidth := cell.Height / imageProportion * percent
		newImageHeight := newImageWidth * imageProportion

		widthCorrection := s.GetCenterCorrection(cell.Width, newImageWidth)
		heightCorrection := s.GetCenterCorrection(cell.Height, newImageHeight)

		rectCell.X = cell.X + left + widthCorrection
		rectCell.Y = top + heightCorrection
		rectCell.Width = newImageWidth
		rectCell.Height = newImageHeight
	} else {
		newImageWidth := cell.Width * percent
		newImageHeight := newImageWidth * imageProportion

		widthCorrection := s.GetCenterCorrection(cell.Width, newImageWidth)
		heightCorrection := s.GetCenterCorrection(cell.Height, newImageHeight)

		rectCell.X = cell.X + left + widthCorrection
		rectCell.Y = top + heightCorrection
		rectCell.Width = newImageWidth
		rectCell.Height = newImageHeight
	}

	return rectCell
}

// GetRectNonCenterColProperties define Width, Height to and rectangle (QrCode, Barcode, Image) inside a cell.
func (s *math) GetRectNonCenterColProperties(rectDimensions *config.Dimensions, cell *core.Cell, margins *config.Margins, prop *props.Rect) *core.Cell {
	percent := prop.Percent / maxPercent
	left, top := margins.Left, margins.Top

	imageProportion := rectDimensions.Height / rectDimensions.Width
	celProportion := cell.Height / cell.Width

	rectCell := &core.Cell{}
	if imageProportion > celProportion {
		newImageWidth := cell.Height / imageProportion * percent
		newImageHeight := newImageWidth * imageProportion

		rectCell.X = cell.X + left + prop.Left
		rectCell.Y = top
		rectCell.Width = newImageWidth
		rectCell.Height = newImageHeight
	} else {
		newImageWidth := cell.Width * percent
		newImageHeight := newImageWidth * imageProportion

		rectCell.X = cell.X + left + prop.Left
		rectCell.Y = top
		rectCell.Width = newImageWidth
		rectCell.Height = newImageHeight
	}

	return rectCell
}

// GetCenterCorrection return the correction of space in X or Y to
// centralize a line in relation with another line.
func (s *math) GetCenterCorrection(outerSize, innerSize float64) float64 {
	const divisorToGetHalf = 2.0
	return (outerSize - innerSize) / divisorToGetHalf
}
