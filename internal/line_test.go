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

func TestNewLine(t *testing.T) {
	// Act
	line := internal.NewLine(nil)

	// Assert
	assert.NotNil(t, line)
	assert.Equal(t, "*internal.line", fmt.Sprintf("%T", line))
}

func TestLine_Draw(t *testing.T) {
	cases := []struct {
		name            string
		Fpdf            func() *mocks.Fpdf
		assertFpdfCalls func(t *testing.T, Fpdf *mocks.Fpdf)
	}{
		{
			"One call",
			func() *mocks.Fpdf {
				Fpdf := &mocks.Fpdf{}
				Fpdf.On("SetDrawColor", mock.Anything, mock.Anything, mock.Anything).Return(nil)
				Fpdf.On("Line", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

				return Fpdf
			},
			func(t *testing.T, Fpdf *mocks.Fpdf) {
				Fpdf.AssertNumberOfCalls(t, "Line", 1)
				Fpdf.AssertNumberOfCalls(t, "SetDrawColor", 2)

				Fpdf.AssertCalled(t, "Line", 1.0, 10.0, 4.0, 5.0)
				Fpdf.AssertCalled(t, "SetDrawColor", 255, 100, 50)
				Fpdf.AssertCalled(t, "SetDrawColor", 0, 0, 0)
			},
		},
	}

	for _, c := range cases {
		// Arrange
		Fpdf := c.Fpdf()

		line := internal.NewLine(Fpdf)
		cell := internal.Cell{
			X:      1.0,
			Y:      10.0,
			Width:  4.0,
			Height: 5.0,
		}

		// Act
		line.Draw(cell, props.Line{
			Color: color.Color{
				Red:   255,
				Green: 100,
				Blue:  50,
			},
		})

		// Assert
		c.assertFpdfCalls(t, Fpdf)
	}
}
