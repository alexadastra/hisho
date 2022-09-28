package pgsql

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/alexadastra/hisho/hisho-core-service/internal/app/models"
	"github.com/pkg/errors"
)

var selectQuery = sq.Select(allColumns...).
	From(tasksTableName).
	OrderBy(createdAtColumnName + " DESC").
	PlaceholderFormat(sq.Dollar)

// GetTaskByID gets task by id
func (s *PGStorage) GetTaskByID(ctx context.Context, id int64) (*models.Task, error) {
	query, args, err := selectQuery.Where(sq.Eq{idColumnName: id}).ToSql()

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

// GetTasksByTerm gets tasks by numeric term from table
func (s *PGStorage) GetTasksByTerm(ctx context.Context, term *models.Term) ([]*models.Task, error) {
	return s.getTasks(ctx, selectQuery.
		Where(sq.Eq{
			termColumnName:       term.ValueInt,
			isArchivedColumnName: false,
		}))
}

// GetTasksByTermAndLimits gets tasks with limit and offset
func (s *PGStorage) GetTasksByTermAndLimits(ctx context.Context, term *models.Term, limit uint64, offset uint64) ([]*models.Task, error) {
	return s.getTasks(ctx, selectQuery.
		Where(sq.Eq{
			termColumnName:       term.ValueInt,
			isArchivedColumnName: false,
		}).
		Limit(limit).
		Offset(offset))
}

// GetClosedTasks gets tasks by numeric term from table
func (s *PGStorage) GetClosedTasks(ctx context.Context) ([]*models.Task, error) {
	return s.getTasks(ctx, selectQuery.
		Where(
			sq.And{
				sq.NotEq{closedAtColumnName: nil},
				sq.Eq{isArchivedColumnName: false},
			},
		))
}

func (s *PGStorage) getTasks(ctx context.Context, q sq.SelectBuilder) ([]*models.Task, error) {
	query, args, err := q.ToSql()

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
