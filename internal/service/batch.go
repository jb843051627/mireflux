package service

import (
	"context"

	"github.com/jb843051627/mireflux/internal/ingest"
)

func runBatch(ctx context.Context, lab *Lab, jobs []func(context.Context) error) error {
	return ingest.RunBatch(ctx, lab.queue, jobs)
}
