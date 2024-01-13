package config_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/consts/provider"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func TestNewBuilder(t *testing.T) {
	// Act
	sut := config.NewBuilder()

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*config.CfgBuilder", fmt.Sprintf("%T", sut))
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
	assert.Equal(t, 0, cfg.WorkersQuantity)
	assert.Equal(t, false, cfg.Debug)
	assert.Equal(t, 12, cfg.MaxGridSize)
	assert.Equal(t, fontfamily.Arial, cfg.DefaultFont.Family)
	assert.Equal(t, 10.0, cfg.DefaultFont.Size)
	assert.Equal(t, fontstyle.Normal, cfg.DefaultFont.Style)
	assert.Equal(t, 0, cfg.DefaultFont.Color.Red)
	assert.Equal(t, 0, cfg.DefaultFont.Color.Green)
	assert.Equal(t, 0, cfg.DefaultFont.Color.Blue)
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
		cfg := sut.WithDefaultFont(nil).Build()

		// Assert
		assert.Equal(t, fontfamily.Arial, cfg.DefaultFont.Family)
		assert.Equal(t, 10.0, cfg.DefaultFont.Size)
		assert.Equal(t, fontstyle.Normal, cfg.DefaultFont.Style)
		assert.Equal(t, 0, cfg.DefaultFont.Color.Red)
		assert.Equal(t, 0, cfg.DefaultFont.Color.Green)
		assert.Equal(t, 0, cfg.DefaultFont.Color.Blue)
	})

	t.Run("when family is filled, should change the default value", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithDefaultFont(&props.Font{
			Family: "new family",
		}).Build()

		// Assert
		assert.Equal(t, "new family", cfg.DefaultFont.Family)
		assert.Equal(t, 10.0, cfg.DefaultFont.Size)
		assert.Equal(t, fontstyle.Normal, cfg.DefaultFont.Style)
		assert.Equal(t, 0, cfg.DefaultFont.Color.Red)
		assert.Equal(t, 0, cfg.DefaultFont.Color.Green)
		assert.Equal(t, 0, cfg.DefaultFont.Color.Blue)
	})

	t.Run("when style is filled, should change the default value", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithDefaultFont(&props.Font{
			Style: fontstyle.Bold,
		}).Build()

		// Assert
		assert.Equal(t, fontfamily.Arial, cfg.DefaultFont.Family)
		assert.Equal(t, 10.0, cfg.DefaultFont.Size)
		assert.Equal(t, fontstyle.Bold, cfg.DefaultFont.Style)
		assert.Equal(t, 0, cfg.DefaultFont.Color.Red)
		assert.Equal(t, 0, cfg.DefaultFont.Color.Green)
		assert.Equal(t, 0, cfg.DefaultFont.Color.Blue)
	})

	t.Run("when size is filled, should change the default value", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithDefaultFont(&props.Font{
			Size: 13,
		}).Build()

		// Assert
		assert.Equal(t, fontfamily.Arial, cfg.DefaultFont.Family)
		assert.Equal(t, 13.0, cfg.DefaultFont.Size)
		assert.Equal(t, fontstyle.Normal, cfg.DefaultFont.Style)
		assert.Equal(t, 0, cfg.DefaultFont.Color.Red)
		assert.Equal(t, 0, cfg.DefaultFont.Color.Green)
		assert.Equal(t, 0, cfg.DefaultFont.Color.Blue)
	})

	t.Run("when color is filled, should change the default value", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithDefaultFont(&props.Font{
			Color: &props.Color{Red: 10, Green: 10, Blue: 10},
		}).Build()

		// Assert
		assert.Equal(t, fontfamily.Arial, cfg.DefaultFont.Family)
		assert.Equal(t, 10.0, cfg.DefaultFont.Size)
		assert.Equal(t, fontstyle.Normal, cfg.DefaultFont.Style)
		assert.Equal(t, 10, cfg.DefaultFont.Color.Red)
		assert.Equal(t, 10, cfg.DefaultFont.Color.Green)
		assert.Equal(t, 10, cfg.DefaultFont.Color.Blue)
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

func TestBuilder_WithWorkerPoolSize(t *testing.T) {
	t.Run("when worker pool size is invalid, should not change the default value", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithWorkerPoolSize(-1).Build()

		// Assert
		assert.Equal(t, 0, cfg.WorkersQuantity)
	})

	t.Run("when worker pool size is valid, should change the default value", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithWorkerPoolSize(7).Build()

		// Assert
		assert.Equal(t, 7, cfg.WorkersQuantity)
	})
}

func TestBuilder_WithDimensions(t *testing.T) {
	t.Run("when dimensions has invalid width, should not change the default value", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithDimensions(0, 80).Build()

		// Assert
		assert.Equal(t, 210.0, cfg.Dimensions.Width)
		assert.Equal(t, 297.0, cfg.Dimensions.Height)
	})

	t.Run("when dimensions has invalid height, should not change the default value", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithDimensions(80, 0).Build()

		// Assert
		assert.Equal(t, 210.0, cfg.Dimensions.Width)
		assert.Equal(t, 297.0, cfg.Dimensions.Height)
	})

	t.Run("when dimensions has valid values, should change the default value", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithDimensions(80, 80).Build()

		// Assert
		assert.Equal(t, 80.0, cfg.Dimensions.Width)
		assert.Equal(t, 80.0, cfg.Dimensions.Height)
	})

	t.Run("when dimensions are set and page size too, should use dimensions values", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithPageSize(pagesize.A1).WithDimensions(80, 80).Build()

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
	t.Run("when margins has invalid left, should not change the default value", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithMargins(-1, 20, 20).Build()

		// Assert
		assert.Equal(t, 10.0, cfg.Margins.Left)
		assert.Equal(t, 10.0, cfg.Margins.Top)
		assert.Equal(t, 10.0, cfg.Margins.Right)
	})

	t.Run("when margins has invalid right, should not change the default value", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithMargins(20, 20, -1).Build()

		// Assert
		assert.Equal(t, 10.0, cfg.Margins.Left)
		assert.Equal(t, 10.0, cfg.Margins.Top)
		assert.Equal(t, 10.0, cfg.Margins.Right)
	})

	t.Run("when margins has invalid top, should not change the default value", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithMargins(20, -1, 20).Build()

		// Assert
		assert.Equal(t, 10.0, cfg.Margins.Left)
		assert.Equal(t, 10.0, cfg.Margins.Top)
		assert.Equal(t, 10.0, cfg.Margins.Right)
	})

	t.Run("when dimensions has valid values, should change the default value", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithMargins(20, 20, 20).Build()

		// Assert
		assert.Equal(t, 20.0, cfg.Margins.Left)
		assert.Equal(t, 20.0, cfg.Margins.Top)
		assert.Equal(t, 20.0, cfg.Margins.Right)
	})
}

func TestBuilder_WithOrientation(t *testing.T) {
	t.Run("when using default page size and orientation is not set, should use vertical", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.Build()

		// Assert
		assert.Equal(t, 210.0, cfg.Dimensions.Width)
		assert.Equal(t, 297.0, cfg.Dimensions.Height)
		assert.True(t, cfg.Dimensions.Height > cfg.Dimensions.Width)
	})
	t.Run("when using default page size and orientation is set to horizontal, should use horizontal", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithOrientation(orientation.Horizontal).Build()

		// Assert
		assert.Equal(t, 297.0, cfg.Dimensions.Width)
		assert.Equal(t, 210.0, cfg.Dimensions.Height)
		assert.True(t, cfg.Dimensions.Width > cfg.Dimensions.Height)
	})
	t.Run("when using default page size and orientation is not set, should use vertical", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithPageSize(pagesize.A5).Build()

		// Assert
		assert.Equal(t, 148.4, cfg.Dimensions.Width)
		assert.Equal(t, 210.0, cfg.Dimensions.Height)
		assert.True(t, cfg.Dimensions.Height > cfg.Dimensions.Width)
	})
	t.Run("when using default page size and orientation is set to horizontal, should use horizontal", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithPageSize(pagesize.A5).WithOrientation(orientation.Horizontal).Build()

		// Assert
		assert.Equal(t, 210.0, cfg.Dimensions.Width)
		assert.Equal(t, 148.4, cfg.Dimensions.Height)
		assert.True(t, cfg.Dimensions.Width > cfg.Dimensions.Height)
	})
}

func TestBuilder_WithAuthor(t *testing.T) {
	t.Run("when author is empty, should ignore", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithAuthor("", true).Build()

		// Assert
		assert.Nil(t, cfg.Metadata.Author)
	})
	t.Run("when author valid, should apply", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithAuthor("author", true).Build()

		// Assert
		assert.Equal(t, "author", cfg.Metadata.Author.Text)
		assert.Equal(t, true, cfg.Metadata.Author.UTF8)
	})
}

func TestBuilder_WithCreator(t *testing.T) {
	t.Run("when creator is empty, should ignore", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithCreator("", true).Build()

		// Assert
		assert.Nil(t, cfg.Metadata.Creator)
	})
	t.Run("when creator valid, should apply", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithCreator("creator", true).Build()

		// Assert
		assert.Equal(t, "creator", cfg.Metadata.Creator.Text)
		assert.Equal(t, true, cfg.Metadata.Creator.UTF8)
	})
}

func TestBuilder_WithSubject(t *testing.T) {
	t.Run("when subject is empty, should ignore", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithSubject("", true).Build()

		// Assert
		assert.Nil(t, cfg.Metadata.Subject)
	})
	t.Run("when subject valid, should apply", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithSubject("subject", true).Build()

		// Assert
		assert.Equal(t, "subject", cfg.Metadata.Subject.Text)
		assert.Equal(t, true, cfg.Metadata.Subject.UTF8)
	})
}

func TestBuilder_WithTitle(t *testing.T) {
	t.Run("when title is empty, should ignore", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithTitle("", true).Build()

		// Assert
		assert.Nil(t, cfg.Metadata.Title)
	})
	t.Run("when title valid, should apply", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithTitle("title", true).Build()

		// Assert
		assert.Equal(t, "title", cfg.Metadata.Title.Text)
		assert.Equal(t, true, cfg.Metadata.Title.UTF8)
	})
}

func TestBuilder_WithCreationDate(t *testing.T) {
	t.Run("when time is zero, should ignore", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()

		// Act
		cfg := sut.WithCreationDate(time.Time{}).Build()

		// Assert
		assert.Nil(t, cfg.Metadata.CreationDate)
	})
	t.Run("when time valid, should apply", func(t *testing.T) {
		// Arrange
		sut := config.NewBuilder()
		timeNow := time.Now()

		// Act
		cfg := sut.WithCreationDate(timeNow).Build()

		// Assert
		assert.Equal(t, &timeNow, cfg.Metadata.CreationDate)
	})
}
