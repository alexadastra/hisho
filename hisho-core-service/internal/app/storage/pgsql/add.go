package pgsql

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/alexadastra/hisho/hisho-core-service/internal/app/models"

	"github.com/pkg/errors"
)

// AddTasks adds task to DB
func (s *PGStorage) AddTasks(ctx context.Context, tasks []*models.Task) ([]*models.Task, error) {
	query := sq.
		Insert(tasksTableName).
		Columns(allColumns...).
		Suffix("on conflict do nothing returning *").
		PlaceholderFormat(sq.Dollar)

	for _, task := range tasks {
		query = query.Values(
			task.ID,
			task.Title,
			task.Term.ValueInt,
			task.IsGreen,
			task.CreatedAt,
			task.UpdatedAt,
			task.ClosedAt,
			task.ClosedReason,
		)
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
