package main

import (
	"github.com/johnfercher/maroto/v2/pkg"
	"github.com/johnfercher/maroto/v2/pkg/config"
)

func main() {
	cfg := config.NewBuilder().
		WithDimensions(&config.Dimensions{
			Width:  200,
			Height: 200,
		}).
		Build()

	_ = pkg.NewMaroto(cfg)

	// Generate
}
