package time_test

import (
	"github.com/johnfercher/maroto/v2/pkg/time"
	"github.com/stretchr/testify/assert"
	"testing"
	buildtinTime "time"
)

func TestGetTimeSpent(t *testing.T) {
	// Act
	timeSpent := time.GetTimeSpent(func() {
		buildtinTime.Sleep(10 * buildtinTime.Millisecond)
	})

	// Assert
	assert.InDelta(t, float64(10*buildtinTime.Millisecond), timeSpent.Value, float64(1*buildtinTime.Millisecond))
}
