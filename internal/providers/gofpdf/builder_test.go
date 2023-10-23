package gofpdf_test

import (
	"fmt"
	"github.com/johnfercher/maroto/v2/internal/fixture"
	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBuilder(t *testing.T) {
	// Act
	sut := gofpdf.NewBuilder()

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*gofpdf.builder", fmt.Sprintf("%T", sut))
}

func TestBuilder_Build(t *testing.T) {
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
	}

	// Act
	dep := sut.Build(cfg, nil)

	// Assert
	assert.NotNil(t, dep)
}
