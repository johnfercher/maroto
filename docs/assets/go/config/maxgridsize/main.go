package main

import (
	"github.com/johnfercher/v2/maroto"
	"github.com/johnfercher/v2/maroto/config"
)

func main() {
	cfg := config.NewBuilder().
		WithMaxGridSize(14).
		Build()

	_ = maroto.NewMaroto(cfg)

	// Generate
}
