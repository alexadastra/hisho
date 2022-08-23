package pgsql

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/alexadastra/hisho/hisho-core-service/internal/app/models"
	"github.com/pkg/errors"
)

// GetTasksByRange gets tasks by time range
func (s *PGStorage) GetTasksByRange(ctx context.Context, from *time.Time, to *time.Time) ([]*models.Task, error) {
	query, args, err := sq.Select(allColumns...).
		From(tasksTableName).
		Where(
			sq.And{sq.GtOrEq{createdAtColumnName: from}, sq.LtOrEq{createdAtColumnName: to}},
		).
		PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		return nil, errors.Wrap(err, "failed to format query")
	}

	rows, err := s.conn.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to request db")
	}

	defer rows.Close()

	return scanTasks(rows)
}
