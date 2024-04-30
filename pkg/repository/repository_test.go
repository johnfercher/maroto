package repository_test

import (
	"os"
	"path"
	"strings"
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/repository"

	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRepository_AddUTF8Font(t *testing.T) {
	t.Run("when fontstyle family is empty, should not add value", func(t *testing.T) {
		// Arrange
		sut := repository.New()

		// Act
		customFonts, err := sut.AddUTF8Font("", fontstyle.Bold, "file").Load()

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, 0, len(customFonts))
	})

	t.Run("when fontstyle style is invalid, should not add value", func(t *testing.T) {
		// Arrange
		sut := repository.New()

		// Act
		customFonts, err := sut.AddUTF8Font("family", "invalid", "file").Load()

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, 0, len(customFonts))
	})

	t.Run("when fontstyle file is empty, should not add value", func(t *testing.T) {
		// Arrange
		sut := repository.New()

		// Act
		customFonts, err := sut.AddUTF8Font("family", fontstyle.Bold, "").Load()

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, 0, len(customFonts))
	})

	t.Run("when fontstyle is valid, should not value", func(t *testing.T) {
		// Arrange
		sut := repository.New()

		// Act
		customFonts, err := sut.AddUTF8Font("family", fontstyle.Bold, buildPath("/docs/assets/fonts/arial-unicode-ms.ttf")).Load()

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, 1, len(customFonts))
		assert.Equal(t, "family", customFonts[0].Family)
		assert.Equal(t, fontstyle.Bold, customFonts[0].Style)
		assert.NotEmpty(t, customFonts[0].File)
		assert.NotEmpty(t, customFonts[0].Bytes)
	})
}

func TestRepository_AddUTF8FontFromBytes(t *testing.T) {
	t.Run("when fontstyle family is empty, should not add value", func(t *testing.T) {
		// Arrange
		sut := repository.New()

		// Act
		customFonts, err := sut.AddUTF8FontFromBytes("", fontstyle.Bold, []byte(``)).Load()

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, 0, len(customFonts))
	})

	t.Run("when fontstyle style is invalid, should not add value", func(t *testing.T) {
		// Arrange
		sut := repository.New()

		// Act
		customFonts, err := sut.AddUTF8FontFromBytes("family", "invalid", []byte(``)).Load()

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, 0, len(customFonts))
	})

	t.Run("when fontstyle bytes is nil, should not add value", func(t *testing.T) {
		// Arrange
		sut := repository.New()

		// Act
		customFonts, err := sut.AddUTF8FontFromBytes("family", fontstyle.Bold, nil).Load()

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, 0, len(customFonts))
	})

	t.Run("when fontstyle is valid, should not value", func(t *testing.T) {
		// Arrange
		sut := repository.New()
		ttf, err := os.ReadFile(buildPath("/docs/assets/fonts/arial-unicode-ms.ttf"))
		require.NoError(t, err)

		// Act
		customFonts, err := sut.AddUTF8FontFromBytes("family", fontstyle.Bold, ttf).Load()

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, 1, len(customFonts))
		assert.Equal(t, "family", customFonts[0].Family)
		assert.Equal(t, fontstyle.Bold, customFonts[0].Style)
		assert.Empty(t, customFonts[0].File)
		assert.NotEmpty(t, customFonts[0].Bytes)
	})
}

func buildPath(file string) string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}

	dir = strings.ReplaceAll(dir, "pkg/repository", "")
	return path.Join(dir, file)
}
