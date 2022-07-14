package pgsql

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/alexadastra/hisho/hisho-core-service/internal/app/models"

	"github.com/pkg/errors"
)

// AddTask adds task to DB
func (s *PGStorage) AddTask(ctx context.Context, task *models.Task) (int64, error) {
	q, args, err := sq.Insert(tasksTableName).
		Values(
			task.ID, task.Title, task.Term, task.CreatedAt, task.UpdatedAt, task.DoneAt).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return 0, errors.Wrap(err, "failed to form the insert query")
	}

	ct, err := s.conn.Exec(ctx, q, args...)
	if err != nil {
		return 0, errors.Wrap(err, "failed to execute the insert query")
	}

	return ct.RowsAffected(), nil
}
