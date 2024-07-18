package math

import (
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
)

type math struct{}

// New create a Math.
func New() *math {
	return &math{}
}

// Resize adjusts the internal dimension of an element to occupy a percentage of the available space
//   - inner: The inner dimensions of the element
//   - outer: The outer dimensions of the element
//   - percent: The percentage of the external dimension that can be occupied
//   - justReferenceWidth: Indicates whether resizing should be done only in relation to width or in relation to width and height
func (s *math) Resize(inner *entity.Dimensions, outer *entity.Dimensions, percent float64, justReferenceWidth bool) *entity.Dimensions {
	percent /= 100.0

	innerProportion := inner.Height / inner.Width
	outerProportion := outer.Height / outer.Width

	newInnerWidth := 0.0

	if innerProportion > outerProportion && !justReferenceWidth {
		newInnerWidth = outer.Height / innerProportion * percent
	} else {
		newInnerWidth = outer.Width * percent
	}

	newInnerHeight := newInnerWidth * innerProportion

	if justReferenceWidth && newInnerHeight > outer.Height {
		newInnerWidth = outer.Height / innerProportion * 1
		newInnerHeight = newInnerWidth * innerProportion
	}

	return &entity.Dimensions{Width: newInnerWidth, Height: newInnerHeight}
}

// GetInnerCenterCell define a inner cell formatted inside outer cell centered.
func (s *math) GetInnerCenterCell(inner *entity.Dimensions, outer *entity.Dimensions) *entity.Cell {
	widthCorrection := s.GetCenterCorrection(outer.Width, inner.Width)
	heightCorrection := s.GetCenterCorrection(outer.Height, inner.Height)

	return &entity.Cell{
		X:      widthCorrection,
		Y:      heightCorrection,
		Width:  inner.Width,
		Height: inner.Height,
	}
}

// GetCenterCorrection return the correction of space in X or Y to
// centralize a line in relation with another line.
func (s *math) GetCenterCorrection(outerSize, innerSize float64) float64 {
	return (outerSize - innerSize) / 2.0
}
