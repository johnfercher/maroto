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
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/propsmapper"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/johnfercher/maroto/v2/pkg/repository"
)

type MarotoBuilder struct {
	cfg        config.Builder
	repository core.ProcessorRepository
}

func NewMarotoBuilder(repository core.ProcessorRepository, cfg config.Builder) *MarotoBuilder {
	return &MarotoBuilder{
		repository: repository,
		cfg:        cfg,
	}
}

func (m *MarotoBuilder) WithPageSize(size string) config.Builder {
	validAndRun(func() { m.cfg.WithPageSize(pagesize.Type(size)) }, size != "")
	return m.cfg
}

// WithDimensions defines custom page dimensions, this overrides page size.
func (m *MarotoBuilder) WithDimensions(dimensions *propsmapper.Dimensions) config.Builder {
	validAndRun(func() { m.cfg.WithDimensions(dimensions.Width, dimensions.Height) }, dimensions != nil)
	return m.cfg
}

// WithLeftMargin customize margin.
func (m *MarotoBuilder) WithMargin(margin *propsmapper.Margins) config.Builder {
	validAndRun(func() { m.cfg.WithBottomMargin(margin.Bottom) }, margin.Bottom >= 0)
	validAndRun(func() { m.cfg.WithTopMargin(margin.Top) }, margin.Top >= 0)
	validAndRun(func() { m.cfg.WithLeftMargin(margin.Left) }, margin.Left >= 0)
	validAndRun(func() { m.cfg.WithRightMargin(margin.Right) }, margin.Right >= 0)
	return m.cfg
}

// WithConcurrentMode defines concurrent generation, chunk workers define how mano chuncks
// will be executed concurrently.
func (m *MarotoBuilder) WithConcurrentMode(chunkWorkers int) config.Builder {
	validAndRun(func() { m.cfg.WithConcurrentMode(chunkWorkers) }, chunkWorkers > 0)
	return m.cfg
}

// WithSequentialMode defines that maroto will run in default mode.
func (m *MarotoBuilder) WithSequentialMode(on bool) config.Builder {
	m.cfg.WithSequentialMode()
	return m.cfg
}

// WithSequentialLowMemoryMode defines that maroto will run focusing in reduce memory consumption,
// chunk workers define how many divisions the work will have.
func (m *MarotoBuilder) WithSequentialLowMemoryMode(chunkWorkers int) config.Builder {
	validAndRun(func() { m.cfg.WithSequentialLowMemoryMode(chunkWorkers) }, chunkWorkers > 0)
	return m.cfg
}

// WithDebug defines a debug behaviour where maroto will draw borders in everything.
func (m *MarotoBuilder) WithDebug(on bool) config.Builder {
	m.cfg.WithDebug(on)
	return m.cfg
}

// WithMaxGridSize defines a custom max grid sum which it will change the sum of column sizes.
func (m *MarotoBuilder) WithMaxGridSize(maxGridSize int) config.Builder {
	validAndRun(func() { m.cfg.WithMaxGridSize(maxGridSize) }, maxGridSize > 0)
	return m.cfg
}

// WithDefaultFont defines a custom font, other than arial. This can be used to define a custom font as default.
func (m *MarotoBuilder) WithDefaultFont(font *propsmapper.Font) config.Builder {
	validAndRun(func() {
		m.cfg.WithDefaultFont(&props.Font{
			Family: font.Family, Style: fontstyle.Type(font.Style),
			Size: font.Size, Color: (*props.Color)(font.Color),
		})
	}, font != nil)
	return m.cfg
}

// WithPageNumber defines a string pattern to write the current page and total.
func (m *MarotoBuilder) WithPageNumber(pageNumber *propsmapper.PageNumber) config.Builder {
	validAndRun(func() {
		m.cfg.WithPageNumber(props.PageNumber{
			Pattern: pageNumber.Pattern, Place: props.Place(pageNumber.Place), Family: pageNumber.Family,
			Style: fontstyle.Type(pageNumber.Style), Size: pageNumber.Size,
			Color: &props.Color{Red: pageNumber.Color.Red, Green: pageNumber.Color.Green, Blue: pageNumber.Color.Blue},
		})
	}, pageNumber != nil)
	return m.cfg
}

// WithProtection defines protection types to the PDF document.
func (m *MarotoBuilder) WithProtection(protectionmapper *propsmapper.Protection) config.Builder {
	validAndRun(func() {
		m.cfg.WithProtection(protection.Type(protectionmapper.Type), protectionmapper.UserPassword, protectionmapper.OwnerPassword)
	}, protectionmapper != nil)
	return m.cfg
}

// WithCompression defines compression.
func (m *MarotoBuilder) WithCompression(compression bool) config.Builder {
	m.cfg.WithCompression(compression)
	return m.cfg
}

// WithOrientation defines the page orientation. The default orientation is vertical,
// if horizontal is defined width and height will be flipped.
func (m *MarotoBuilder) WithOrientation(orientationMapper string) config.Builder {
	validAndRun(func() { m.cfg.WithOrientation(orientation.Type(orientationMapper)) }, orientationMapper != "")
	return m.cfg
}

// WithAuthor defines the author name metadata.
func (m *MarotoBuilder) WithMetadata(metadata *propsmapper.Metadata) config.Builder {
	if metadata != nil {
		validAndRun(func() { m.cfg.WithAuthor(metadata.Author.Text, metadata.Author.UTF8) }, metadata.Author != nil)
		validAndRun(func() { m.cfg.WithCreationDate(*metadata.CreationDate) }, metadata.CreationDate != nil)
		validAndRun(func() { m.cfg.WithCreator(metadata.Creator.Text, metadata.Creator.UTF8) }, metadata.Creator != nil)
		validAndRun(func() { m.cfg.WithKeywords(metadata.KeywordsStr.Text, metadata.KeywordsStr.UTF8) }, metadata.KeywordsStr != nil)
		validAndRun(func() { m.cfg.WithSubject(metadata.Subject.Text, metadata.Subject.UTF8) }, metadata.Subject != nil)
		validAndRun(func() { m.cfg.WithTitle(metadata.Title.Text, metadata.Title.UTF8) }, metadata.Title != nil)
	}
	return m.cfg
}

// WithCustomFonts add custom fonts.
func (m *MarotoBuilder) WithCustomFonts(customFonts []*propsmapper.CustomFont) (config.Builder, error) {
	if len(customFonts) == 0 {
		return m.cfg, nil
	}
	newFonts, err := m.loadFonts(customFonts)
	if err != nil {
		return nil, err
	}
	m.cfg.WithCustomFonts(newFonts)
	return m.cfg, nil
}

// WithBackgroundImage defines the background image that will be applied in every page.
func (m *MarotoBuilder) WithBackgroundImage(backgroundImage string) (config.Builder, error) {
	if backgroundImage == "" {
		return m.cfg, nil
	}

	ext, imageBytes, err := m.repository.GetDocument(backgroundImage)
	if err != nil {
		return nil, err
	}
	m.cfg.WithBackgroundImage(imageBytes, extension.Type(ext))
	return m.cfg, nil
}

// WithDisableAutoPageBreak defines the option to disable automatic page breaks.
func (m *MarotoBuilder) WithDisableAutoPageBreak(disabled bool) config.Builder {
	m.cfg.WithDisableAutoPageBreak(disabled)
	return m.cfg
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
