package props_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/johnfercher/maroto/v2/internal/fixture"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/breakline"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/props"
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
		assert.Equal(t, breakline.EmptySpaceStrategy, textProp.BreakLineStrategy)
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
		assert.Equal(t, breakline.EmptySpaceStrategy, textProp.BreakLineStrategy)
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
		assert.Equal(t, breakline.EmptySpaceStrategy, textProp.BreakLineStrategy)
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
		assert.Equal(t, breakline.EmptySpaceStrategy, textProp.BreakLineStrategy)
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
		assert.Equal(t, breakline.EmptySpaceStrategy, textProp.BreakLineStrategy)
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
		assert.Equal(t, breakline.EmptySpaceStrategy, textProp.BreakLineStrategy)
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
		assert.Equal(t, breakline.EmptySpaceStrategy, textProp.BreakLineStrategy)
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

func TestPageNumber_WithFont(t *testing.T) {
	t.Run("when font already defined, should keep it", func(t *testing.T) {
		// Arrange
		pageNumber := &props.PageNumber{
			Color:  &props.RedColor,
			Size:   15,
			Style:  fontstyle.Bold,
			Family: fontfamily.Helvetica,
		}

		font := &props.Font{
			Color:  &props.BlueColor,
			Size:   13,
			Style:  fontstyle.Italic,
			Family: fontfamily.Arial,
		}

		// Act
		pageNumber.WithFont(font)

		// Assert
		assert.Equal(t, &props.RedColor, pageNumber.Color)
		assert.Equal(t, 15.0, pageNumber.Size)
		assert.Equal(t, fontstyle.Bold, pageNumber.Style)
		assert.Equal(t, fontfamily.Helvetica, pageNumber.Family)
	})
	t.Run("when font not defined, should apply", func(t *testing.T) {
		// Arrange
		pageNumber := &props.PageNumber{}

		font := &props.Font{
			Color:  &props.BlueColor,
			Size:   13,
			Style:  fontstyle.Italic,
			Family: fontfamily.Arial,
		}

		// Act
		pageNumber.WithFont(font)

		// Assert
		assert.Equal(t, &props.BlueColor, pageNumber.Color)
		assert.Equal(t, 13.0, pageNumber.Size)
		assert.Equal(t, fontstyle.Italic, pageNumber.Style)
		assert.Equal(t, fontfamily.Arial, pageNumber.Family)
	})
}

func TestPageNumber_AppendMap(t *testing.T) {
	t.Run("when append map, should append correctly", func(t *testing.T) {
		// Arrange
		pageNumber := &props.PageNumber{
			Pattern: "pattern",
			Place:   props.Bottom,
			Color:   &props.RedColor,
			Size:    15,
			Style:   fontstyle.Bold,
			Family:  fontfamily.Helvetica,
		}

		m := make(map[string]interface{})

		// Act
		m = pageNumber.AppendMap(m)

		// Assert
		assert.Equal(t, "pattern", m["page_number_pattern"])
		assert.Equal(t, props.Bottom, m["page_number_place"])
		assert.Equal(t, fontfamily.Helvetica, m["page_number_family"])
		assert.Equal(t, fontstyle.Bold, m["page_number_style"])
		assert.Equal(t, 15.0, m["page_number_size"])
		assert.Equal(t, "RGB(255, 0, 0)", m["page_number_color"])
	})
}
