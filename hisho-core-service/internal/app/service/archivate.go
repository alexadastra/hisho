package service

import (
	"context"

	"github.com/pkg/errors"
)

// ArchivateTasks gathers tasks that need to be archived and archive them
func (s *Service) ArchivateTasks(ctx context.Context) error {
	tasks, err := s.storage.GetClosedTasks(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to get closed tasks")
	}

	err = s.storage.ArchivateTasks(ctx, tasks)
	if err != nil {
		return errors.Wrap(err, "failed to archivate closed tasks")
	}

	return nil
}
