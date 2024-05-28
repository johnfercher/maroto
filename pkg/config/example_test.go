package config_test

import (
	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
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

// ExampleCfgBuilder_WithMargins demonstrates how to customize margins
func ExampleCfgBuilder_WithMargins() {
	// Ignore all margin customizations that are less than 0.
	cfg := config.NewBuilder().
		WithMargins(&entity.Margins{}).
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
