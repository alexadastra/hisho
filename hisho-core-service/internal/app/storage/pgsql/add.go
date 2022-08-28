package pgsql

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/alexadastra/hisho/hisho-core-service/internal/app/models"
	"github.com/jmoiron/sqlx"

	"github.com/pkg/errors"
)

var (
	insertQuery = sq.
		Insert(tasksTableName).
		Columns(titleColumnName, termColumnName, createdAtColumnName, updatedAtColumnName, doneAtColumnName).
		Suffix("on conflict do nothing returning *").
		PlaceholderFormat(sq.Dollar)
)

// AddTasks adds task to DB
func (s *PGStorage) AddTasks(ctx context.Context, tasks []*models.Task) ([]*models.Task, error) {
	query := insertQuery
	for _, task := range tasks {
		query = query.Values(task.Title, task.Term.ValueInt, task.CreatedAt, task.UpdatedAt, task.DoneAt)
	}

	q, args, err := query.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to form the insert query")
	}

	rows, err := s.conn.QueryxContext(ctx, q, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute the insert query")
	}

	defer rows.Close()

	return scanTasks(rows)
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
			&task.CreatedAt,
			&task.UpdatedAt,
			&task.DoneAt,
		); err != nil {
			return nil, errors.Wrap(err, "falied to scan row into struct")
		}
		term, err := models.NewTermFromInt(termVal)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create new term")
		}
		task.Term = term
		tasks = append(tasks, task)
	}

	return tasks, nil
}
