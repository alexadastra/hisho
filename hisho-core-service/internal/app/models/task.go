package models

import (
	"time"

	"github.com/alexadastra/hisho/hisho-core-service/pkg/api"
	"github.com/pkg/errors"
)

// Task represents task entity
type Task struct {
	ID        int64  `db:"id"`
	Title     string `db:"title"`
	Term      *Term
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DoneAt    *time.Time `db:"created_at"`
}

// NewTask converts proto struct to the internal one
func NewTask(apiTask *api.Task) (*Task, error) {
	term, err := NewTermFromString(apiTask.Term.String())
	if err != nil {
		return nil, errors.Wrap(err, "failed to create new term")
	}
	return &Task{
		ID:        int64(apiTask.Id),
		Title:     apiTask.Title,
		Term:      term,
		CreatedAt: fromTimestampb(apiTask.CreatedAt),
		UpdatedAt: fromTimestampb(apiTask.UpdatedAt),
		DoneAt:    fromTimestampb(apiTask.UpdatedAt),
	}, nil
}

// ToProto converts internal struct from the one from proto-file
func (task *Task) ToProto() *api.Task {
	return &api.Task{
		Id:        uint64(task.ID),
		Title:     task.Title,
		Term:      api.Term(api.Term_value[task.Term.Value]),
		CreatedAt: toTimestampb(task.CreatedAt),
		UpdatedAt: toTimestampb(task.UpdatedAt),
		DoneAt:    toTimestampb(task.DoneAt),
	}
}
