package buildermapper

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/propsmapper"
)

type Builder struct {
	PageSize                string
	Dimensions              *propsmapper.Dimensions
	Margins                 *propsmapper.Margins
	ConcurrentMode          int
	SequentialMode          bool
	SequentialLowMemoryMode int
	Debug                   bool
	MaxGridSize             int
	DefaultFont             *propsmapper.Font
	CustomFonts             []*propsmapper.CustomFont
	PageNumber              *propsmapper.PageNumber
	Protection              *propsmapper.Protection
	Compression             bool
	Orientation             string
	Metadata                *propsmapper.Metadata
	DisableAutoPageBreak    bool
	GenerationMode          string
	BackgroundImage         string
}

// NewBuilder is responsible for creating Builder properties. If an invalid property is provided, a default value will be assigned.
func NewBuilder(builder interface{}) (*Builder, error) {
	builderMap, ok := builder.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("builder settings could not be deserialized")
	}

	return &Builder{
		PageSize:                factoryField(builderMap["page_size"], ""),
		Dimensions:              propsmapper.NewDimensions(builderMap["dimensions"]),
		Margins:                 propsmapper.NewMargins(builderMap["margins"]),
		ConcurrentMode:          int(factoryField(builderMap["concurrent_mode"], -1.0)),
		SequentialMode:          factoryField(builderMap["sequential_mode"], false),
		SequentialLowMemoryMode: int(factoryField(builderMap["sequential_low_memory_mode"], -1.0)),
		Debug:                   factoryField(builderMap["debug"], false),
		MaxGridSize:             int(factoryField(builderMap["max_grid_size"], -1.0)),
		DefaultFont:             propsmapper.NewFont(builderMap["default_font"]),
		CustomFonts:             propsmapper.NewCustomFonts(builderMap["custom_font"]),
		PageNumber:              propsmapper.NewPageNumber(builderMap["page_number"]),
		Protection:              propsmapper.NewProtection(builderMap["protection"]),
		Compression:             factoryField(builderMap["compression"], false),
		Orientation:             propsmapper.NewOrientation(factoryField(builderMap["orientation"], "")),
		Metadata:                propsmapper.NewMetadata(builderMap["metadata"]),
		DisableAutoPageBreak:    factoryField(builderMap["disable_auto_page_break"], false),
		GenerationMode:          factoryField(builderMap["generation_mode"], ""),
		BackgroundImage:         factoryField(builderMap["background_image"], ""),
	}, nil
}

func factoryField[T any](val interface{}, defaultValue T) T {
	result, ok := val.(T)
	if !ok {
		return defaultValue
	}
	return result
}
