package border_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/johnfercher/maroto/v2/pkg/consts/border"
)

func TestType_IsValid(t *testing.T) {
	t.Run("When type is None, should not be valid", func(t *testing.T) {
		// Arrange
		borderType := border.None

		// Act & Assert
		assert.False(t, borderType.IsValid())
	})
	t.Run("When type is full, should be valid", func(t *testing.T) {
		// Arrange
		borderType := border.Full

		// Act & Assert
		assert.True(t, borderType.IsValid())
	})
	t.Run("When type is left, should be valid", func(t *testing.T) {
		// Arrange
		borderType := border.Left

		// Act & Assert
		assert.True(t, borderType.IsValid())
	})
	t.Run("When type is top, should be valid", func(t *testing.T) {
		// Arrange
		borderType := border.Top

		// Act & Assert
		assert.True(t, borderType.IsValid())
	})
	t.Run("When type is right, should be valid", func(t *testing.T) {
		// Arrange
		borderType := border.Right

		// Act & Assert
		assert.True(t, borderType.IsValid())
	})
	t.Run("When type is bottom, should be valid", func(t *testing.T) {
		// Arrange
		borderType := border.Bottom

		// Act & Assert
		assert.True(t, borderType.IsValid())
	})
}

func TestType_HasBorders(t *testing.T) {
	t.Run("When type is None, should not have any border", func(t *testing.T) {
		// Arrange
		borderType := border.None

		// Act & Assert
		assert.False(t, borderType.HasLeft())
		assert.False(t, borderType.HasTop())
		assert.False(t, borderType.HasRight())
		assert.False(t, borderType.HasBottom())
	})

	t.Run("When type is Full, should have all borders", func(t *testing.T) {
		// Arrange
		borderType := border.Full

		// Act & Assert
		assert.True(t, borderType.HasLeft())
		assert.True(t, borderType.HasTop())
		assert.True(t, borderType.HasRight())
		assert.True(t, borderType.HasBottom())
	})

	t.Run("When type is Left, should have only left border", func(t *testing.T) {
		// Arrange
		borderType := border.Left

		// Act & Assert
		assert.True(t, borderType.HasLeft())
		assert.False(t, borderType.HasTop())
		assert.False(t, borderType.HasRight())
		assert.False(t, borderType.HasBottom())
	})

	t.Run("When type is combined (Left|Top), should have left and top borders", func(t *testing.T) {
		// Arrange
		borderType := border.Left | border.Top

		// Act & Assert
		assert.True(t, borderType.HasLeft())
		assert.True(t, borderType.HasTop())
		assert.False(t, borderType.HasRight())
		assert.False(t, borderType.HasBottom())
	})

	t.Run("When type is combined (Left|Right), should have left and right borders", func(t *testing.T) {
		// Arrange
		borderType := border.Left | border.Right

		// Act & Assert
		assert.True(t, borderType.HasLeft())
		assert.False(t, borderType.HasTop())
		assert.True(t, borderType.HasRight())
		assert.False(t, borderType.HasBottom())
	})

	t.Run("When type is combined (Top|Bottom), should have top and bottom borders", func(t *testing.T) {
		// Arrange
		borderType := border.Top | border.Bottom

		// Act & Assert
		assert.False(t, borderType.HasLeft())
		assert.True(t, borderType.HasTop())
		assert.False(t, borderType.HasRight())
		assert.True(t, borderType.HasBottom())
	})

	t.Run("When type is combined (Left|Top|Right), should have left, top and right borders", func(t *testing.T) {
		// Arrange
		borderType := border.Left | border.Top | border.Right

		// Act & Assert
		assert.True(t, borderType.HasLeft())
		assert.True(t, borderType.HasTop())
		assert.True(t, borderType.HasRight())
		assert.False(t, borderType.HasBottom())
	})
}

func TestType_String(t *testing.T) {
	t.Run("When type is None, should return empty string", func(t *testing.T) {
		// Arrange
		borderType := border.None

		// Act & Assert
		assert.Equal(t, "", borderType.String())
	})

	t.Run("When type is Full, should return '1'", func(t *testing.T) {
		// Arrange
		borderType := border.Full

		// Act & Assert
		assert.Equal(t, "1", borderType.String())
	})

	t.Run("When type is Left, should return 'L'", func(t *testing.T) {
		// Arrange
		borderType := border.Left

		// Act & Assert
		assert.Equal(t, "L", borderType.String())
	})

	t.Run("When type is combined (Left|Top), should return 'LT'", func(t *testing.T) {
		// Arrange
		borderType := border.Left | border.Top

		// Act & Assert
		assert.Equal(t, "LT", borderType.String())
	})

	t.Run("When type is combined (Left|Right), should return 'LR'", func(t *testing.T) {
		// Arrange
		borderType := border.Left | border.Right

		// Act & Assert
		assert.Equal(t, "LR", borderType.String())
	})

	t.Run("When type is combined (Left|Top|Right), should return 'LTR'", func(t *testing.T) {
		// Arrange
		borderType := border.Left | border.Top | border.Right

		// Act & Assert
		assert.Equal(t, "LTR", borderType.String())
	})
}
