package pagemapper_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/internal/fixture"
	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/listmapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/pagemapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/rowmapper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewPage(t *testing.T) {
	t.Run("When an invalid field is submitted, should return an error", func(t *testing.T) {
		var invalidInterface interface{} = 1
		factory := mocks.NewAbstractFactoryMaps(t)

		doc, err := pagemapper.NewPage(invalidInterface, "test", factory)

		assert.Nil(t, doc)
		assert.NotNil(t, err)
	})

	t.Run("When 2 rows are sent, should set the 2 rows", func(t *testing.T) {
		templateRows := map[string]interface{}{
			"row_template_1": nil,
			"row_template_2": nil,
		}

		validRow := fixture.MapperRow()
		factory := mocks.NewAbstractFactoryMaps(t)
		factory.On("NewRow", mock.Anything, "row_template_1").Return(validRow, nil)
		factory.On("NewRow", mock.Anything, "row_template_2").Return(validRow, nil)

		doc, err := pagemapper.NewPage(templateRows, "test", factory)

		assert.Nil(t, err)
		assert.Equal(t, 2, len(doc.Rows))
		assert.IsType(t, &rowmapper.Row{}, doc.Rows[0])
	})

	t.Run("when 1 list is sent, it should add 1 list to the document", func(t *testing.T) {
		templateRows := map[string]interface{}{
			"list_rows_1": nil,
		}

		validPage := fixture.MapperList()
		factory := mocks.NewAbstractFactoryMaps(t)
		factory.On("NewList", mock.Anything, "list_rows_1", mock.Anything).Return(validPage, nil)

		doc, err := pagemapper.NewPage(templateRows, "test", factory)

		assert.Nil(t, err)
		assert.Equal(t, len(doc.Rows), 1)
		assert.IsType(t, &listmapper.List{}, doc.Rows[0])
	})
}

func TestGenerate(t *testing.T) {
	t.Run("when content no has source_key, should return an error", func(t *testing.T) {
		content := map[string]interface{}{}
		factory := mocks.NewAbstractFactoryMaps(t)
		provider := mocks.NewProcessorProvider(t)

		page := pagemapper.Page{Rows: make([]mappers.Componentmapper, 0), Factory: factory, SourceKey: "test"}
		newPage, err := page.Generate(content, provider)

		assert.NotNil(t, err)
		assert.Nil(t, newPage)
	})
	t.Run("when page no has rows, it should no sent rows", func(t *testing.T) {
		content := map[string]interface{}{"content": map[string]interface{}{"row_1": "any"}}
		factory := mocks.NewAbstractFactoryMaps(t)
		provider := mocks.NewProcessorProvider(t)
		provider.EXPECT().CreatePage().Return(nil, nil)

		page := pagemapper.Page{Rows: make([]mappers.Componentmapper, 0), Factory: factory, SourceKey: "content"}
		_, err := page.Generate(content, provider)

		assert.Nil(t, err)
	})
	t.Run("when is not possible generate components, should return an error", func(t *testing.T) {
		content := map[string]interface{}{"content": map[string]interface{}{"text": "value"}}
		factory := mocks.NewAbstractFactoryMaps(t)
		provider := mocks.NewProcessorProvider(t)
		component := mocks.NewComponentmapper(t)
		component.EXPECT().Generate(mock.Anything, provider).Return(nil, fmt.Errorf("any"))

		page := pagemapper.Page{Rows: []mappers.Componentmapper{component}, Factory: factory, SourceKey: "content"}
		_, err := page.Generate(content, provider)

		assert.NotNil(t, err)
	})
	t.Run("when is not possible generate page, should return an error", func(t *testing.T) {
		content := map[string]interface{}{"content": map[string]interface{}{"text": "value"}}
		factory := mocks.NewAbstractFactoryMaps(t)
		provider := mocks.NewProcessorProvider(t)
		provider.EXPECT().CreatePage().Return(nil, fmt.Errorf("any"))

		page := pagemapper.Page{Rows: make([]mappers.Componentmapper, 0), Factory: factory, SourceKey: "content"}
		_, err := page.Generate(content, provider)

		assert.NotNil(t, err)
	})
}
