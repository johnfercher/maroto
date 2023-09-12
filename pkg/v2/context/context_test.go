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
	t.Run("Should create a new root context without dimensions, and with correctly set page dimensions", func(t *testing.T) {
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
		assert.Equal(t, width, ctx.pageDimensions.Width)
		assert.Equal(t, height, ctx.pageDimensions.Height)
		assert.Zerof(t, ctx.Dimensions.Width, "Dimensions.Width should be zero")
		assert.Zerof(t, ctx.Dimensions.Height, "Dimensions.Height should be zero")
		assert.Zerof(t, ctx.Coordinate.X, "Coordinate.X should be zero")
		assert.Zerof(t, ctx.Coordinate.Y, "Coordinate.Y should be zero")
	})
}

func TestContext_MaxWidth(t *testing.T) {
	t.Run("Should return the maximum available space inside margins", func(t *testing.T) {
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
		assert.Equal(t, width-MarginSize-MarginSize, maxWidth)
	})
	t.Run("Should return the maximum available space inside margins when child context is modified", func(t *testing.T) {
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
		childCtx := ctx.WithDimension(100, 100)
		maxWidth := childCtx.MaxWidth()

		// Assert
		assert.Equal(t, width-MarginSize-MarginSize, maxWidth)
	})
}

func TestContext_MaxHeight(t *testing.T) {
	t.Run("Should return the maximum available space inside margins", func(t *testing.T) {
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

		// Act_
		maxHeight := ctx.MaxHeight()

		// Assert
		assert.Equal(t, height-MarginSize-MarginSize, maxHeight)
	})
	t.Run("Should return the maximum available space inside margins when child context is modified", func(t *testing.T) {
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
		childCtx := ctx.WithDimension(100, 100)
		maxHeight := childCtx.MaxHeight()

		// Assert
		assert.Equal(t, height-MarginSize-MarginSize, maxHeight)
	})
}

func TestContext_WithCoordinateOffset(t *testing.T) {
	t.Run("Should return a new context with the correct coordinate offset", func(t *testing.T) {
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
		childCtx := ctx.WithCoordinates(100, 100)

		// Assert
		assert.Equal(t, 100.0, childCtx.Coordinate.X)
		assert.Equal(t, 100.0, childCtx.Coordinate.Y)
	})
	t.Run("Should not modify the base context when creating a child context", func(t *testing.T) {
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
		childCtx := ctx.WithCoordinates(100, 100)

		// Assert
		// region Base Context Assertions
		assert.Equal(t, 0.0, ctx.Coordinate.X)
		assert.Equal(t, 0.0, ctx.Coordinate.Y)
		assert.Equal(t, 0.0, ctx.Dimensions.Width)
		assert.Equal(t, 0.0, ctx.Dimensions.Height)
		assert.Equal(t, width, ctx.pageDimensions.Width)
		assert.Equal(t, height, ctx.pageDimensions.Height)
		assert.Equal(t, margins, ctx.Margins)
		// endregion
		// region Child Context Assertions
		assert.Equal(t, margins, childCtx.Margins)
		assert.Equal(t, width, childCtx.pageDimensions.Width)
		assert.Equal(t, height, childCtx.pageDimensions.Height)
		assert.Equal(t, 100.0, childCtx.Coordinate.X)
		assert.Equal(t, 100.0, childCtx.Coordinate.Y)
		assert.Equal(t, 0.0, childCtx.Dimensions.Width)
		assert.Equal(t, 0.0, childCtx.Dimensions.Height)
		// endregion

	})
}

func TestContext_WithDimension(t *testing.T) {
	t.Run("Should return a new context with the correct dimensions", func(t *testing.T) {
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
		childWidth := 100.0
		childHeight := 100.0

		// Act
		childCtx := ctx.WithDimension(childWidth, childHeight)

		// Assert
		assert.Equal(t, childWidth, childCtx.Dimensions.Width)
		assert.Equal(t, childHeight, childCtx.Dimensions.Height)
	})
	t.Run("Should not modify the base context when creating a child context", func(t *testing.T) {
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
		childWidth := 100.0
		childHeight := 100.0

		// Act
		childCtx := ctx.WithDimension(childWidth, childHeight)

		// Assert
		// region Base Context Assertions
		assert.Equal(t, 0.0, ctx.Coordinate.X)
		assert.Equal(t, 0.0, ctx.Coordinate.Y)
		assert.Equal(t, 0.0, ctx.Dimensions.Width)
		assert.Equal(t, 0.0, ctx.Dimensions.Height)
		assert.Equal(t, width, ctx.pageDimensions.Width)
		assert.Equal(t, height, ctx.pageDimensions.Height)
		assert.Equal(t, margins, ctx.Margins)
		// endregion
		// region Child Context Assertions
		assert.Equal(t, margins, childCtx.Margins)
		assert.Equal(t, width, childCtx.pageDimensions.Width)
		assert.Equal(t, height, childCtx.pageDimensions.Height)
		assert.Equal(t, 0.0, childCtx.Coordinate.X)
		assert.Equal(t, 0.0, childCtx.Coordinate.Y)
		assert.Equal(t, childWidth, childCtx.Dimensions.Width)
		assert.Equal(t, childHeight, childCtx.Dimensions.Height)
		// endregion
	})
}

func TestContext_GetXOffset(t *testing.T) {
	t.Run("Should return the maximum possible offset when smaller than max width", func(t *testing.T) {
		// Arrange
		width := SideSize
		height := SideSize
		xOffset := 100.0

		margins := &Margins{
			Left:   MarginSize,
			Right:  MarginSize,
			Top:    MarginSize,
			Bottom: MarginSize,
		}

		ctx := NewRootContext(width, height, margins)
		ctx = ctx.WithCoordinates(xOffset, 0.0)

		// Act
		result := ctx.GetXOffset()

		// Assert
		assert.Equal(t, xOffset, result)
		assert.Zerof(t, ctx.Coordinate.Y, "Y coordinate should be 0")
	})
	t.Run("Should return parent coordinate and break to the next line when larger than max width", func(t *testing.T) {
		// Arrange
		width := SideSize
		height := SideSize
		xOffset := 100.0

		margins := &Margins{
			Left:   MarginSize,
			Right:  MarginSize,
			Top:    MarginSize,
			Bottom: MarginSize,
		}

		ctx := NewRootContext(width, height, margins).WithDimension(1100, 1100).WithCoordinates(xOffset, 0.0)

		// Act
		result := ctx.GetXOffset()

		// Assert
		assert.Equal(t, xOffset, result)
		assert.Equal(t, float64(LineHeight), ctx.Coordinate.Y)
	})
}

func TestContext_GetYOffset(t *testing.T) {
	t.Run("Should return the maximum possible offset when smaller than max height", func(t *testing.T) {
		// Arrange
		width := SideSize
		height := SideSize
		yOffset := 100.0

		margins := &Margins{
			Left:   MarginSize,
			Right:  MarginSize,
			Top:    MarginSize,
			Bottom: MarginSize,
		}

		ctx := NewRootContext(width, height, margins)
		ctx = ctx.WithCoordinates(0.0, yOffset)

		// Act
		result := ctx.GetYOffset()

		// Assert
		assert.Equal(t, yOffset, result)
		assert.Zerof(t, ctx.Coordinate.X, "X coordinate should be 0")
	})
	t.Run("Should return 0 and break to the next page when larger than max height", func(t *testing.T) {
		// Arrange
		width := SideSize
		height := SideSize
		yOffset := 100.0

		margins := &Margins{
			Left:   MarginSize,
			Right:  MarginSize,
			Top:    MarginSize,
			Bottom: MarginSize,
		}

		ctx := NewRootContext(width, height, margins).WithDimension(1100, 1100).WithCoordinates(0.0, yOffset)

		// Act
		result := ctx.GetYOffset()

		// Assert
		assert.Equal(t, float64(0), result)
		assert.Equal(t, float64(0), ctx.Coordinate.X)
		assert.Equal(t, 1, ctx.CurrentPage)
	})
}
