package pagemapper_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/pagemapper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetOrder(t *testing.T) {
	t.Run("when getOrder is called, should return defined order", func(t *testing.T) {
		templateRows := map[string]interface{}{
			"order": 10.0,
		}
		doc, _ := pagemapper.NewPage(templateRows, "test", mocks.NewAbstractFactoryMaps(t))

		assert.Equal(t, 10, doc.GetOrder())
	})
}

func TestNewPage(t *testing.T) {

	t.Run("when the component order is greater than the number of components, an error should be returned", func(t *testing.T) {
		templateRows := map[string]interface{}{
			"row_template_1": nil,
			"order":          1.0,
		}

		orderedComponent := mocks.NewOrderedComponents(t)
		orderedComponent.EXPECT().GetOrder().Return(2)
		factory := mocks.NewAbstractFactoryMaps(t)
		factory.On("NewRow", mock.Anything, "row_template_1").Return(orderedComponent, nil)

		_, err := pagemapper.NewPage(templateRows, "test", factory)

		assert.NotNil(t, err)
	})

	t.Run("when the template order is repeated, an error should be returned", func(t *testing.T) {
		templateRows := map[string]interface{}{
			"row_template_1": nil,
			"row_template_2": nil,
			"order":          1.0,
		}

		orderedComponent := mocks.NewOrderedComponents(t)
		orderedComponent.EXPECT().GetOrder().Return(1)

		factory := mocks.NewAbstractFactoryMaps(t)
		factory.On("NewRow", mock.Anything, "row_template_1").Return(orderedComponent, nil)
		factory.On("NewRow", mock.Anything, "row_template_2").Return(orderedComponent, nil)

		_, err := pagemapper.NewPage(templateRows, "test", factory)

		assert.NotNil(t, err)
	})

	t.Run("when 2 rows are submitted, should add the 2 rows in the correct order", func(t *testing.T) {
		templateRows := map[string]interface{}{
			"row_template_2": nil,
			"row_template_1": nil,
			"order":          1.0,
		}

		orderedComponent1 := mocks.NewOrderedComponents(t)
		orderedComponent1.EXPECT().GetOrder().Return(1)
		orderedComponent2 := mocks.NewOrderedComponents(t)
		orderedComponent2.EXPECT().GetOrder().Return(2)

		factory := mocks.NewAbstractFactoryMaps(t)
		factory.On("NewRow", mock.Anything, "row_template_1").Return(orderedComponent1, nil)
		factory.On("NewRow", mock.Anything, "row_template_2").Return(orderedComponent2, nil)

		doc, err := pagemapper.NewPage(templateRows, "test", factory)

		assert.Nil(t, err)
		assert.Equal(t, len(doc.Rows), 2)
		assert.Equal(t, orderedComponent1, doc.Rows[0])
	})

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
			"order":          1.0,
		}

		orderedComponent1 := mocks.NewOrderedComponents(t)
		orderedComponent1.EXPECT().GetOrder().Return(1)
		orderedComponent2 := mocks.NewOrderedComponents(t)
		orderedComponent2.EXPECT().GetOrder().Return(2)

		factory := mocks.NewAbstractFactoryMaps(t)
		factory.On("NewRow", mock.Anything, "row_template_1").Return(orderedComponent1, nil)
		factory.On("NewRow", mock.Anything, "row_template_2").Return(orderedComponent2, nil)

		doc, err := pagemapper.NewPage(templateRows, "test", factory)

		assert.Nil(t, err)
		assert.Equal(t, 2, len(doc.Rows))
	})

	t.Run("when 1 list is sent, it should add 1 list to the document", func(t *testing.T) {
		templateRows := map[string]interface{}{
			"list_rows_1": nil,
			"order":       1.0,
		}

		orderedComponent1 := mocks.NewOrderedComponents(t)
		orderedComponent1.EXPECT().GetOrder().Return(1)
		factory := mocks.NewAbstractFactoryMaps(t)
		factory.On("NewList", mock.Anything, "list_rows_1", mock.Anything).Return(orderedComponent1, nil)

		doc, err := pagemapper.NewPage(templateRows, "test", factory)

		assert.Nil(t, err)
		assert.Equal(t, len(doc.Rows), 1)
	})

	t.Run("when the order field is not sent, should return an error", func(t *testing.T) {
		templateRows := map[string]interface{}{
			"row_template_1": nil,
		}
		factory := mocks.NewAbstractFactoryMaps(t)

		_, err := pagemapper.NewPage(templateRows, "test", factory)

		assert.NotNil(t, err)
	})

	t.Run("when the order field is less than 1, should return an error", func(t *testing.T) {
		templateRows := map[string]interface{}{
			"row_template_1": nil,
			"order":          0,
		}
		factory := mocks.NewAbstractFactoryMaps(t)

		_, err := pagemapper.NewPage(templateRows, "test", factory)

		assert.NotNil(t, err)
	})
}

func TestGenerate(t *testing.T) {
	t.Run("when content no has source_key, should return an error", func(t *testing.T) {
		content := map[string]interface{}{}
		factory := mocks.NewAbstractFactoryMaps(t)
		provider := mocks.NewProcessorProvider(t)

		page := pagemapper.Page{Rows: make([]mappers.OrderedComponents, 0), Factory: factory, SourceKey: "test"}
		newPage, err := page.Generate(content, provider)

		assert.NotNil(t, err)
		assert.Nil(t, newPage)
	})
	t.Run("when page no has rows, it should no sent rows", func(t *testing.T) {
		content := map[string]interface{}{"content": map[string]interface{}{"row_1": "any"}}
		factory := mocks.NewAbstractFactoryMaps(t)
		provider := mocks.NewProcessorProvider(t)
		provider.EXPECT().CreatePage().Return(nil, nil)

		page := pagemapper.Page{Rows: make([]mappers.OrderedComponents, 0), Factory: factory, SourceKey: "content"}
		_, err := page.Generate(content, provider)

		assert.Nil(t, err)
	})
	t.Run("when is not possible generate components, should return an error", func(t *testing.T) {
		content := map[string]interface{}{"content": map[string]interface{}{"text": "value"}}
		factory := mocks.NewAbstractFactoryMaps(t)
		provider := mocks.NewProcessorProvider(t)
		component := mocks.NewOrderedComponents(t)
		component.EXPECT().Generate(mock.Anything, provider).Return(nil, fmt.Errorf("any"))

		page := pagemapper.Page{Rows: []mappers.OrderedComponents{component}, Factory: factory, SourceKey: "content"}
		_, err := page.Generate(content, provider)

		assert.NotNil(t, err)
	})
	t.Run("when is not possible generate page, should return an error", func(t *testing.T) {
		content := map[string]interface{}{"content": map[string]interface{}{"text": "value"}}
		factory := mocks.NewAbstractFactoryMaps(t)
		provider := mocks.NewProcessorProvider(t)
		provider.EXPECT().CreatePage().Return(nil, fmt.Errorf("any"))

		page := pagemapper.Page{Rows: make([]mappers.OrderedComponents, 0), Factory: factory, SourceKey: "content"}
		_, err := page.Generate(content, provider)

		assert.NotNil(t, err)
	})
}
