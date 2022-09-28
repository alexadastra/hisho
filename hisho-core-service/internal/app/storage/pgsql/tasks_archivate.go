package pgsql

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/alexadastra/hisho/hisho-core-service/internal/app/models"
	"github.com/pkg/errors"
)

// ArchivateTasks sets tasks as archived (soft delete)
func (s *PGStorage) ArchivateTasks(ctx context.Context, tasks []*models.Task) error {
	q := sq.Update(tasksTableName).
		SetMap(map[string]interface{}{
			isArchivedColumnName: true,
		}).
		Suffix(returningAll).
		PlaceholderFormat(sq.Dollar)

	// select by task IDs
	or := sq.Or{}
	for _, task := range tasks {
		or = append(or, sq.Eq{idColumnName: task.ID})
	}
	q = q.Where(or)

	query, args, err := q.ToSql()
	if err != nil {
		return errors.Wrap(err, "failed to format query")
	}

	rows, err := s.conn.QueryxContext(ctx, query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to request db")
	}

	defer rows.Close()

	return nil
}
