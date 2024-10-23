package codemapper_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/codemapper"
	"github.com/stretchr/testify/assert"
)

func TestNewQrCode(t *testing.T) {
	t.Run("when invalid qrcode is sent, should return an error", func(t *testing.T) {
		qrcodeTemplate := 1

		qrcode, err := codemapper.NewQrcode(qrcodeTemplate)

		assert.Nil(t, qrcode)
		assert.NotNil(t, err)
	})
	t.Run("when props is not sent, should create qrcode", func(t *testing.T) {
		qrcodeTemplate := map[string]interface{}{
			"code": "123456789",
		}

		qrcode, err := codemapper.NewQrcode(qrcodeTemplate)

		assert.Nil(t, err)
		assert.NotNil(t, qrcode)
	})
	t.Run("when invalid props is sent, should return an error", func(t *testing.T) {
		qrcodeTemplate := map[string]interface{}{
			"props": 1,
			"code":  "123456789",
		}

		qrcode, err := codemapper.NewQrcode(qrcodeTemplate)

		assert.Nil(t, qrcode)
		assert.NotNil(t, err)
	})
	t.Run("when invalid field is sent, should return an error", func(t *testing.T) {
		qrcodeTemplate := map[string]interface{}{
			"invalid_field": 1,
			"code":          "123456789",
		}

		qrcode, err := codemapper.NewQrcode(qrcodeTemplate)

		assert.Nil(t, qrcode)
		assert.NotNil(t, err)
	})
	t.Run("when source_key and code are not sent, should return an error", func(t *testing.T) {
		qrcodeTemplate := map[string]interface{}{}

		qrcode, err := codemapper.NewQrcode(qrcodeTemplate)

		assert.Nil(t, qrcode)
		assert.NotNil(t, err)
	})
	t.Run("when invalid code is sent, should return an error", func(t *testing.T) {
		qrcodeTemplate := map[string]interface{}{
			"code": 123,
		}

		qrcode, err := codemapper.NewQrcode(qrcodeTemplate)

		assert.Nil(t, qrcode)
		assert.NotNil(t, err)
	})
	t.Run("when invalid source_key is sent, should return an error", func(t *testing.T) {
		qrcodeTemplate := map[string]interface{}{
			"source_key": 123,
		}

		qrcode, err := codemapper.NewQrcode(qrcodeTemplate)

		assert.Nil(t, qrcode)
		assert.NotNil(t, err)
	})
	t.Run("when code is not sent, should add source key", func(t *testing.T) {
		qrcodeTemplate := map[string]interface{}{
			"source_key": "source",
		}

		qrcode, err := codemapper.NewQrcode(qrcodeTemplate)

		assert.Nil(t, err)
		assert.Equal(t, qrcode.SourceKey, "source")
	})

	t.Run("when source_key is not sent, should add code", func(t *testing.T) {
		qrcodeTemplate := map[string]interface{}{
			"code": "code",
		}

		qrcode, err := codemapper.NewQrcode(qrcodeTemplate)

		assert.Nil(t, err)
		assert.Equal(t, qrcode.Code, "code")
	})
}
