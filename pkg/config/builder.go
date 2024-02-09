// Package config implements custom configuration builder.
package config

import (
	"strings"
	"time"

	"github.com/johnfercher/maroto/v2/pkg/consts/extension"

	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"

	"github.com/johnfercher/maroto/v2/pkg/consts/protection"

	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/consts/provider"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

// Builder is the abstraction responsible for global customizations on the document.
type Builder interface {
	WithPageSize(size pagesize.Type) Builder
	WithDimensions(width float64, height float64) Builder
	WithMargins(left float64, top float64, right float64) Builder
	WithWorkerPoolSize(poolSize int) Builder
	WithDebug(on bool) Builder
	WithMaxGridSize(maxGridSize int) Builder
	WithDefaultFont(font *props.Font) Builder
	WithPageNumber(pattern string, place props.Place) Builder
	WithProtection(protectionType protection.Type, userPassword, ownerPassword string) Builder
	WithCompression(compression bool) Builder
	WithOrientation(orientation orientation.Type) Builder
	WithAuthor(author string, isUTF8 bool) Builder
	WithCreator(creator string, isUTF8 bool) Builder
	WithSubject(subject string, isUTF8 bool) Builder
	WithTitle(title string, isUTF8 bool) Builder
	WithCreationDate(time time.Time) Builder
	WithCustomFonts([]*entity.CustomFont) Builder
	WithBackgroundImage([]byte, extension.Type) Builder
	WithDisableAutoPageBreak(disabled bool) Builder
	Build() *entity.Config
}

type CfgBuilder struct {
	providerType         provider.Type
	dimensions           *entity.Dimensions
	margins              *entity.Margins
	workerPoolSize       int
	debug                bool
	maxGridSize          int
	defaultFont          *props.Font
	customFonts          []*entity.CustomFont
	pageNumberPattern    string
	pageNumberPlace      props.Place
	protection           *entity.Protection
	compression          bool
	pageSize             *pagesize.Type
	orientation          orientation.Type
	metadata             *entity.Metadata
	backgroundImage      *entity.Image
	disableAutoPageBreak bool
}

// NewBuilder is responsible to create an instance of Builder.
func NewBuilder() Builder {
	return &CfgBuilder{
		providerType: provider.Gofpdf,
		margins: &entity.Margins{
			Left:   pagesize.DefaultLeftMargin,
			Right:  pagesize.DefaultRightMargin,
			Top:    pagesize.DefaultTopMargin,
			Bottom: pagesize.DefaultBottomMargin,
		},
		maxGridSize: pagesize.DefaultMaxGridSum,
		defaultFont: &props.Font{
			Size:   pagesize.DefaultFontSize,
			Family: fontfamily.Arial,
			Style:  fontstyle.Normal,
			Color:  &props.BlackColor,
		},
		metadata: &entity.Metadata{},
	}
}

// WithPageSize defines the page size, ex: A4, A4 and etc.
func (b *CfgBuilder) WithPageSize(size pagesize.Type) Builder {
	if size == "" {
		return b
	}

	b.pageSize = &size
	return b
}

// WithDimensions defines custom page dimensions, this overrides page size.
func (b *CfgBuilder) WithDimensions(width float64, height float64) Builder {
	if width <= 0 || height <= 0 {
		return b
	}

	b.dimensions = &entity.Dimensions{
		Width:  width,
		Height: height,
	}

	return b
}

// WithMargins defines custom margins, bottom margin is not customizable due to gofpdf limitations.
func (b *CfgBuilder) WithMargins(left float64, top float64, right float64) Builder {
	if left < pagesize.MinLeftMargin {
		return b
	}

	if top < pagesize.MinRightMargin {
		return b
	}

	if right < pagesize.MinTopMargin {
		return b
	}

	b.margins.Left = left
	b.margins.Top = top
	b.margins.Right = right

	return b
}

// WithWorkerPoolSize defines go routine workers, when defined this will execute maroto concurrently.
func (b *CfgBuilder) WithWorkerPoolSize(poolSize int) Builder {
	if poolSize < 0 {
		return b
	}

	b.workerPoolSize = poolSize
	return b
}

// WithDebug defines a debug behaviour where maroto will draw borders in everything.
func (b *CfgBuilder) WithDebug(on bool) Builder {
	b.debug = on
	return b
}

// WithMaxGridSize defines a custom max grid sum which it will change the sum of column sizes.
func (b *CfgBuilder) WithMaxGridSize(maxGridSize int) Builder {
	if maxGridSize < 0 {
		return b
	}

	b.maxGridSize = maxGridSize
	return b
}

// WithDefaultFont defines a custom font, other than arial. This can be used to define a custom font as default.
func (b *CfgBuilder) WithDefaultFont(font *props.Font) Builder {
	if font == nil {
		return b
	}

	if font.Family != "" {
		b.defaultFont.Family = font.Family
	}

	if font.Size != 0 {
		b.defaultFont.Size = font.Size
	}

	if font.Style != "" {
		b.defaultFont.Style = font.Style
	}

	if font.Color != nil {
		b.defaultFont.Color = font.Color
	}

	return b
}

// WithCustomFonts add custom fonts.
func (b *CfgBuilder) WithCustomFonts(customFonts []*entity.CustomFont) Builder {
	if customFonts == nil {
		return b
	}

	b.customFonts = customFonts
	return b
}

// WithPageNumber defines a string pattern to write the current page and total.
func (b *CfgBuilder) WithPageNumber(pattern string, place props.Place) Builder {
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

// WithProtection defines protection types to the PDF document.
func (b *CfgBuilder) WithProtection(protectionType protection.Type, userPassword, ownerPassword string) Builder {
	b.protection = &entity.Protection{
		Type:          protectionType,
		UserPassword:  userPassword,
		OwnerPassword: ownerPassword,
	}

	return b
}

// WithCompression defines compression.
func (b *CfgBuilder) WithCompression(compression bool) Builder {
	b.compression = compression
	return b
}

// WithOrientation defines the page orientation. The default orientation is vertical,
// if horizontal is defined width and height will be flipped.
func (b *CfgBuilder) WithOrientation(orientation orientation.Type) Builder {
	b.orientation = orientation
	return b
}

// WithAuthor defines the author name metadata.
func (b *CfgBuilder) WithAuthor(author string, isUTF8 bool) Builder {
	if author == "" {
		return b
	}

	b.metadata.Author = &entity.Utf8Text{
		Text: author,
		UTF8: isUTF8,
	}

	return b
}

// WithCreator defines the creator name metadata.
func (b *CfgBuilder) WithCreator(creator string, isUTF8 bool) Builder {
	if creator == "" {
		return b
	}

	b.metadata.Creator = &entity.Utf8Text{
		Text: creator,
		UTF8: isUTF8,
	}

	return b
}

// WithSubject defines the subject metadata.
func (b *CfgBuilder) WithSubject(subject string, isUTF8 bool) Builder {
	if subject == "" {
		return b
	}

	b.metadata.Subject = &entity.Utf8Text{
		Text: subject,
		UTF8: isUTF8,
	}

	return b
}

// WithTitle defines the title metadata.
func (b *CfgBuilder) WithTitle(title string, isUTF8 bool) Builder {
	if title == "" {
		return b
	}

	b.metadata.Title = &entity.Utf8Text{
		Text: title,
		UTF8: isUTF8,
	}

	return b
}

// WithCreationDate defines the creation date metadata.
func (b *CfgBuilder) WithCreationDate(time time.Time) Builder {
	if time.IsZero() {
		return b
	}

	b.metadata.CreationDate = &time

	return b
}

// WithBackgroundImage defines the background image that will be applied in every page.
func (b *CfgBuilder) WithBackgroundImage(bytes []byte, ext extension.Type) Builder {
	b.backgroundImage = &entity.Image{
		Bytes:     bytes,
		Extension: ext,
	}

	return b
}

// WithDisableAutoPageBreak defines the option to disable automatic page breaks.
func (b *CfgBuilder) WithDisableAutoPageBreak(disabled bool) Builder {
	b.disableAutoPageBreak = disabled
	return b
}

// Build finalizes the customization returning the entity.Config.
func (b *CfgBuilder) Build() *entity.Config {
	return &entity.Config{
		ProviderType:         b.providerType,
		Dimensions:           b.getDimensions(),
		Margins:              b.margins,
		WorkersQuantity:      b.workerPoolSize,
		Debug:                b.debug,
		MaxGridSize:          b.maxGridSize,
		DefaultFont:          b.defaultFont,
		PageNumberPattern:    b.pageNumberPattern,
		PageNumberPlace:      b.pageNumberPlace,
		Protection:           b.protection,
		Compression:          b.compression,
		Metadata:             b.metadata,
		CustomFonts:          b.customFonts,
		BackgroundImage:      b.backgroundImage,
		DisableAutoPageBreak: b.disableAutoPageBreak,
	}
}

func (b *CfgBuilder) getDimensions() *entity.Dimensions {
	if b.dimensions != nil {
		return b.dimensions
	}

	pageSize := pagesize.A4
	if b.pageSize != nil {
		pageSize = *b.pageSize
	}

	width, height := pagesize.GetDimensions(pageSize)
	dimensions := &entity.Dimensions{
		Width:  width,
		Height: height,
	}

	if b.orientation == orientation.Horizontal && height > width {
		dimensions.Width, dimensions.Height = dimensions.Height, dimensions.Width
	}

	return dimensions
}
