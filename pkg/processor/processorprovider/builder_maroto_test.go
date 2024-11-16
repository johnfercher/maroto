package processorprovider_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/internal/fixture"
	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/consts/protection"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/propsmapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/processorprovider"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func createBuilderMocks(t *testing.T) *mocks.Builder {
	build := mocks.NewBuilder(t)
	build.EXPECT().WithPageSize(mock.Anything).Return(config.NewBuilder())
	build.EXPECT().WithDimensions(mock.Anything, mock.Anything).Return(config.NewBuilder())
	build.EXPECT().WithTopMargin(mock.Anything).Return(config.NewBuilder())
	build.EXPECT().WithBottomMargin(mock.Anything).Return(config.NewBuilder())
	build.EXPECT().WithLeftMargin(mock.Anything).Return(config.NewBuilder())
	build.EXPECT().WithRightMargin(mock.Anything).Return(config.NewBuilder())
	build.EXPECT().WithConcurrentMode(mock.Anything).Return(config.NewBuilder())
	build.EXPECT().WithSequentialMode().Return(config.NewBuilder())
	build.EXPECT().WithDebug(mock.Anything).Return(config.NewBuilder())
	build.EXPECT().WithMaxGridSize(mock.Anything).Return(config.NewBuilder())
	build.EXPECT().WithDefaultFont(mock.Anything).Return(config.NewBuilder())
	build.EXPECT().WithPageNumber(mock.Anything).Return(config.NewBuilder())
	build.EXPECT().WithProtection(mock.Anything, mock.Anything, mock.Anything).Return(config.NewBuilder())
	build.EXPECT().WithCompression(mock.Anything).Return(config.NewBuilder())
	build.EXPECT().WithOrientation(mock.Anything).Return(config.NewBuilder())
	build.EXPECT().WithAuthor(mock.Anything, mock.Anything).Return(config.NewBuilder())
	build.EXPECT().WithCreationDate(mock.Anything).Return(config.NewBuilder())
	build.EXPECT().WithCreator(mock.Anything, mock.Anything).Return(config.NewBuilder())
	build.EXPECT().WithKeywords(mock.Anything, mock.Anything).Return(config.NewBuilder())
	build.EXPECT().WithSubject(mock.Anything, mock.Anything).Return(config.NewBuilder())
	build.EXPECT().WithTitle(mock.Anything, mock.Anything).Return(config.NewBuilder())
	build.EXPECT().WithDisableAutoPageBreak(mock.Anything).Return(config.NewBuilder())
	build.EXPECT().WithCustomFonts(mock.Anything).Return(config.NewBuilder())
	return build
}

func validateCallOfAllMethods(t *testing.T, builder *mocks.Builder) {
	builder.AssertNumberOfCalls(t, "WithPageSize", 1)
	builder.AssertNumberOfCalls(t, "WithDimensions", 1)
	builder.AssertNumberOfCalls(t, "WithTopMargin", 1)
	builder.AssertNumberOfCalls(t, "WithBottomMargin", 1)
	builder.AssertNumberOfCalls(t, "WithLeftMargin", 1)
	builder.AssertNumberOfCalls(t, "WithRightMargin", 1)
	builder.AssertNumberOfCalls(t, "WithConcurrentMode", 1)
	builder.AssertNumberOfCalls(t, "WithSequentialMode", 1)
	builder.AssertNumberOfCalls(t, "WithDebug", 1)
	builder.AssertNumberOfCalls(t, "WithMaxGridSize", 1)
	builder.AssertNumberOfCalls(t, "WithDefaultFont", 1)
	builder.AssertNumberOfCalls(t, "WithPageNumber", 1)
	builder.AssertNumberOfCalls(t, "WithProtection", 1)
	builder.AssertNumberOfCalls(t, "WithCompression", 1)
	builder.AssertNumberOfCalls(t, "WithOrientation", 1)
	builder.AssertNumberOfCalls(t, "WithAuthor", 1)
	builder.AssertNumberOfCalls(t, "WithCreationDate", 1)
	builder.AssertNumberOfCalls(t, "WithCreator", 1)
	builder.AssertNumberOfCalls(t, "WithKeywords", 1)
	builder.AssertNumberOfCalls(t, "WithSubject", 1)
	builder.AssertNumberOfCalls(t, "WithTitle", 1)
	builder.AssertNumberOfCalls(t, "WithDisableAutoPageBreak", 1)
	builder.AssertNumberOfCalls(t, "WithCustomFonts", 1)
}

func TestCreateMarotoBuilder(t *testing.T) {
	t.Run("when all props are sent, should add all props", func(t *testing.T) {
		repository := mocks.NewProcessorRepository(t)
		repository.EXPECT().GetDocument(mock.Anything).Return(string(extension.Png), []byte{123}, nil)
		build := createBuilderMocks(t)
		fixProps := fixture.BuilderProps()

		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg, err := builder.CreateMarotoBuilder(fixProps)

		assert.Nil(t, err)
		assert.NotNil(t, cfg)
		validateCallOfAllMethods(t, build)
	})
}

func TestWithPageSize(t *testing.T) {
	t.Run("when page size is null, should not set page size", func(t *testing.T) {
		repository := mocks.NewProcessorRepository(t)
		build := mocks.NewBuilder(t)

		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg := builder.WithPageSize("")

		assert.NotNil(t, cfg)
	})
	t.Run("when page size is sent, should set page size", func(t *testing.T) {
		repository := mocks.NewProcessorRepository(t)
		build := mocks.NewBuilder(t)
		build.EXPECT().WithPageSize(pagesize.A1).Return(config.NewBuilder())

		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg := builder.WithPageSize("a1")

		assert.NotNil(t, cfg)
		build.AssertNumberOfCalls(t, "WithPageSize", 1)
	})
}

func TestWithDimensions(t *testing.T) {
	t.Run("when dimensions is nil, should not set dimensions", func(t *testing.T) {
		repository := mocks.NewProcessorRepository(t)
		build := mocks.NewBuilder(t)

		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg := builder.WithDimensions(nil)

		assert.NotNil(t, cfg)
	})
	t.Run("when dimensions is sent, should set dimensions", func(t *testing.T) {
		fixDimensions := propsmapper.Dimensions{Width: 10.0, Height: 10.0}
		repository := mocks.NewProcessorRepository(t)
		build := mocks.NewBuilder(t)
		build.EXPECT().WithDimensions(fixDimensions.Width, fixDimensions.Height).Return(config.NewBuilder())

		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg := builder.WithDimensions(&fixDimensions)

		assert.NotNil(t, cfg)
		build.AssertNumberOfCalls(t, "WithDimensions", 1)
	})
}

func TestWithMargins(t *testing.T) {
	t.Run("when only left margin is sent, should set only left margin", func(t *testing.T) {
		fixMargins := propsmapper.Margins{Left: 10, Right: -1, Top: -1, Bottom: -1}
		repository := mocks.NewProcessorRepository(t)
		build := mocks.NewBuilder(t)
		build.EXPECT().WithLeftMargin(fixMargins.Left).Return(config.NewBuilder())

		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg := builder.WithMargin(&fixMargins)

		assert.NotNil(t, cfg)
		build.AssertNumberOfCalls(t, "WithLeftMargin", 1)
	})
	t.Run("when only right margin is sent, should set only right margin", func(t *testing.T) {
		fixMargins := propsmapper.Margins{Left: -1, Right: 10, Top: -1, Bottom: -1}
		repository := mocks.NewProcessorRepository(t)
		build := mocks.NewBuilder(t)
		build.EXPECT().WithRightMargin(fixMargins.Right).Return(config.NewBuilder())

		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg := builder.WithMargin(&fixMargins)

		assert.NotNil(t, cfg)
		build.AssertNumberOfCalls(t, "WithRightMargin", 1)
	})
	t.Run("when only bottom margin is sent, should set only bottom margin", func(t *testing.T) {
		fixMargins := propsmapper.Margins{Left: -1, Right: -1, Top: 10, Bottom: -1}
		repository := mocks.NewProcessorRepository(t)
		build := mocks.NewBuilder(t)
		build.EXPECT().WithTopMargin(fixMargins.Top).Return(config.NewBuilder())

		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg := builder.WithMargin(&fixMargins)

		assert.NotNil(t, cfg)
		build.AssertNumberOfCalls(t, "WithTopMargin", 1)
	})
	t.Run("when only top margin is sent, should set only top margin", func(t *testing.T) {
		fixMargins := propsmapper.Margins{Left: -1, Right: -1, Top: -1, Bottom: 10}
		repository := mocks.NewProcessorRepository(t)
		build := mocks.NewBuilder(t)
		build.EXPECT().WithBottomMargin(fixMargins.Bottom).Return(config.NewBuilder())

		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg := builder.WithMargin(&fixMargins)

		assert.NotNil(t, cfg)
		build.AssertNumberOfCalls(t, "WithBottomMargin", 1)
	})
	t.Run("When margins is not set, should no set margin", func(t *testing.T) {
		fixMargins := propsmapper.Margins{Left: 10, Right: 11, Top: 12, Bottom: 13}
		repository := mocks.NewProcessorRepository(t)
		build := mocks.NewBuilder(t)
		build.EXPECT().WithBottomMargin(fixMargins.Bottom).Return(config.NewBuilder())
		build.EXPECT().WithTopMargin(fixMargins.Top).Return(config.NewBuilder())
		build.EXPECT().WithLeftMargin(fixMargins.Left).Return(config.NewBuilder())
		build.EXPECT().WithRightMargin(fixMargins.Right).Return(config.NewBuilder())

		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg := builder.WithMargin(&fixMargins)

		assert.NotNil(t, cfg)
		build.AssertNumberOfCalls(t, "WithBottomMargin", 1)
		build.AssertNumberOfCalls(t, "WithTopMargin", 1)
		build.AssertNumberOfCalls(t, "WithLeftMargin", 1)
		build.AssertNumberOfCalls(t, "WithRightMargin", 1)
	})
}

func TestWithConcurrentMode(t *testing.T) {
	t.Run("when 0 works is sent, should  not set concurrent Mode", func(t *testing.T) {
		repository := mocks.NewProcessorRepository(t)
		build := mocks.NewBuilder(t)

		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg := builder.WithConcurrentMode(0)

		assert.NotNil(t, cfg)
	})
	t.Run("when 2 works are sent, should set concurrent mode", func(t *testing.T) {
		repository := mocks.NewProcessorRepository(t)
		build := mocks.NewBuilder(t)
		build.EXPECT().WithConcurrentMode(2).Return(config.NewBuilder())

		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg := builder.WithConcurrentMode(2)

		assert.NotNil(t, cfg)
		build.AssertNumberOfCalls(t, "WithConcurrentMode", 1)
	})
}

func TestWithSequentialMode(t *testing.T) {
	t.Run("when sequential mode is true, should set sequential mode", func(t *testing.T) {
		repository := mocks.NewProcessorRepository(t)
		build := mocks.NewBuilder(t)
		build.EXPECT().WithSequentialMode().Return(config.NewBuilder())

		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg := builder.WithSequentialMode(true)

		assert.NotNil(t, cfg)
		build.AssertNumberOfCalls(t, "WithSequentialMode", 1)
	})
	t.Run("when sequential mode is false, should not set sequential mode", func(t *testing.T) {
		repository := mocks.NewProcessorRepository(t)
		build := mocks.NewBuilder(t)
		build.EXPECT().WithSequentialMode().Return(config.NewBuilder())

		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg := builder.WithSequentialMode(false)

		assert.NotNil(t, cfg)
		build.AssertNumberOfCalls(t, "WithSequentialMode", 1)
	})
}

func TestWithSequentialLowMemoryMode(t *testing.T) {
	t.Run("when sequential low memory mode is 0, should not set sequential low memory", func(t *testing.T) {
		repository := mocks.NewProcessorRepository(t)
		build := mocks.NewBuilder(t)

		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg := builder.WithSequentialLowMemoryMode(0)

		assert.NotNil(t, cfg)
	})

	t.Run("when sequential low memory mode is 2, should set sequential low memory", func(t *testing.T) {
		repository := mocks.NewProcessorRepository(t)
		build := mocks.NewBuilder(t)
		build.EXPECT().WithSequentialLowMemoryMode(2).Return(config.NewBuilder())

		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg := builder.WithSequentialLowMemoryMode(2)

		assert.NotNil(t, cfg)
		build.AssertNumberOfCalls(t, "WithSequentialLowMemoryMode", 1)
	})
}

func TestWithDebug(t *testing.T) {
	t.Run("when debug is true, should set debug", func(t *testing.T) {
		repository := mocks.NewProcessorRepository(t)
		build := mocks.NewBuilder(t)
		build.EXPECT().WithDebug(true).Return(config.NewBuilder())

		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg := builder.WithDebug(true)

		assert.NotNil(t, cfg)
		build.AssertNumberOfCalls(t, "WithDebug", 1)
	})
	t.Run("when debug is false, should  not set debug", func(t *testing.T) {
		repository := mocks.NewProcessorRepository(t)
		build := mocks.NewBuilder(t)
		build.EXPECT().WithDebug(false).Return(config.NewBuilder())

		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg := builder.WithDebug(false)

		assert.NotNil(t, cfg)
		build.AssertNumberOfCalls(t, "WithDebug", 1)
	})
}

func TestWithMaxGridSize(t *testing.T) {
	t.Run("when max grid size is 0, should not set max grid size", func(t *testing.T) {
		repository := mocks.NewProcessorRepository(t)
		build := mocks.NewBuilder(t)

		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg := builder.WithMaxGridSize(0)

		assert.NotNil(t, cfg)
	})

	t.Run("when max grid size is 2, should set max grid size with 2", func(t *testing.T) {
		repository := mocks.NewProcessorRepository(t)
		build := mocks.NewBuilder(t)
		build.EXPECT().WithMaxGridSize(2).Return(config.NewBuilder())

		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg := builder.WithMaxGridSize(2)

		assert.NotNil(t, cfg)
		build.AssertNumberOfCalls(t, "WithMaxGridSize", 1)
	})
}

func TestWithDefaultFont(t *testing.T) {
	t.Run("when font is nil, should not set font", func(t *testing.T) {
		repository := mocks.NewProcessorRepository(t)
		build := mocks.NewBuilder(t)

		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg := builder.WithDefaultFont(nil)

		assert.NotNil(t, cfg)
	})
	t.Run("when invalid style is call, should set font", func(t *testing.T) {
		fixFont := fixture.FontProp()
		fixFont.Style = "invalid"
		repository := mocks.NewProcessorRepository(t)
		build := mocks.NewBuilder(t)
		build.EXPECT().WithDefaultFont(&fixFont).Return(config.NewBuilder())

		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg := builder.WithDefaultFont(&propsmapper.Font{
			Family: fixFont.Family,
			Style:  string(fixFont.Style), Size: fixFont.Size, Color: (*propsmapper.Color)(fixFont.Color),
		})

		assert.NotNil(t, cfg)
		build.AssertNumberOfCalls(t, "WithDefaultFont", 1)
	})
	t.Run("when font is sent, should set font", func(t *testing.T) {
		fixFont := fixture.FontProp()
		repository := mocks.NewProcessorRepository(t)
		build := mocks.NewBuilder(t)
		build.EXPECT().WithDefaultFont(&fixFont).Return(config.NewBuilder())

		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg := builder.WithDefaultFont(&propsmapper.Font{
			Family: fixFont.Family,
			Style:  string(fixFont.Style), Size: fixFont.Size, Color: (*propsmapper.Color)(fixFont.Color),
		})

		assert.NotNil(t, cfg)
		build.AssertNumberOfCalls(t, "WithDefaultFont", 1)
	})
}

func TestWithPageNumber(t *testing.T) {
	t.Run("when page number is nil, should not set page number", func(t *testing.T) {
		repository := mocks.NewProcessorRepository(t)
		build := mocks.NewBuilder(t)

		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg := builder.WithPageNumber(nil)

		assert.NotNil(t, cfg)
	})
	t.Run("when page number is sent, should set page number", func(t *testing.T) {
		fixPageNumber := fixture.PageProp()
		repository := mocks.NewProcessorRepository(t)
		build := mocks.NewBuilder(t)
		build.EXPECT().WithPageNumber(fixPageNumber).Return(config.NewBuilder())

		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg := builder.WithPageNumber(&propsmapper.PageNumber{
			Pattern: fixPageNumber.Pattern, Place: string(fixPageNumber.Place), Family: fixPageNumber.Family,
			Style: string(fixPageNumber.Style), Size: fixPageNumber.Size, Color: (*propsmapper.Color)(fixPageNumber.Color),
		})

		assert.NotNil(t, cfg)
		build.AssertNumberOfCalls(t, "WithPageNumber", 1)
	})
}

func TestWithProtection(t *testing.T) {
	t.Run("when protection is nil, should not set protection", func(t *testing.T) {
		repository := mocks.NewProcessorRepository(t)
		build := mocks.NewBuilder(t)

		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg := builder.WithProtection(nil)

		assert.NotNil(t, cfg)
	})

	t.Run("when protection is sent, should set protection", func(t *testing.T) {
		fixProtection := propsmapper.Protection{Type: 0, UserPassword: "user.password", OwnerPassword: "pass"}
		repository := mocks.NewProcessorRepository(t)
		build := mocks.NewBuilder(t)
		build.EXPECT().WithProtection(protection.Type(fixProtection.Type), fixProtection.UserPassword,
			fixProtection.OwnerPassword).Return(config.NewBuilder())

		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg := builder.WithProtection(&fixProtection)

		assert.NotNil(t, cfg)
		build.AssertNumberOfCalls(t, "WithProtection", 1)
	})
}

func TestWithCompression(t *testing.T) {
	t.Run("when compression is true, should set compression", func(t *testing.T) {
		repository := mocks.NewProcessorRepository(t)
		build := mocks.NewBuilder(t)
		build.EXPECT().WithCompression(true).Return(config.NewBuilder())
		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg := builder.WithCompression(true)

		assert.NotNil(t, cfg)
		build.AssertNumberOfCalls(t, "WithCompression", 1)
	})
}

func TestWithOrientation(t *testing.T) {
	t.Run("when orientation is null, should not set orientation", func(t *testing.T) {
		repository := mocks.NewProcessorRepository(t)
		build := mocks.NewBuilder(t)

		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg := builder.WithOrientation("")

		assert.NotNil(t, cfg)
	})
	t.Run("when orientation is sent, should set orientation", func(t *testing.T) {
		repository := mocks.NewProcessorRepository(t)
		build := mocks.NewBuilder(t)
		build.EXPECT().WithOrientation(orientation.Vertical).Return(config.NewBuilder())

		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg := builder.WithOrientation("vertical")

		assert.NotNil(t, cfg)
		build.AssertNumberOfCalls(t, "WithOrientation", 1)
	})
}

func TestWithMetadata(t *testing.T) {
	t.Run("when metadata is nil, should not set any metadata", func(t *testing.T) {
		repository := mocks.NewProcessorRepository(t)
		build := mocks.NewBuilder(t)

		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg := builder.WithMetadata(nil)

		assert.NotNil(t, cfg)
	})
	t.Run("when only author is sent, should set only author", func(t *testing.T) {
		fixMetadata := fixture.Metadata()
		fixMetadata.CreationDate = nil
		fixMetadata.Creator = nil
		fixMetadata.Subject = nil
		fixMetadata.KeywordsStr = nil
		fixMetadata.Title = nil

		repository := mocks.NewProcessorRepository(t)
		build := mocks.NewBuilder(t)
		build.EXPECT().WithAuthor(fixMetadata.Author.Text, fixMetadata.Author.UTF8).Return(config.NewBuilder())

		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg := builder.WithMetadata(fixMetadata)

		assert.NotNil(t, cfg)
		build.AssertNumberOfCalls(t, "WithAuthor", 1)
	})
	t.Run("when creation date is nil, should not set creation date", func(t *testing.T) {
		fixMetadata := fixture.Metadata()
		fixMetadata.Author = nil
		fixMetadata.Creator = nil
		fixMetadata.Subject = nil
		fixMetadata.KeywordsStr = nil
		fixMetadata.Title = nil

		repository := mocks.NewProcessorRepository(t)
		build := mocks.NewBuilder(t)
		build.EXPECT().WithCreationDate(*fixMetadata.CreationDate).Return(config.NewBuilder())

		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg := builder.WithMetadata(fixMetadata)

		assert.NotNil(t, cfg)
		build.AssertNumberOfCalls(t, "WithCreationDate", 1)
	})
	t.Run("when creator is nil, should not set creator", func(t *testing.T) {
		fixMetadata := fixture.Metadata()
		fixMetadata.Author = nil
		fixMetadata.CreationDate = nil
		fixMetadata.Subject = nil
		fixMetadata.KeywordsStr = nil
		fixMetadata.Title = nil

		repository := mocks.NewProcessorRepository(t)
		build := mocks.NewBuilder(t)
		build.EXPECT().WithCreator(fixMetadata.Creator.Text, fixMetadata.Creator.UTF8).Return(config.NewBuilder())

		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg := builder.WithMetadata(fixMetadata)

		assert.NotNil(t, cfg)
		build.AssertNumberOfCalls(t, "WithCreator", 1)
	})
	t.Run("when keywords is nil, should not set keywords", func(t *testing.T) {
		fixMetadata := fixture.Metadata()
		fixMetadata.Author = nil
		fixMetadata.CreationDate = nil
		fixMetadata.Subject = nil
		fixMetadata.Creator = nil
		fixMetadata.Title = nil

		repository := mocks.NewProcessorRepository(t)
		build := mocks.NewBuilder(t)
		build.EXPECT().WithKeywords(fixMetadata.KeywordsStr.Text, fixMetadata.KeywordsStr.UTF8).Return(config.NewBuilder())

		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg := builder.WithMetadata(fixMetadata)

		assert.NotNil(t, cfg)
		build.AssertNumberOfCalls(t, "WithKeywords", 1)
	})
	t.Run("when subject is nil, should not set subject", func(t *testing.T) {
		fixMetadata := fixture.Metadata()
		fixMetadata.Author = nil
		fixMetadata.CreationDate = nil
		fixMetadata.KeywordsStr = nil
		fixMetadata.Creator = nil
		fixMetadata.Title = nil

		repository := mocks.NewProcessorRepository(t)
		build := mocks.NewBuilder(t)
		build.EXPECT().WithSubject(fixMetadata.Subject.Text, fixMetadata.Subject.UTF8).Return(config.NewBuilder())

		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg := builder.WithMetadata(fixMetadata)

		assert.NotNil(t, cfg)
		build.AssertNumberOfCalls(t, "WithSubject", 1)
	})
	t.Run("when title is nil, should not set title", func(t *testing.T) {
		fixMetadata := fixture.Metadata()
		fixMetadata.Author = nil
		fixMetadata.CreationDate = nil
		fixMetadata.KeywordsStr = nil
		fixMetadata.Creator = nil
		fixMetadata.Subject = nil

		repository := mocks.NewProcessorRepository(t)
		build := mocks.NewBuilder(t)
		build.EXPECT().WithTitle(fixMetadata.Title.Text, fixMetadata.Title.UTF8).Return(config.NewBuilder())

		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg := builder.WithMetadata(fixMetadata)

		assert.NotNil(t, cfg)
		build.AssertNumberOfCalls(t, "WithTitle", 1)
	})
	t.Run("when all metadatas is sent, should set all metadatas", func(t *testing.T) {
		fixMetadata := fixture.Metadata()
		repository := mocks.NewProcessorRepository(t)
		build := mocks.NewBuilder(t)
		build.EXPECT().WithAuthor(fixMetadata.Author.Text, fixMetadata.Author.UTF8).Return(config.NewBuilder())
		build.EXPECT().WithCreationDate(*fixMetadata.CreationDate).Return(config.NewBuilder())
		build.EXPECT().WithCreator(fixMetadata.Creator.Text, fixMetadata.Creator.UTF8).Return(config.NewBuilder())
		build.EXPECT().WithKeywords(fixMetadata.KeywordsStr.Text, fixMetadata.KeywordsStr.UTF8).Return(config.NewBuilder())
		build.EXPECT().WithSubject(fixMetadata.Subject.Text, fixMetadata.Subject.UTF8).Return(config.NewBuilder())
		build.EXPECT().WithTitle(fixMetadata.Title.Text, fixMetadata.Title.UTF8).Return(config.NewBuilder())

		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg := builder.WithMetadata(fixMetadata)

		assert.NotNil(t, cfg)
		build.AssertNumberOfCalls(t, "WithAuthor", 1)
		build.AssertNumberOfCalls(t, "WithCreationDate", 1)
		build.AssertNumberOfCalls(t, "WithCreator", 1)
		build.AssertNumberOfCalls(t, "WithKeywords", 1)
		build.AssertNumberOfCalls(t, "WithSubject", 1)
		build.AssertNumberOfCalls(t, "WithTitle", 1)
	})
}

func TestWithCustomFonts(t *testing.T) {
	t.Run("when fonts is not sent, should not set fonts", func(t *testing.T) {
		repository := mocks.NewProcessorRepository(t)
		build := mocks.NewBuilder(t)

		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg, err := builder.WithCustomFonts(nil)

		assert.Nil(t, err)
		assert.NotNil(t, cfg)
	})
	t.Run("when is not possible load font file, should return an error", func(t *testing.T) {
		fixCustomFont := propsmapper.CustomFont{Family: "family", Style: "bold", File: "file"}
		repository := mocks.NewProcessorRepository(t)
		repository.EXPECT().GetDocument(fixCustomFont.File).Return("", nil, fmt.Errorf("any"))
		build := mocks.NewBuilder(t)

		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg, err := builder.WithCustomFonts([]*propsmapper.CustomFont{&fixCustomFont})

		assert.NotNil(t, err)
		assert.Nil(t, cfg)
	})
	t.Run("when 2 fonts is sent, should set 2 fonts", func(t *testing.T) {
		fixCustomFont := propsmapper.CustomFont{Family: "family", Style: "bold", File: "file"}
		repository := mocks.NewProcessorRepository(t)
		repository.EXPECT().GetDocument(fixCustomFont.File).Return(".ttt", []byte{}, nil)
		build := mocks.NewBuilder(t)
		build.EXPECT().WithCustomFonts(mock.Anything).Return(config.NewBuilder())

		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg, err := builder.WithCustomFonts([]*propsmapper.CustomFont{&fixCustomFont, &fixCustomFont})

		assert.Nil(t, err)
		assert.NotNil(t, cfg)
		repository.AssertNumberOfCalls(t, "GetDocument", 2)
		build.AssertNumberOfCalls(t, "WithCustomFonts", 1)
	})
}

func TestWithBackgroundImage(t *testing.T) {
	t.Run("when background image is not sent, should not set background image", func(t *testing.T) {
		repository := mocks.NewProcessorRepository(t)
		build := mocks.NewBuilder(t)

		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg, err := builder.WithCustomFonts(nil)

		assert.Nil(t, err)
		assert.NotNil(t, cfg)
	})
	t.Run("when is not possible load image, should return an error", func(t *testing.T) {
		repository := mocks.NewProcessorRepository(t)
		repository.EXPECT().GetDocument("img").Return("", nil, fmt.Errorf("any"))
		build := mocks.NewBuilder(t)

		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg, err := builder.WithBackgroundImage("img")

		assert.NotNil(t, err)
		assert.Nil(t, cfg)
	})
	t.Run("when backgroun image is sent, should set background", func(t *testing.T) {
		repository := mocks.NewProcessorRepository(t)
		repository.EXPECT().GetDocument("img").Return(string(extension.Png), []byte{123}, nil)
		build := mocks.NewBuilder(t)
		build.EXPECT().WithBackgroundImage([]byte{123}, extension.Png).Return(config.NewBuilder())

		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg, err := builder.WithBackgroundImage("img")

		assert.Nil(t, err)
		assert.NotNil(t, cfg)
		repository.AssertNumberOfCalls(t, "GetDocument", 1)
		build.AssertNumberOfCalls(t, "WithBackgroundImage", 1)
	})
}

func TestWithDisableAutoPageBreak(t *testing.T) {
	t.Run("when disable auto page break is true, should disable auto page break", func(t *testing.T) {
		repository := mocks.NewProcessorRepository(t)
		build := mocks.NewBuilder(t)
		build.EXPECT().WithDisableAutoPageBreak(true).Return(config.NewBuilder())

		builder := processorprovider.NewMarotoBuilder(repository, build)
		cfg := builder.WithDisableAutoPageBreak(true)

		assert.NotNil(t, cfg)
		build.AssertNumberOfCalls(t, "WithDisableAutoPageBreak", 1)
	})
}
