package props_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/johnfercher/maroto/v2/internal/fixture"
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func TestLine_MakeValid(t *testing.T) {
	t.Run("when style is empty, should apply solid", func(t *testing.T) {
		// Arrange
		prop := props.Line{
			Style: "",
		}

		// Act
		prop.MakeValid()

		// Assert
		assert.Equal(t, linestyle.Solid, prop.Style)
	})
	t.Run("when thickness is 0.0, should apply default", func(t *testing.T) {
		// Arrange
		prop := props.Line{
			Thickness: 0.0,
		}

		// Act
		prop.MakeValid()

		// Assert
		assert.Equal(t, 0.2, prop.Thickness)
	})
	t.Run("when orientation is empty, should apply horizontal", func(t *testing.T) {
		// Arrange
		prop := props.Line{
			Orientation: "",
		}

		// Act
		prop.MakeValid()

		// Assert
		assert.Equal(t, orientation.Horizontal, prop.Orientation)
	})
	t.Run("when offset percent is less than 5, should apply 5", func(t *testing.T) {
		// Arrange
		prop := props.Line{
			OffsetPercent: 4,
		}

		// Act
		prop.MakeValid()

		// Assert
		assert.Equal(t, 5.0, prop.OffsetPercent)
	})
	t.Run("when offset percent is greater than 95, should apply 95", func(t *testing.T) {
		// Arrange
		prop := props.Line{
			OffsetPercent: 96,
		}

		// Act
		prop.MakeValid()

		// Assert
		assert.Equal(t, 95.0, prop.OffsetPercent)
	})
	t.Run("when size percent is less than 1, should apply 90", func(t *testing.T) {
		// Arrange
		prop := props.Line{
			SizePercent: 0,
		}

		// Act
		prop.MakeValid()

		// Assert
		assert.Equal(t, 90.0, prop.SizePercent)
	})
	t.Run("when size percent is greater than 100, should apply 100", func(t *testing.T) {
		// Arrange
		prop := props.Line{
			SizePercent: 101,
		}

		// Act
		prop.MakeValid()

		// Assert
		assert.Equal(t, 100.0, prop.SizePercent)
	})
}

func TestLine_ToMap(t *testing.T) {
	t.Run("when line is nil, should return nil", func(t *testing.T) {
		// Arrange
		var prop *props.Line

		// Act
		m := prop.ToMap()

		// Assert
		assert.Nil(t, m)
	})
	t.Run("when line is filled, should return map filled", func(t *testing.T) {
		// Arrange
		prop := fixture.LineProp()

		// Act
		m := prop.ToMap()

		// Assert
		assert.Equal(t, "RGB(100, 50, 200)", m["prop_color"])
		assert.Equal(t, linestyle.Dashed, m["prop_style"])
		assert.Equal(t, 1.1, m["prop_thickness"])
		assert.Equal(t, orientation.Vertical, m["prop_orientation"])
		assert.Equal(t, 50.0, m["prop_offset_percent"])
		assert.Equal(t, 20.0, m["prop_size_percent"])
	})
}
