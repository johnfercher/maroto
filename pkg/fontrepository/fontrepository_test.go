package fontrepository_test

import (
	"os"
	"path"
	"strings"
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/fontrepository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRepository_AddUTF8Font(t *testing.T) {
	t.Parallel()
	t.Run("when fontstyle family is empty, should not add value", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := fontrepository.New()

		// Act
		customFonts, err := sut.AddUTF8Font("", fontstyle.Bold, "file").Load()

		// Assert
		assert.Nil(t, err)
		assert.Empty(t, customFonts)
	})
	t.Run("when fontstyle style is invalid, should not add value", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := fontrepository.New()

		// Act
		customFonts, err := sut.AddUTF8Font("family", "invalid", "file").Load()

		// Assert
		assert.Nil(t, err)
		assert.Empty(t, customFonts)
	})
	t.Run("when fontstyle file is empty, should not add value", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := fontrepository.New()

		// Act
		customFonts, err := sut.AddUTF8Font("family", fontstyle.Bold, "").Load()

		// Assert
		assert.Nil(t, err)
		assert.Empty(t, customFonts)
	})
	t.Run("when fontstyle is valid, should add value", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := fontrepository.New()

		// Act
		customFonts, err := sut.AddUTF8Font("family", fontstyle.Bold, buildPath("/docs/assets/fonts/arial-unicode-ms.ttf")).Load()

		// Assert
		assert.Nil(t, err)
		assert.Len(t, customFonts, 1)
		assert.Equal(t, "family", customFonts[0].GetFamily())
		assert.Equal(t, fontstyle.Bold, customFonts[0].GetStyle())
		assert.NotEmpty(t, customFonts[0].GetFile())
		assert.NotEmpty(t, customFonts[0].GetBytes())
	})
}

func TestRepository_AddUTF8FontFromBytes(t *testing.T) {
	t.Parallel()
	t.Run("when fontstyle family is empty, should not add value", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := fontrepository.New()

		// Act
		customFonts, err := sut.AddUTF8FontFromBytes("", fontstyle.Bold, []byte(``)).Load()

		// Assert
		assert.Nil(t, err)
		assert.Empty(t, customFonts)
	})
	t.Run("when fontstyle style is invalid, should not add value", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := fontrepository.New()

		// Act
		customFonts, err := sut.AddUTF8FontFromBytes("family", "invalid", []byte(``)).Load()

		// Assert
		assert.Nil(t, err)
		assert.Empty(t, customFonts)
	})
	t.Run("when fontstyle bytes is nil, should not add value", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := fontrepository.New()

		// Act
		customFonts, err := sut.AddUTF8FontFromBytes("family", fontstyle.Bold, nil).Load()

		// Assert
		assert.Nil(t, err)
		assert.Empty(t, customFonts)
	})
	t.Run("when fontstyle is valid, should add value", func(t *testing.T) {
		t.Parallel()
		// Arrange
		sut := fontrepository.New()
		ttf, err := os.ReadFile(buildPath("/docs/assets/fonts/arial-unicode-ms.ttf"))
		require.NoError(t, err)

		// Act
		customFonts, err := sut.AddUTF8FontFromBytes("family", fontstyle.Bold, ttf).Load()

		// Assert
		assert.Nil(t, err)
		assert.Len(t, customFonts, 1)
		assert.Equal(t, "family", customFonts[0].GetFamily())
		assert.Equal(t, fontstyle.Bold, customFonts[0].GetStyle())
		assert.Empty(t, customFonts[0].GetFile())
		assert.NotEmpty(t, customFonts[0].GetBytes())
	})
}

func buildPath(file string) string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}

	dir = strings.ReplaceAll(dir, "pkg/fontrepository", "")
	return path.Join(dir, file)
}
