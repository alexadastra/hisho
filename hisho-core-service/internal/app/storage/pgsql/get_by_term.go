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
		PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		return nil, errors.Wrap(err, "failed to format query")
	}

	rows, err := s.conn.Query(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to request db")
	}

	defer rows.Close()

	tasks := make([]*models.Task, 0)
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Term, &task.CreatedAt, &task.UpdatedAt, &task.DoneAt); err != nil {
			return nil, errors.Wrap(err, "failed to scan row")
		}
		tasks = append(tasks, &task)
	}
	return tasks, nil
}
