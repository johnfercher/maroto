package main

import (
	"github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/config"
)

func main() {
	cfg := config.NewBuilder().
		WithDimensions(&config.Dimensions{
			Width:  200,
			Height: 200,
		}).
		Build()

	m := v2.NewMaroto(cfg)

	// Generate
}
