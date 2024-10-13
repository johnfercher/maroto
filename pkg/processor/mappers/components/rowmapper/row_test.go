package rowmapper_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/rowmapper"
	"github.com/stretchr/testify/assert"
)

func TestNewRow(t *testing.T) {
	t.Run("when invalid interface is sent, should return an error", func(t *testing.T) {
		factory := mocks.NewAbstractFactoryMaps(t)
		var templateRow interface{} = 1

		doc, err := rowmapper.NewRow(templateRow, "", factory)

		assert.Nil(t, doc)
		assert.NotNil(t, err)
	})
	t.Run("when row height is not sent, should set height to 0", func(t *testing.T) {
		factory := mocks.NewAbstractFactoryMaps(t)
		var templateRow map[string]interface{}

		doc, err := rowmapper.NewRow(templateRow, "", factory)

		assert.Nil(t, err)
		assert.Equal(t, 0.0, doc.Height)
	})
	t.Run("when invalid height is sent, should return an error", func(t *testing.T) {
		factory := mocks.NewAbstractFactoryMaps(t)
		templateRow := map[string]interface{}{
			"height": "invalid",
		}

		doc, err := rowmapper.NewRow(templateRow, "", factory)

		assert.Nil(t, doc)
		assert.NotNil(t, err)
	})
	t.Run("when an invalid field is sent, should return an error", func(t *testing.T) {
		factory := mocks.NewAbstractFactoryMaps(t)
		templateRow := map[string]interface{}{
			"invalid_field": "invalid",
		}

		doc, err := rowmapper.NewRow(templateRow, "", factory)

		assert.Nil(t, doc)
		assert.NotNil(t, err)
	})
	t.Run("when an invalid col is sent, should return an error", func(t *testing.T) {
		factory := mocks.NewAbstractFactoryMaps(t)
		templateRow := map[string]interface{}{
			"cols": "invalid",
		}

		doc, err := rowmapper.NewRow(templateRow, "", factory)

		assert.Nil(t, doc)
		assert.NotNil(t, err)
	})
}
