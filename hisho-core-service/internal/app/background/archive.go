package background

import (
	"context"
	"sync"
	"time"

	"github.com/alexadastra/hisho/hisho-core-service/internal/app/worker"
)

// ArchiveBackroundJob inits archivation process with given ticker duration
type ArchiveBackroundJob struct {
	ticker     *time.Ticker // TODO replace with cron job
	workerPool *worker.ArchiveWorkerPool
	cancel     context.CancelFunc
	wg         *sync.WaitGroup
}

// NewArchiveBackroundJob constructs new ArchiveBackroundJob
func NewArchiveBackroundJob(workerPool *worker.ArchiveWorkerPool, duration time.Duration) *ArchiveBackroundJob {
	return &ArchiveBackroundJob{
		ticker:     time.NewTicker(duration),
		workerPool: workerPool,
		wg:         &sync.WaitGroup{},
	}
}

// Start starts ticker checking in another goroutine
func (bgJob *ArchiveBackroundJob) Start(ctx context.Context) {
	ctx, bgJob.cancel = context.WithCancel(ctx)
	bgJob.wg.Add(1)

	go func() {
		defer bgJob.wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case <-bgJob.ticker.C:
				bgJob.workerPool.Send()
			}
		}
	}()
}

// Stop stops ticker by context cancelling
func (bgJob *ArchiveBackroundJob) Stop() error {
	bgJob.cancel()
	bgJob.wg.Wait()

	return nil
}
