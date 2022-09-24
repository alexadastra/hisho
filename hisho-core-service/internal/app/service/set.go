package service

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"github.com/alexadastra/hisho/hisho-core-service/internal/app/models"
)

// ChangeTaskStatus changes task status
func (s *Service) ChangeTaskStatus(ctx context.Context, id uint64, event models.ChangeStatusEvent, reason string) (*models.Task, error) {
	task, err := s.storage.GetTaskByID(ctx, int64(id))
	if err != nil {
		return nil, errors.Wrap(err, "failed to find task by id")
	}
	if task == nil {
		return nil, errors.New("task by id not found")
	}

	now := time.Now().UTC()

	var closedAt *time.Time
	closedReason := ""

	switch event {
	case models.LeaveUnchangedEvent:
		return task, nil
	case models.SetDoneEvent:
		closedAt = &now
		closedReason = ""
	case models.CloseEvent:
		closedAt = &now
		closedReason = reason
	case models.ReopenEvent:
		closedAt = nil
		closedReason = ""
	}

	updatedTask, err := s.storage.UpdateTaskStatus(ctx, int64(id), closedAt, closedReason)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update task status")
	}
	return updatedTask, nil
}

// EditTask changes task changable fields, if they are non-empty
func (s *Service) EditTask(ctx context.Context, newTask *models.Task) (*models.Task, error) {
	task, err := s.storage.GetTaskByID(ctx, newTask.ID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find task by id")
	}
	if task == nil {
		return nil, errors.New("task by id not found")
	}

	taskToUpdate := task
	hasChanges := true
	if newTask.Title != task.Title {
		taskToUpdate.Title = newTask.Title
		hasChanges = true
	}
	if newTask.IsGreen != task.IsGreen {
		taskToUpdate.IsGreen = newTask.IsGreen
		hasChanges = true
	}
	if newTask.Term != task.Term {
		taskToUpdate.Term = newTask.Term
		hasChanges = true
	}

	if !hasChanges {
		return taskToUpdate, nil
	}

	updatedTask, err := s.storage.UpdateTask(ctx, taskToUpdate)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update task status")
	}
	return updatedTask, nil
}
