package config_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"

	"github.com/johnfercher/maroto/v2/pkg/color"

	"github.com/johnfercher/maroto/v2/pkg/props"

	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/provider"

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
	assert.Equal(t, 210.0, cfg.Dimensions.Width)
	assert.Equal(t, 297.0, cfg.Dimensions.Height)
	assert.Equal(t, 10.0, cfg.Margins.Top)
	assert.Equal(t, 10.0, cfg.Margins.Left)
	assert.Equal(t, 10.0, cfg.Margins.Right)
	assert.Equal(t, 0, cfg.Workers)
	assert.Equal(t, false, cfg.Debug)
	assert.Equal(t, 12, cfg.MaxGridSize)
	assert.Equal(t, fontfamily.Arial, cfg.Font.Family)
	assert.Equal(t, 10.0, cfg.Font.Size)
	assert.Equal(t, fontstyle.Normal, cfg.Font.Style)
	assert.Equal(t, 0, cfg.Font.Color.Red)
	assert.Equal(t, 0, cfg.Font.Color.Green)
	assert.Equal(t, 0, cfg.Font.Color.Blue)
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
	t.Run("when fontstyle is nil, should not change the default value", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithFont(nil).Build()

		// Assert
		assert.Equal(t, fontfamily.Arial, cfg.Font.Family)
		assert.Equal(t, 10.0, cfg.Font.Size)
		assert.Equal(t, fontstyle.Normal, cfg.Font.Style)
		assert.Equal(t, 0, cfg.Font.Color.Red)
		assert.Equal(t, 0, cfg.Font.Color.Green)
		assert.Equal(t, 0, cfg.Font.Color.Blue)
	})

	t.Run("when family is filled, should change the default value", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithFont(&props.Font{
			Family: "new family",
		}).Build()

		// Assert
		assert.Equal(t, "new family", cfg.Font.Family)
		assert.Equal(t, 10.0, cfg.Font.Size)
		assert.Equal(t, fontstyle.Normal, cfg.Font.Style)
		assert.Equal(t, 0, cfg.Font.Color.Red)
		assert.Equal(t, 0, cfg.Font.Color.Green)
		assert.Equal(t, 0, cfg.Font.Color.Blue)
	})

	t.Run("when style is filled, should change the default value", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithFont(&props.Font{
			Style: fontstyle.Bold,
		}).Build()

		// Assert
		assert.Equal(t, fontfamily.Arial, cfg.Font.Family)
		assert.Equal(t, 10.0, cfg.Font.Size)
		assert.Equal(t, fontstyle.Bold, cfg.Font.Style)
		assert.Equal(t, 0, cfg.Font.Color.Red)
		assert.Equal(t, 0, cfg.Font.Color.Green)
		assert.Equal(t, 0, cfg.Font.Color.Blue)
	})

	t.Run("when size is filled, should change the default value", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithFont(&props.Font{
			Size: 13,
		}).Build()

		// Assert
		assert.Equal(t, fontfamily.Arial, cfg.Font.Family)
		assert.Equal(t, 13.0, cfg.Font.Size)
		assert.Equal(t, fontstyle.Normal, cfg.Font.Style)
		assert.Equal(t, 0, cfg.Font.Color.Red)
		assert.Equal(t, 0, cfg.Font.Color.Green)
		assert.Equal(t, 0, cfg.Font.Color.Blue)
	})

	t.Run("when color is filled, should change the default value", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithFont(&props.Font{
			Color: &color.Color{Red: 10, Green: 10, Blue: 10},
		}).Build()

		// Assert
		assert.Equal(t, fontfamily.Arial, cfg.Font.Family)
		assert.Equal(t, 10.0, cfg.Font.Size)
		assert.Equal(t, fontstyle.Normal, cfg.Font.Style)
		assert.Equal(t, 10, cfg.Font.Color.Red)
		assert.Equal(t, 10, cfg.Font.Color.Green)
		assert.Equal(t, 10, cfg.Font.Color.Blue)
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
		cfg := sut.WithPageSize(pagesize.A2).Build()

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
		cfg := sut.WithPageSize(pagesize.A1).WithDimensions(&config.Dimensions{Width: 80, Height: 80}).Build()

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

func TestBuilder_AddUTF8Font(t *testing.T) {
	t.Run("when fontstyle family is empty, should not add value", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.AddUTF8Font("", fontstyle.Bold, "file").Build()

		// Assert
		assert.Equal(t, 0, len(cfg.CustomFonts))
	})

	t.Run("when fontstyle style is invalid, should not add value", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.AddUTF8Font("family", "invalid", "file").Build()

		// Assert
		assert.Equal(t, 0, len(cfg.CustomFonts))
	})

	t.Run("when fontstyle file is empty, should not add value", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.AddUTF8Font("family", fontstyle.Bold, "").Build()

		// Assert
		assert.Equal(t, 0, len(cfg.CustomFonts))
	})

	t.Run("when fontstyle is valid, should not value", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.AddUTF8Font("family", fontstyle.Bold, "file").Build()

		// Assert
		assert.Equal(t, 1, len(cfg.CustomFonts))
		assert.Equal(t, "family", cfg.CustomFonts[0].Family)
		assert.Equal(t, fontstyle.Bold, cfg.CustomFonts[0].Style)
		assert.Equal(t, "file", cfg.CustomFonts[0].File)
	})
}
