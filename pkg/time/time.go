package time

import (
	"time"

	"github.com/johnfercher/maroto/v2/pkg/metrics"
)

func GetTimeSpent(closure func()) *metrics.Time {
	start := time.Now()
	closure()
	return &metrics.Time{
		Value: float64(time.Since(start).Nanoseconds()),
		Scale: metrics.Nano,
	}
}
