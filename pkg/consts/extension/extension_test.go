package extension_test

import (
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestType_IsValid(t *testing.T) {
	t.Run("when type is empty, should not be valid", func(t *testing.T) {
		// Act
		extensionType := extension.Type("")

		// Act & Assert
		assert.False(t, extensionType.IsValid())
	})
	t.Run("when type is jpg, should be valid", func(t *testing.T) {
		// Act
		extensionType := extension.Jpg

		// Act & Assert
		assert.True(t, extensionType.IsValid())
	})
	t.Run("when type is png, should be valid", func(t *testing.T) {
		// Act
		extensionType := extension.Png

		// Act & Assert
		assert.True(t, extensionType.IsValid())
	})
}
