package main

import (
	"github.com/johnfercher/v2/maroto"
	"github.com/johnfercher/v2/maroto/config"
)

func main() {
	cfg := config.NewBuilder().
		WithDebug(true).
		Build()

	_ = maroto.NewMaroto(cfg)

	// Generate
}
