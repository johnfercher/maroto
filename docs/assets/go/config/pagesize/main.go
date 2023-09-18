package main

import (
	"github.com/johnfercher/maroto/v2/pkg"
	"github.com/johnfercher/maroto/v2/pkg/config"
)

func main() {
	cfg := config.NewBuilder().
		WithPageSize(config.A4).
		Build()

	_ = pkg.NewMaroto(cfg)

	// Generate
}
