package service

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"

	"github.com/alexadastra/hisho/hisho-core-service/internal/app/models"
	"github.com/alexadastra/hisho/hisho-core-service/internal/app/storage"
)

const timeFormat = "2006/01/02"

// Service handles everything
type Service struct {
	storage storage.Storage
}

// NewService constructs new Service
func NewService(storage storage.Storage) *Service {
	return &Service{
		storage: storage,
	}
}

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
func (s *Service) GetTasksByTerm(ctx context.Context, term string) ([]*models.Task, error) {
	termInt, err := matchTerm(term)
	if err != nil {
		return nil, errors.Wrap(err, "failed to match term")
	}

	tasks, err := s.storage.GetTasksByTerm(ctx, termInt)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch tasks")
	}
	return tasks, nil
}

func matchTerm(termStr string) (int64, error) {
	switch termStr {
	case "WEEK":
		return 1, nil
	case "TODAY":
		return 2, nil
	case "OTHER":
		return 0, nil
	default:
		return 0, fmt.Errorf("unexpected term: %s", termStr)
	}
}
