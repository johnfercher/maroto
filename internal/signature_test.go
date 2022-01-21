package internal_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/internal/mocks"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewSignature(t *testing.T) {
	signature := internal.NewSignature(&mocks.Fpdf{}, &mocks.Math{}, &mocks.Text{})

	assert.NotNil(t, signature)
	assert.Equal(t, fmt.Sprintf("%T", signature), "*internal.signature")
}

// nolint:dupl // better this way
func TestSignature_AddSpaceFor_DefaultMargins(t *testing.T) {
	// Arrange
	pdf := &mocks.Fpdf{}
	pdf.On("GetMargins").Return(10.0, 10.0, 10.0, 10.0)
	pdf.On("Line", mock.Anything, mock.Anything, mock.Anything, mock.Anything)

	math := &mocks.Math{}
	math.On("GetWidthPerCol", mock.Anything).Return(50.0)

	text := &mocks.Text{}
	text.On("Add", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)

	signature := internal.NewSignature(pdf, math, text)

	// Act
	signature.AddSpaceFor("label", internal.Cell{5, 5, 2, 0}, props.Text{Size: 10.0, Color: color.Color{Red: 20, Green: 20, Blue: 20}})

	// Assert
	pdf.AssertNumberOfCalls(t, "Line", 1)
	pdf.AssertCalled(t, "Line", 19.0, 15.0, 13.0, 15.0)
	text.AssertNumberOfCalls(t, "Add", 1)
	text.AssertCalled(t, "Add", "label", internal.Cell{
		5.0,
		7.0,
		2.0,
		0.0,
	}, props.Text{
		Size: 10.0,
		Color: color.Color{
			Red:   20,
			Green: 20,
			Blue:  20,
		},
	})
}

// nolint:dupl // better this way
func TestSignature_AddSpaceFor_NotDefaultMargins(t *testing.T) {
	// Arrange
	pdf := &mocks.Fpdf{}
	pdf.On("GetMargins").Return(20.0, 10.0, 10.0, 10.0)
	pdf.On("Line", mock.Anything, mock.Anything, mock.Anything, mock.Anything)

	math := &mocks.Math{}
	math.On("GetWidthPerCol", mock.Anything).Return(50.0)

	text := &mocks.Text{}
	text.On("Add", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)

	signature := internal.NewSignature(pdf, math, text)

	// Act
	signature.AddSpaceFor("label", internal.Cell{2, 5, 5, 0}, props.Text{Size: 10.0, Color: color.Color{Red: 20, Green: 20, Blue: 20}})

	// Assert
	pdf.AssertNumberOfCalls(t, "Line", 1)
	pdf.AssertCalled(t, "Line", 26.0, 15.0, 23.0, 15.0)
	text.AssertNumberOfCalls(t, "Add", 1)
	text.AssertCalled(t, "Add", "label", internal.Cell{
		2.0,
		7.0,
		5.0,
		0.0,
	}, props.Text{
		Size: 10.0,
		Color: color.Color{
			Red:   20,
			Green: 20,
			Blue:  20,
		},
	})
}
