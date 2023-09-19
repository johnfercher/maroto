package config

import (
	"github.com/johnfercher/maroto/v2/pkg/consts/provider"
	"github.com/johnfercher/maroto/v2/pkg/props"
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
