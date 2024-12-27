package imagemapper_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/imagemapper"
	"github.com/stretchr/testify/assert"
)

func TestGetOrder(t *testing.T) {
	t.Run("when getOrder is called, should return defined order", func(t *testing.T) {
		templateRows := map[string]interface{}{
			"order": 10.0,
			"value": "img",
		}

		doc, _ := imagemapper.NewImage(templateRows)

		assert.Equal(t, 10, doc.GetOrder())
	})
}

func TestNewImage(t *testing.T) {
	t.Run("when invalid image is sent, should return an error", func(t *testing.T) {
		imageTemplate := 1

		image, err := imagemapper.NewImage(imageTemplate)

		assert.Nil(t, image)
		assert.NotNil(t, err)
	})
	t.Run("when props is not sent, image is created", func(t *testing.T) {
		imageTemplate := map[string]interface{}{
			"source_key": "image",
			"order":      1.0,
		}

		image, err := imagemapper.NewImage(imageTemplate)

		assert.Nil(t, err)
		assert.NotNil(t, image)
	})
	t.Run("when invalid props is sent, should return an error", func(t *testing.T) {
		imageTemplate := map[string]interface{}{
			"props":      1,
			"source_key": "name",
			"order":      1.0,
		}

		image, err := imagemapper.NewImage(imageTemplate)

		assert.Nil(t, image)
		assert.NotNil(t, err)
	})
	t.Run("when invalid field is sent, should return an error", func(t *testing.T) {
		imageTemplate := map[string]interface{}{
			"invalid_field": 1,
			"source_key":    "name",
			"order":         1.0,
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
			"order":      1.0,
		}

		image, err := imagemapper.NewImage(imageTemplate)

		assert.Nil(t, image)
		assert.NotNil(t, err)
	})
	t.Run("when source_key and path are not sent, should return an error", func(t *testing.T) {
		imageTemplate := map[string]interface{}{}

		image, err := imagemapper.NewImage(imageTemplate)

		assert.Nil(t, image)
		assert.NotNil(t, err)
	})
	t.Run("when source_key is sent, should add source_key", func(t *testing.T) {
		imageTemplate := map[string]interface{}{
			"source_key": "icon",
			"order":      1.0,
		}

		image, err := imagemapper.NewImage(imageTemplate)

		assert.Nil(t, err)
		assert.Equal(t, image.SourceKey, "icon")
	})
	t.Run("when props is sent, should add props", func(t *testing.T) {
		imageTemplate := map[string]interface{}{
			"source_key": "name",
			"order":      1.0,
			"props": map[string]interface{}{
				"left": 10.0,
			},
		}

		image, err := imagemapper.NewImage(imageTemplate)

		assert.Nil(t, err)
		assert.Equal(t, 10.0, image.Props.Left)
	})
}

func TestImageGenerate(t *testing.T) {
	t.Run("if image is not found, should return an error", func(t *testing.T) {
		content := map[string]interface{}{}
		provider := mocks.NewProcessorProvider(t)

		image := imagemapper.Image{SourceKey: "code"}
		component, err := image.Generate(content, provider)

		assert.Nil(t, component)
		assert.NotNil(t, err)
	})
	t.Run("if image content is not valid, should return an error", func(t *testing.T) {
		content := map[string]interface{}{
			"code": 1,
		}
		provider := mocks.NewProcessorProvider(t)

		image := imagemapper.Image{SourceKey: "code"}
		component, err := image.Generate(content, provider)

		assert.Nil(t, component)
		assert.NotNil(t, err)
	})
	t.Run("If the image has no props, the props will not be sent", func(t *testing.T) {
		content := map[string]interface{}{
			"Path": "path.png",
		}
		provider := mocks.NewProcessorProvider(t)
		provider.EXPECT().CreateImage("path.png").Return(nil, nil)

		image := imagemapper.Image{SourceKey: "Path"}
		_, err := image.Generate(content, provider)

		assert.Nil(t, err)
		provider.AssertNumberOfCalls(t, "CreateImage", 1)
	})
}
