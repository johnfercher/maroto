// nolint:dupl
package codemapper_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/mocks"
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

func TestQrCodeGenerate(t *testing.T) {
	t.Run("if source key is not found, should return an error", func(t *testing.T) {
		content := map[string]interface{}{}
		provider := mocks.NewProcessorProvider(t)

		QrCode := codemapper.Qrcode{SourceKey: "code"}
		component, err := QrCode.Generate(content, provider)

		assert.Nil(t, component)
		assert.NotNil(t, err)
	})
	t.Run("if source key content is not valid, should return an error", func(t *testing.T) {
		content := map[string]interface{}{
			"code": 1,
		}
		provider := mocks.NewProcessorProvider(t)

		QrCode := codemapper.Qrcode{SourceKey: "code"}
		component, err := QrCode.Generate(content, provider)

		assert.Nil(t, component)
		assert.NotNil(t, err)
	})
	t.Run("If the QrCode has no props, the props will not be sent", func(t *testing.T) {
		content := map[string]interface{}{
			"code": "code",
		}

		provider := mocks.NewProcessorProvider(t)
		provider.EXPECT().CreateQrCode("code").Return(nil)

		QrCode := codemapper.Qrcode{SourceKey: "code"}
		_, err := QrCode.Generate(content, provider)

		assert.Nil(t, err)
		provider.AssertNumberOfCalls(t, "CreateQrCode", 1)
	})
	t.Run("when valid code is sent, should generate QrCode", func(t *testing.T) {
		content := map[string]interface{}{}

		provider := mocks.NewProcessorProvider(t)
		provider.EXPECT().CreateQrCode("code").Return(nil)

		QrCode := codemapper.Qrcode{Code: "code"}
		_, err := QrCode.Generate(content, provider)

		assert.Nil(t, err)
		provider.AssertNumberOfCalls(t, "CreateQrCode", 1)
	})
}
