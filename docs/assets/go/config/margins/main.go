package main

import (
	v2 "github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/config"
)

func main() {
	cfg := config.NewBuilder().
		WithMargins(&config.Margins{
			Left:  10,
			Right: 10,
			Top:   10,
		}).
		Build()

	_ = v2.NewMaroto(cfg)

	// Generate
}
