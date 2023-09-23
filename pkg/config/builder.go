package config

import (
	"strings"
	"time"

	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"

	"github.com/johnfercher/maroto/v2/pkg/consts/protection"

	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/consts/provider"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type Builder interface {
	WithPageSize(size pagesize.Type) Builder
	WithDimensions(dimensions *Dimensions) Builder
	WithMargins(margins *Margins) Builder
	WithProvider(providerType provider.Type) Builder
	WithWorkerPoolSize(poolSize int) Builder
	WithDebug(on bool) Builder
	WithMaxGridSize(maxGridSize int) Builder
	WithFont(font *props.Font) Builder
	AddUTF8Font(family string, style fontstyle.Type, file string) Builder
	WithPageNumber(pattern string, place props.Place) Builder
	WithProtection(protectionType protection.Type, userPassword, ownerPassword string) Builder
	WithCompression(compression bool) Builder
	WithOrientation(orientation orientation.Type) Builder
	WithAuthor(author string, isUTF8 bool) Builder
	WithCreator(creator string, isUTF8 bool) Builder
	WithSubject(subject string, isUTF8 bool) Builder
	WithTitle(title string, isUTF8 bool) Builder
	WithCreationDate(time time.Time) Builder
	Build() *Config
}

type builder struct {
	providerType      provider.Type
	dimensions        *Dimensions
	margins           *Margins
	workerPoolSize    int
	debug             bool
	maxGridSize       int
	font              *props.Font
	customFonts       []*CustomFont
	pageNumberPattern string
	pageNumberPlace   props.Place
	protection        *Protection
	compression       bool
	orientation       orientation.Type
	metadata          *Metadata
}

func NewBuilder() Builder {
	return &builder{
		providerType: provider.Gofpdf,
		margins: &Margins{
			Left:  pagesize.MinLeftMargin,
			Right: pagesize.MinRightMargin,
			Top:   pagesize.MinTopMargin,
		},
		maxGridSize: pagesize.DefaultMaxGridSum,
		font: &props.Font{
			Size:   pagesize.DefaultFontSize,
			Family: fontfamily.Arial,
			Style:  fontstyle.Normal,
			Color:  props.NewBlack(),
		},
		metadata: &Metadata{},
	}
}

func (b *builder) WithPageSize(size pagesize.Type) Builder {
	if size == "" {
		return b
	}

	width, height := pagesize.GetDimensions(size)
	b.dimensions = &Dimensions{
		Width:  width,
		Height: height,
	}

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

	if margins.Left < pagesize.MinLeftMargin {
		return b
	}

	if margins.Right < pagesize.MinRightMargin {
		return b
	}

	if margins.Top < pagesize.MinTopMargin {
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

func (b *builder) AddUTF8Font(family string, style fontstyle.Type, file string) Builder {
	if family == "" {
		return b
	}

	if !style.IsValid() {
		return b
	}

	if file == "" {
		return b
	}

	b.customFonts = append(b.customFonts, &CustomFont{
		Family: family,
		Style:  style,
		File:   file,
	})

	return b
}

func (b *builder) WithPageNumber(pattern string, place props.Place) Builder {
	if !strings.Contains(pattern, "{current}") && !strings.Contains(pattern, "{total}") {
		return b
	}

	if !place.IsValid() {
		return b
	}

	b.pageNumberPattern = pattern
	b.pageNumberPlace = place

	return b
}

func (b *builder) WithProtection(protectionType protection.Type, userPassword, ownerPassword string) Builder {
	b.protection = &Protection{
		Type:          protectionType,
		UserPassword:  userPassword,
		OwnerPassword: ownerPassword,
	}

	return b
}

func (b *builder) WithCompression(compression bool) Builder {
	b.compression = compression
	return b
}

func (b *builder) WithOrientation(orientation orientation.Type) Builder {
	b.orientation = orientation
	return b
}

func (b *builder) WithAuthor(author string, isUTF8 bool) Builder {
	if author == "" {
		return b
	}

	b.metadata.Author = &Utf8Text{
		Text: author,
		UTF8: isUTF8,
	}

	return b
}

func (b *builder) WithCreator(creator string, isUTF8 bool) Builder {
	if creator == "" {
		return b
	}

	b.metadata.Creator = &Utf8Text{
		Text: creator,
		UTF8: isUTF8,
	}

	return b
}

func (b *builder) WithSubject(subject string, isUTF8 bool) Builder {
	if subject == "" {
		return b
	}

	b.metadata.Subject = &Utf8Text{
		Text: subject,
		UTF8: isUTF8,
	}

	return b
}

func (b *builder) WithTitle(title string, isUTF8 bool) Builder {
	if title == "" {
		return b
	}

	b.metadata.Title = &Utf8Text{
		Text: title,
		UTF8: isUTF8,
	}

	return b
}

func (b *builder) WithCreationDate(time time.Time) Builder {
	if time.IsZero() {
		return b
	}

	b.metadata.CreationDate = time

	return b
}

func (b *builder) Build() *Config {
	return &Config{
		ProviderType:      b.providerType,
		Dimensions:        b.getDimensions(),
		Margins:           b.margins,
		Workers:           b.workerPoolSize,
		Debug:             b.debug,
		MaxGridSize:       b.maxGridSize,
		DefaultFont:       b.font,
		CustomFonts:       b.customFonts,
		PageNumberPattern: b.pageNumberPattern,
		PageNumberPlace:   b.pageNumberPlace,
		Protection:        b.protection,
		Compression:       b.compression,
		Metadata:          b.metadata,
	}
}

func (b *builder) getDimensions() *Dimensions {
	if b.dimensions != nil {
		return b.dimensions
	}

	width, height := pagesize.GetDimensions(pagesize.A4)
	dimensions := &Dimensions{
		Width:  width,
		Height: height,
	}

	if b.orientation == orientation.Landscape && height > width {
		dimensions.Width, dimensions.Height = dimensions.Height, dimensions.Width
	}

	return dimensions
}
