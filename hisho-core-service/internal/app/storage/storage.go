package storage

import (
	"context"
	"time"

	"github.com/alexadastra/hisho/hisho-core-service/internal/app/models"
)

// Storage stores tasks
type Storage interface {
	GetTasksByRange(ctx context.Context, from *time.Time, to *time.Time) ([]*models.Task, error)
	GetTasksByTerm(ctx context.Context, term int64) ([]*models.Task, error)
	SetTaskAsDone(ctx context.Context, id int64) (*models.Task, error)
}
