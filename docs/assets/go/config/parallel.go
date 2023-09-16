package main

import (
	"github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/config"
)

func main() {
	cfg := config.NewBuilder().
		WithWorkerPoolSize(10).
		Build()

	m := v2.NewMaroto(cfg)

	// Generate
}
