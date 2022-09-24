package service

import (
	"context"
	"time"

	"github.com/alexadastra/hisho/hisho-core-service/internal/app/models"
)

// Storage stores tasks
type Storage interface {
	AddTasks(ctx context.Context, tasks []*models.Task) ([]*models.Task, error)
	GetTaskByID(ctx context.Context, id int64) (*models.Task, error)
	GetTasksByTerm(ctx context.Context, term *models.Term) ([]*models.Task, error)
	GetTasksByTermAndLimits(ctx context.Context, term *models.Term, limit uint64, offset uint64) ([]*models.Task, error)
	UpdateTask(ctx context.Context, task *models.Task) (*models.Task, error)
	UpdateTaskStatus(ctx context.Context, id int64, closedAt *time.Time, reason string) (*models.Task, error)
}

// Service handles everything
type Service struct {
	storage Storage
}

// NewService constructs new Service
func NewService(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}
