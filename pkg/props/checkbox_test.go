package props_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/johnfercher/maroto/v2/internal/fixture"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func TestCheckbox_MakeValid(t *testing.T) {
	t.Parallel()
	t.Run("when size is zero, should apply default 5.0", func(t *testing.T) {
		t.Parallel()
		// Arrange
		prop := props.Checkbox{Size: 0}

		// Act
		prop.MakeValid()

		// Assert
		assert.Equal(t, 5.0, prop.Size)
	})
	t.Run("when size is negative, should apply default 5.0", func(t *testing.T) {
		t.Parallel()
		// Arrange
		prop := props.Checkbox{Size: -1}

		// Act
		prop.MakeValid()

		// Assert
		assert.Equal(t, 5.0, prop.Size)
	})
	t.Run("when size is positive, should keep value", func(t *testing.T) {
		t.Parallel()
		// Arrange
		prop := props.Checkbox{Size: 10}

		// Act
		prop.MakeValid()

		// Assert
		assert.Equal(t, 10.0, prop.Size)
	})
	t.Run("when top is negative, should apply 0", func(t *testing.T) {
		t.Parallel()
		// Arrange
		prop := props.Checkbox{Top: -5, Size: 5}

		// Act
		prop.MakeValid()

		// Assert
		assert.Equal(t, 0.0, prop.Top)
	})
	t.Run("when top is zero, should keep value", func(t *testing.T) {
		t.Parallel()
		// Arrange
		prop := props.Checkbox{Top: 0, Size: 5}

		// Act
		prop.MakeValid()

		// Assert
		assert.Equal(t, 0.0, prop.Top)
	})
	t.Run("when top is positive, should keep value", func(t *testing.T) {
		t.Parallel()
		// Arrange
		prop := props.Checkbox{Top: 3, Size: 5}

		// Act
		prop.MakeValid()

		// Assert
		assert.Equal(t, 3.0, prop.Top)
	})
	t.Run("when left is negative, should apply 0", func(t *testing.T) {
		t.Parallel()
		// Arrange
		prop := props.Checkbox{Left: -3, Size: 5}

		// Act
		prop.MakeValid()

		// Assert
		assert.Equal(t, 0.0, prop.Left)
	})
	t.Run("when left is zero, should keep value", func(t *testing.T) {
		t.Parallel()
		// Arrange
		prop := props.Checkbox{Left: 0, Size: 5}

		// Act
		prop.MakeValid()

		// Assert
		assert.Equal(t, 0.0, prop.Left)
	})
	t.Run("when left is positive, should keep value", func(t *testing.T) {
		t.Parallel()
		// Arrange
		prop := props.Checkbox{Left: 7, Size: 5}

		// Act
		prop.MakeValid()

		// Assert
		assert.Equal(t, 7.0, prop.Left)
	})
}

func TestCheckbox_ToMap(t *testing.T) {
	t.Parallel()
	t.Run("when checked is false and all fields are zero, should return empty map", func(t *testing.T) {
		t.Parallel()
		// Arrange
		prop := props.Checkbox{}

		// Act
		m := prop.ToMap()

		// Assert
		assert.Empty(t, m)
	})
	t.Run("when checkbox is filled, should return map filled", func(t *testing.T) {
		t.Parallel()
		// Arrange
		prop := fixture.CheckboxProp()

		// Act
		m := prop.ToMap()

		// Assert
		assert.Equal(t, true, m["prop_checked"])
		assert.Equal(t, 5.0, m["prop_top"])
		assert.Equal(t, 5.0, m["prop_left"])
		assert.Equal(t, 10.0, m["prop_size"])
	})
	t.Run("when checked is false, should not include checked in map", func(t *testing.T) {
		t.Parallel()
		// Arrange
		prop := props.Checkbox{Checked: false, Top: 2, Left: 3, Size: 5}

		// Act
		m := prop.ToMap()

		// Assert
		assert.NotContains(t, m, "prop_checked")
		assert.Equal(t, 2.0, m["prop_top"])
		assert.Equal(t, 3.0, m["prop_left"])
		assert.Equal(t, 5.0, m["prop_size"])
	})
	t.Run("when top is zero, should not include top in map", func(t *testing.T) {
		t.Parallel()
		// Arrange
		prop := props.Checkbox{Top: 0, Left: 1, Size: 5}

		// Act
		m := prop.ToMap()

		// Assert
		assert.NotContains(t, m, "prop_top")
	})
	t.Run("when left is zero, should not include left in map", func(t *testing.T) {
		t.Parallel()
		// Arrange
		prop := props.Checkbox{Left: 0, Top: 1, Size: 5}

		// Act
		m := prop.ToMap()

		// Assert
		assert.NotContains(t, m, "prop_left")
	})
	t.Run("when size is zero, should not include size in map", func(t *testing.T) {
		t.Parallel()
		// Arrange
		prop := props.Checkbox{Size: 0, Top: 1, Left: 1}

		// Act
		m := prop.ToMap()

		// Assert
		assert.NotContains(t, m, "prop_size")
	})
}
