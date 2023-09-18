package main

import (
	"github.com/johnfercher/maroto/v2/pkg"
	"github.com/johnfercher/maroto/v2/pkg/config"
)

func main() {
	cfg := config.NewBuilder().
		WithMaxGridSize(14).
		Build()

	_ = pkg.NewMaroto(cfg)

	// Generate
}
