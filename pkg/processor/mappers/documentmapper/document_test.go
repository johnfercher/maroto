package documentmapper

import (
	"testing"

	"github.com/johnfercher/maroto/v2/internal/fixture"
	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/processor/deserializer"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/listmapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/pagemapper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewPdf(t *testing.T) {
	t.Run("when builder is sent, should set builder", func(t *testing.T) {
		builderDocument := `
			{
				"builder": {"chunk_workers": 10}
			}
		`
		factory := mocks.NewAbstractFactoryMaps(t)

		template, err := deserializer.NewJsonDeserializer().Deserialize(builderDocument)
		assert.Nil(t, err)

		doc, err := NewPdf(template, factory)
		assert.Nil(t, err)
		assert.Equal(t, doc.Builder.ChunkWorkers, 10)
	})

	t.Run("when an invalid builder is passed, should return an error", func(t *testing.T) {
		builderDocument := `{"builder": 10}`
		factory := mocks.NewAbstractFactoryMaps(t)

		template, err := deserializer.NewJsonDeserializer().Deserialize(builderDocument)
		assert.Nil(t, err)

		doc, err := NewPdf(template, factory)

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

		template, err := deserializer.NewJsonDeserializer().Deserialize(builderDocument)
		assert.Nil(t, err)

		doc, err := NewPdf(template, factory)

		assert.Nil(t, err)
		assert.Equal(t, 2, len(doc.Header))
	})

	t.Run("when an invalid header is passed, should return an error", func(t *testing.T) {
		builderDocument := `{"header": 1}`
		factory := mocks.NewAbstractFactoryMaps(t)

		template, err := deserializer.NewJsonDeserializer().Deserialize(builderDocument)
		assert.Nil(t, err)

		_, err = NewPdf(template, factory)
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

		template, err := deserializer.NewJsonDeserializer().Deserialize(builderDocument)
		assert.Nil(t, err)

		doc, err := NewPdf(template, factory)

		assert.Nil(t, err)
		assert.Equal(t, 2, len(doc.Footer))
	})

	t.Run("when an invalid footer is passed, should return an error", func(t *testing.T) {
		builderDocument := `{"footer": 1}`
		factory := mocks.NewAbstractFactoryMaps(t)

		template, err := deserializer.NewJsonDeserializer().Deserialize(builderDocument)
		assert.Nil(t, err)

		_, err = NewPdf(template, factory)

		assert.NotNil(t, err)
	})

	t.Run("when 2 pages are sent, it should add 2 pages to the document", func(t *testing.T) {
		builderDocument := `
			{"pages": {
				"page_template_1":{},
				"page_template_2":{} 
			}}
		`
		validPage := fixture.MapperPage()
		factory := mocks.NewAbstractFactoryMaps(t)
		factory.On("NewPage", mock.Anything, "page_template_1").Return(validPage, nil)
		factory.On("NewPage", mock.Anything, "page_template_2").Return(validPage, nil)

		template, err := deserializer.NewJsonDeserializer().Deserialize(builderDocument)
		assert.Nil(t, err)

		doc, err := NewPdf(template, factory)

		assert.Nil(t, err)
		assert.Equal(t, len(doc.pages), 2)
		assert.IsType(t, &pagemapper.Page{}, doc.pages[0])
	})

	t.Run("when 1 list is sent, it should add 1 list to the document", func(t *testing.T) {
		builderDocument := `
			{"pages": {
				"list_template_1":{}
			}}
		`
		validPage := fixture.MapperList()
		factory := mocks.NewAbstractFactoryMaps(t)
		factory.On("NewList", mock.Anything, "list_template_1", mock.Anything).Return(validPage, nil)

		template, err := deserializer.NewJsonDeserializer().Deserialize(builderDocument)
		assert.Nil(t, err)

		doc, err := NewPdf(template, factory)

		assert.Nil(t, err)
		assert.Equal(t, len(doc.pages), 1)
		assert.IsType(t, &listmapper.List{}, doc.pages[0])
	})

	t.Run("when an invalid page is sent, it should return an error", func(t *testing.T) {
		builderDocument := `{"pages": 1}`
		factory := mocks.NewAbstractFactoryMaps(t)

		template, err := deserializer.NewJsonDeserializer().Deserialize(builderDocument)
		assert.Nil(t, err)

		_, err = NewPdf(template, factory)
		assert.NotNil(t, err)
	})
}
