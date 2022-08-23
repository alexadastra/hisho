package storage

import (
	"context"
	"time"

	"github.com/alexadastra/hisho/hisho-core-service/internal/app/models"
)

// Storage stores tasks
type Storage interface {
	AddTasks(ctx context.Context, tasks []*models.Task) ([]*models.Task, error)
	GetTasksByRange(ctx context.Context, from *time.Time, to *time.Time) ([]*models.Task, error)
	GetTasksByTerm(ctx context.Context, term *models.Term) ([]*models.Task, error)
	SetTaskAsDone(ctx context.Context, id int64) (*models.Task, error)
}
