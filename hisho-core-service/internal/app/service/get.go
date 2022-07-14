package service

import (
	"context"
	"time"

	"github.com/alexadastra/hisho/hisho-core-service/internal/app/models"
	"github.com/pkg/errors"
)

// GetTodaysTasks returns tasks that were created today
func (s *Service) GetTodaysTasks(ctx context.Context) ([]*models.Task, error) {
	from, err := time.Parse(timeFormat, time.Now().Format(timeFormat))
	if err != nil {
		return nil, errors.Wrap(err, "failed to format 'from' date")
	}
	to := from.Add(24 * time.Hour)
	tasks, err := s.storage.GetTasksByRange(ctx, &from, &to)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch tasks")
	}
	return tasks, nil
}

// GetTasksByTerm gets tasks by term
func (s *Service) GetTasksByTerm(ctx context.Context, term *models.Term) ([]*models.Task, error) {
	tasks, err := s.storage.GetTasksByTerm(ctx, term)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch tasks")
	}
	return tasks, nil
}
