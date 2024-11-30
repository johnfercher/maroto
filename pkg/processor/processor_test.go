package processor_test

import (
	"encoding/json"
	"testing"

	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/processor"
	"github.com/johnfercher/maroto/v2/pkg/processor/deserializer"
	processortest "github.com/johnfercher/maroto/v2/pkg/processor/test"
	"github.com/johnfercher/maroto/v2/pkg/test"
	"github.com/stretchr/testify/assert"
)

func TestRegisterTemplate(t *testing.T) {
	t.Run("when template is recorded, should no return error", func(t *testing.T) {
	})
	t.Run("when is not possible deserialize template, should return an error", func(t *testing.T) {
	})
	t.Run("when is not possible register template, should return an error", func(t *testing.T) {
	})
}

func loadTemplate(templatePath string) map[string]interface{} {
	var template map[string]interface{}
	file, _ := processortest.NewFileReader().LoadFile(templatePath)
	if err := json.Unmarshal(file, &template); err != nil {
		return nil
	}
	return template
}

func TestGenerateTemplate(t *testing.T) {
	t.Run("when valid template is found, should return valid pdf", func(t *testing.T) {
		fixtemplate := loadTemplate("processor/json/simple_pdf_templates.json")
		fixContent, _ := processortest.NewFileReader().LoadFile("processor/json/simple_pdf_content.json")

		repository := mocks.NewProcessorRepository(t)
		repository.EXPECT().ReadTemplate("simple_pdf").Return(fixtemplate, nil)

		provider, err := processor.NewProcessor(repository, deserializer.NewJSONDeserializer()).GenerateDocument("simple_pdf", string(fixContent))
		assert.Nil(t, err)
		test.New(t).Assert((*provider).GetStructure()).Equals("processor/simple_pdf.json")
	})

	t.Run("when template with Addpage is found, should set template", func(t *testing.T) {
		fixtemplate := loadTemplate("processor/json/add_page_template.json")
		fixContent, _ := processortest.NewFileReader().LoadFile("processor/json/add_page_content.json")

		repository := mocks.NewProcessorRepository(t)
		repository.EXPECT().ReadTemplate("add_page").Return(fixtemplate, nil)

		provider, err := processor.NewProcessor(repository, deserializer.NewJSONDeserializer()).GenerateDocument("add_page", string(fixContent))
		assert.Nil(t, err)
		test.New(t).Assert((*provider).GetStructure()).Equals("examples/addpage.json")

		doc, _ := (*provider).Generate()
		err = doc.Save("test.pdf")
		assert.Nil(t, err)
	})

	t.Run("when sent template is not found, should reuturn an error", func(t *testing.T) {
	})
	t.Run("when template with invalid field is found, should return an error", func(t *testing.T) {
	})
	t.Run("when invalid content is sent, should return an error", func(t *testing.T) {
	})
}
