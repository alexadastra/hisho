package models

import "time"

// Task represents task entity
type Task struct {
	ID        int64      `db:"id"`
	Title     string     `db:"title"`
	Term      int64      `db:"term"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DoneAt    *time.Time `db:"created_at"`
}
