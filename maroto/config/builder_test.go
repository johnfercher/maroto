package config_test

import (
	"fmt"
	"github.com/johnfercher/v2/maroto/config"
	"github.com/johnfercher/v2/maroto/consts"
	"github.com/johnfercher/v2/maroto/provider"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBuilder(t *testing.T) {
	// Act
	sut := config.NewBuilder()

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*config.builder", fmt.Sprintf("%T", sut))
}

func TestBuilder_Build(t *testing.T) {
	// Arrange
	sut := config.NewBuilder()

	// Act
	cfg := sut.Build()

	// Assert
	assert.Equal(t, provider.Gofpdf, cfg.ProviderType)
	assert.Equal(t, 12, cfg.MaxGridSize)
	assert.Equal(t, false, cfg.Debug)
	assert.Equal(t, 0, cfg.Workers)
	assert.Equal(t, 10.0, cfg.Margins.Top)
	assert.Equal(t, 10.0, cfg.Margins.Left)
	assert.Equal(t, 10.0, cfg.Margins.Right)
	assert.Equal(t, 210.0, cfg.Dimensions.Width)
	assert.Equal(t, 297.0, cfg.Dimensions.Height)
	assert.Equal(t, consts.Arial, cfg.Font.Family)
	assert.Equal(t, 10.0, cfg.Font.Size)
	assert.Equal(t, consts.Normal, cfg.Font.Style)
}

func TestBuilder_WithDebug(t *testing.T) {
	// Arrange
	sut := config.NewBuilder()

	// Act
	cfg := sut.WithDebug(true).Build()

	// Assert
	assert.Equal(t, true, cfg.Debug)
}

func TestBuilder_WithFont(t *testing.T) {
	t.Run("when font is nil, should not change the default value", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithFont(nil).Build()

		// Assert
		assert.Equal(t, consts.Arial, cfg.Font.Family)
		assert.Equal(t, 10.0, cfg.Font.Size)
		assert.Equal(t, consts.Normal, cfg.Font.Style)
	})

	t.Run("when family is filled, should change the default value", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithFont(&config.Font{
			Family: "new family",
		}).Build()

		// Assert
		assert.Equal(t, "new family", cfg.Font.Family)
		assert.Equal(t, 10.0, cfg.Font.Size)
		assert.Equal(t, consts.Normal, cfg.Font.Style)
	})

	t.Run("when style is filled, should change the default value", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithFont(&config.Font{
			Style: consts.Bold,
		}).Build()

		// Assert
		assert.Equal(t, consts.Arial, cfg.Font.Family)
		assert.Equal(t, 10.0, cfg.Font.Size)
		assert.Equal(t, consts.Bold, cfg.Font.Style)
	})

	t.Run("when size is filled, should change the default value", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithFont(&config.Font{
			Size: 13,
		}).Build()

		// Assert
		assert.Equal(t, consts.Arial, cfg.Font.Family)
		assert.Equal(t, 13.0, cfg.Font.Size)
		assert.Equal(t, consts.Normal, cfg.Font.Style)
	})
}

func TestBuilder_WithPageSize(t *testing.T) {
	t.Run("when page size is empty, should not change the default value", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithPageSize("").Build()

		// Assert
		assert.Equal(t, 210.0, cfg.Dimensions.Width)
		assert.Equal(t, 297.0, cfg.Dimensions.Height)
	})
	t.Run("when page size is filled, should change the default value", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithPageSize(config.A2).Build()

		// Assert
		assert.Equal(t, 419.9, cfg.Dimensions.Width)
		assert.Equal(t, 594.0, cfg.Dimensions.Height)
	})
}

func TestBuilder_WithProvider(t *testing.T) {
	t.Run("when provider is empty, should not change the default value", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithProvider("").Build()

		// Assert
		assert.Equal(t, provider.Gofpdf, cfg.ProviderType)
	})

	t.Run("when provider is filled, should change the default value", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithProvider(provider.HTML).Build()

		// Assert
		assert.Equal(t, provider.HTML, cfg.ProviderType)
	})
}

func TestBuilder_WithWorkerPoolSize(t *testing.T) {
	t.Run("when worker pool size is invalid, should not change the default value", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithWorkerPoolSize(-1).Build()

		// Assert
		assert.Equal(t, 0, cfg.Workers)
	})

	t.Run("when worker pool size is valid, should change the default value", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithWorkerPoolSize(7).Build()

		// Assert
		assert.Equal(t, 7, cfg.Workers)
	})
}

func TestBuilder_WithDimensions(t *testing.T) {
	t.Run("when dimensions is nil, should not change the default value", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithDimensions(nil).Build()

		// Assert
		assert.Equal(t, 210.0, cfg.Dimensions.Width)
		assert.Equal(t, 297.0, cfg.Dimensions.Height)
	})

	t.Run("when dimensions has invalid width, should not change the default value", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithDimensions(&config.Dimensions{Width: 0, Height: 80}).Build()

		// Assert
		assert.Equal(t, 210.0, cfg.Dimensions.Width)
		assert.Equal(t, 297.0, cfg.Dimensions.Height)
	})

	t.Run("when dimensions has invalid height, should not change the default value", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithDimensions(&config.Dimensions{Width: 80, Height: 0}).Build()

		// Assert
		assert.Equal(t, 210.0, cfg.Dimensions.Width)
		assert.Equal(t, 297.0, cfg.Dimensions.Height)
	})

	t.Run("when dimensions has valid values, should change the default value", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithDimensions(&config.Dimensions{Width: 80, Height: 80}).Build()

		// Assert
		assert.Equal(t, 80.0, cfg.Dimensions.Width)
		assert.Equal(t, 80.0, cfg.Dimensions.Height)
	})

	t.Run("when dimensions are set and page size too, should use dimensions values", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithPageSize(config.A1).WithDimensions(&config.Dimensions{Width: 80, Height: 80}).Build()

		// Assert
		assert.Equal(t, 80.0, cfg.Dimensions.Width)
		assert.Equal(t, 80.0, cfg.Dimensions.Height)
	})
}

func TestBuilder_WithMaxGridSize(t *testing.T) {
	t.Run("when max grid size is invalid, should not change the default value", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithMaxGridSize(-1).Build()

		// Assert
		assert.Equal(t, 12, cfg.MaxGridSize)
	})

	t.Run("when max grid size is valid, should change the default value", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithMaxGridSize(8).Build()

		// Assert
		assert.Equal(t, 8, cfg.MaxGridSize)
	})
}

func TestBuilder_WithMargins(t *testing.T) {
	t.Run("when margins is nil, should not change the default value", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithMargins(nil).Build()

		// Assert
		assert.Equal(t, 10.0, cfg.Margins.Left)
		assert.Equal(t, 10.0, cfg.Margins.Top)
		assert.Equal(t, 10.0, cfg.Margins.Right)
	})

	t.Run("when margins has invalid left, should not change the default value", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithMargins(&config.Margins{Left: 8, Right: 20, Top: 20}).Build()

		// Assert
		assert.Equal(t, 10.0, cfg.Margins.Left)
		assert.Equal(t, 10.0, cfg.Margins.Top)
		assert.Equal(t, 10.0, cfg.Margins.Right)
	})

	t.Run("when margins has invalid right, should not change the default value", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithMargins(&config.Margins{Left: 20, Right: 8, Top: 20}).Build()

		// Assert
		assert.Equal(t, 10.0, cfg.Margins.Left)
		assert.Equal(t, 10.0, cfg.Margins.Top)
		assert.Equal(t, 10.0, cfg.Margins.Right)
	})

	t.Run("when margins has invalid top, should not change the default value", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithMargins(&config.Margins{Left: 20, Right: 20, Top: 8}).Build()

		// Assert
		assert.Equal(t, 10.0, cfg.Margins.Left)
		assert.Equal(t, 10.0, cfg.Margins.Top)
		assert.Equal(t, 10.0, cfg.Margins.Right)
	})

	t.Run("when dimensions has valid values, should change the default value", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithMargins(&config.Margins{Left: 20, Right: 20, Top: 20}).Build()

		// Assert
		assert.Equal(t, 20.0, cfg.Margins.Left)
		assert.Equal(t, 20.0, cfg.Margins.Top)
		assert.Equal(t, 20.0, cfg.Margins.Right)
	})
}
