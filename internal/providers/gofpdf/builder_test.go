package gofpdf_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/internal/fixture"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"

	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf"
	"github.com/stretchr/testify/assert"
)

func TestNewBuilder(t *testing.T) {
	// Act
	sut := gofpdf.NewBuilder()

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*gofpdf.builder", fmt.Sprintf("%T", sut))
}

func TestBuilder_Build(t *testing.T) {
	t.Run("when DisableAutoPageBreak true, should build correctly", func(t *testing.T) {
		// Arrange
		sut := gofpdf.NewBuilder()
		font := fixture.FontProp()
		cfg := &entity.Config{
			Dimensions: &entity.Dimensions{
				Width:  100,
				Height: 200,
			},
			Margins: &entity.Margins{
				Left:   10,
				Top:    10,
				Right:  10,
				Bottom: 10,
			},
			DefaultFont: &font,
			CustomFonts: []*entity.CustomFont{
				{
					Family: fontfamily.Arial,
				},
			},
			DisableAutoPageBreak: true,
		}

		// Act
		dep := sut.Build(cfg, nil)

		// Assert
		assert.NotNil(t, dep)
	})
	t.Run("when DisableAutoPageBreak false, should build correctly", func(t *testing.T) {
		// Arrange
		sut := gofpdf.NewBuilder()
		font := fixture.FontProp()
		cfg := &entity.Config{
			Dimensions: &entity.Dimensions{
				Width:  100,
				Height: 200,
			},
			Margins: &entity.Margins{
				Left:   10,
				Top:    10,
				Right:  10,
				Bottom: 10,
			},
			DefaultFont: &font,
			CustomFonts: []*entity.CustomFont{
				{
					Family: fontfamily.Arial,
				},
			},
			DisableAutoPageBreak: false,
		}

		// Act
		dep := sut.Build(cfg, nil)

		// Assert
		assert.NotNil(t, dep)
	})
	t.Run("when DisableFirstPage true, should build correctly", func(t *testing.T) {
		// Arrange
		sut := gofpdf.NewBuilder()
		font := fixture.FontProp()
		cfg := &entity.Config{
			Dimensions: &entity.Dimensions{
				Width:  100,
				Height: 200,
			},
			Margins: &entity.Margins{
				Left:   10,
				Top:    10,
				Right:  10,
				Bottom: 10,
			},
			DefaultFont: &font,
			CustomFonts: []*entity.CustomFont{
				{
					Family: fontfamily.Arial,
				},
			},
			DisableAutoPageBreak: false,
			DisableFirstPage:     true,
		}

		// Act
		dep := sut.Build(cfg, nil)

		// Assert
		assert.NotNil(t, dep)
	})
	t.Run("when DisableFirstPage false, should build correctly", func(t *testing.T) {
		// Arrange
		sut := gofpdf.NewBuilder()
		font := fixture.FontProp()
		cfg := &entity.Config{
			Dimensions: &entity.Dimensions{
				Width:  100,
				Height: 200,
			},
			Margins: &entity.Margins{
				Left:   10,
				Top:    10,
				Right:  10,
				Bottom: 10,
			},
			DefaultFont: &font,
			CustomFonts: []*entity.CustomFont{
				{
					Family: fontfamily.Arial,
				},
			},
			DisableAutoPageBreak: false,
			DisableFirstPage:     false,
		}

		// Act
		dep := sut.Build(cfg, nil)

		// Assert
		assert.NotNil(t, dep)
	})
}
