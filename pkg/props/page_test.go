package props_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/internal/fixture"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/breakline"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/stretchr/testify/assert"
)

func TestPlace_IsValid(t *testing.T) {
	t.Run("when place is left_top, should return valid", func(t *testing.T) {
		// Arrange
		sut := props.LeftTop

		// Act/Assert
		assert.True(t, sut.IsValid())
	})
	t.Run("when place is top, should return valid", func(t *testing.T) {
		// Arrange
		sut := props.Top

		// Act/Assert
		assert.True(t, sut.IsValid())
	})
	t.Run("when place is right_top, should return valid", func(t *testing.T) {
		// Arrange
		sut := props.RightTop

		// Act/Assert
		assert.True(t, sut.IsValid())
	})
	t.Run("when place is left_bottom, should return valid", func(t *testing.T) {
		// Arrange
		sut := props.LeftBottom

		// Act/Assert
		assert.True(t, sut.IsValid())
	})
	t.Run("when place is bottom, should return valid", func(t *testing.T) {
		// Arrange
		sut := props.Bottom

		// Act/Assert
		assert.True(t, sut.IsValid())
	})
	t.Run("when place is right_bottom, should return valid", func(t *testing.T) {
		// Arrange
		sut := props.RightBottom

		// Act/Assert
		assert.True(t, sut.IsValid())
	})
	t.Run("when place is invalid should return invalid", func(t *testing.T) {
		// Arrange
		sut := props.Place("invalid")

		// Act/Assert
		assert.False(t, sut.IsValid())
	})
}

// nolint:dupl
func TestPage_GetNumberTextProp(t *testing.T) {
	t.Run("when place is left bottom, should map correctly", func(t *testing.T) {
		// Arrange
		prop := fixture.PageProp()
		prop.Place = props.LeftBottom

		// Act
		textProp := prop.GetNumberTextProp(100)

		// Assert
		assert.Equal(t, 100.0, textProp.Top)
		assert.Equal(t, 0.0, textProp.Left)
		assert.Equal(t, 0.0, textProp.Right)
		assert.Equal(t, fontfamily.Helvetica, textProp.Family)
		assert.Equal(t, fontstyle.Bold, textProp.Style)
		assert.Equal(t, 14.0, textProp.Size)
		assert.Equal(t, align.Left, textProp.Align)
		assert.Equal(t, breakline.EmptyLineStrategy, textProp.BreakLineStrategy)
		assert.Equal(t, 0.0, textProp.VerticalPadding)
		assert.Equal(t, &props.Color{Red: 100, Green: 50, Blue: 200}, textProp.Color)
	})
	t.Run("when place is left top, should map correctly", func(t *testing.T) {
		// Arrange
		prop := fixture.PageProp()
		prop.Place = props.LeftTop

		// Act
		textProp := prop.GetNumberTextProp(100)

		// Assert
		assert.Equal(t, 0.0, textProp.Top)
		assert.Equal(t, 0.0, textProp.Left)
		assert.Equal(t, 0.0, textProp.Right)
		assert.Equal(t, fontfamily.Helvetica, textProp.Family)
		assert.Equal(t, fontstyle.Bold, textProp.Style)
		assert.Equal(t, 14.0, textProp.Size)
		assert.Equal(t, align.Left, textProp.Align)
		assert.Equal(t, breakline.EmptyLineStrategy, textProp.BreakLineStrategy)
		assert.Equal(t, 0.0, textProp.VerticalPadding)
		assert.Equal(t, &props.Color{Red: 100, Green: 50, Blue: 200}, textProp.Color)
	})
	t.Run("when place is right bottom, should map correctly", func(t *testing.T) {
		// Arrange
		prop := fixture.PageProp()
		prop.Place = props.RightBottom

		// Act
		textProp := prop.GetNumberTextProp(100)

		// Assert
		assert.Equal(t, 100.0, textProp.Top)
		assert.Equal(t, 0.0, textProp.Left)
		assert.Equal(t, 0.0, textProp.Right)
		assert.Equal(t, fontfamily.Helvetica, textProp.Family)
		assert.Equal(t, fontstyle.Bold, textProp.Style)
		assert.Equal(t, 14.0, textProp.Size)
		assert.Equal(t, align.Right, textProp.Align)
		assert.Equal(t, breakline.EmptyLineStrategy, textProp.BreakLineStrategy)
		assert.Equal(t, 0.0, textProp.VerticalPadding)
		assert.Equal(t, &props.Color{Red: 100, Green: 50, Blue: 200}, textProp.Color)
	})
	t.Run("when place is right top, should map correctly", func(t *testing.T) {
		// Arrange
		prop := fixture.PageProp()
		prop.Place = props.RightTop

		// Act
		textProp := prop.GetNumberTextProp(100)

		// Assert
		assert.Equal(t, 0.0, textProp.Top)
		assert.Equal(t, 0.0, textProp.Left)
		assert.Equal(t, 0.0, textProp.Right)
		assert.Equal(t, fontfamily.Helvetica, textProp.Family)
		assert.Equal(t, fontstyle.Bold, textProp.Style)
		assert.Equal(t, 14.0, textProp.Size)
		assert.Equal(t, align.Right, textProp.Align)
		assert.Equal(t, breakline.EmptyLineStrategy, textProp.BreakLineStrategy)
		assert.Equal(t, 0.0, textProp.VerticalPadding)
		assert.Equal(t, &props.Color{Red: 100, Green: 50, Blue: 200}, textProp.Color)
	})
	t.Run("when place is right bottom, should map correctly", func(t *testing.T) {
		// Arrange
		prop := fixture.PageProp()
		prop.Place = props.RightBottom

		// Act
		textProp := prop.GetNumberTextProp(100)

		// Assert
		assert.Equal(t, 100.0, textProp.Top)
		assert.Equal(t, 0.0, textProp.Left)
		assert.Equal(t, 0.0, textProp.Right)
		assert.Equal(t, fontfamily.Helvetica, textProp.Family)
		assert.Equal(t, fontstyle.Bold, textProp.Style)
		assert.Equal(t, 14.0, textProp.Size)
		assert.Equal(t, align.Right, textProp.Align)
		assert.Equal(t, breakline.EmptyLineStrategy, textProp.BreakLineStrategy)
		assert.Equal(t, 0.0, textProp.VerticalPadding)
		assert.Equal(t, &props.Color{Red: 100, Green: 50, Blue: 200}, textProp.Color)
	})
	t.Run("when place is bottom, should map correctly", func(t *testing.T) {
		// Arrange
		prop := fixture.PageProp()
		prop.Place = props.Bottom

		// Act
		textProp := prop.GetNumberTextProp(100)

		// Assert
		assert.Equal(t, 100.0, textProp.Top)
		assert.Equal(t, 0.0, textProp.Left)
		assert.Equal(t, 0.0, textProp.Right)
		assert.Equal(t, fontfamily.Helvetica, textProp.Family)
		assert.Equal(t, fontstyle.Bold, textProp.Style)
		assert.Equal(t, 14.0, textProp.Size)
		assert.Equal(t, align.Center, textProp.Align)
		assert.Equal(t, breakline.EmptyLineStrategy, textProp.BreakLineStrategy)
		assert.Equal(t, 0.0, textProp.VerticalPadding)
		assert.Equal(t, &props.Color{Red: 100, Green: 50, Blue: 200}, textProp.Color)
	})
	t.Run("when place is left bottom, should map correctly", func(t *testing.T) {
		// Arrange
		prop := fixture.PageProp()
		prop.Place = props.LeftBottom

		// Act
		textProp := prop.GetNumberTextProp(100)

		// Assert
		assert.Equal(t, 100.0, textProp.Top)
		assert.Equal(t, 0.0, textProp.Left)
		assert.Equal(t, 0.0, textProp.Right)
		assert.Equal(t, fontfamily.Helvetica, textProp.Family)
		assert.Equal(t, fontstyle.Bold, textProp.Style)
		assert.Equal(t, 14.0, textProp.Size)
		assert.Equal(t, align.Left, textProp.Align)
		assert.Equal(t, breakline.EmptyLineStrategy, textProp.BreakLineStrategy)
		assert.Equal(t, 0.0, textProp.VerticalPadding)
		assert.Equal(t, &props.Color{Red: 100, Green: 50, Blue: 200}, textProp.Color)
	})
}

func TestPage_GetPageString(t *testing.T) {
	// Arrange
	prop := fixture.PageProp()

	// Act
	s := prop.GetPageString(10, 101)

	// Assert
	assert.Equal(t, "10 / 101", s)
}
