package models

import (
	"time"

	"github.com/alexadastra/hisho/hisho-core-service/pkg/api"
)

// Task represents task entity
type Task struct {
	ID        int64      `db:"id"`
	Title     string     `db:"title"`
	Term      int64      `db:"term"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DoneAt    *time.Time `db:"created_at"`
}

// FromProto converts proto struct to the internal one
func FromProto(apiTask *api.Task) *Task {
	return &Task{
		ID:        int64(apiTask.Id),
		Title:     apiTask.Title,
		Term:      int64(api.Term_value[apiTask.Term]),
		CreatedAt: fromTimestampb(apiTask.CreatedAt),
		UpdatedAt: fromTimestampb(apiTask.UpdatedAt),
		DoneAt:    fromTimestampb(apiTask.UpdatedAt),
	}
}

// ToProto converts internal struct from the one from proto-file
func ToProto(task *Task) *api.Task {
	return &api.Task{
		Id:        uint64(task.ID),
		Title:     task.Title,
		Term:      api.Term_name[int32(task.Term)],
		CreatedAt: toTimestampb(task.CreatedAt),
		UpdatedAt: toTimestampb(task.UpdatedAt),
		DoneAt:    toTimestampb(task.DoneAt),
	}
}
