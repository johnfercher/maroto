package listmapper_test

import (
	"errors"
	"testing"

	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/listmapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/processorprovider"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetOrder(t *testing.T) {
	t.Run("when getOrder is called, should return defined order", func(t *testing.T) {
		templateRows := map[string]interface{}{
			"order": 10.0,
		}
		factory := mocks.NewAbstractFactoryMaps(t)

		doc, _ := listmapper.NewList(templateRows, "test", factory.NewRow)

		assert.Equal(t, 10, doc.GetOrder())
	})
}

func TestNewList(t *testing.T) {
	t.Run("when invalid interface is sent, it should return an error", func(t *testing.T) {
		var invalidInterface interface{} = 1
		factory := mocks.NewAbstractFactoryMaps(t)

		doc, err := listmapper.NewList(invalidInterface, "test", factory.NewPage)

		assert.Nil(t, doc)
		assert.NotNil(t, err)
	})

	t.Run("when the component order is greater than the number of components, an error should be returned", func(t *testing.T) {
		templateList := map[string]interface{}{
			"page_template_1": nil,
			"order":           1.0,
		}

		orderedComponent := mocks.NewOrderedComponents(t)
		orderedComponent.EXPECT().GetOrder().Return(2)
		factory := mocks.NewAbstractFactoryMaps(t)
		factory.On("NewPage", mock.Anything, "page_template_1").Return(orderedComponent, nil)

		_, err := listmapper.NewList(templateList, "test", factory.NewPage)

		assert.NotNil(t, err)
	})

	t.Run("when the template order is repeated, an error should be returned", func(t *testing.T) {
		templateList := map[string]interface{}{
			"page_template_1": nil,
			"page_template_2": nil,
			"order":           1.0,
		}

		orderedComponent := mocks.NewOrderedComponents(t)
		orderedComponent.EXPECT().GetOrder().Return(2)
		factory := mocks.NewAbstractFactoryMaps(t)
		factory.On("NewPage", mock.Anything, "page_template_1").Return(orderedComponent, nil)
		factory.On("NewPage", mock.Anything, "page_template_2").Return(orderedComponent, nil)

		_, err := listmapper.NewList(templateList, "test", factory.NewPage)

		assert.NotNil(t, err)
	})

	t.Run("when 2 pages are submitted, should add the 2 pages in the correct order", func(t *testing.T) {
		templatePages := map[string]interface{}{
			"page_template_2": nil,
			"page_template_1": nil,
			"order":           1.0,
		}
		orderedComponent1 := mocks.NewOrderedComponents(t)
		orderedComponent1.EXPECT().GetOrder().Return(1)
		orderedComponent2 := mocks.NewOrderedComponents(t)
		orderedComponent2.EXPECT().GetOrder().Return(2)

		factory := mocks.NewAbstractFactoryMaps(t)
		factory.On("NewPage", mock.Anything, "page_template_1").Return(orderedComponent1, nil)
		factory.On("NewPage", mock.Anything, "page_template_2").Return(orderedComponent2, nil)

		doc, err := listmapper.NewList(templatePages, "test", factory.NewPage)

		assert.Nil(t, err)
		assert.Equal(t, len(doc.Templates), 2)
		assert.Equal(t, orderedComponent1, doc.Templates[0])
	})

	t.Run("when component not can generate, it should return an error", func(t *testing.T) {
		templatePages := map[string]interface{}{
			"page_template_1": nil,
			"page_template_2": nil,
			"order":           1.0,
		}

		factory := mocks.NewAbstractFactoryMaps(t)
		factory.On("NewPage", mock.Anything, "page_template_1").Return(nil, errors.New(""))
		doc, err := listmapper.NewList(templatePages, "test", factory.NewPage)

		assert.Nil(t, doc)
		assert.NotNil(t, err)
	})

	t.Run("when the order field is not sent, should return an error", func(t *testing.T) {
		templateRows := map[string]interface{}{
			"row_template_1": nil,
		}
		factory := mocks.NewAbstractFactoryMaps(t)

		_, err := listmapper.NewList(templateRows, "test", factory.NewRow)

		assert.NotNil(t, err)
	})

	t.Run("when the order field is less than 1, should return an error", func(t *testing.T) {
		templateRows := map[string]interface{}{
			"row_template_1": nil,
			"order":          0,
		}
		factory := mocks.NewAbstractFactoryMaps(t)

		_, err := listmapper.NewList(templateRows, "test", factory.NewRow)

		assert.NotNil(t, err)
	})

	t.Run("when 2-components are sent, it should add 2 components in list", func(t *testing.T) {
		templatePages := map[string]interface{}{
			"page_template_1": nil,
			"page_template_2": nil,
			"order":           1.0,
		}
		orderedComponent1 := mocks.NewOrderedComponents(t)
		orderedComponent1.EXPECT().GetOrder().Return(1)
		orderedComponent2 := mocks.NewOrderedComponents(t)
		orderedComponent2.EXPECT().GetOrder().Return(2)

		factory := mocks.NewAbstractFactoryMaps(t)
		factory.On("NewPage", mock.Anything, "page_template_1").Return(orderedComponent1, nil)
		factory.On("NewPage", mock.Anything, "page_template_2").Return(orderedComponent2, nil)

		doc, err := listmapper.NewList(templatePages, "test", factory.NewPage)

		assert.Nil(t, err)
		assert.Equal(t, len(doc.Templates), 2)
	})
}

func TestGenerate(t *testing.T) {
	t.Run("when source_key is not found, should return an error", func(t *testing.T) {
		provider := mocks.NewProcessorProvider(t)
		content := map[string]interface{}{}
		list := listmapper.List{SourceKey: "list", Templates: make([]mappers.OrderedComponents, 0)}

		components, err := list.Generate(content, provider)

		assert.Nil(t, components)
		assert.NotNil(t, err)
	})
	t.Run("when invalid content is sent, should return an error", func(t *testing.T) {
		provider := mocks.NewProcessorProvider(t)
		content := map[string]interface{}{"list": 1}
		list := listmapper.List{SourceKey: "list", Templates: make([]mappers.OrderedComponents, 0)}

		components, err := list.Generate(content, provider)

		assert.Nil(t, components)
		assert.NotNil(t, err)
	})

	t.Run("when 2 templates are added, it should generate 4 components", func(t *testing.T) {
		content1 := map[string]interface{}{"row_1": nil, "row_2": nil}
		content2 := map[string]interface{}{"row_1": nil, "row_2": nil}
		listContent := map[string]interface{}{"list": []interface{}{content1, content2}}
		provider := mocks.NewProcessorProvider(t)
		providerComponent := mocks.NewProviderComponent(t)
		component := mocks.NewOrderedComponents(t)
		component.EXPECT().Generate(mock.Anything, provider).Return([]processorprovider.ProviderComponent{providerComponent}, nil)

		list := listmapper.List{SourceKey: "list", Templates: []mappers.OrderedComponents{component, component}}
		components, err := list.Generate(listContent, provider)

		assert.NotNil(t, components)
		assert.Nil(t, err)
		component.AssertNumberOfCalls(t, "Generate", 4)
		assert.Len(t, components, 4)
	})
}
