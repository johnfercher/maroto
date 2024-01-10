package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/johnfercher/maroto/v2/pkg/core/entity"
)

func TestCell_GetDimensions(t *testing.T) {
	// Arrange
	cell := entity.Cell{
		X:      10,
		Y:      10,
		Width:  100,
		Height: 100,
	}

	// Act
	dimensions := cell.GetDimensions()

	// Assert
	assert.Equal(t, 100.0, dimensions.Width)
	assert.Equal(t, 100.0, dimensions.Height)
}

func TestCell_Copy(t *testing.T) {
	t.Run("copy should return same values", func(t *testing.T) {
		// Arrange
		cell := entity.Cell{
			X:      10,
			Y:      10,
			Width:  100,
			Height: 100,
		}

		// Act
		copyCell := cell.Copy()

		// Assert
		assert.Equal(t, cell.X, copyCell.X)
		assert.Equal(t, cell.Y, copyCell.Y)
		assert.Equal(t, cell.Width, copyCell.Width)
		assert.Equal(t, cell.Height, copyCell.Height)
	})
	t.Run("copy should not allow side-effects", func(t *testing.T) {
		// Arrange
		cell := entity.Cell{
			X:      10,
			Y:      10,
			Width:  100,
			Height: 100,
		}

		// Act
		copyCell := cell.Copy()
		copyCell.X = 15
		copyCell.Y = 15
		copyCell.Width = 90
		copyCell.Height = 90

		// Assert
		assert.Equal(t, 10.0, cell.X)
		assert.Equal(t, 10.0, cell.Y)
		assert.Equal(t, 100.0, cell.Width)
		assert.Equal(t, 100.0, cell.Height)
	})
}

func TestNewRootContext(t *testing.T) {
	// Arrange
	width := 100.0
	height := 300.0
	margins := entity.Margins{
		Left:   10,
		Right:  10,
		Top:    10,
		Bottom: 20,
	}

	// Act
	cell := entity.NewRootContext(width, height, margins)

	// Assert
	assert.Equal(t, 0.0, cell.X)
	assert.Equal(t, 0.0, cell.Y)
	assert.Equal(t, 80.0, cell.Width)
	assert.Equal(t, 270.0, cell.Height)
}
