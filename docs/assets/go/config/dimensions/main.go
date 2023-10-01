package main

import (
	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/config"
)

func main() {
	cfg := config.NewBuilder().
		WithDimensions(200, 200).
		Build()

	_ = maroto.New(cfg)

	// Generate
}
