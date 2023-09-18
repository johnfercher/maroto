package config

import (
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/johnfercher/maroto/v2/pkg/provider"
)

type Maroto struct {
	ProviderType provider.Type
	Dimensions   *Dimensions
	Margins      *Margins
	Font         *props.Font
	CustomFonts  []*CustomFont
	Workers      int
	Debug        bool
	MaxGridSize  int
}
