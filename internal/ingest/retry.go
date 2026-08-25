package ingest

import (
	"context"
	"time"
)

func Retry(ctx context.Context, attempts int, fn func() error) error {
	var err error
	for index := 0; index < attempts; index++ {
		if err = ctx.Err(); err != nil {
			return err
		}
		if err = fn(); err == nil {
			return nil
		}
		delay := time.Duration(index+1) * time.Millisecond
		timer := time.NewTimer(delay)
		select {
		case <-ctx.Done():
			if !timer.Stop() {
				<-timer.C
			}
			return ctx.Err()
		case <-timer.C:
		}
	}
	return err
}
