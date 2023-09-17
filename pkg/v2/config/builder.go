package config

import (
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/v2/provider"
)

type builder struct {
	providerType   provider.Type
	dimensions     *Dimensions
	margins        *Margins
	workerPoolSize int
	debug          bool
	maxGridSize    int
	font           *Font
}

type Builder interface {
	WithPageSize(size PageSize) Builder
	WithDimensions(dimensions *Dimensions) Builder
	WithMargins(margins *Margins) Builder
	WithProvider(providerType provider.Type) Builder
	WithWorkerPoolSize(poolSize int) Builder
	WithDebug(on bool) Builder
	WithMaxGridSize(maxGridSize int) Builder
	WithFont(font *Font) Builder
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
		font: &Font{
			Size:   DefaultFontSize,
			Family: consts.Arial,
			Style:  consts.Normal,
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

func (b *builder) WithFont(font *Font) Builder {
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
	}
}

func (b *builder) getDimensions() *Dimensions {
	if b.dimensions != nil {
		return b.dimensions
	}

	return GetDimensions(A4)
}
