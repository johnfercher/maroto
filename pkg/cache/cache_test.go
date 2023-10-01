package cache_test

import (
	"fmt"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/cache"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	// Act
	sut := cache.New()

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*cache.cache", fmt.Sprintf("%T", sut))
}

func TestCache_GetImage(t *testing.T) {
	t.Run("when cannot get image, should return error", func(t *testing.T) {
		// Arrange
		sut := cache.New()

		// Act
		img, err := sut.GetImage("image", extension.Jpg)

		// Assert
		assert.Nil(t, img)
		assert.NotNil(t, err)
	})
	t.Run("when can get image, should return image", func(t *testing.T) {
		// Arrange
		sut := cache.New()
		sut.AddImage("image", &entity.Image{
			Extension: extension.Jpg,
		})

		// Act
		img, err := sut.GetImage("image", extension.Jpg)

		// Assert
		assert.NotNil(t, img)
		assert.Nil(t, err)
	})
}

func TestCache_AddImage(t *testing.T) {
	t.Run("when add image, return works", func(t *testing.T) {
		// Arrange
		sut := cache.New()

		// Act
		sut.AddImage("image", &entity.Image{
			Extension: extension.Jpg,
		})

		// Assert
		img, err := sut.GetImage("image", extension.Jpg)
		assert.NotNil(t, img)
		assert.Nil(t, err)
	})
}

func TestCache_LoadImage(t *testing.T) {
	t.Run("when cannot find image, should return error", func(t *testing.T) {
		// Arrange
		sut := cache.New()

		// Act
		err := sut.LoadImage("image", extension.Jpg)

		// Assert
		assert.NotNil(t, err)
	})
	t.Run("when can find image, should not return error and find image", func(t *testing.T) {
		// Arrange
		sut := cache.New()

		// Act
		err := sut.LoadImage(buildPath("/docs/assets/images/biplane.jpg"), extension.Jpg)

		// Assert
		assert.Nil(t, err)
		img, err := sut.GetImage(buildPath("/docs/assets/images/biplane.jpg"), extension.Jpg)
		assert.Nil(t, err)
		assert.NotNil(t, img)
	})
}

func buildPath(file string) string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}

	dir = strings.ReplaceAll(dir, "pkg/cache", "")
	return path.Join(dir, file)
}
