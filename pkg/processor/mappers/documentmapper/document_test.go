package documentmapper_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/internal/fixture"
	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/processor/deserializer"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/documentmapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/processorprovider"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewPdf(t *testing.T) {
	t.Run("when builder is sent, should set builder", func(t *testing.T) {
		builderDocument := `
			{
				"builder": {"concurrent_mode": 10}
			}
		`
		factory := mocks.NewAbstractFactoryMaps(t)

		template, err := deserializer.NewJSONDeserializer().Deserialize(builderDocument)
		assert.Nil(t, err)

		doc, err := documentmapper.NewPdf(template, factory)
		assert.Nil(t, err)
		assert.Equal(t, doc.Builder.ConcurrentMode, 10)
	})

	t.Run("when an invalid builder is passed, should return an error", func(t *testing.T) {
		builderDocument := `{"builder": 10}`
		factory := mocks.NewAbstractFactoryMaps(t)

		template, err := deserializer.NewJSONDeserializer().Deserialize(builderDocument)
		assert.Nil(t, err)

		doc, err := documentmapper.NewPdf(template, factory)

		assert.NotNil(t, err)
		assert.Nil(t, doc)
	})

	t.Run("When a 2-rows header is sent, should set the header", func(t *testing.T) {
		builderDocument := `
			{"header": {
				"row_template_1": {},
				"row_template_2": {}
			}}
		`
		validRow := fixture.MapperRow()
		factory := mocks.NewAbstractFactoryMaps(t)
		factory.On("NewRow", mock.Anything, mock.Anything).Return(validRow, nil)

		template, err := deserializer.NewJSONDeserializer().Deserialize(builderDocument)
		assert.Nil(t, err)

		doc, err := documentmapper.NewPdf(template, factory)

		assert.Nil(t, err)
		assert.Equal(t, 2, len(doc.Header))
	})

	t.Run("when an invalid header is passed, should return an error", func(t *testing.T) {
		builderDocument := `{"header": 1}`
		factory := mocks.NewAbstractFactoryMaps(t)

		template, err := deserializer.NewJSONDeserializer().Deserialize(builderDocument)
		assert.Nil(t, err)

		_, err = documentmapper.NewPdf(template, factory)
		assert.NotNil(t, err)
	})

	t.Run("When a 2-rows footer is sent, should set the footer", func(t *testing.T) {
		builderDocument := `
			{"footer": {
				"row_template_1": {},
				"row_template_2": {}
			}}
		`
		validRow := fixture.MapperRow()
		factory := mocks.NewAbstractFactoryMaps(t)
		factory.On("NewRow", mock.Anything, mock.Anything).Return(validRow, nil)

		template, err := deserializer.NewJSONDeserializer().Deserialize(builderDocument)
		assert.Nil(t, err)

		doc, err := documentmapper.NewPdf(template, factory)

		assert.Nil(t, err)
		assert.Equal(t, 2, len(doc.Footer))
	})

	t.Run("when an invalid footer is passed, should return an error", func(t *testing.T) {
		builderDocument := `{"footer": 1}`
		factory := mocks.NewAbstractFactoryMaps(t)

		template, err := deserializer.NewJSONDeserializer().Deserialize(builderDocument)
		assert.Nil(t, err)

		_, err = documentmapper.NewPdf(template, factory)

		assert.NotNil(t, err)
	})

	t.Run("when the template order is greater than the number of pages, an error should be returned", func(t *testing.T) {
		builderDocument := `{"pages": {"page_template_1":{}}}`
		orderedComponent := mocks.NewOrderedComponents(t)
		orderedComponent.EXPECT().GetOrder().Return(2)

		factory := mocks.NewAbstractFactoryMaps(t)
		factory.On("NewPage", mock.Anything, "page_template_1").Return(orderedComponent, nil)

		template, err := deserializer.NewJSONDeserializer().Deserialize(builderDocument)
		assert.Nil(t, err)

		_, err = documentmapper.NewPdf(template, factory)

		assert.NotNil(t, err)
	})

	t.Run("when the template order is repeated, an error should be returned", func(t *testing.T) {
		builderDocument := `{"pages": {"page_template_2":{}, "page_template_1":{}}}`
		orderedComponent := mocks.NewOrderedComponents(t)
		orderedComponent.EXPECT().GetOrder().Return(2)

		factory := mocks.NewAbstractFactoryMaps(t)
		factory.On("NewPage", mock.Anything, "page_template_1").Return(orderedComponent, nil)
		factory.On("NewPage", mock.Anything, "page_template_2").Return(orderedComponent, nil)

		template, err := deserializer.NewJSONDeserializer().Deserialize(builderDocument)
		assert.Nil(t, err)

		_, err = documentmapper.NewPdf(template, factory)

		assert.NotNil(t, err)
	})

	t.Run("when 2 pages are submitted, should add the 2 pages in the correct order", func(t *testing.T) {
		builderDocument := `
			{"pages": {
				"page_template_1":{},
				"page_template_2":{} 
			}}
		`
		orderedComponent1 := mocks.NewOrderedComponents(t)
		orderedComponent1.EXPECT().GetOrder().Return(1)
		orderedComponent2 := mocks.NewOrderedComponents(t)
		orderedComponent2.EXPECT().GetOrder().Return(2)

		factory := mocks.NewAbstractFactoryMaps(t)
		factory.On("NewPage", mock.Anything, "page_template_1").Return(orderedComponent1, nil)
		factory.On("NewPage", mock.Anything, "page_template_2").Return(orderedComponent2, nil)

		template, err := deserializer.NewJSONDeserializer().Deserialize(builderDocument)
		assert.Nil(t, err)

		doc, err := documentmapper.NewPdf(template, factory)

		assert.Nil(t, err)
		assert.Equal(t, len(doc.Pages), 2)
		assert.Equal(t, orderedComponent1, doc.Pages[0])
	})

	t.Run("when 2 pages are sent, it should add 2 pages to the document", func(t *testing.T) {
		builderDocument := `
			{"pages": {
				"page_template_1":{},
				"page_template_2":{} 
			}}
		`
		orderedComponent1 := mocks.NewOrderedComponents(t)
		orderedComponent1.EXPECT().GetOrder().Return(1)
		orderedComponent2 := mocks.NewOrderedComponents(t)
		orderedComponent2.EXPECT().GetOrder().Return(2)

		factory := mocks.NewAbstractFactoryMaps(t)
		factory.On("NewPage", mock.Anything, "page_template_1").Return(orderedComponent1, nil)
		factory.On("NewPage", mock.Anything, "page_template_2").Return(orderedComponent2, nil)

		template, err := deserializer.NewJSONDeserializer().Deserialize(builderDocument)
		assert.Nil(t, err)

		doc, err := documentmapper.NewPdf(template, factory)

		assert.Nil(t, err)
		assert.Equal(t, len(doc.Pages), 2)
	})

	t.Run("when 1 list is sent, it should add 1 list to the document", func(t *testing.T) {
		builderDocument := `
			{"pages": {
				"list_template_1":{}
			}}
		`
		orderedComponent1 := mocks.NewOrderedComponents(t)
		orderedComponent1.EXPECT().GetOrder().Return(1)

		factory := mocks.NewAbstractFactoryMaps(t)
		factory.On("NewList", mock.Anything, "list_template_1", mock.Anything).Return(orderedComponent1, nil)

		template, err := deserializer.NewJSONDeserializer().Deserialize(builderDocument)
		assert.Nil(t, err)

		doc, err := documentmapper.NewPdf(template, factory)

		assert.Nil(t, err)
		assert.Equal(t, len(doc.Pages), 1)
	})

	t.Run("when an invalid page is sent, it should return an error", func(t *testing.T) {
		builderDocument := `{"pages": 1}`
		factory := mocks.NewAbstractFactoryMaps(t)

		template, err := deserializer.NewJSONDeserializer().Deserialize(builderDocument)
		assert.Nil(t, err)

		_, err = documentmapper.NewPdf(template, factory)
		assert.NotNil(t, err)
	})
}

func TestGenerate(t *testing.T) {
	t.Run("when document have no template, should not set template ", func(t *testing.T) {
		fixContent := make(map[string]interface{})
		mockProvider := mocks.NewProcessorProvider(t)

		doc := documentmapper.Document{
			Pages: make([]mappers.Componentmapper, 0), Header: make([]mappers.Componentmapper, 0),
			Footer: make([]mappers.Componentmapper, 0),
		}
		provider, err := doc.Generate(fixContent, mockProvider)

		assert.Nil(t, err)
		assert.NotNil(t, provider)
	})
	t.Run("when templates have no content, should generate templates with empty content", func(t *testing.T) {
		fixContent := map[string]interface{}{"test": "any"}

		mockProviderComponents := []processorprovider.ProviderComponent{mocks.NewProviderComponent(t)}
		mockComponent := mocks.NewComponentmapper(t)
		mockComponent.EXPECT().Generate(map[string]interface{}{}, mock.Anything).Return(mockProviderComponents, nil)
		mockProvider := mocks.NewProcessorProvider(t)
		mockProvider.EXPECT().AddPages(mockProviderComponents[0]).Return(mockProvider, nil)
		mockProvider.EXPECT().AddHeader(mockProviderComponents[0]).Return(mockProvider, nil)
		mockProvider.EXPECT().AddFooter(mockProviderComponents[0]).Return(mockProvider, nil)

		doc := documentmapper.Document{
			Pages: []mappers.Componentmapper{mockComponent}, Header: []mappers.Componentmapper{mockComponent},
			Footer: []mappers.Componentmapper{mockComponent},
		}
		provider, err := doc.Generate(fixContent, mockProvider)

		assert.Nil(t, err)
		assert.NotNil(t, provider)
	})
	t.Run("when it is not possible to generate the page, it should return an error", func(t *testing.T) {
		fixContent := map[string]interface{}{}

		mockComponent := mocks.NewComponentmapper(t)
		mockComponent.EXPECT().Generate(map[string]interface{}{}, mock.Anything).Return(nil, fmt.Errorf("any"))
		mockProvider := mocks.NewProcessorProvider(t)

		doc := documentmapper.Document{
			Pages: []mappers.Componentmapper{mockComponent}, Header: []mappers.Componentmapper{},
			Footer: []mappers.Componentmapper{},
		}
		provider, err := doc.Generate(fixContent, mockProvider)

		assert.Nil(t, provider)
		assert.NotNil(t, err)
	})
	t.Run("when it is not possible to generate the header, it should return an error", func(t *testing.T) {
		fixContent := map[string]interface{}{}

		mockComponent := mocks.NewComponentmapper(t)
		mockComponent.EXPECT().Generate(map[string]interface{}{}, mock.Anything).Return(nil, fmt.Errorf("any"))
		mockProvider := mocks.NewProcessorProvider(t)

		doc := documentmapper.Document{
			Pages: []mappers.Componentmapper{}, Header: []mappers.Componentmapper{mockComponent},
			Footer: []mappers.Componentmapper{},
		}
		provider, err := doc.Generate(fixContent, mockProvider)

		assert.Nil(t, provider)
		assert.NotNil(t, err)
	})
	t.Run("when it is not possible to generate the footer, it should return an error", func(t *testing.T) {
		fixContent := map[string]interface{}{}

		mockComponent := mocks.NewComponentmapper(t)
		mockComponent.EXPECT().Generate(map[string]interface{}{}, mock.Anything).Return(nil, fmt.Errorf("any"))
		mockProvider := mocks.NewProcessorProvider(t)

		doc := documentmapper.Document{
			Pages: []mappers.Componentmapper{}, Header: []mappers.Componentmapper{},
			Footer: []mappers.Componentmapper{mockComponent},
		}
		provider, err := doc.Generate(fixContent, mockProvider)

		assert.Nil(t, provider)
		assert.NotNil(t, err)
	})
	t.Run("when it is not possible add the page to the document, it should return an error", func(t *testing.T) {
		fixContent := map[string]interface{}{}

		mockProviderComponents := []processorprovider.ProviderComponent{mocks.NewProviderComponent(t)}
		mockComponent := mocks.NewComponentmapper(t)
		mockComponent.EXPECT().Generate(map[string]interface{}{}, mock.Anything).Return(mockProviderComponents, nil)
		mockProvider := mocks.NewProcessorProvider(t)
		mockProvider.EXPECT().AddPages(mockProviderComponents[0]).Return(nil, fmt.Errorf("any"))

		doc := documentmapper.Document{
			Pages: []mappers.Componentmapper{mockComponent}, Header: []mappers.Componentmapper{},
			Footer: []mappers.Componentmapper{},
		}
		provider, err := doc.Generate(fixContent, mockProvider)

		assert.Nil(t, provider)
		assert.NotNil(t, err)
	})
	t.Run("when it is not possible add the header to the document, it should return an error", func(t *testing.T) {
		fixContent := map[string]interface{}{}

		mockProviderComponents := []processorprovider.ProviderComponent{mocks.NewProviderComponent(t)}
		mockComponent := mocks.NewComponentmapper(t)
		mockComponent.EXPECT().Generate(map[string]interface{}{}, mock.Anything).Return(mockProviderComponents, nil)
		mockProvider := mocks.NewProcessorProvider(t)
		mockProvider.EXPECT().AddHeader(mockProviderComponents[0]).Return(nil, fmt.Errorf("any"))

		doc := documentmapper.Document{
			Pages: []mappers.Componentmapper{}, Header: []mappers.Componentmapper{mockComponent},
			Footer: []mappers.Componentmapper{},
		}
		provider, err := doc.Generate(fixContent, mockProvider)

		assert.Nil(t, provider)
		assert.NotNil(t, err)
	})
	t.Run("when it is not possible add the footer to the document, it should return an error", func(t *testing.T) {
		fixContent := map[string]interface{}{}

		mockProviderComponents := []processorprovider.ProviderComponent{mocks.NewProviderComponent(t)}
		mockComponent := mocks.NewComponentmapper(t)
		mockComponent.EXPECT().Generate(map[string]interface{}{}, mock.Anything).Return(mockProviderComponents, nil)
		mockProvider := mocks.NewProcessorProvider(t)
		mockProvider.EXPECT().AddFooter(mockProviderComponents[0]).Return(nil, fmt.Errorf("any"))

		doc := documentmapper.Document{
			Pages: []mappers.Componentmapper{}, Header: []mappers.Componentmapper{},
			Footer: []mappers.Componentmapper{mockComponent},
		}
		provider, err := doc.Generate(fixContent, mockProvider)

		assert.Nil(t, provider)
		assert.NotNil(t, err)
	})
	t.Run("when document with page, header and footer is call, should generate document", func(t *testing.T) {
		fixContent := map[string]interface{}{
			"header": map[string]interface{}{"row_header": "test"},
			"footer": map[string]interface{}{"row_footer": "test"},
			"pages":  map[string]interface{}{"template_page_1": "test"},
		}

		mockProviderComponents := []processorprovider.ProviderComponent{mocks.NewProviderComponent(t)}
		mockComponent := mocks.NewComponentmapper(t)
		mockComponent.EXPECT().Generate(fixContent["header"], mock.Anything).Return(mockProviderComponents, nil)
		mockComponent.EXPECT().Generate(fixContent["footer"], mock.Anything).Return(mockProviderComponents, nil)
		mockComponent.EXPECT().Generate(fixContent["pages"], mock.Anything).Return(mockProviderComponents, nil)
		mockProvider := mocks.NewProcessorProvider(t)
		mockProvider.EXPECT().AddPages(mockProviderComponents[0]).Return(mockProvider, nil)
		mockProvider.EXPECT().AddHeader(mockProviderComponents[0]).Return(mockProvider, nil)
		mockProvider.EXPECT().AddFooter(mockProviderComponents[0]).Return(mockProvider, nil)

		doc := documentmapper.Document{
			Pages: []mappers.Componentmapper{mockComponent}, Header: []mappers.Componentmapper{mockComponent},
			Footer: []mappers.Componentmapper{mockComponent},
		}
		provider, err := doc.Generate(fixContent, mockProvider)

		assert.Nil(t, err)
		assert.NotNil(t, provider)
		mockComponent.AssertNumberOfCalls(t, "Generate", 3)
	})
}
