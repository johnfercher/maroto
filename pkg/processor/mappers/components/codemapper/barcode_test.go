// nolint:dupl
package codemapper_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/codemapper"
	"github.com/stretchr/testify/assert"
)

func TestNewBarcode(t *testing.T) {
	t.Run("when invalid barcode is sent, should return an error", func(t *testing.T) {
		barcodeTemplate := 1

		barcode, err := codemapper.NewMatrixcode(barcodeTemplate)

		assert.Nil(t, barcode)
		assert.NotNil(t, err)
	})
	t.Run("when props is not sent, barcode is created", func(t *testing.T) {
		barcodeTemplate := map[string]interface{}{
			"code": "123456789",
		}

		barcode, err := codemapper.NewBarcode(barcodeTemplate)

		assert.Nil(t, err)
		assert.NotNil(t, barcode)
	})
	t.Run("when invalid props is sent, should return an error", func(t *testing.T) {
		barcodeTemplate := map[string]interface{}{
			"props": 1,
			"code":  "123456789",
		}

		barcode, err := codemapper.NewBarcode(barcodeTemplate)

		assert.Nil(t, barcode)
		assert.NotNil(t, err)
	})
	t.Run("when invalid field is sent, should return an error", func(t *testing.T) {
		barcodeTemplate := map[string]interface{}{
			"invalid_field": 1,
			"code":          "123456789",
		}

		barcode, err := codemapper.NewBarcode(barcodeTemplate)

		assert.Nil(t, barcode)
		assert.NotNil(t, err)
	})
	t.Run("when source_key and code are not sent, should return an error", func(t *testing.T) {
		barcodeTemplate := map[string]interface{}{}

		barcode, err := codemapper.NewBarcode(barcodeTemplate)

		assert.Nil(t, barcode)
		assert.NotNil(t, err)
	})
	t.Run("when invalid code is sent, should return an error", func(t *testing.T) {
		barcodeTemplate := map[string]interface{}{
			"code": 123,
		}

		barcode, err := codemapper.NewBarcode(barcodeTemplate)

		assert.Nil(t, barcode)
		assert.NotNil(t, err)
	})
	t.Run("when invalid source_key is sent, should return an error", func(t *testing.T) {
		barcodeTemplate := map[string]interface{}{
			"source_key": 123,
		}

		barcode, err := codemapper.NewBarcode(barcodeTemplate)

		assert.Nil(t, barcode)
		assert.NotNil(t, err)
	})
	t.Run("when code is not sent, should add source key", func(t *testing.T) {
		barcodeTemplate := map[string]interface{}{
			"source_key": "source",
		}

		barcode, err := codemapper.NewBarcode(barcodeTemplate)

		assert.Nil(t, err)
		assert.Equal(t, barcode.SourceKey, "source")
	})

	t.Run("when source_key is not sent, should add code", func(t *testing.T) {
		barcodeTemplate := map[string]interface{}{
			"code": "code",
		}

		barcode, err := codemapper.NewBarcode(barcodeTemplate)

		assert.Nil(t, err)
		assert.Equal(t, barcode.Code, "code")
	})
}

func TestGenerate(t *testing.T) {
	t.Run("if source key is not found, should return an error", func(t *testing.T) {
		content := map[string]interface{}{}
		provider := mocks.NewProcessorProvider(t)

		barcode := codemapper.Barcode{SourceKey: "code"}
		component, err := barcode.Generate(content, provider)

		assert.Nil(t, component)
		assert.NotNil(t, err)
	})
	t.Run("if source key content is not valid, should return an error", func(t *testing.T) {
		content := map[string]interface{}{
			"code": 1,
		}
		provider := mocks.NewProcessorProvider(t)

		barcode := codemapper.Barcode{SourceKey: "code"}
		component, err := barcode.Generate(content, provider)

		assert.Nil(t, component)
		assert.NotNil(t, err)
	})
	t.Run("If the barcode has no props, the props will not be sent", func(t *testing.T) {
		content := map[string]interface{}{
			"code": "code",
		}

		provider := mocks.NewProcessorProvider(t)
		provider.EXPECT().CreateBarCode("code").Return(nil)

		barcode := codemapper.Barcode{SourceKey: "code"}
		_, err := barcode.Generate(content, provider)

		assert.Nil(t, err)
		provider.AssertNumberOfCalls(t, "CreateBarCode", 1)
	})
	t.Run("when valid code is sent, should generate barcode", func(t *testing.T) {
		content := map[string]interface{}{}

		provider := mocks.NewProcessorProvider(t)
		provider.EXPECT().CreateBarCode("code").Return(nil)

		barcode := codemapper.Barcode{Code: "code"}
		_, err := barcode.Generate(content, provider)

		assert.Nil(t, err)
		provider.AssertNumberOfCalls(t, "CreateBarCode", 1)
	})
}
