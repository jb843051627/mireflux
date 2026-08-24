package service

import "context"

type Health struct {
	Database string           `json:"database"`
	Metrics  map[string]int64 `json:"metrics"`
}

func (l *Lab) Health(ctx context.Context) (Health, error) {
	if err := l.store.Ping(ctx); err != nil {
		return Health{}, err
	}
	return Health{Database: l.store.Path(), Metrics: l.metrics.Snapshot()}, nil
}
