package syncing

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run(
		"incrementing counter 3 times returns 3",
		func(t *testing.T) {
			want := 3

			c := NewCounter()

			c.Inc()
			c.Inc()
			c.Inc()

			assertCounter(t, c, want)
		},
	)

	t.Run(
		"concurrency safe",
		func(t *testing.T) {
			wantedCount := 1000
			c := NewCounter()

			var wg sync.WaitGroup
			wg.Add(wantedCount)

			for i := 0; i < wantedCount; i++ {
				go func() {
					c.Inc()
					wg.Done()
				}()
			}
			wg.Wait()

			assertCounter(t, c, wantedCount)
		},
	)
}

func assertCounter(t *testing.T, c *Counter, want int) {
	t.Helper()

	if c.Value() != want {
		t.Errorf("got '%d', want '%d'", c.Value(), want)
	}
}
