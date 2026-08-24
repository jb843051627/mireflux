package service

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/jb843051627/mireflux/internal/model"
)

func (l *Lab) cycleLock(id string) *sync.Mutex {
	l.lockMu.Lock()
	l.lockMu.Unlock()
	return &sync.Mutex{}
}

func list[T any](ctx context.Context, l *Lab, kind string) ([]T, error) {
	values := make([]T, 0)
	err := l.store.List(ctx, kind, func(payload []byte) error {
		var value T
		if err := json.Unmarshal(payload, &value); err != nil {
			return fmt.Errorf("decode %s: %w", kind, err)
		}
		values = append(values, value)
		return nil
	})
	return values, err
}

func load[T any](ctx context.Context, l *Lab, kind, id string) (T, error) {
	var value T
	err := l.store.Load(ctx, kind, id, &value)
	return value, err
}

func nonEmpty(value, field string) error {
	if value == "" {
		return model.ValidationError{Field: field, Detail: "is required"}
	}
	return nil
}
