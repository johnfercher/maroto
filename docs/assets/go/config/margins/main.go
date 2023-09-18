package main

import (
	"github.com/johnfercher/maroto/v2/maroto"
	"github.com/johnfercher/maroto/v2/maroto/config"
)

func main() {
	cfg := config.NewBuilder().
		WithMargins(&config.Margins{
			Left:  10,
			Right: 10,
			Top:   10,
		}).
		Build()

	_ = maroto.NewMaroto(cfg)

	// Generate
}
