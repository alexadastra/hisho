package service

import (
	"context"

	"github.com/alexadastra/hisho/hisho-core-service/internal/app/models"
	"github.com/pkg/errors"
)

// GetTasksByTerm gets tasks by term
func (s *Service) GetTasksByTerm(ctx context.Context, term *models.Term, paging *models.Paging) ([]*models.Task, error) {
	if paging.IsEnabled {
		limit, offset := paging.Page.ToLimitOffset()
		tasks, err := s.storage.GetTasksByTermAndLimits(ctx, term, uint64(limit), uint64(offset))
		if err != nil {
			return nil, errors.Wrap(err, "failed to fetch tasks")
		}
		return tasks, nil
	}
	tasks, err := s.storage.GetTasksByTerm(ctx, term)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch tasks")
	}
	return tasks, nil
}
