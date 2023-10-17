package metrics_test

import (
	"github.com/johnfercher/maroto/v2/pkg/metrics"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTime_Normalize(t *testing.T) {
	t.Run("when scale is nano, should divide by 1000 and change to micro", func(t *testing.T) {
		// Arrange
		time := metrics.Time{
			Value: 3000,
			Scale: metrics.Nano,
		}

		// Act
		ok := time.Normalize()

		// Assert
		assert.True(t, ok)
		assert.Equal(t, 3.0, time.Value)
		assert.Equal(t, metrics.Micro, time.Scale)
	})
	t.Run("when scale is nano, should divide by 1000 and change to micro", func(t *testing.T) {
		// Arrange
		time := metrics.Time{
			Value: 2000,
			Scale: metrics.Micro,
		}

		// Act
		ok := time.Normalize()

		// Assert
		assert.True(t, ok)
		assert.Equal(t, 2.0, time.Value)
		assert.Equal(t, metrics.Milli, time.Scale)
	})
	t.Run("when scale is milli, should return false", func(t *testing.T) {
		// Arrange
		time := metrics.Time{
			Value: 2000,
			Scale: metrics.Milli,
		}

		// Act
		ok := time.Normalize()

		// Assert
		assert.False(t, ok)
		assert.Equal(t, 2000.0, time.Value)
		assert.Equal(t, metrics.Milli, time.Scale)
	})
}

func TestTime_String(t *testing.T) {
	// Arrange
	time := metrics.Time{
		Value: 2000,
		Scale: metrics.Milli,
	}

	// Act
	s := time.String()

	// Assert
	assert.Equal(t, "2000.00ms", s)
}
