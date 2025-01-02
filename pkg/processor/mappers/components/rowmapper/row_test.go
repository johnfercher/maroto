package rowmapper_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/rowmapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/propsmapper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetOrder(t *testing.T) {
	t.Run("when getOrder is called, should return defined order", func(t *testing.T) {
		templateRows := map[string]interface{}{
			"order": 10.0,
		}
		factory := mocks.NewAbstractFactoryMaps(t)

		doc, _ := rowmapper.NewRow(templateRows, "test", factory)

		assert.Equal(t, 10, doc.GetOrder())
	})
}

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
		templateRow := map[string]interface{}{
			"order": 1.0,
		}

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
	t.Run("when the order field is not sent, should return an error", func(t *testing.T) {
		templateRows := map[string]interface{}{}
		factory := mocks.NewAbstractFactoryMaps(t)

		_, err := rowmapper.NewRow(templateRows, "test", factory)

		assert.NotNil(t, err)
	})
	t.Run("when the order field is less than 1, should return an error", func(t *testing.T) {
		templateRows := map[string]interface{}{
			"order": 0.0,
		}
		factory := mocks.NewAbstractFactoryMaps(t)

		_, err := rowmapper.NewRow(templateRows, "test", factory)

		assert.NotNil(t, err)
	})
	t.Run("when invalid props is sent, should return an erros", func(t *testing.T) {
		templateRow := map[string]interface{}{
			"height": 10.0,
			"props":  1,
			"order":  1.0,
		}
		factory := mocks.NewAbstractFactoryMaps(t)

		_, err := rowmapper.NewRow(templateRow, "test", factory)

		assert.NotNil(t, err)
	})
	t.Run("when valid props is sent, should create row", func(t *testing.T) {
		templateRow := map[string]interface{}{
			"height": 10.0,
			"props": map[string]interface{}{
				"line_style":  "solid",
				"border_type": "top",
			},
			"order": 1.0,
		}
		factory := mocks.NewAbstractFactoryMaps(t)

		_, err := rowmapper.NewRow(templateRow, "test_row", factory)

		assert.Nil(t, err)
	})
}

func TestGenerate(t *testing.T) {
	t.Run("when content no has source_key, should send an empty list", func(t *testing.T) {
		content := map[string]interface{}{"source_key_test": 1}
		factory := mocks.NewAbstractFactoryMaps(t)
		provider := mocks.NewProcessorProvider(t)
		provider.EXPECT().CreateRow(10.0, (*propsmapper.Cell)(nil)).Return(nil, nil)
		component := mocks.NewComponentmapper(t)
		component.EXPECT().Generate(map[string]interface{}{}, provider).Return(nil, nil)

		row := rowmapper.Row{Height: 10, Cols: []mappers.Componentmapper{component}, Factory: factory, SourceKey: "test"}
		newRow, err := row.Generate(content, provider)

		assert.NotNil(t, newRow)
		assert.Nil(t, err)
	})
	t.Run("when row no has row, it should no sent row", func(t *testing.T) {
		content := map[string]interface{}{"content": map[string]interface{}{"text": "value"}}
		factory := mocks.NewAbstractFactoryMaps(t)
		provider := mocks.NewProcessorProvider(t)
		provider.EXPECT().CreateRow(10.0, (*propsmapper.Cell)(nil)).Return(nil, nil)

		row := rowmapper.Row{Height: 10.0, Cols: make([]mappers.Componentmapper, 0), Factory: factory, SourceKey: "content"}
		_, err := row.Generate(content, provider)

		assert.Nil(t, err)
	})
	t.Run("when is not possible generate components, should return an error", func(t *testing.T) {
		content := map[string]interface{}{"content": map[string]interface{}{"text": "value"}}
		factory := mocks.NewAbstractFactoryMaps(t)
		provider := mocks.NewProcessorProvider(t)
		component := mocks.NewComponentmapper(t)
		component.EXPECT().Generate(mock.Anything, provider).Return(nil, fmt.Errorf("any"))

		row := rowmapper.Row{Height: 10.0, Cols: []mappers.Componentmapper{component}, Factory: factory, SourceKey: "content"}
		_, err := row.Generate(content, provider)

		assert.NotNil(t, err)
	})
	t.Run("when is not possible generate row, should return an error", func(t *testing.T) {
		content := map[string]interface{}{"content": map[string]interface{}{"text": "value"}}
		factory := mocks.NewAbstractFactoryMaps(t)
		provider := mocks.NewProcessorProvider(t)
		provider.EXPECT().CreateRow(10.0, (*propsmapper.Cell)(nil)).Return(nil, fmt.Errorf("any"))

		row := rowmapper.Row{Height: 10.0, Cols: make([]mappers.Componentmapper, 0), Factory: factory, SourceKey: "content"}
		_, err := row.Generate(content, provider)

		assert.NotNil(t, err)
	})
}
