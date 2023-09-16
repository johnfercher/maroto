package config

import (
	"github.com/johnfercher/maroto/pkg/v2/provider"
)

type builder struct {
	providerType provider.Type
	pageSize     PageSize
	dimensions   *Dimensions
	margins      *Margins
	threadPool   int
	debug        bool
}

type Builder interface {
	WithPageSize(size PageSize) Builder
	WithDimensions(dimensions *Dimensions) Builder
	WithMargins(margins *Margins) Builder
	WithProvider(providerType provider.Type) Builder
	WithThreadPool(pool int) Builder
	WithDebug(on bool) Builder
	Build() *Maroto
}

func NewBuilder() Builder {
	return &builder{
		providerType: provider.Gofpdf,
		pageSize:     A4,
		dimensions:   GetDimensions(A4),
		margins: &Margins{
			Left:   MinLeftMargin,
			Right:  MinRightMargin,
			Top:    MinTopMargin,
			Bottom: MinBottomMargin,
		},
	}
}

func (b *builder) WithPageSize(size PageSize) Builder {
	if size == "" {
		return b
	}

	b.pageSize = size
	b.dimensions = GetDimensions(size)

	return b
}

func (b *builder) WithDimensions(dimensions *Dimensions) Builder {
	if dimensions == nil {
		return b
	}
	if dimensions.Width != 0 && dimensions.Height != 0 {
		return b
	}

	b.dimensions = dimensions

	return b
}

func (b *builder) WithMargins(margins *Margins) Builder {
	if margins == nil {
		return b
	}

	/*
		We need to warrant that the margins are in a certain limit due to gofpdf limitations
	*/
	if margins.Left < MinLeftMargin {
		margins.Left = MinLeftMargin
	}

	if margins.Right < MinRightMargin {
		margins.Right = MinRightMargin
	}

	if margins.Top < MinTopMargin {
		margins.Top = MinTopMargin
	}

	if margins.Bottom < MinBottomMargin {
		margins.Bottom = MinBottomMargin
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

func (b *builder) WithThreadPool(pool int) Builder {
	if pool <= 0 {
		return b
	}

	b.threadPool = pool
	return b
}

func (b *builder) WithDebug(on bool) Builder {
	b.debug = on
	return b
}

func (b *builder) Build() *Maroto {
	return &Maroto{
		ProviderType: b.providerType,
		Dimensions:   b.dimensions,
		Margins:      b.margins,
		ThreadPool:   b.threadPool,
		Debug:        b.debug,
	}
}