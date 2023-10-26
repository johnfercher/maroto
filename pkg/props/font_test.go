package props_test

import (
	"github.com/johnfercher/maroto/v2/internal/fixture"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/stretchr/testify/assert"
)

func TestFont_MakeValid(t *testing.T) {
	t.Run("when family is not defined, should define default", func(t *testing.T) {
		// Arrange
		prop := props.Font{
			Family: "",
		}

		// Act
		prop.MakeValid(fontfamily.Arial)

		// Assert
		assert.Equal(t, prop.Family, fontfamily.Arial)
	})
	t.Run("when style is not defined, should define normal", func(t *testing.T) {
		// Arrange
		prop := props.Font{
			Style: "",
		}

		// Act
		prop.MakeValid(fontfamily.Arial)

		// Assert
		assert.Equal(t, prop.Style, fontstyle.Normal)
	})
	t.Run("", func(t *testing.T) {
		// Arrange
		prop := props.Font{
			Size: 0.0,
		}

		// Act
		prop.MakeValid(fontfamily.Arial)

		// Assert
		assert.Equal(t, prop.Size, 8.0)
	})
}

func TestFont_ToTextProp(t *testing.T) {
	// Arrange
	prop := fixture.FontProp()

	// Act
	textProp := prop.ToTextProp(align.Center, 10, 5)

	// Assert
	assert.Equal(t, prop.Family, textProp.Family)
	assert.Equal(t, prop.Style, textProp.Style)
	assert.Equal(t, prop.Size, textProp.Size)
	assert.Equal(t, prop.Color, textProp.Color)
	assert.Equal(t, align.Center, textProp.Align)
	assert.Equal(t, 10.0, textProp.Top)
	assert.Equal(t, 5.0, textProp.VerticalPadding)
}

func TestFont_AppendMap(t *testing.T) {
	// Arrange
	prop := fixture.FontProp()
	m := make(map[string]interface{})

	// Act
	m = prop.AppendMap(m)

	// Assert
	assert.Equal(t, fontfamily.Helvetica, m["prop_font_family"])
	assert.Equal(t, fontstyle.Bold, m["prop_font_style"])
	assert.Equal(t, 14.0, m["prop_font_size"])
	assert.Equal(t, "RGB(100, 50, 200)", m["prop_font_color"])
}
