package main

import (
	v2 "github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/config"
)

func main() {
	cfg := config.NewBuilder().
		WithPageSize(config.A4).
		Build()

	_ = v2.NewMaroto(cfg)

	// Generate
}
