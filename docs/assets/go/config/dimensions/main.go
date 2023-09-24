package main

import (
	"github.com/johnfercher/maroto/v2/pkg"
	"github.com/johnfercher/maroto/v2/pkg/config"
)

func main() {
	cfg := config.NewBuilder().
		WithDimensions(200, 200).
		Build()

	_ = pkg.NewMaroto(cfg)

	// Generate
}
