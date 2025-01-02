package signaturemapper_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/signaturemapper"
	"github.com/stretchr/testify/assert"
)

func TestQrcodeGetOrder(t *testing.T) {
	t.Run("when getOrder is called, should return defined order", func(t *testing.T) {
		templateRows := map[string]interface{}{
			"order": 10.0,
			"value": "test",
		}

		doc, _ := signaturemapper.NewSignature(templateRows)

		assert.Equal(t, 10, doc.GetOrder())
	})
}

func TestNewSignature(t *testing.T) {
	t.Run("when invalid signature is sent, should return an error", func(t *testing.T) {
		signatureTemplate := 1

		signature, err := signaturemapper.NewSignature(signatureTemplate)

		assert.Nil(t, signature)
		assert.NotNil(t, err)
	})
	t.Run("when props is not sent, signature is created", func(t *testing.T) {
		signatureTemplate := map[string]interface{}{
			"value": "123456789",
			"order": 1.0,
		}

		signature, err := signaturemapper.NewSignature(signatureTemplate)

		assert.Nil(t, err)
		assert.NotNil(t, signature)
	})
	t.Run("when invalid props is sent, should return an error", func(t *testing.T) {
		signatureTemplate := map[string]interface{}{
			"props": 1,
			"value": "123456789",
			"order": 1.0,
		}

		signature, err := signaturemapper.NewSignature(signatureTemplate)

		assert.Nil(t, signature)
		assert.NotNil(t, err)
	})
	t.Run("when invalid field is sent, should return an error", func(t *testing.T) {
		signatureTemplate := map[string]interface{}{
			"invalid_field": 1,
			"order":         1.0,
		}

		signature, err := signaturemapper.NewSignature(signatureTemplate)

		assert.Nil(t, signature)
		assert.NotNil(t, err)
	})
	t.Run("when source_key and value are not sent, should return an error", func(t *testing.T) {
		signatureTemplate := map[string]interface{}{}

		signature, err := signaturemapper.NewSignature(signatureTemplate)

		assert.Nil(t, signature)
		assert.NotNil(t, err)
	})
	t.Run("when invalid value is sent, should return an error", func(t *testing.T) {
		signatureTemplate := map[string]interface{}{
			"value": 123,
			"order": 1.0,
		}

		signature, err := signaturemapper.NewSignature(signatureTemplate)

		assert.Nil(t, signature)
		assert.NotNil(t, err)
	})
	t.Run("when invalid source_key is sent, should return an error", func(t *testing.T) {
		signatureTemplate := map[string]interface{}{
			"source_key": 123,
			"order":      1.0,
		}

		signature, err := signaturemapper.NewSignature(signatureTemplate)

		assert.Nil(t, signature)
		assert.NotNil(t, err)
	})
	t.Run("when value is not sent, should add source key", func(t *testing.T) {
		signatureTemplate := map[string]interface{}{
			"source_key": "source",
			"order":      1.0,
		}

		signature, err := signaturemapper.NewSignature(signatureTemplate)

		assert.Nil(t, err)
		assert.Equal(t, signature.SourceKey, "source")
	})

	t.Run("when source_key is not sent, should add code", func(t *testing.T) {
		signatureTemplate := map[string]interface{}{
			"value": "value",
			"order": 1.0,
		}

		signature, err := signaturemapper.NewSignature(signatureTemplate)

		assert.Nil(t, err)
		assert.Equal(t, signature.Value, "value")
	})
}

func TestGenerate(t *testing.T) {
	t.Run("if source key is not found, should return an error", func(t *testing.T) {
		content := map[string]interface{}{}
		provider := mocks.NewProcessorProvider(t)

		signature := signaturemapper.Signature{SourceKey: "code"}
		component, err := signature.Generate(content, provider)

		assert.Nil(t, component)
		assert.NotNil(t, err)
	})
	t.Run("if source key content is not valid, should return an error", func(t *testing.T) {
		content := map[string]interface{}{
			"value": 1,
		}
		provider := mocks.NewProcessorProvider(t)

		signature := signaturemapper.Signature{SourceKey: "value"}
		component, err := signature.Generate(content, provider)

		assert.Nil(t, component)
		assert.NotNil(t, err)
	})
	t.Run("If the signature has no props, the props will not be sent", func(t *testing.T) {
		content := map[string]interface{}{
			"value": "signature",
		}

		provider := mocks.NewProcessorProvider(t)
		provider.EXPECT().CreateSignature("signature").Return(nil)

		signature := signaturemapper.Signature{SourceKey: "value"}
		_, err := signature.Generate(content, provider)

		assert.Nil(t, err)
		provider.AssertNumberOfCalls(t, "CreateSignature", 1)
	})
	t.Run("when valid code is sent, should generate signature", func(t *testing.T) {
		content := map[string]interface{}{}

		provider := mocks.NewProcessorProvider(t)
		provider.EXPECT().CreateSignature("signature").Return(nil)

		signature := signaturemapper.Signature{Value: "signature"}
		_, err := signature.Generate(content, provider)

		assert.Nil(t, err)
		provider.AssertNumberOfCalls(t, "CreateSignature", 1)
	})
}
