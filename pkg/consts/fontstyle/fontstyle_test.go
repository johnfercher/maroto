package fontstyle_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/stretchr/testify/assert"
)

func TestType_IsValid(t *testing.T) {
	t.Run("when style is invalid, should be invalid", func(t *testing.T) {
		// Arrange
		fontStyle := fontstyle.Type("invalid")

		// Act & Assert
		assert.False(t, fontStyle.IsValid())
	})
	t.Run("when style is normal, should be valid", func(t *testing.T) {
		// Arrange
		fontStyle := fontstyle.Normal

		// Act & Assert
		assert.True(t, fontStyle.IsValid())
	})
	t.Run("when style is bold, should be valid", func(t *testing.T) {
		// Arrange
		fontStyle := fontstyle.Bold

		// Act & Assert
		assert.True(t, fontStyle.IsValid())
	})
	t.Run("when style is bold italic, should be valid", func(t *testing.T) {
		// Arrange
		fontStyle := fontstyle.BoldItalic

		// Act & Assert
		assert.True(t, fontStyle.IsValid())
	})
}
