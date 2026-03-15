package entity_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/core/entity"

	"github.com/stretchr/testify/assert"
)

func TestDimensions_AppendMap(t *testing.T) {
	t.Parallel()
	// Arrange
	sut := fixtureDimensions()
	m := make(map[string]any)

	// Act
	m = sut.AppendMap("label", m)

	// Assert
	assert.Equal(t, 100.0, m["label_dimension_width"])
	assert.Equal(t, 200.0, m["label_dimension_height"])
}

func fixtureDimensions() entity.Dimensions {
	return entity.Dimensions{
		Width:  100,
		Height: 200,
	}
}
