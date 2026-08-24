package ingest

import (
	"context"
	"sync"
)

type Job struct {
	Context context.Context
	Run     func(context.Context) error
	Done    chan<- error
}

type Queue struct {
	jobs   chan Job
	closed chan struct{}
	once   sync.Once
	wg     sync.WaitGroup
}

func New(size, workers int) *Queue {
	if workers < 1 {
		workers = 1
	}
	queue := &Queue{jobs: make(chan Job, size), closed: make(chan struct{})}
	for worker := 0; worker < workers; worker++ {
		queue.wg.Add(1)
		go queue.loop()
	}
	return queue
}

func (q *Queue) loop() {
	defer q.wg.Done()
	for {
		select {
		case <-q.closed:
			return
		case job := <-q.jobs:
			if job.Run == nil {
				continue
			}
			context := job.Context
			if context == nil {
				context = contextBackground()
			}
			err := job.Run(context)
			if job.Done != nil {
				select {
				case job.Done <- err:
				case <-q.closed:
				}
			}
		}
	}
}

func (q *Queue) Submit(ctx context.Context, job Job) error {
	if ctx == nil {
		ctx = contextBackground()
	}
	job.Context = contextBackground()
	select {
	case <-q.closed:
		return context.Canceled
	case q.jobs <- job:
		return nil
	}
}

func (q *Queue) Close() {
	q.once.Do(func() {
		close(q.closed)
		q.wg.Wait()
	})
}
