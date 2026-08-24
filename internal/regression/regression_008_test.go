package regression

import (
	"github.com/jb843051627/mireflux/internal/metrics"
	"sync"
	"testing"
)

func TestBug08_ConcurrentMetricsRemainSynchronized(t *testing.T) {
	registry := metrics.New()
	var workers sync.WaitGroup
	workers.Add(2)
	go func() {
		defer workers.Done()
		for i := 0; i < 2000; i++ {
			registry.Add("readings", 1)
		}
	}()
	go func() {
		defer workers.Done()
		for i := 0; i < 2000; i++ {
			_ = registry.Snapshot()
		}
	}()
	workers.Wait()
	if got := registry.Get("readings"); got != 2000 {
		t.Fatalf("metric total = %d, want 2000", got)
	}
}
