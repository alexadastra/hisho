package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/zeebo/xxh3"

	"github.com/alexadastra/hisho/hisho-core-service/internal/app/models"
)

// AddTask handles task adding
func (s *Service) AddTask(ctx context.Context, requests []*models.TaskRequest) (*models.Task, error) {
	now := time.Now().UTC()

	tasks := make([]*models.Task, 0, len(requests))

	for _, request := range requests {
		tasks = append(tasks, &models.Task{
			ID:        int64(xxh3.HashString(uuid.New().String())),
			Title:     request.Title,
			Term:      request.Term,
			IsGreen:   request.IsGreen,
			CreatedAt: &now,
			UpdatedAt: &now,
			ClosedAt:  nil,
		})
	}

	rows, err := s.storage.AddTasks(ctx, tasks)
	if err != nil {
		return nil, errors.Wrap(err, "failed to save task")
	}
	if len(rows) == 0 {
		return nil, errors.New("task wasn't added")
	}
	return rows[0], nil
}
