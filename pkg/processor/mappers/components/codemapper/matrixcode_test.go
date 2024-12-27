// nolint:dupl
package codemapper_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/codemapper"
	"github.com/stretchr/testify/assert"
)

func TestMatrixCodeGetOrder(t *testing.T) {
	t.Run("when getOrder is called, should return defined order", func(t *testing.T) {
		templateRows := map[string]interface{}{
			"order": 10.0,
			"value": "code",
		}

		doc, _ := codemapper.NewMatrixcode(templateRows)

		assert.Equal(t, 10, doc.GetOrder())
	})
}

func TestNewMatrixcode(t *testing.T) {
	t.Run("when invalid matrixcode is sent, should return an error", func(t *testing.T) {
		matrixcodeTemplate := 1

		matrixcode, err := codemapper.NewMatrixcode(matrixcodeTemplate)

		assert.Nil(t, matrixcode)
		assert.NotNil(t, err)
	})
	t.Run("when props is not sent, matrixcode is created", func(t *testing.T) {
		matrixcodeTemplate := map[string]interface{}{
			"order": 1.0,
			"value": "123456789",
		}

		matrixcode, err := codemapper.NewMatrixcode(matrixcodeTemplate)

		assert.Nil(t, err)
		assert.NotNil(t, matrixcode)
	})
	t.Run("when invalid props is sent, should return an error", func(t *testing.T) {
		matrixcodeTemplate := map[string]interface{}{
			"order": 1.0,
			"props": 1,
			"code":  "123456789",
		}

		matrixcode, err := codemapper.NewMatrixcode(matrixcodeTemplate)

		assert.Nil(t, matrixcode)
		assert.NotNil(t, err)
	})
	t.Run("when invalid field is sent, should return an error", func(t *testing.T) {
		matrixcodeTemplate := map[string]interface{}{
			"order":         1.0,
			"invalid_field": 1,
			"code":          "123456789",
		}

		matrixcode, err := codemapper.NewMatrixcode(matrixcodeTemplate)

		assert.Nil(t, matrixcode)
		assert.NotNil(t, err)
	})
	t.Run("when source_key and code are not sent, should return an error", func(t *testing.T) {
		matrixcodeTemplate := map[string]interface{}{
			"order": 1.0,
		}

		matrixcode, err := codemapper.NewMatrixcode(matrixcodeTemplate)

		assert.Nil(t, matrixcode)
		assert.NotNil(t, err)
	})
	t.Run("when invalid code is sent, should return an error", func(t *testing.T) {
		matrixcodeTemplate := map[string]interface{}{
			"order": 1.0,
			"value": 123,
		}

		matrixcode, err := codemapper.NewMatrixcode(matrixcodeTemplate)

		assert.Nil(t, matrixcode)
		assert.NotNil(t, err)
	})
	t.Run("when invalid source_key is sent, should return an error", func(t *testing.T) {
		matrixcodeTemplate := map[string]interface{}{
			"order":      1.0,
			"source_key": 123,
		}

		matrixcode, err := codemapper.NewMatrixcode(matrixcodeTemplate)

		assert.Nil(t, matrixcode)
		assert.NotNil(t, err)
	})
	t.Run("when code is not sent, should add source key", func(t *testing.T) {
		matrixcodeTemplate := map[string]interface{}{
			"order":      1.0,
			"source_key": "source",
		}

		matrixcode, err := codemapper.NewMatrixcode(matrixcodeTemplate)

		assert.Nil(t, err)
		assert.Equal(t, matrixcode.SourceKey, "source")
	})

	t.Run("when source_key is not sent, should add code", func(t *testing.T) {
		matrixcodeTemplate := map[string]interface{}{
			"order": 1.0,
			"value": "code",
		}

		matrixcode, err := codemapper.NewMatrixcode(matrixcodeTemplate)

		assert.Nil(t, err)
		assert.Equal(t, matrixcode.Code, "code")
	})
}

func TestMatrixcodeGenerate(t *testing.T) {
	t.Run("if source key is not found, should return an error", func(t *testing.T) {
		content := map[string]interface{}{
			"order": 1.0,
		}
		provider := mocks.NewProcessorProvider(t)

		matrixcode := codemapper.Matrixcode{SourceKey: "code"}
		component, err := matrixcode.Generate(content, provider)

		assert.Nil(t, component)
		assert.NotNil(t, err)
	})
	t.Run("if source key content is not valid, should return an error", func(t *testing.T) {
		content := map[string]interface{}{
			"order": 1.0,
			"code":  1,
		}
		provider := mocks.NewProcessorProvider(t)

		matrixcode := codemapper.Matrixcode{SourceKey: "code"}
		component, err := matrixcode.Generate(content, provider)

		assert.Nil(t, component)
		assert.NotNil(t, err)
	})
	t.Run("If the matrixcode has no props, the props will not be sent", func(t *testing.T) {
		content := map[string]interface{}{
			"order": 1.0,
			"code":  "code",
		}

		provider := mocks.NewProcessorProvider(t)
		provider.EXPECT().CreateMatrixCode("code").Return(nil)

		matrixcode := codemapper.Matrixcode{SourceKey: "code"}
		_, err := matrixcode.Generate(content, provider)

		assert.Nil(t, err)
		provider.AssertNumberOfCalls(t, "CreateMatrixCode", 1)
	})
	t.Run("when valid code is sent, should generate matrixcode", func(t *testing.T) {
		content := map[string]interface{}{
			"order": 1.0,
		}

		provider := mocks.NewProcessorProvider(t)
		provider.EXPECT().CreateMatrixCode("code").Return(nil)

		matrixcode := codemapper.Matrixcode{Code: "code"}
		_, err := matrixcode.Generate(content, provider)

		assert.Nil(t, err)
		provider.AssertNumberOfCalls(t, "CreateMatrixCode", 1)
	})
}
