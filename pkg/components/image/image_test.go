package image_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/stretchr/testify/assert"
)

func TestFromBytes(t *testing.T) {
	t.Run("when extension is not valid, should return error", func(t *testing.T) {
		// Act
		img, err := image.FromBytes([]byte{1, 2, 3}, "invalid")

		// Assert
		assert.Nil(t, img)
		assert.NotNil(t, err)
	})
	t.Run("when extension is not valid, should return error", func(t *testing.T) {
		// Act
		img, err := image.FromBytes([]byte{1, 2, 3}, extension.Jpg)

		// Assert
		assert.NotNil(t, img)
		assert.Nil(t, err)
	})
}
