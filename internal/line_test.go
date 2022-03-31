package internal_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"

	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/internal/mocks"
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
		props           func() props.Line
		spaceHeight     float64
		assertFpdfCalls func(t *testing.T, Fpdf *mocks.Fpdf)
	}{
		{
			"Solid thin black line",
			func() *mocks.Fpdf {
				Fpdf := &mocks.Fpdf{}
				Fpdf.On("SetDrawColor", mock.Anything, mock.Anything, mock.Anything).Return(nil)
				Fpdf.On("Line", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
				Fpdf.On("SetLineWidth", mock.Anything).Return(nil)

				return Fpdf
			},
			func() props.Line {
				return props.Line{}
			},
			1.0,
			func(t *testing.T, Fpdf *mocks.Fpdf) {
				Fpdf.AssertNumberOfCalls(t, "Line", 1)
				Fpdf.AssertNumberOfCalls(t, "SetDrawColor", 2)
				Fpdf.AssertNumberOfCalls(t, "SetLineWidth", 2)

				Fpdf.AssertCalled(t, "Line", 1.0, 10.0, 40.0, 10.0)
				Fpdf.AssertCalled(t, "SetDrawColor", 0, 0, 0)
				Fpdf.AssertCalled(t, "SetDrawColor", 0, 0, 0)
				Fpdf.AssertCalled(t, "SetLineWidth", 0.1)
				Fpdf.AssertCalled(t, "SetLineWidth", 0.1)
			},
		},
		{
			"Solid thin colored line",
			func() *mocks.Fpdf {
				Fpdf := &mocks.Fpdf{}
				Fpdf.On("SetDrawColor", mock.Anything, mock.Anything, mock.Anything).Return(nil)
				Fpdf.On("Line", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
				Fpdf.On("SetLineWidth", mock.Anything).Return(nil)

				return Fpdf
			},
			func() props.Line {
				return props.Line{
					Color: color.Color{
						Red:   255,
						Green: 100,
						Blue:  50,
					},
				}
			},
			1.0,
			func(t *testing.T, Fpdf *mocks.Fpdf) {
				Fpdf.AssertNumberOfCalls(t, "Line", 1)
				Fpdf.AssertNumberOfCalls(t, "SetDrawColor", 2)
				Fpdf.AssertNumberOfCalls(t, "SetLineWidth", 2)

				Fpdf.AssertCalled(t, "Line", 1.0, 10.0, 40.0, 10.0)
				Fpdf.AssertCalled(t, "SetDrawColor", 255, 100, 50)
				Fpdf.AssertCalled(t, "SetDrawColor", 0, 0, 0)
				Fpdf.AssertCalled(t, "SetLineWidth", 0.1)
				Fpdf.AssertCalled(t, "SetLineWidth", 0.1)
			},
		},
		{
			"Solid thick black line",
			func() *mocks.Fpdf {
				Fpdf := &mocks.Fpdf{}
				Fpdf.On("SetDrawColor", mock.Anything, mock.Anything, mock.Anything).Return(nil)
				Fpdf.On("Line", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
				Fpdf.On("SetLineWidth", mock.Anything).Return(nil)

				return Fpdf
			},
			func() props.Line {
				return props.Line{
					Width: 1.0,
				}
			},
			1.0,
			func(t *testing.T, Fpdf *mocks.Fpdf) {
				Fpdf.AssertNumberOfCalls(t, "Line", 1)
				Fpdf.AssertNumberOfCalls(t, "SetDrawColor", 2)
				Fpdf.AssertNumberOfCalls(t, "SetLineWidth", 2)

				Fpdf.AssertCalled(t, "Line", 1.0, 10.0, 40.0, 10.0)
				Fpdf.AssertCalled(t, "SetDrawColor", 0, 0, 0)
				Fpdf.AssertCalled(t, "SetDrawColor", 0, 0, 0)
				Fpdf.AssertCalled(t, "SetLineWidth", 1.0)
				Fpdf.AssertCalled(t, "SetLineWidth", 0.1)
			},
		},
		// nolint:dupl // better this way
		{
			"Dashed thin black line",
			func() *mocks.Fpdf {
				Fpdf := &mocks.Fpdf{}
				Fpdf.On("SetDrawColor", mock.Anything, mock.Anything, mock.Anything).Return(nil)
				Fpdf.On("Line", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
				Fpdf.On("SetLineWidth", mock.Anything).Return(nil)

				return Fpdf
			},
			func() props.Line {
				return props.Line{
					Style: consts.Dashed,
				}
			},
			1.0,
			func(t *testing.T, Fpdf *mocks.Fpdf) {
				Fpdf.AssertNumberOfCalls(t, "Line", 8)
				Fpdf.AssertNumberOfCalls(t, "SetDrawColor", 2)
				Fpdf.AssertNumberOfCalls(t, "SetLineWidth", 2)

				Fpdf.AssertCalled(t, "SetDrawColor", 0, 0, 0)
				Fpdf.AssertCalled(t, "SetDrawColor", 0, 0, 0)
				Fpdf.AssertCalled(t, "SetLineWidth", 0.1)
				Fpdf.AssertCalled(t, "SetLineWidth", 0.1)
			},
		},
		// nolint:dupl // better this way
		{
			"Dashed thin black line",
			func() *mocks.Fpdf {
				Fpdf := &mocks.Fpdf{}
				Fpdf.On("SetDrawColor", mock.Anything, mock.Anything, mock.Anything).Return(nil)
				Fpdf.On("Line", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
				Fpdf.On("SetLineWidth", mock.Anything).Return(nil)

				return Fpdf
			},
			func() props.Line {
				return props.Line{
					Style: consts.Dotted,
				}
			},
			1.0,
			func(t *testing.T, Fpdf *mocks.Fpdf) {
				Fpdf.AssertNumberOfCalls(t, "Line", 13)
				Fpdf.AssertNumberOfCalls(t, "SetDrawColor", 2)
				Fpdf.AssertNumberOfCalls(t, "SetLineWidth", 2)

				Fpdf.AssertCalled(t, "SetDrawColor", 0, 0, 0)
				Fpdf.AssertCalled(t, "SetDrawColor", 0, 0, 0)
				Fpdf.AssertCalled(t, "SetLineWidth", 0.1)
				Fpdf.AssertCalled(t, "SetLineWidth", 0.1)
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
			Width:  40.0,
			Height: 10.0,
		}

		// Act
		prop := c.props()
		prop.MakeValid(c.spaceHeight)
		line.Draw(cell, prop)

		// Assert
		c.assertFpdfCalls(t, Fpdf)
	}
}
