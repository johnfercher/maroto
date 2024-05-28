package config_test

import (
	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
)

// ExampleNewBuilder demonstrates how to use builder.
func ExampleNewBuilder() {
	cfg := config.NewBuilder().Build()

	_ = maroto.New(cfg)

	// generate document
}

// ExampleCfgBuilder_Build demonstrates how to build configs.
func ExampleCfgBuilder_Build() {
	cfg := config.NewBuilder().Build()

	_ = maroto.New(cfg)

	// generate document
}

// ExampleCfgBuilder_WithPageSize demonstrates how to customize page size.
func ExampleCfgBuilder_WithPageSize() {
	// If pagesize is invalid, then ignore customization.
	cfg := config.NewBuilder().
		WithPageSize(pagesize.A5).
		Build()

	_ = maroto.New(cfg)

	// generate document
}

// ExampleCfgBuilder_WithTopMargin demonstrates how to customize margin.
func ExampleCfgBuilder_WithTopMargin() {
	// If top less than minimum, ignore customization.
	cfg := config.NewBuilder().
		WithTopMargin(15).
		Build()

	_ = maroto.New(cfg)

	// generate document
}

// ExampleCfgBuilder_WithRightMargin demonstrates how to customize margin.
func ExampleCfgBuilder_WithRightMargin() {
	// If top less than minimum, ignore customization.
	cfg := config.NewBuilder().
		WithRightMargin(15).
		Build()

	_ = maroto.New(cfg)

	// generate document
}

// ExampleCfgBuilder_WithLeftMargin demonstrates how to customize margin.
func ExampleCfgBuilder_WithLeftMargin() {
	// If top less than minimum, ignore customization.
	cfg := config.NewBuilder().
		WithLeftMargin(15).
		Build()

	_ = maroto.New(cfg)

	// generate document
}

// ExampleCfgBuilder_WithBottomMargin demonstrates how to customize margin.
func ExampleCfgBuilder_WithBottomMargin() {
	// If top less than minimum, ignore customization.
	cfg := config.NewBuilder().
		WithBottomMargin(15).
		Build()

	_ = maroto.New(cfg)

	// generate document
}

// ExampleCfgBuilder_WithConcurrentMode demonstrates how to enable concurrent generation.
func ExampleCfgBuilder_WithConcurrentMode() {
	// if chunkWorkers is less than 1, then ignore customization.
	chunkWorkers := 7
	cfg := config.NewBuilder().
		WithConcurrentMode(chunkWorkers).
		Build()

	_ = maroto.New(cfg)

	// generate document
}
