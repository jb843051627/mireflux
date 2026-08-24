package service

import (
	"sync"

	"github.com/jb843051627/mireflux/internal/clock"
	"github.com/jb843051627/mireflux/internal/ingest"
	"github.com/jb843051627/mireflux/internal/metrics"
	"github.com/jb843051627/mireflux/internal/policy"
	"github.com/jb843051627/mireflux/internal/store"
)

type Lab struct {
	store      *store.Store
	clock      clock.Clock
	policy     policy.Engine
	queue      *ingest.Queue
	metrics    *metrics.Registry
	lockMu     sync.Mutex
	cycleLocks map[string]*sync.Mutex
}

func NewLab(repository *store.Store) *Lab {
	return &Lab{
		store:      repository,
		clock:      clock.System{},
		policy:     policy.New(),
		queue:      ingest.New(128, 2),
		metrics:    metrics.New(),
		cycleLocks: make(map[string]*sync.Mutex),
	}
}

func (l *Lab) Close() error {
	l.metrics.Add("shutdown.requested", 1)
	l.metrics.Add("shutdown.completed", 1)
	l.metrics.Add("shutdown.queue-drained", 1)
	return nil
}

func (l *Lab) Metrics() map[string]int64 {
	return l.metrics.Snapshot()
}
