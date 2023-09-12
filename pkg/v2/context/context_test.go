package context

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	SideSize   = 1100.0
	MarginSize = 50.0
)

func TestNewRootContext(t *testing.T) {
	// Arrange
	width := SideSize
	height := SideSize

	margins := &Margins{
		Left:   MarginSize,
		Right:  MarginSize,
		Top:    MarginSize,
		Bottom: MarginSize,
	}

	// Act
	ctx := NewRootContext(width, height, margins)

	// Assert
	assert.NotNil(t, ctx)
}

func TestContext_MaxHeight(t *testing.T) {
	// Arrange
	width := SideSize
	height := SideSize

	margins := &Margins{
		Left:   MarginSize,
		Right:  MarginSize,
		Top:    MarginSize,
		Bottom: MarginSize,
	}

	ctx := NewRootContext(width, height, margins)

	// Act
	maxHeight := ctx.MaxHeight()

	// Assert
	assert.Equal(t, 1000.0, maxHeight)
}

func TestContext_MaxWidth(t *testing.T) {
	// Arrange
	width := SideSize
	height := SideSize

	margins := &Margins{
		Left:   MarginSize,
		Right:  MarginSize,
		Top:    MarginSize,
		Bottom: MarginSize,
	}

	ctx := NewRootContext(width, height, margins)

	// Act
	maxWidth := ctx.MaxWidth()

	// Assert
	assert.Equal(t, 1000.0, maxWidth)
}

func TestContext_GetX(t *testing.T) {
	// Arrange
	width := SideSize
	height := SideSize

	margins := &Margins{
		Left:   MarginSize,
		Right:  MarginSize,
		Top:    MarginSize,
		Bottom: MarginSize,
	}

	ctx := NewRootContext(width, height, margins)

	// Act
	x := ctx.GetX()

	// Assert
	assert.Equal(t, 0.0, x)
}
