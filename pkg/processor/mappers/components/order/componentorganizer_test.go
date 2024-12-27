package order_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/order"
	"github.com/stretchr/testify/assert"
)

func TestN(t *testing.T) {
	t.Run("when the order field is not sent, should return an error", func(t *testing.T) {
		template := map[string]interface{}{
			"template_1": nil,
		}

		_, err := order.SetPageOrder(&template, "resource_name", "resource_key")

		assert.NotNil(t, err)
	})
	t.Run("when the order field is less than 1, should return an error", func(t *testing.T) {
		template := map[string]interface{}{
			"row_template_1": nil,
			"order":          0.0,
		}

		_, err := order.SetPageOrder(&template, "resource_name", "resource_key")

		assert.NotNil(t, err)
	})
	t.Run("when the order field found, should remove order field", func(t *testing.T) {
		template := map[string]interface{}{
			"row_template_1": nil,
			"order":          1.0,
		}

		_, err := order.SetPageOrder(&template, "resource_name", "resource_key")
		_, ok := template["order"]

		assert.Nil(t, err)
		assert.False(t, ok)
	})
}
