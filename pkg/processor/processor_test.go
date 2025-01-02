package processor_test

import (
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
		err := processor.RegisterTemplate("simple_pdf", string(fixtemplate))
		require.NoError(t, err)

		provider, err := processor.GenerateDocument("simple_pdf", string(fixContent))

		assert.Nil(t, err)
		test.New(t).Assert((*provider).GetStructure()).Equals("processor/simple_pdf.json")
	})

	t.Run("when template with Addpage example is found, should set template", func(t *testing.T) {
		fixtemplate, _ := processortest.NewFileReader().LoadFile("processor/json/add_page_template.json")
		fixContent, _ := processortest.NewFileReader().LoadFile("processor/json/add_page_content.json")
		processor := processor.NewProcessor(repository.NewMemoryStorage(loader.NewLoader()), deserializer.NewJSONDeserializer())
		err := processor.RegisterTemplate("add_page", string(fixtemplate))
		require.NoError(t, err)

		provider, err := processor.GenerateDocument("add_page", string(fixContent))

		assert.Nil(t, err)
		test.New(t).Assert((*provider).GetStructure()).Equals("examples/addpage.json")
	})

	t.Run("when template with AutoRow example is found, should set template", func(t *testing.T) {
		test.SetupTestDir(t)

		fixtemplate, _ := processortest.NewFileReader().LoadFile("processor/json/auto_row_template.json")
		fixContent, _ := processortest.NewFileReader().LoadFile("processor/json/auto_row_content.json")
		processor := processor.NewProcessor(repository.NewMemoryStorage(loader.NewLoader()), deserializer.NewJSONDeserializer())
		err := processor.RegisterTemplate("auto_row", string(fixtemplate))
		require.NoError(t, err)

		provider, err := processor.GenerateDocument("auto_row", string(fixContent))

		assert.Nil(t, err)
		test.New(t).Assert((*provider).GetStructure()).Equals("examples/autorow.json")
	})
	t.Run("when template with background example is found, should set template", func(t *testing.T) {
		test.SetupTestDir(t)

		fixtemplate, _ := processortest.NewFileReader().LoadFile("processor/json/background_template.json")
		fixContent, _ := processortest.NewFileReader().LoadFile("processor/json/background_content.json")
		processor := processor.NewProcessor(repository.NewMemoryStorage(loader.NewLoader()), deserializer.NewJSONDeserializer())
		err := processor.RegisterTemplate("background", string(fixtemplate))
		require.NoError(t, err)

		provider, err := processor.GenerateDocument("background", string(fixContent))

		assert.Nil(t, err)
		test.New(t).Assert((*provider).GetStructure()).Equals("examples/background.json")
	})
	t.Run("when template with barcodegrid example is found, should set template", func(t *testing.T) {
		test.SetupTestDir(t)

		fixtemplate, _ := processortest.NewFileReader().LoadFile("processor/json/barcodegrid_template.json")
		fixContent, _ := processortest.NewFileReader().LoadFile("processor/json/barcodegrid_content.json")
		processor := processor.NewProcessor(repository.NewMemoryStorage(loader.NewLoader()), deserializer.NewJSONDeserializer())
		err := processor.RegisterTemplate("barcodegrid", string(fixtemplate))
		require.NoError(t, err)

		provider, err := processor.GenerateDocument("barcodegrid", string(fixContent))

		assert.Nil(t, err)
		test.New(t).Assert((*provider).GetStructure()).Equals("examples/barcodegrid.json")
	})
	t.Run("when template with billing example is found, should set template", func(t *testing.T) {
		test.SetupTestDir(t)

		fixtemplate, _ := processortest.NewFileReader().LoadFile("processor/json/billing_template.json")
		fixContent, _ := processortest.NewFileReader().LoadFile("processor/json/billing_content.json")
		processor := processor.NewProcessor(repository.NewMemoryStorage(loader.NewLoader()), deserializer.NewJSONDeserializer())
		err := processor.RegisterTemplate("billing", string(fixtemplate))
		require.NoError(t, err)

		provider, err := processor.GenerateDocument("billing", string(fixContent))

		assert.Nil(t, err)
		test.New(t).Assert((*provider).GetStructure()).Equals("examples/billing.json")
	})
	t.Run("when template with cellstyle example is found, should set template", func(t *testing.T) {
		test.SetupTestDir(t)

		fixtemplate, _ := processortest.NewFileReader().LoadFile("processor/json/cell_template.json")
		processor := processor.NewProcessor(repository.NewMemoryStorage(loader.NewLoader()), deserializer.NewJSONDeserializer())
		err := processor.RegisterTemplate("cellstyle", string(fixtemplate))
		require.NoError(t, err)

		provider, err := processor.GenerateDocument("cellstyle", "{}")

		assert.Nil(t, err)
		test.New(t).Assert((*provider).GetStructure()).Equals("examples/cellstyle.json")
	})
	t.Run("when template with compression example is found, should set template", func(t *testing.T) {
		test.SetupTestDir(t)

		fixtemplate, _ := processortest.NewFileReader().LoadFile("processor/json/compressionv2_template.json")
		processor := processor.NewProcessor(repository.NewMemoryStorage(loader.NewLoader()), deserializer.NewJSONDeserializer())
		err := processor.RegisterTemplate("compression", string(fixtemplate))
		require.NoError(t, err)

		provider, err := processor.GenerateDocument("compression", "{}")

		assert.Nil(t, err)
		test.New(t).Assert((*provider).GetStructure()).Equals("examples/compression.json")
	})
	t.Run("when template with customdimensions example is found, should set template", func(t *testing.T) {
		test.SetupTestDir(t)

		fixtemplate, _ := processortest.NewFileReader().LoadFile("processor/json/customdimensions_template.json")
		processor := processor.NewProcessor(repository.NewMemoryStorage(loader.NewLoader()), deserializer.NewJSONDeserializer())
		err := processor.RegisterTemplate("dimensions", string(fixtemplate))
		require.NoError(t, err)

		provider, err := processor.GenerateDocument("dimensions", "{}")

		assert.Nil(t, err)
		test.New(t).Assert((*provider).GetStructure()).Equals("examples/customdimensions.json")
	})
	t.Run("when template with customfont example is found, should set template", func(t *testing.T) {
		test.SetupTestDir(t)

		fixtemplate, _ := processortest.NewFileReader().LoadFile("processor/json/customfont_template.json")
		fixcontent, _ := processortest.NewFileReader().LoadFile("processor/json/customfont_content.json")
		processor := processor.NewProcessor(repository.NewMemoryStorage(loader.NewLoader()), deserializer.NewJSONDeserializer())
		err := processor.RegisterTemplate("font", string(fixtemplate))
		require.NoError(t, err)

		provider, err := processor.GenerateDocument("font", string(fixcontent))

		assert.Nil(t, err)
		test.New(t).Assert((*provider).GetStructure()).Equals("examples/customfont.json")
	})
	t.Run("when template with custompage example is found, should set template", func(t *testing.T) {
		test.SetupTestDir(t)

		fixtemplate, _ := processortest.NewFileReader().LoadFile("processor/json/custompage_template.json")
		processor := processor.NewProcessor(repository.NewMemoryStorage(loader.NewLoader()), deserializer.NewJSONDeserializer())
		err := processor.RegisterTemplate("page", string(fixtemplate))
		require.NoError(t, err)

		provider, err := processor.GenerateDocument("page", "{}")

		assert.Nil(t, err)
		test.New(t).Assert((*provider).GetStructure()).Equals("examples/custompage.json")
	})
	t.Run("when template with datamatrixgrid example is found, should set template", func(t *testing.T) {
		test.SetupTestDir(t)

		fixtemplate, _ := processortest.NewFileReader().LoadFile("processor/json/datamatrixgrid_template.json")
		processor := processor.NewProcessor(repository.NewMemoryStorage(loader.NewLoader()), deserializer.NewJSONDeserializer())
		err := processor.RegisterTemplate("datamatrixgrid", string(fixtemplate))
		require.NoError(t, err)

		provider, err := processor.GenerateDocument("datamatrixgrid", "{}")

		assert.Nil(t, err)
		test.New(t).Assert((*provider).GetStructure()).Equals("examples/datamatrixgrid.json")
	})
	t.Run("when template with autopagebreak example is found, should set template", func(t *testing.T) {
		test.SetupTestDir(t)

		fixtemplate, _ := processortest.NewFileReader().LoadFile("processor/json/disablepagebreak_template.json")
		processor := processor.NewProcessor(repository.NewMemoryStorage(loader.NewLoader()), deserializer.NewJSONDeserializer())
		err := processor.RegisterTemplate("disablepagebreak", string(fixtemplate))
		require.NoError(t, err)

		provider, err := processor.GenerateDocument("disablepagebreak", "{}")

		assert.Nil(t, err)
		test.New(t).Assert((*provider).GetStructure()).Equals("examples/disablepagebreak.json")
	})
	t.Run("when template with footer example is found, should set template", func(t *testing.T) {
		test.SetupTestDir(t)

		fixtemplate, _ := processortest.NewFileReader().LoadFile("processor/json/footer_template.json")
		fixcontent, _ := processortest.NewFileReader().LoadFile("processor/json/footer_content.json")
		processor := processor.NewProcessor(repository.NewMemoryStorage(loader.NewLoader()), deserializer.NewJSONDeserializer())
		err := processor.RegisterTemplate("footer", string(fixtemplate))
		require.NoError(t, err)

		provider, err := processor.GenerateDocument("footer", string(fixcontent))

		assert.Nil(t, err)
		test.New(t).Assert((*provider).GetStructure()).Equals("examples/footer.json")
	})
	t.Run("when template with header example is found, should set template", func(t *testing.T) {
		test.SetupTestDir(t)

		fixtemplate, _ := processortest.NewFileReader().LoadFile("processor/json/header_template.json")
		fixcontent, _ := processortest.NewFileReader().LoadFile("processor/json/header_content.json")
		processor := processor.NewProcessor(repository.NewMemoryStorage(loader.NewLoader()), deserializer.NewJSONDeserializer())
		err := processor.RegisterTemplate("header", string(fixtemplate))
		require.NoError(t, err)

		provider, err := processor.GenerateDocument("header", string(fixcontent))

		assert.Nil(t, err)
		test.New(t).Assert((*provider).GetStructure()).Equals("examples/header.json")
	})
	t.Run("when template with imagegrid example is found, should set template", func(t *testing.T) {
		test.SetupTestDir(t)

		fixtemplate, _ := processortest.NewFileReader().LoadFile("processor/json/imagegrid_template.json")
		processor := processor.NewProcessor(repository.NewMemoryStorage(loader.NewLoader()), deserializer.NewJSONDeserializer())
		err := processor.RegisterTemplate("imagegrid", string(fixtemplate))
		require.NoError(t, err)

		provider, err := processor.GenerateDocument("imagegrid", "{}")

		assert.Nil(t, err)
		test.New(t).Assert((*provider).GetStructure()).Equals("examples/imagegrid.json")
	})
	t.Run("when template with linegrid example is found, should set template", func(t *testing.T) {
		test.SetupTestDir(t)

		fixtemplate, _ := processortest.NewFileReader().LoadFile("processor/json/linegrid_template.json")
		processor := processor.NewProcessor(repository.NewMemoryStorage(loader.NewLoader()), deserializer.NewJSONDeserializer())
		err := processor.RegisterTemplate("line", string(fixtemplate))
		require.NoError(t, err)

		provider, err := processor.GenerateDocument("line", "{}")

		assert.Nil(t, err)
		test.New(t).Assert((*provider).GetStructure()).Equals("examples/line.json")
	})
	t.Run("when template with invalid field is found, should return an error", func(t *testing.T) {

	})
	t.Run("when invalid content is sent, should return an error", func(t *testing.T) {
	})
}
