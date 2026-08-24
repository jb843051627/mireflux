package ingest

import (
	"context"
	"fmt"
)

func RunBatch(ctx context.Context, queue *Queue, jobs []func(context.Context) error) error {
	done := make(chan error, len(jobs))
	for _, work := range jobs {
		if err := queue.Submit(ctx, Job{Run: work, Done: done}); err != nil {
			return err
		}
	}
	for range jobs {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-done:
			if err != nil {
				return fmt.Errorf("batch ingestion: %w", err)
			}
		}
	}
	return nil
}
