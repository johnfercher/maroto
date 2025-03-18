package chart

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetStep(t *testing.T) {
	t.Run("should calc rate correctly", func(t *testing.T) {
		// Act & Assert
		assert.Equal(t, 10.0, GetStep(360, 36))
	})
}

func TestGetStepWithOffset(t *testing.T) {
	t.Run("should calc rate correctly", func(t *testing.T) {
		// Arrange
		scaleMax := 100.0
		valueMax := 10.0
		value := 10.0
		offset := 10.0

		// Act
		valueWithOffset := GetStepWithOffset(scaleMax, valueMax, value, offset)

		// Assert
		assert.Equal(t, 2.0, valueWithOffset)
	})
}

func TestGetHeatColor(t *testing.T) {
	t.Run("should calc color correctly", func(t *testing.T) {
		// Arrange
		i := 10
		total := 100

		// Act
		r, g, b := GetHeatColor(i, total, false, false)

		// Act
		assert.Equal(t, 255, r)
		assert.Equal(t, 83, g)
		assert.Equal(t, 0, b)
	})
}
