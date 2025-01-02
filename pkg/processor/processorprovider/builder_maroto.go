package processorprovider

import (
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/consts/protection"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/processor/core"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/buildermapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/propsmapper"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/johnfercher/maroto/v2/pkg/repository"
)

// MarotoBuilder is responsible for creating Maroto builder props from buildermapper
type MarotoBuilder struct {
	cfg        config.Builder
	repository core.ProcessorRepository
}

// NewMarotoBuilder is responsible for creating an object MarotoBuilder
//   - It will use repository for search files like image and font
func NewMarotoBuilder(repository core.ProcessorRepository, cfg config.Builder) *MarotoBuilder {
	return &MarotoBuilder{
		repository: repository,
		cfg:        cfg,
	}
}

// CreateMarotoBuilder is responsible for facilitating the creation of the builder
func (m *MarotoBuilder) CreateMarotoBuilder(builder *buildermapper.Builder) (config.Builder, error) {
	m = m.WithPageSize(builder.PageSize).WithDimensions(builder.Dimensions).WithMargin(builder.Margins).
		WithConcurrentMode(builder.ConcurrentMode).WithSequentialMode(builder.SequentialMode).
		WithSequentialLowMemoryMode(builder.SequentialLowMemoryMode).WithDebug(builder.Debug).WithMaxGridSize(builder.MaxGridSize).
		WithDefaultFont(builder.DefaultFont).WithPageNumber(builder.PageNumber).WithProtection(builder.Protection).
		WithCompression(builder.Compression).WithOrientation(builder.Orientation).WithMetadata(builder.Metadata).
		WithDisableAutoPageBreak(builder.DisableAutoPageBreak)

	if _, err := m.WithBackgroundImage(builder.BackgroundImage); err != nil {
		return nil, err
	}
	if _, err := m.WithCustomFonts(builder.CustomFonts); err != nil {
		return nil, err
	}

	return m.cfg, nil
}

// WithPageSize will add page size properties
//   - if size is null, size will not be added
func (m *MarotoBuilder) WithPageSize(size string) *MarotoBuilder {
	validAndRun(func() { m.cfg.WithPageSize(pagesize.Type(size)) }, size != "")
	return m
}

// WithDimensions defines custom page dimensions, this overrides page size.
//   - if dimensions is null, dimensions will not be added
func (m *MarotoBuilder) WithDimensions(dimensions *propsmapper.Dimensions) *MarotoBuilder {
	validAndRun(func() { m.cfg.WithDimensions(dimensions.Width, dimensions.Height) }, dimensions != nil)
	return m
}

// WithMargin customizes each margin individually - Left, Right, top and booton
//   - if margin is null or individual margin is less than 0, margin will not be added
func (m *MarotoBuilder) WithMargin(margin *propsmapper.Margins) *MarotoBuilder {
	if margin != nil {
		validAndRun(func() { m.cfg.WithBottomMargin(margin.Bottom) }, margin.Bottom >= 0)
		validAndRun(func() { m.cfg.WithTopMargin(margin.Top) }, margin.Top >= 0)
		validAndRun(func() { m.cfg.WithLeftMargin(margin.Left) }, margin.Left >= 0)
		validAndRun(func() { m.cfg.WithRightMargin(margin.Right) }, margin.Right >= 0)
	}
	return m
}

// WithConcurrentMode defines concurrent generation, chunk workers define how mano chuncks will be executed concurrently.
//   - if chunkWorkers is less than 0, will not be added
func (m *MarotoBuilder) WithConcurrentMode(chunkWorkers int) *MarotoBuilder {
	validAndRun(func() { m.cfg.WithConcurrentMode(chunkWorkers) }, chunkWorkers > 0)
	return m
}

// WithSequentialMode defines that maroto will run in default mode.
func (m *MarotoBuilder) WithSequentialMode(on bool) *MarotoBuilder {
	validAndRun(func() { m.cfg.WithSequentialMode() }, on)
	return m
}

// WithSequentialLowMemoryMode defines that maroto will run focusing in reduce memory consumption,
// chunk workers define how many divisions the work will have.
//   - if chunkWorkers is less than 0, will not be added
func (m *MarotoBuilder) WithSequentialLowMemoryMode(chunkWorkers int) *MarotoBuilder {
	validAndRun(func() { m.cfg.WithSequentialLowMemoryMode(chunkWorkers) }, chunkWorkers > 0)
	return m
}

// WithDebug defines a debug behaviour where maroto will draw borders in everything.
func (m *MarotoBuilder) WithDebug(on bool) *MarotoBuilder {
	m.cfg.WithDebug(on)
	return m
}

// WithMaxGridSize defines a custom max grid sum which it will change the sum of column sizes.
//   - if maxGridSize is less than 0, will not be added
func (m *MarotoBuilder) WithMaxGridSize(maxGridSize int) *MarotoBuilder {
	validAndRun(func() { m.cfg.WithMaxGridSize(maxGridSize) }, maxGridSize > 0)
	return m
}

// WithDefaultFont defines a custom font, other than arial. This can be used to define a custom font as default.
//   - if font is nill, will not be added
func (m *MarotoBuilder) WithDefaultFont(font *propsmapper.Font) *MarotoBuilder {
	validAndRun(func() {
		m.cfg.WithDefaultFont(&props.Font{
			Family: font.Family, Style: fontstyle.Type(font.Style),
			Size: font.Size, Color: (*props.Color)(font.Color),
		})
	}, font != nil)
	return m
}

// WithPageNumber defines a string pattern to write the current page and total.
//   - if pageNumber is nill, will not be added
func (m *MarotoBuilder) WithPageNumber(pageNumber *propsmapper.PageNumber) *MarotoBuilder {
	validAndRun(func() {
		m.cfg.WithPageNumber(props.PageNumber{
			Pattern: pageNumber.Pattern, Place: props.Place(pageNumber.Place), Family: pageNumber.Family,
			Style: fontstyle.Type(pageNumber.Style), Size: pageNumber.Size,
			Color: (*props.Color)(pageNumber.Color),
		})
	}, pageNumber != nil)
	return m
}

// WithProtection defines protection types to the PDF document.
//   - if protectionmapper is nill, will not be added
func (m *MarotoBuilder) WithProtection(protectionmapper *propsmapper.Protection) *MarotoBuilder {
	validAndRun(func() {
		m.cfg.WithProtection(protection.Type(protectionmapper.Type), protectionmapper.UserPassword, protectionmapper.OwnerPassword)
	}, protectionmapper != nil)
	return m
}

// WithCompression defines compression.
func (m *MarotoBuilder) WithCompression(compression bool) *MarotoBuilder {
	m.cfg.WithCompression(compression)
	return m
}

// WithOrientation defines the page orientation. The default orientation is vertical,
// if horizontal is defined width and height will be flipped.
//   - if orientationMapper is nill, will not be added
func (m *MarotoBuilder) WithOrientation(orientationMapper string) *MarotoBuilder {
	validAndRun(func() { m.cfg.WithOrientation(orientation.Type(orientationMapper)) }, orientationMapper != "")
	return m
}

// WithMetadata customizes each metadata individually - Author, CreationDate, Creator, Keywords, subject and title
//   - if metadata is null or individual metadata is less than 0, metadata will not be added
func (m *MarotoBuilder) WithMetadata(metadata *propsmapper.Metadata) *MarotoBuilder {
	if metadata != nil {
		validAndRun(func() { m.cfg.WithAuthor(metadata.Author.Text, metadata.Author.UTF8) }, metadata.Author != nil)
		validAndRun(func() { m.cfg.WithCreationDate(*metadata.CreationDate) }, metadata.CreationDate != nil)
		validAndRun(func() { m.cfg.WithCreator(metadata.Creator.Text, metadata.Creator.UTF8) }, metadata.Creator != nil)
		validAndRun(func() { m.cfg.WithKeywords(metadata.KeywordsStr.Text, metadata.KeywordsStr.UTF8) }, metadata.KeywordsStr != nil)
		validAndRun(func() { m.cfg.WithSubject(metadata.Subject.Text, metadata.Subject.UTF8) }, metadata.Subject != nil)
		validAndRun(func() { m.cfg.WithTitle(metadata.Title.Text, metadata.Title.UTF8) }, metadata.Title != nil)
	}
	return m
}

// WithCustomFonts add custom fonts.
//   - if the font file cannot be loaded, an error will be returned
func (m *MarotoBuilder) WithCustomFonts(customFonts []*propsmapper.CustomFont) (*MarotoBuilder, error) {
	if len(customFonts) == 0 {
		return m, nil
	}
	newFonts, err := m.loadFonts(customFonts)
	if err != nil {
		return nil, err
	}
	m.cfg.WithCustomFonts(newFonts)
	return m, nil
}

// WithBackgroundImage defines the background image that will be applied in every page.
//   - if the image file cannot be loaded, an error will be returned
func (m *MarotoBuilder) WithBackgroundImage(backgroundImage string) (*MarotoBuilder, error) {
	if backgroundImage == "" {
		return m, nil
	}

	ext, imageBytes, err := m.repository.GetDocument(backgroundImage)
	if err != nil {
		return nil, err
	}
	m.cfg.WithBackgroundImage(imageBytes, extension.Type(ext))
	return m, nil
}

// WithDisableAutoPageBreak defines the option to disable automatic page breaks.
func (m *MarotoBuilder) WithDisableAutoPageBreak(disabled bool) *MarotoBuilder {
	m.cfg.WithDisableAutoPageBreak(disabled)
	return m
}

func validAndRun(setParam func(), parameter bool) {
	if parameter {
		setParam()
	}
}

func (m *MarotoBuilder) loadFonts(customFonts []*propsmapper.CustomFont) ([]*entity.CustomFont, error) {
	fontRepository := repository.New()

	for _, customFont := range customFonts {
		_, fontBytes, err := m.repository.GetDocument(customFont.File)
		if err != nil {
			return nil, err
		}
		fontRepository.AddUTF8FontFromBytes(customFont.Family, fontstyle.Type(customFont.Style), fontBytes)
	}
	return fontRepository.Load()
}
