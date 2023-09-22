package config

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
}
