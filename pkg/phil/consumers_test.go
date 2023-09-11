package phil

import (
	"testing"
	"time"
)

func TestObserver(t *testing.T) {
	t.Run("Test Observer", func(t *testing.T) {
		s := &subject{}

		observer := NewComplianceRequestedObserver(s)
		defer observer.Close()

		DoService(s)

		time.Sleep(1 * time.Second)
	})
}
