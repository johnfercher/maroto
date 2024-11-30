package listmapper_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/internal/fixture"
	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/listmapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/processorprovider"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewList(t *testing.T) {
	t.Run("when invalid interface is sent, it should return an error", func(t *testing.T) {
		var invalidInterface interface{} = 1
		factory := mocks.NewAbstractFactoryMaps(t)

		doc, err := listmapper.NewList(invalidInterface, "test", factory.NewPage)

		assert.Nil(t, doc)
		assert.NotNil(t, err)
	})

	t.Run("when component not can generate, it should return an error", func(t *testing.T) {
		templatePages := map[string]interface{}{
			"page_template_1": nil,
			"page_template_2": nil,
		}

		factory := mocks.NewAbstractFactoryMaps(t)
		factory.On("NewPage", mock.Anything, "page_template_1").Return(nil, errors.New(""))
		doc, err := listmapper.NewList(templatePages, "test", factory.NewPage)

		assert.Nil(t, doc)
		assert.NotNil(t, err)
	})

	t.Run("when 2-components are sent, it should add 2 components in list", func(t *testing.T) {
		validPage := fixture.MapperPage()
		templatePages := map[string]interface{}{
			"page_template_1": nil,
			"page_template_2": nil,
		}
		factory := mocks.NewAbstractFactoryMaps(t)
		factory.On("NewPage", mock.Anything, "page_template_1").Return(validPage, nil)
		factory.On("NewPage", mock.Anything, "page_template_2").Return(validPage, nil)

		doc, err := listmapper.NewList(templatePages, "test", factory.NewPage)

		assert.Nil(t, err)
		assert.Equal(t, len(doc.Templates), 2)
	})
}

func TestGenerate(t *testing.T) {
	t.Run("when source_key is not found, should return an error", func(t *testing.T) {
		provider := mocks.NewProcessorProvider(t)
		content := map[string]interface{}{}
		list := listmapper.List{SourceKey: "list", Templates: make([]mappers.Componentmapper, 0)}

		components, err := list.Generate(content, provider)

		assert.Nil(t, components)
		assert.NotNil(t, err)
	})
	t.Run("when invalid content is sent, should return an error", func(t *testing.T) {
		provider := mocks.NewProcessorProvider(t)
		content := map[string]interface{}{"list": 1}
		list := listmapper.List{SourceKey: "list", Templates: make([]mappers.Componentmapper, 0)}

		components, err := list.Generate(content, provider)

		assert.Nil(t, components)
		assert.NotNil(t, err)
	})

	t.Run("when components is not generate, should return an error", func(t *testing.T) {
		contentRow1 := map[string]interface{}{"row_1": nil}
		listContent := map[string]interface{}{"list": []interface{}{contentRow1}}
		provider := mocks.NewProcessorProvider(t)
		component := mocks.NewComponentmapper(t)
		component.EXPECT().Generate(mock.Anything, provider).Return(nil, fmt.Errorf("any"))

		list := listmapper.List{SourceKey: "list", Templates: []mappers.Componentmapper{component}}
		components, err := list.Generate(listContent, provider)

		assert.Nil(t, components)
		assert.NotNil(t, err)
	})

	t.Run("when 2 templates are added, it should generate 4 components", func(t *testing.T) {
		content1 := map[string]interface{}{"row_1": nil, "row_2": nil}
		content2 := map[string]interface{}{"row_1": nil, "row_2": nil}
		listContent := map[string]interface{}{"list": []interface{}{content1, content2}}
		provider := mocks.NewProcessorProvider(t)
		providerComponent := mocks.NewProviderComponent(t)
		component := mocks.NewComponentmapper(t)
		component.EXPECT().Generate(mock.Anything, provider).Return([]processorprovider.ProviderComponent{providerComponent}, nil)

		list := listmapper.List{SourceKey: "list", Templates: []mappers.Componentmapper{component, component}}
		components, err := list.Generate(listContent, provider)

		assert.NotNil(t, components)
		assert.Nil(t, err)
		component.AssertNumberOfCalls(t, "Generate", 4)
		assert.Len(t, components, 4)
	})
}
