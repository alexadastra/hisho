package worker

import (
	"context"

	service "github.com/alexadastra/hisho/hisho-core-service/internal/app/service"
	"golang.org/x/sync/errgroup"
)

// ArchiveWorkerPool handles workers that processes tasks archivation
type ArchiveWorkerPool struct {
	g           *errgroup.Group
	cancel      context.CancelFunc
	inCh        chan struct{}
	service     *service.Service
	workerCount int
}

// NewArchiveWorkerPool creates new ArchiveWorkerPool
func NewArchiveWorkerPool(service *service.Service, workerCount int) *ArchiveWorkerPool {
	return &ArchiveWorkerPool{
		workerCount: workerCount,
		service:     service,
		inCh:        make(chan struct{}),
	}
}

// Start inits and starts up processor goroutines with error group
func (w *ArchiveWorkerPool) Start(ctx context.Context) {
	w.g, ctx = errgroup.WithContext(ctx)
	ctx, w.cancel = context.WithCancel(ctx)

	for idx := 0; idx < w.workerCount; idx++ {
		w.g.Go(func() error { return w.loop(ctx) })
	}
}

// Send is an interface for starting up archivation
func (w *ArchiveWorkerPool) Send() {
	w.inCh <- struct{}{}
}

// Stop stops all goroutines gracefully
func (w *ArchiveWorkerPool) Stop() error {
	w.cancel()
	return w.g.Wait()
}

func (w *ArchiveWorkerPool) loop(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			// TODO: log context cancellation
			return nil
		case _, ok := <-w.inCh:
			if !ok {
				// TODO: log channel closure
				return nil
			}
			err := w.service.ArchivateTasks(ctx)
			if err != nil {
				// TODO: log error here
				panic(err)
			}
		}
	}
}
