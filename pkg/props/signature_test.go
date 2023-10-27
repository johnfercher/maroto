package props_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/internal/fixture"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/stretchr/testify/assert"
)

func TestSignature_ToMap(t *testing.T) {
	t.Run("when prop is nil, should return nil", func(t *testing.T) {
		// Arrange
		var sut *props.Signature

		// Act
		m := sut.ToMap()

		// Assert
		assert.Nil(t, m)
	})
	t.Run("when prop is filled, should return filled map", func(t *testing.T) {
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

func TestSignature_ToLineProp(t *testing.T) {
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
