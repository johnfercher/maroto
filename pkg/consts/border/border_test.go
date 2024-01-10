package border_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/johnfercher/maroto/v2/pkg/consts/border"
)

func TestType_IsValid(t *testing.T) {
	t.Run("When type is empty, should not be valid", func(t *testing.T) {
		// Arrange
		borderType := border.Type("")

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
