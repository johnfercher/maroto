package props_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/stretchr/testify/assert"
)

func TestCheckbox_ToMap(t *testing.T) {
	t.Run("when all properties are set, should return complete map", func(t *testing.T) {
		// Arrange
		sut := &props.Checkbox{
			Checked: true,
			BoxSize: 5.0,
			Left:    2.0,
			Top:     3.0,
			Right:   4.0,
			Bottom:  5.0,
		}

		// Act
		m := sut.ToMap()

		// Assert
		assert.Equal(t, true, m["prop_checked"])
		assert.Equal(t, 5.0, m["prop_box_size"])
		assert.Equal(t, 2.0, m["prop_left"])
		assert.Equal(t, 3.0, m["prop_top"])
		assert.Equal(t, 4.0, m["prop_right"])
		assert.Equal(t, 5.0, m["prop_bottom"])
	})

	t.Run("when properties are default, should return minimal map", func(t *testing.T) {
		// Arrange
		sut := &props.Checkbox{}

		// Act
		m := sut.ToMap()

		// Assert
		_, hasChecked := m["prop_checked"]
		assert.False(t, hasChecked)
		_, hasBoxSize := m["prop_box_size"]
		assert.False(t, hasBoxSize)
	})
}

func TestCheckbox_MakeValid(t *testing.T) {
	t.Run("when box size is too small, should set to default", func(t *testing.T) {
		// Arrange
		sut := &props.Checkbox{
			BoxSize: 0.5,
		}

		// Act
		sut.MakeValid(nil)

		// Assert
		assert.Equal(t, 4.0, sut.BoxSize)
	})

	t.Run("when box size is too large, should set to default", func(t *testing.T) {
		// Arrange
		sut := &props.Checkbox{
			BoxSize: 25.0,
		}

		// Act
		sut.MakeValid(nil)

		// Assert
		assert.Equal(t, 4.0, sut.BoxSize)
	})

	t.Run("when box size is valid, should keep it", func(t *testing.T) {
		// Arrange
		sut := &props.Checkbox{
			BoxSize: 6.0,
		}

		// Act
		sut.MakeValid(nil)

		// Assert
		assert.Equal(t, 6.0, sut.BoxSize)
	})

	t.Run("when negative padding, should set to zero", func(t *testing.T) {
		// Arrange
		sut := &props.Checkbox{
			Left:   -1.0,
			Top:    -2.0,
			Right:  -3.0,
			Bottom: -4.0,
		}

		// Act
		sut.MakeValid(nil)

		// Assert
		assert.Equal(t, 0.0, sut.Left)
		assert.Equal(t, 0.0, sut.Top)
		assert.Equal(t, 0.0, sut.Right)
		assert.Equal(t, 0.0, sut.Bottom)
	})

	t.Run("when all properties are valid, should keep them", func(t *testing.T) {
		// Arrange
		sut := &props.Checkbox{
			Checked: true,
			BoxSize: 5.0,
			Left:    1.0,
			Top:     2.0,
			Right:   3.0,
			Bottom:  4.0,
		}

		// Act
		sut.MakeValid(nil)

		// Assert
		assert.Equal(t, true, sut.Checked)
		assert.Equal(t, 5.0, sut.BoxSize)
		assert.Equal(t, 1.0, sut.Left)
		assert.Equal(t, 2.0, sut.Top)
		assert.Equal(t, 3.0, sut.Right)
		assert.Equal(t, 4.0, sut.Bottom)
	})
}
