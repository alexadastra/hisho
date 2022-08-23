package service

import (
	"context"

	"github.com/pkg/errors"

	"github.com/alexadastra/hisho/hisho-core-service/internal/app/models"
)

// AddTask handles task adding
func (s *Service) AddTask(ctx context.Context, task *models.Task) (*models.Task, error) {
	rows, err := s.storage.AddTasks(ctx, []*models.Task{task})
	if err != nil {
		return nil, errors.Wrap(err, "failed to save task")
	}
	if len(rows) == 0 {
		return nil, errors.New("task wasn't added")
	}
	return rows[0], nil
}
