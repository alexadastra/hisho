package pgsql

import (
	"context"

	"github.com/alexadastra/hisho/hisho-core-service/internal/app/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL database driver
	"github.com/pkg/errors"
)

const (
	tasksTableName         = "tasks"
	idColumnName           = "id"
	titleColumnName        = "title"
	termColumnName         = "term"
	isGreenColumnName      = "is_green"
	createdAtColumnName    = "created_at"
	updatedAtColumnName    = "updated_at"
	closedAtColumnName     = "closed_at"
	closedReasonColumnName = "closed_reason"
	isArchivedColumnName   = "is_archived"

	onConflictDoNothing = "on conflict do nothing"
	returningAll        = "returning *"
)

var allColumns = []string{
	idColumnName, titleColumnName, termColumnName, isGreenColumnName, createdAtColumnName,
	updatedAtColumnName, closedAtColumnName, closedReasonColumnName, isArchivedColumnName,
}

// PGStorage implements Storage with Postgres database
type PGStorage struct {
	conn *sqlx.DB
}

// NewStorage constructs new PGStorage
func NewStorage(ctx context.Context, host string) (*PGStorage, error) {
	conn, err := sqlx.Connect("postgres", host)
	if err != nil {
		return nil, errors.Wrap(err, "unable to connect to database")
	}

	return &PGStorage{
		conn: conn,
	}, nil
}

func scanTasks(rows *sqlx.Rows) ([]*models.Task, error) {
	tasks := make([]*models.Task, 0)
	for rows.Next() {
		task := &models.Task{}
		var termVal int64
		if err := rows.Scan(
			&task.ID,
			&task.Title,
			&termVal,
			&task.IsGreen,
			&task.CreatedAt,
			&task.UpdatedAt,
			&task.ClosedAt,
			&task.ClosedReason,
			&task.IsArchived,
		); err != nil {
			return nil, errors.Wrap(err, "falied to scan row into struct")
		}
		term, err := models.NewTermFromInt(termVal)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create new term")
		}
		task.Term = *term
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func scanTask(rows *sqlx.Rows) (*models.Task, error) {
	tasks, err := scanTasks(rows)
	if err != nil {
		return nil, err
	}

	if len(tasks) > 1 {
		return nil, errors.New("more than one rows when one was requested")
	}

	if len(tasks) == 0 {
		return nil, nil
	}

	return tasks[0], nil
}
