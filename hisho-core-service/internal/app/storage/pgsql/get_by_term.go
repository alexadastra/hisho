package pgsql

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/alexadastra/hisho/hisho-core-service/internal/app/models"
	"github.com/pkg/errors"
)

// GetTasksByTerm gets tasks by numeric term from table
func (s *PGStorage) GetTasksByTerm(ctx context.Context, term *models.Term) ([]*models.Task, error) {
	query, args, err := sq.Select(allColumns...).
		From(tasksTableName).
		Where(
			sq.Eq{termColumnName: term.ValueInt},
		).
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

	return scanTasks(rows)
}

// GetTasksByTermAndLimits gets task with limit and offset
func (s *PGStorage) GetTasksByTermAndLimits(ctx context.Context, term *models.Term, limit uint64, offset uint64) ([]*models.Task, error) {
	query, args, err := sq.Select(allColumns...).
		From(tasksTableName).
		Where(
			sq.Eq{termColumnName: term.ValueInt},
		).
		OrderBy(createdAtColumnName + " DESC").
		Limit(limit).
		Offset(offset).
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

	return scanTasks(rows)
}
