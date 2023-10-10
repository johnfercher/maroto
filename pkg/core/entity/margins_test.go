package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMargins_AppendMap(t *testing.T) {
	// Arrange
	sut := fixtureMargins()
	m := make(map[string]interface{})

	// Act
	m = sut.AppendMap(m)

	// Assert
	assert.Equal(t, 20.0, m["config_margin_left"])
	assert.Equal(t, 30.0, m["config_margin_top"])
	assert.Equal(t, 40.0, m["config_margin_right"])
	assert.Equal(t, 50.0, m["config_margin_bottom"])
}

func fixtureMargins() Margins {
	return Margins{
		Left:   20,
		Top:    30,
		Right:  40,
		Bottom: 50,
	}
}
