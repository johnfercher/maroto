package props_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/internal/fixture"

	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/stretchr/testify/assert"
)

func TestBarcode_ToMap(t *testing.T) {
	t.Run("when barcode is nil, should return nil", func(t *testing.T) {
		// Arrange
		var sut *props.Barcode

		// Act
		m := sut.ToMap()

		// Assert
		assert.Nil(t, m)
	})
	t.Run("when barcode is filled, should return map filled correctly", func(t *testing.T) {
		// Arrange
		sut := fixture.BarcodeProp()
		sut.Center = true

		// Act
		m := sut.ToMap()

		// Assert
		assert.Equal(t, 10.0, m["prop_left"])
		assert.Equal(t, 10.0, m["prop_top"])
		assert.Equal(t, 98.0, m["prop_percent"])
		assert.Equal(t, 16.0, m["prop_proportion_width"])
		assert.Equal(t, 3.2, m["prop_proportion_height"])
		assert.Equal(t, true, m["prop_center"])
	})
}

func TestBarcode_MakeValid(t *testing.T) {
	t.Run("when percent is less than zero, should become 100", func(t *testing.T) {
		// Arrange
		prop := props.Barcode{
			Percent: -2,
		}

		// Act
		prop.MakeValid()

		// Assert
		assert.Equal(t, prop.Percent, 100.0)
	})
	t.Run("when percent is greater than 100, should become 100", func(t *testing.T) {
		// Arrange
		prop := props.Barcode{
			Percent: 102,
		}

		// Act
		prop.MakeValid()

		// Assert
		assert.Equal(t, prop.Percent, 100.0)
	})
	t.Run("when is center, top and left should become 0", func(t *testing.T) {
		// Arrange
		prop := props.Barcode{
			Center: true,
			Top:    5,
			Left:   5,
		}

		// Act
		prop.MakeValid()

		// Assert
		assert.Equal(t, prop.Top, 0.0)
		assert.Equal(t, prop.Left, 0.0)
	})
	t.Run("when left is less than 0, should become 0", func(t *testing.T) {
		// Arrange
		prop := props.Barcode{
			Left: -5,
		}

		// Act
		prop.MakeValid()

		// Assert
		assert.Equal(t, prop.Left, 0.0)
	})
	t.Run("when top is less than 0, should become 0", func(t *testing.T) {
		// Arrange
		prop := props.Barcode{
			Top: -5,
		}

		// Act
		prop.MakeValid()

		// Assert
		assert.Equal(t, prop.Top, 0.0)
	})
	t.Run("when proportion.width less than 0", func(t *testing.T) {
		// Arrange
		prop := props.Barcode{
			Proportion: props.Proportion{
				Width: -5,
			},
		}

		// Act
		prop.MakeValid()

		// Assert
		assert.Equal(t, prop.Proportion.Width, 1.0)
	})
	t.Run("when proportion.height less than 0", func(t *testing.T) {
		// Arrange
		prop := props.Barcode{
			Proportion: props.Proportion{
				Height: -5,
			},
		}

		// Act
		prop.MakeValid()

		// Assert
		assert.Equal(t, prop.Proportion.Height, 0.20)
	})
	t.Run("when height is smaller than 10% of width", func(t *testing.T) {
		// Arrange
		prop := props.Barcode{
			Proportion: props.Proportion{
				Width:  11,
				Height: 1,
			},
		}

		// Act
		prop.MakeValid()

		// Assert
		assert.Equal(t, prop.Proportion.Height, prop.Proportion.Width*0.10)
	})
	t.Run("when height is grather than 20% of width", func(t *testing.T) {
		// Arrange
		prop := props.Barcode{
			Proportion: props.Proportion{
				Width:  5,
				Height: 5,
			},
		}

		// Act
		prop.MakeValid()

		// Assert
		assert.Equal(t, prop.Proportion.Height, prop.Proportion.Width*0.20)
	})
}

func TestBarcode_ToRectProp(t *testing.T) {
	// Arrange
	prop := fixture.BarcodeProp()

	// Act
	rect := prop.ToRectProp()

	// Assert
	assert.Equal(t, prop.Left, rect.Left)
	assert.Equal(t, prop.Top, rect.Top)
	assert.Equal(t, prop.Percent, rect.Percent)
	assert.Equal(t, prop.Center, rect.Center)
}
