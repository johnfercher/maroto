package entity

import (
	"github.com/johnfercher/maroto/v2/pkg/consts/provider"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type Config struct {
	ProviderType      provider.Type
	Dimensions        *Dimensions
	Margins           *Margins
	DefaultFont       *props.Font
	CustomFonts       []*CustomFont
	Workers           int
	Debug             bool
	MaxGridSize       int
	PageNumberPattern string
	PageNumberPlace   props.Place
	Protection        *Protection
	Compression       bool
	Metadata          *Metadata
	BackgroundImage   *Image
}

func (c *Config) ToMap() map[string]interface{} {
	m := make(map[string]interface{})

	if c.ProviderType != "" {
		m["config_provider_type"] = c.ProviderType
	}

	if c.Dimensions != nil {
		m = c.Dimensions.AppendMap(m)
	}

	if c.Margins != nil {
		m = c.Margins.AppendMap(m)
	}

	if c.DefaultFont != nil {
		m = c.DefaultFont.AppendMap(m)
	}

	if c.Workers != 0 {
		m["config_workers"] = c.Workers
	}

	if c.Debug {
		m["config_debug"] = c.Debug
	}

	if c.MaxGridSize != 0 {
		m["config_max_grid_sum"] = c.MaxGridSize
	}

	if c.PageNumberPattern != "" {
		m["config_page_number_pattern"] = c.PageNumberPattern
	}

	if c.PageNumberPlace != "" {
		m["config_page_number_place"] = c.PageNumberPlace
	}

	if c.Protection != nil {
		m = c.Protection.AppendMap(m)
	}

	if c.Compression {
		m["config_compression"] = c.Compression
	}

	if c.Metadata != nil {
		m = c.Metadata.AppendMap(m)
	}

	if c.BackgroundImage != nil {
		m = c.BackgroundImage.AppendMap(m)
	}

	return m
}
