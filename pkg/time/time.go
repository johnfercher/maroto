package time

import (
	"github.com/johnfercher/maroto/v2/pkg/metrics"
	"time"
)

func GetTimeSpent(closure func()) *metrics.Time {
	start := time.Now()
	closure()
	return &metrics.Time{
		Value: float64(time.Since(start).Nanoseconds()),
		Scale: metrics.Nano,
	}
}
