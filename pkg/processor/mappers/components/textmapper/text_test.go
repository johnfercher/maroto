package textmapper_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/textmapper"
	"github.com/stretchr/testify/assert"
)

func TestNewText(t *testing.T) {
	t.Run("when invalid text is sent, should return an error", func(t *testing.T) {
		textTemplate := 1

		text, err := textmapper.NewText(textTemplate)

		assert.Nil(t, text)
		assert.NotNil(t, err)
	})
	t.Run("when props is not sent, text is created", func(t *testing.T) {
		textTemplate := map[string]interface{}{
			"value": "123456789",
		}

		text, err := textmapper.NewText(textTemplate)

		assert.Nil(t, err)
		assert.NotNil(t, text)
	})
	t.Run("when invalid props is sent, should return an error", func(t *testing.T) {
		textTemplate := map[string]interface{}{
			"props": 1,
			"value": "123456789",
		}

		text, err := textmapper.NewText(textTemplate)

		assert.Nil(t, text)
		assert.NotNil(t, err)
	})
	t.Run("when invalid field is sent, should return an error", func(t *testing.T) {
		textTemplate := map[string]interface{}{
			"invalid_field": 1,
			"value":         "123456789",
		}

		text, err := textmapper.NewText(textTemplate)

		assert.Nil(t, text)
		assert.NotNil(t, err)
	})
	t.Run("when source_key and value are not sent, should return an error", func(t *testing.T) {
		textTemplate := map[string]interface{}{}

		text, err := textmapper.NewText(textTemplate)

		assert.Nil(t, text)
		assert.NotNil(t, err)
	})
	t.Run("when invalid value is sent, should return an error", func(t *testing.T) {
		textTemplate := map[string]interface{}{
			"value": 123,
		}

		text, err := textmapper.NewText(textTemplate)

		assert.Nil(t, text)
		assert.NotNil(t, err)
	})
	t.Run("when invalid source_key is sent, should return an error", func(t *testing.T) {
		textTemplate := map[string]interface{}{
			"source_key": 123,
		}

		text, err := textmapper.NewText(textTemplate)

		assert.Nil(t, text)
		assert.NotNil(t, err)
	})
	t.Run("when value is not sent, should add source key", func(t *testing.T) {
		textTemplate := map[string]interface{}{
			"source_key": "source",
		}

		text, err := textmapper.NewText(textTemplate)

		assert.Nil(t, err)
		assert.Equal(t, text.SourceKey, "source")
	})

	t.Run("when source_key is not sent, should add value", func(t *testing.T) {
		textTemplate := map[string]interface{}{
			"value": "value",
		}

		text, err := textmapper.NewText(textTemplate)

		assert.Nil(t, err)
		assert.Equal(t, text.Value, "value")
	})
}

func TestGenerate(t *testing.T) {
	t.Run("if source key is not found, should return an error", func(t *testing.T) {
		content := map[string]interface{}{}
		provider := mocks.NewProcessorProvider(t)

		text := textmapper.Text{SourceKey: "text"}
		component, err := text.Generate(content, provider)

		assert.Nil(t, component)
		assert.NotNil(t, err)
	})
	t.Run("if source key content is not valid, should return an error", func(t *testing.T) {
		content := map[string]interface{}{
			"value": 1,
		}
		provider := mocks.NewProcessorProvider(t)

		text := textmapper.Text{SourceKey: "value"}
		component, err := text.Generate(content, provider)

		assert.Nil(t, component)
		assert.NotNil(t, err)
	})
	t.Run("If the text has no props, the props will not be sent", func(t *testing.T) {
		content := map[string]interface{}{
			"value": "text",
		}

		provider := mocks.NewProcessorProvider(t)
		provider.EXPECT().CreateText("text").Return(nil)

		text := textmapper.Text{SourceKey: "value"}
		_, err := text.Generate(content, provider)

		assert.Nil(t, err)
		provider.AssertNumberOfCalls(t, "CreateText", 1)
	})
	t.Run("when valid text is sent, should generate text", func(t *testing.T) {
		content := map[string]interface{}{}

		provider := mocks.NewProcessorProvider(t)
		provider.EXPECT().CreateText("text").Return(nil)

		text := textmapper.Text{Value: "text"}
		_, err := text.Generate(content, provider)

		assert.Nil(t, err)
		provider.AssertNumberOfCalls(t, "CreateText", 1)
	})
}
