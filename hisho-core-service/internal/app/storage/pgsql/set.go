package pgsql

import (
	"context"
	"time"

	"github.com/alexadastra/hisho/hisho-core-service/internal/app/models"
	"github.com/pkg/errors"

	sq "github.com/Masterminds/squirrel"
)

// ChangeTaskStatus changes task 'closed_at' fiels
func (s *PGStorage) ChangeTaskStatus(ctx context.Context, id int64, closedAt *time.Time, reason string) (*models.Task, error) {
	query, args, err := sq.Update(tasksTableName).
		SetMap(map[string]interface{}{
			closedAtColumnName:     closedAt,
			closedReasonColumnName: reason,
			updatedAtColumnName:    time.Now().UTC(),
		}).
		Where(idColumnName, id).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, errors.Wrap(err, "failed to format query")
	}

	rows, err := s.conn.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to request db")
	}

	defer rows.Close()

	return scanTask(rows)
}

// EditTask sets editable task fields
func (s *PGStorage) EditTask(ctx context.Context, task models.Task) (*models.Task, error) {
	query, args, err := sq.Update(tasksTableName).
		SetMap(map[string]interface{}{
			titleColumnName:     task.Title,
			termColumnName:      task.Term.ValueInt,
			isGreenColumnName:   task.IsGreen,
			updatedAtColumnName: time.Now().UTC(),
		}).
		Where(idColumnName, task.ID).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, errors.Wrap(err, "failed to format query")
	}

	rows, err := s.conn.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to request db")
	}

	defer rows.Close()

	return scanTask(rows)
}
