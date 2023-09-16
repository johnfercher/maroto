package main

import (
	"github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/config"
)

func main() {
	cfg := config.NewBuilder().
		WithMargins(&config.Margins{
			Left:   10,
			Right:  10,
			Top:    10,
			Bottom: 20,
		}).
		Build()

	m := v2.NewMaroto(cfg)

	// Generate
}
