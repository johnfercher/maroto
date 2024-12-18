package processor_test

import (
	"os"
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/processor"
	"github.com/johnfercher/maroto/v2/pkg/processor/deserializer"
	"github.com/johnfercher/maroto/v2/pkg/processor/loader"
	"github.com/johnfercher/maroto/v2/pkg/processor/repository"
	processortest "github.com/johnfercher/maroto/v2/pkg/processor/test"
	"github.com/johnfercher/maroto/v2/pkg/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestDir(t *testing.T) {
	originalDir, err := os.Getwd()
	require.NoError(t, err)

	err = os.Chdir("../../")
	require.NoError(t, err)

	t.Cleanup(func() {
		err := os.Chdir(originalDir)
		require.NoError(t, err)
	})
}

func TestRegisterTemplate(t *testing.T) {
	t.Run("when template is recorded, should no return error", func(t *testing.T) {
	})
	t.Run("when is not possible deserialize template, should return an error", func(t *testing.T) {
	})
	t.Run("when is not possible register template, should return an error", func(t *testing.T) {
	})
}

func TestGenerateTemplate(t *testing.T) {
	t.Run("when valid template is found, should return valid pdf", func(t *testing.T) {
		fixtemplate, _ := processortest.NewFileReader().LoadFile("processor/json/simple_pdf_templates.json")
		fixContent, _ := processortest.NewFileReader().LoadFile("processor/json/simple_pdf_content.json")

		processor := processor.NewProcessor(repository.NewMemoryStorage(loader.NewLoader()), deserializer.NewJSONDeserializer())
		processor.RegisterTemplate("simple_pdf", string(fixtemplate))
		provider, err := processor.GenerateDocument("simple_pdf", string(fixContent))

		assert.Nil(t, err)
		test.New(t).Assert((*provider).GetStructure()).Equals("processor/simple_pdf.json")
	})

	t.Run("when template with Addpage is found, should set template", func(t *testing.T) {
		fixtemplate, _ := processortest.NewFileReader().LoadFile("processor/json/add_page_template.json")
		fixContent, _ := processortest.NewFileReader().LoadFile("processor/json/add_page_content.json")

		processor := processor.NewProcessor(repository.NewMemoryStorage(loader.NewLoader()), deserializer.NewJSONDeserializer())
		processor.RegisterTemplate("add_page", string(fixtemplate))
		provider, err := processor.GenerateDocument("add_page", string(fixContent))

		assert.Nil(t, err)
		test.New(t).Assert((*provider).GetStructure()).Equals("examples/addpage.json")
	})

	t.Run("when template with AutoRow is found, should set template", func(t *testing.T) {
		setupTestDir(t)
		fixtemplate, _ := processortest.NewFileReader().LoadFile("processor/json/auto_row_template.json")
		fixContent, _ := processortest.NewFileReader().LoadFile("processor/json/auto_row_content.json")

		processor := processor.NewProcessor(repository.NewMemoryStorage(loader.NewLoader()), deserializer.NewJSONDeserializer())
		processor.RegisterTemplate("auto_row", string(fixtemplate))
		provider, err := processor.GenerateDocument("auto_row", string(fixContent))

		assert.Nil(t, err)
		test.New(t).Assert((*provider).GetStructure()).Equals("examples/autorow.json")
	})
	t.Run("when sent template is not found, should reuturn an error", func(t *testing.T) {

	})
	t.Run("when template with invalid field is found, should return an error", func(t *testing.T) {
	})
	t.Run("when invalid content is sent, should return an error", func(t *testing.T) {
	})
}
