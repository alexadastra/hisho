package service

import (
	"context"

	"github.com/pkg/errors"

	"github.com/alexadastra/hisho/hisho-core-service/internal/app/models"
)

// AddTask handles task adding
func (s *Service) AddTask(ctx context.Context, task *models.Task) error {
	rowsAffected, err := s.storage.AddTask(ctx, task)
	if err != nil {
		return errors.Wrap(err, "failed to save task")
	}
	if rowsAffected == 0 {
		return errors.New("task wasn't added")
	}
	return nil
}
