package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDimensions_AppendMap(t *testing.T) {
	// Arrange
	sut := fixtureDimensions()
	m := make(map[string]interface{})

	// Act
	m = sut.AppendMap("label", m)

	// Assert
	assert.Equal(t, 100.0, m["label_dimension_width"])
	assert.Equal(t, 200.0, m["label_dimension_height"])
}

func fixtureDimensions() Dimensions {
	return Dimensions{
		Width:  100,
		Height: 200,
	}
}
