package main

import (
	"github.com/johnfercher/v2/maroto"
	"github.com/johnfercher/v2/maroto/config"
)

func main() {
	cfg := config.NewBuilder().
		WithDimensions(&config.Dimensions{
			Width:  200,
			Height: 200,
		}).
		Build()

	_ = maroto.NewMaroto(cfg)

	// Generate
}
