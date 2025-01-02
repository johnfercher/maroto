// Package line implements creation of lines.
package linemapper_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/linemapper"
	"github.com/stretchr/testify/assert"
)

func TestQrcodeGetOrder(t *testing.T) {
	t.Run("when getOrder is called, should return defined order", func(t *testing.T) {
		templateRows := map[string]interface{}{
			"order": 10.0,
		}

		doc, _ := linemapper.NewLine(templateRows)

		assert.Equal(t, 10, doc.GetOrder())
	})
}

func TestNewLine(t *testing.T) {
	t.Run("when invalid line is sent, should return an error", func(t *testing.T) {
		lineTemplate := 1

		line, err := linemapper.NewLine(lineTemplate)

		assert.Nil(t, line)
		assert.NotNil(t, err)
	})
	t.Run("when props is not sent, line is created", func(t *testing.T) {
		lineTemplate := map[string]interface{}{
			"order": 1.0,
		}

		line, err := linemapper.NewLine(lineTemplate)

		assert.Nil(t, err)
		assert.NotNil(t, line)
	})
	t.Run("when invalid props is sent, should return an error", func(t *testing.T) {
		lineTemplate := map[string]interface{}{
			"props": 1,
			"order": 1.0,
		}

		line, err := linemapper.NewLine(lineTemplate)

		assert.Nil(t, line)
		assert.NotNil(t, err)
	})
	t.Run("when invalid field is sent, should return an error", func(t *testing.T) {
		lineTemplate := map[string]interface{}{
			"invalid_field": 1,
			"code":          "123456789",
			"order":         1.0,
		}

		line, err := linemapper.NewLine(lineTemplate)

		assert.Nil(t, line)
		assert.NotNil(t, err)
	})
}
