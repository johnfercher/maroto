package codemapper_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/codemapper"
	"github.com/stretchr/testify/assert"
)

func TestNewBarcode(t *testing.T) {
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
