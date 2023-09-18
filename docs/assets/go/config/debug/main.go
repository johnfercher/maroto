package main

import (
	"github.com/johnfercher/maroto/v2/maroto"
	"github.com/johnfercher/maroto/v2/maroto/config"
)

func main() {
	cfg := config.NewBuilder().
		WithDebug(true).
		Build()

	_ = maroto.NewMaroto(cfg)

	// Generate
}
