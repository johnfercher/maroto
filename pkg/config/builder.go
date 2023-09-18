package config

import (
	"github.com/johnfercher/maroto/v2/pkg/color"
	"github.com/johnfercher/maroto/v2/pkg/consts"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/johnfercher/maroto/v2/pkg/provider"
)

type builder struct {
	providerType   provider.Type
	dimensions     *Dimensions
	margins        *Margins
	workerPoolSize int
	debug          bool
	maxGridSize    int
	font           *props.Font
	customFonts    []*CustomFont
}

type Builder interface {
	WithPageSize(size PageSize) Builder
	WithDimensions(dimensions *Dimensions) Builder
	WithMargins(margins *Margins) Builder
	WithProvider(providerType provider.Type) Builder
	WithWorkerPoolSize(poolSize int) Builder
	WithDebug(on bool) Builder
	WithMaxGridSize(maxGridSize int) Builder
	WithFont(font *props.Font) Builder
	AddUTF8Font(customFont *CustomFont) Builder
	Build() *Maroto
}

func NewBuilder() Builder {
	return &builder{
		providerType: provider.Gofpdf,
		margins: &Margins{
			Left:  MinLeftMargin,
			Right: MinRightMargin,
			Top:   MinTopMargin,
		},
		maxGridSize: DefaultMaxGridSum,
		font: &props.Font{
			Size:   DefaultFontSize,
			Family: consts.Arial,
			Style:  consts.Normal,
			Color:  color.NewBlack(),
		},
	}
}

func (b *builder) WithPageSize(size PageSize) Builder {
	if size == "" {
		return b
	}

	b.dimensions = GetDimensions(size)

	return b
}

func (b *builder) WithDimensions(dimensions *Dimensions) Builder {
	if dimensions == nil {
		return b
	}
	if dimensions.Width <= 0 || dimensions.Height <= 0 {
		return b
	}

	b.dimensions = dimensions

	return b
}

func (b *builder) WithMargins(margins *Margins) Builder {
	if margins == nil {
		return b
	}

	if margins.Left < MinLeftMargin {
		return b
	}

	if margins.Right < MinRightMargin {
		return b
	}

	if margins.Top < MinTopMargin {
		return b
	}

	b.margins = margins

	return b
}

func (b *builder) WithProvider(providerType provider.Type) Builder {
	if providerType == "" {
		return b
	}

	b.providerType = providerType
	return b
}

func (b *builder) WithWorkerPoolSize(poolSize int) Builder {
	if poolSize < 0 {
		return b
	}

	b.workerPoolSize = poolSize
	return b
}

func (b *builder) WithDebug(on bool) Builder {
	b.debug = on
	return b
}

func (b *builder) WithMaxGridSize(maxGridSize int) Builder {
	if maxGridSize < 0 {
		return b
	}

	b.maxGridSize = maxGridSize
	return b
}

func (b *builder) WithFont(font *props.Font) Builder {
	if font == nil {
		return b
	}

	if font.Family != "" {
		b.font.Family = font.Family
	}

	if font.Size != 0 {
		b.font.Size = font.Size
	}

	if font.Style != "" {
		b.font.Style = font.Style
	}

	if font.Color != nil {
		b.font.Color = font.Color
	}

	return b
}

func (b *builder) AddUTF8Font(customFont *CustomFont) Builder {
	if customFont == nil {
		return b
	}

	if customFont.Family == "" {
		return b
	}

	if !customFont.Style.IsValid() {
		return b
	}

	if customFont.File == "" {
		return b
	}

	b.customFonts = append(b.customFonts, customFont)
	return b
}

func (b *builder) Build() *Maroto {
	return &Maroto{
		ProviderType: b.providerType,
		Dimensions:   b.getDimensions(),
		Margins:      b.margins,
		Workers:      b.workerPoolSize,
		Debug:        b.debug,
		MaxGridSize:  b.maxGridSize,
		Font:         b.font,
		CustomFonts:  b.customFonts,
	}
}

func (b *builder) getDimensions() *Dimensions {
	if b.dimensions != nil {
		return b.dimensions
	}

	return GetDimensions(A4)
}
