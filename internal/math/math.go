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
	GetRectCenterColProperties(rectDimensions *config.Dimensions, cellDimensions *config.Dimensions, percent float64) *core.Cell
	GetRectNonCenterColProperties(rectDimensions *config.Dimensions, cellDimensions *config.Dimensions, prop *props.Rect) *core.Cell
	GetCenterCorrection(outerSize, innerSize float64) float64
}

type math struct{}

// New create a Math.
func New() *math {
	return &math{}
}

func (s *math) GetRectCenterColProperties(rectDimensions *config.Dimensions, cellDimensions *config.Dimensions, percent float64) *core.Cell {
	percent /= 100.0

	imageProportion := rectDimensions.Height / rectDimensions.Width
	celProportion := cellDimensions.Height / cellDimensions.Width

	rectCell := &core.Cell{}
	if imageProportion > celProportion {
		newImageWidth := cellDimensions.Height / imageProportion * percent
		newImageHeight := newImageWidth * imageProportion

		widthCorrection := s.GetCenterCorrection(cellDimensions.Width, newImageWidth)
		heightCorrection := s.GetCenterCorrection(cellDimensions.Height, newImageHeight)

		rectCell.X = widthCorrection
		rectCell.Y = heightCorrection
		rectCell.Width = newImageWidth
		rectCell.Height = newImageHeight
	} else {
		newImageWidth := cellDimensions.Width * percent
		newImageHeight := newImageWidth * imageProportion

		widthCorrection := s.GetCenterCorrection(cellDimensions.Width, newImageWidth)
		heightCorrection := s.GetCenterCorrection(cellDimensions.Height, newImageHeight)

		rectCell.X = widthCorrection
		rectCell.Y = heightCorrection
		rectCell.Width = newImageWidth
		rectCell.Height = newImageHeight
	}

	return rectCell
}

// GetRectNonCenterColProperties define Width, Height to and rectangle (QrCode, Barcode, Image) inside a cell.
func (s *math) GetRectNonCenterColProperties(rectDimensions *config.Dimensions, cellDimensions *config.Dimensions, prop *props.Rect) *core.Cell {
	percent := prop.Percent / maxPercent

	imageProportion := rectDimensions.Height / rectDimensions.Width
	celProportion := cellDimensions.Height / cellDimensions.Width

	rectCell := &core.Cell{}
	if imageProportion > celProportion {
		newImageWidth := cellDimensions.Height / imageProportion * percent
		newImageHeight := newImageWidth * imageProportion

		rectCell.X = prop.Left
		rectCell.Y = 0
		rectCell.Width = newImageWidth
		rectCell.Height = newImageHeight
	} else {
		newImageWidth := cellDimensions.Width * percent
		newImageHeight := newImageWidth * imageProportion

		rectCell.X = prop.Left
		rectCell.Y = 0
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
