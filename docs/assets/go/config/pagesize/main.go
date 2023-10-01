package main

import (
	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
)

func main() {
	cfg := config.NewBuilder().
		WithPageSize(pagesize.A4).
		Build()

	_ = maroto.New(cfg)

	// Generate
}
