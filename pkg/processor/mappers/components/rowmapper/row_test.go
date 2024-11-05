package rowmapper_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/rowmapper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

func TestGenerate(t *testing.T) {
	t.Run("when content no has source_key, should return an error", func(t *testing.T) {
		content := map[string]interface{}{}
		factory := mocks.NewAbstractFactoryMaps(t)
		provider := mocks.NewProcessorProvider(t)

		row := rowmapper.Row{Height: 10, Cols: make([]mappers.Componentmapper, 0), Factory: factory, SourceKey: "test"}
		newRow, err := row.Generate(content, provider)

		assert.NotNil(t, err)
		assert.Nil(t, newRow)
	})
	t.Run("when row no has row, it should no sent row", func(t *testing.T) {
		content := map[string]interface{}{"content": map[string]interface{}{"text": "value"}}
		factory := mocks.NewAbstractFactoryMaps(t)
		provider := mocks.NewProcessorProvider(t)
		provider.EXPECT().CreateRow(10.0).Return(nil, nil)

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
		provider.EXPECT().CreateRow(10.0).Return(nil, fmt.Errorf("any"))

		row := rowmapper.Row{Height: 10.0, Cols: make([]mappers.Componentmapper, 0), Factory: factory, SourceKey: "content"}
		_, err := row.Generate(content, provider)

		assert.NotNil(t, err)
	})
}
