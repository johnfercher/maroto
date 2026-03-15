package entity_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/core/entity"

	"github.com/stretchr/testify/assert"

	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
)

func TestImage_AppendMap(t *testing.T) {
	t.Parallel()
	// Arrange
	sut := fixtureImage()
	m := make(map[string]any)

	// Act
	m = sut.AppendMap(m)

	// Assert
	assert.Equal(t, "[1 2 3]", m["entity_image_bytes"])
	assert.Equal(t, extension.Png, m["entity_extension"])
	assert.Equal(t, 100.0, m["background_dimension_width"])
	assert.Equal(t, 200.0, m["background_dimension_height"])
}

func fixtureImage() entity.Image {
	dimensions := fixtureDimensions()
	return entity.Image{
		Bytes:      []byte{1, 2, 3},
		Extension:  extension.Png,
		Dimensions: &dimensions,
	}
}
