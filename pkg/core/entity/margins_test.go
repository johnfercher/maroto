package entity_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/core/entity"

	"github.com/stretchr/testify/assert"
)

func TestMargins_AppendMap(t *testing.T) {
	t.Parallel()
	// Arrange
	sut := fixtureMargins()
	m := make(map[string]any)

	// Act
	m = sut.AppendMap(m)

	// Assert
	assert.Equal(t, 20.0, m["config_margin_left"])
	assert.Equal(t, 30.0, m["config_margin_top"])
	assert.Equal(t, 40.0, m["config_margin_right"])
	assert.Equal(t, 50.0, m["config_margin_bottom"])
}

func fixtureMargins() entity.Margins {
	return entity.Margins{
		Left:   20,
		Top:    30,
		Right:  40,
		Bottom: 50,
	}
}
