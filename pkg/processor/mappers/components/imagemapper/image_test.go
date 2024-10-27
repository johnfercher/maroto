package imagemapper_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/imagemapper"
	"github.com/stretchr/testify/assert"
)

func TestNewImage(t *testing.T) {
	t.Run("when invalid image is sent, should return an error", func(t *testing.T) {
		imageTemplate := 1

		image, err := imagemapper.NewImage(imageTemplate)

		assert.Nil(t, image)
		assert.NotNil(t, err)
	})
	t.Run("when props is not sent, image is created", func(t *testing.T) {
		imageTemplate := map[string]interface{}{
			"source_key": "source key",
		}

		image, err := imagemapper.NewImage(imageTemplate)

		assert.Nil(t, err)
		assert.NotNil(t, image)
	})
	t.Run("when invalid props is sent, should return an error", func(t *testing.T) {
		imageTemplate := map[string]interface{}{
			"props":      1,
			"source_key": "name",
		}

		image, err := imagemapper.NewImage(imageTemplate)

		assert.Nil(t, image)
		assert.NotNil(t, err)
	})
	t.Run("when invalid field is sent, should return an error", func(t *testing.T) {
		imageTemplate := map[string]interface{}{
			"invalid_field": 1,
			"source_key":    "name",
		}

		image, err := imagemapper.NewImage(imageTemplate)

		assert.Nil(t, image)
		assert.NotNil(t, err)
	})
	t.Run("when source_key is not sent, should return an error", func(t *testing.T) {
		imageTemplate := map[string]interface{}{}

		image, err := imagemapper.NewImage(imageTemplate)

		assert.Nil(t, image)
		assert.NotNil(t, err)
	})
	t.Run("when invalid source_key is sent, should return an error", func(t *testing.T) {
		imageTemplate := map[string]interface{}{
			"source_key": 123,
		}

		image, err := imagemapper.NewImage(imageTemplate)

		assert.Nil(t, image)
		assert.NotNil(t, err)
	})
	t.Run("when source_key is sent, should add source_key", func(t *testing.T) {
		imageTemplate := map[string]interface{}{
			"source_key": "icon",
		}

		image, err := imagemapper.NewImage(imageTemplate)

		assert.Nil(t, err)
		assert.Equal(t, image.SourceKey, "icon")
	})
	t.Run("when props is sent, should add props", func(t *testing.T) {
		imageTemplate := map[string]interface{}{
			"source_key": "name",
			"props": map[string]interface{}{
				"left": 10.0,
			},
		}

		image, err := imagemapper.NewImage(imageTemplate)

		assert.Nil(t, err)
		assert.Equal(t, 10.0, image.Props.Left)
	})
}
