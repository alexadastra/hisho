package models

import (
	"time"

	"github.com/alexadastra/hisho/hisho-core-service/pkg/api"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/zeebo/xxh3"
)

// Task represents task entity
type Task struct {
	ID           int64      `db:"id"`
	Title        string     `db:"title"`
	Term         Term       `db:"term"`
	IsGreen      bool       `db:"is_greeen"`
	CreatedAt    time.Time  `db:"created_at"`
	UpdatedAt    time.Time  `db:"updated_at"`
	ClosedAt     *time.Time `db:"created_at"`
	ClosedReason string     `db:"closed_reason"`
	IsArchived   bool       `db:"is_archived"`
}

// NewTask converts proto struct to the internal one
func NewTask(apiTask *api.Task) (*Task, error) {
	if apiTask.Title == "" {
		return nil, errors.New("empty title")
	}
	term, err := NewTermFromString(apiTask.Term.String())
	if err != nil {
		return nil, errors.Wrap(err, "failed to create new term")
	}
	return &Task{
		ID:           int64(apiTask.Id),
		Title:        apiTask.Title,
		Term:         *term,
		IsGreen:      apiTask.IsGreen,
		CreatedAt:    *fromTimestampb(apiTask.CreatedAt),
		UpdatedAt:    *fromTimestampb(apiTask.UpdatedAt),
		ClosedAt:     fromTimestampb(apiTask.ClosedAt),
		ClosedReason: apiTask.ClosedReason,
		IsArchived:   false,
	}, nil
}

// NewTaskFromRequest creates new Task from the request
func NewTaskFromRequest(request *TaskRequest) *Task {
	now := time.Now().UTC()
	return &Task{
		ID:        int64(xxh3.HashString(uuid.New().String())),
		Title:     request.Title,
		Term:      request.Term,
		IsGreen:   request.IsGreen,
		CreatedAt: now,
		UpdatedAt: now,
		ClosedAt:  nil,
	}
}

// ToProto converts internal struct from the one from proto-file
func (task *Task) ToProto() *api.Task {
	return &api.Task{
		Id:           uint64(task.ID),
		Title:        task.Title,
		Term:         api.Term(api.Term_value[task.Term.Value]),
		IsGreen:      task.IsGreen,
		CreatedAt:    toTimestampb(&task.CreatedAt),
		UpdatedAt:    toTimestampb(&task.UpdatedAt),
		ClosedAt:     toTimestampb(task.ClosedAt),
		ClosedReason: task.ClosedReason,
		Status:       newStatus(task.ClosedAt, task.ClosedReason),
	}
}

// TaskRequest represents request for creating the task
type TaskRequest struct {
	Title   string `db:"title"`
	Term    Term   `db:"term"`
	IsGreen bool   `db:"is_greeen"`
}

// NewTaskRequest constructs new TaskRequest
func NewTaskRequest(request *api.AddTaskRequest) (*TaskRequest, error) {
	if request.Title == "" {
		return nil, errors.New("empty title")
	}

	term, err := NewTermFromString(request.Term.String())
	if err != nil {
		return nil, errors.Wrap(err, "failed to create new term")
	}

	return &TaskRequest{
		Title:   request.Title,
		Term:    *term,
		IsGreen: request.IsGreen,
	}, nil
}
