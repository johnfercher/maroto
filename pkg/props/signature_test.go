package props_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/johnfercher/maroto/v2/internal/fixture"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func TestSignature_ToMap(t *testing.T) {
	t.Parallel()
	t.Run("when prop is nil, should return nil", func(t *testing.T) {
		t.Parallel()
		// Arrange
		var sut *props.Signature

		// Act
		m := sut.ToMap()

		// Assert
		assert.Nil(t, m)
	})
	t.Run("when prop is filled, should return filled map", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := fixture.SignatureProp()

		// Act
		m := sut.ToMap()

		// Assert
		assert.Equal(t, fontfamily.Helvetica, m["prop_font_family"])
		assert.Equal(t, fontstyle.Bold, m["prop_font_style"])
		assert.Equal(t, 14.0, m["prop_font_size"])
		assert.Equal(t, linestyle.Dashed, m["prop_line_style"])
		assert.Equal(t, 1.1, m["prop_line_thickness"])
		assert.Equal(t, "RGB(100, 50, 200)", m["prop_font_color"])
		assert.Equal(t, "RGB(100, 50, 200)", m["prop_line_color"])
	})
}

func TestSignature_MakeValid(t *testing.T) {
	t.Parallel()
	t.Run("when font family is empty, should apply default font family", func(t *testing.T) {
		t.Parallel()
		// Arrange
		prop := props.Signature{FontFamily: ""}

		// Act
		prop.MakeValid(fontfamily.Arial)

		// Assert
		assert.Equal(t, fontfamily.Arial, prop.FontFamily)
	})
	t.Run("when font style is empty, should apply bold", func(t *testing.T) {
		t.Parallel()
		// Arrange
		prop := props.Signature{FontStyle: ""}

		// Act
		prop.MakeValid(fontfamily.Arial)

		// Assert
		assert.Equal(t, fontstyle.Bold, prop.FontStyle)
	})
	t.Run("when font size is zero, should apply 8.0", func(t *testing.T) {
		t.Parallel()
		// Arrange
		prop := props.Signature{FontSize: 0}

		// Act
		prop.MakeValid(fontfamily.Arial)

		// Assert
		assert.Equal(t, 8.0, prop.FontSize)
	})
	t.Run("when line style is empty, should apply solid", func(t *testing.T) {
		t.Parallel()
		// Arrange
		prop := props.Signature{LineStyle: ""}

		// Act
		prop.MakeValid(fontfamily.Arial)

		// Assert
		assert.Equal(t, linestyle.Solid, prop.LineStyle)
	})
	t.Run("when line thickness is zero, should apply default thickness", func(t *testing.T) {
		t.Parallel()
		// Arrange
		prop := props.Signature{LineThickness: 0}

		// Act
		prop.MakeValid(fontfamily.Arial)

		// Assert
		assert.Equal(t, linestyle.DefaultLineThickness, prop.LineThickness)
	})
	t.Run("when safe padding is zero, should apply 1.5", func(t *testing.T) {
		t.Parallel()
		// Arrange
		prop := props.Signature{SafePadding: 0}

		// Act
		prop.MakeValid(fontfamily.Arial)

		// Assert
		assert.Equal(t, 1.5, prop.SafePadding)
	})
	t.Run("when safe padding is negative, should apply 1.5", func(t *testing.T) {
		t.Parallel()
		// Arrange
		prop := props.Signature{SafePadding: -1}

		// Act
		prop.MakeValid(fontfamily.Arial)

		// Assert
		assert.Equal(t, 1.5, prop.SafePadding)
	})
}

func TestSignature_ToLineProp(t *testing.T) {
	t.Parallel()
	// Arrange
	prop := fixture.SignatureProp()

	// Act
	lineProp := prop.ToLineProp(55)

	// Assert
	assert.Equal(t, prop.LineColor, lineProp.Color)
	assert.Equal(t, prop.LineStyle, lineProp.Style)
	assert.Equal(t, prop.LineThickness, lineProp.Thickness)
	assert.Equal(t, orientation.Horizontal, lineProp.Orientation)
	assert.Equal(t, 55.0, lineProp.OffsetPercent)
	assert.Equal(t, 90.0, lineProp.SizePercent)
}

func TestSignature_ToFontProp(t *testing.T) {
	t.Parallel()
	// Arrange
	prop := fixture.SignatureProp()

	// Act
	fontProp := prop.ToFontProp()

	// Assert
	assert.Equal(t, prop.FontFamily, fontProp.Family)
	assert.Equal(t, prop.FontStyle, fontProp.Style)
	assert.Equal(t, prop.FontSize, fontProp.Size)
	assert.Equal(t, &props.Color{Red: 100, Green: 50, Blue: 200}, fontProp.Color)
}

func TestSignature_ToTextProp(t *testing.T) {
	t.Parallel()
	// Arrange
	prop := fixture.SignatureProp()

	// Act
	textProp := prop.ToTextProp(align.Center, 5.0, 2.0)

	// Assert
	assert.Equal(t, prop.FontFamily, textProp.Family)
	assert.Equal(t, prop.FontStyle, textProp.Style)
	assert.Equal(t, prop.FontSize, textProp.Size)
	assert.Equal(t, align.Center, textProp.Align)
	assert.Equal(t, 5.0, textProp.Top)
	assert.Equal(t, 2.0, textProp.VerticalPadding)
	assert.Equal(t, prop.FontColor, textProp.Color)
}
