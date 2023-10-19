package gofpdf_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf"

	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	// Act
	sut := gofpdf.New(&entity.Config{
		Dimensions:  &entity.Dimensions{},
		Margins:     &entity.Margins{},
		DefaultFont: &props.Font{},
	}, nil)

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*gofpdf.provider", fmt.Sprintf("%T", sut))
}
