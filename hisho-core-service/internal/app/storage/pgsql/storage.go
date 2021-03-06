package pgsql

import (
	"context"

	"github.com/alexadastra/hisho/hisho-core-service/internal/app/storage"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

const (
	tasksTableName      = "tasks"
	idColumnName        = "id"
	titleColumnName     = "title"
	termColumnName      = "term"
	createdAtColumnName = "created_at"
	updatedAtColumnName = "updated_at"
	doneAtColumnName    = "done_at"
)

var allColumns = []string{idColumnName, titleColumnName, termColumnName, createdAtColumnName, updatedAtColumnName, doneAtColumnName}

// PGStorage implements Storage with Postgres database
type PGStorage struct {
	conn *pgx.Conn
}

// NewStorage constructs new PGStorage
func NewStorage(ctx context.Context, host string) (storage.Storage, error) {
	conn, err := pgx.Connect(ctx, host)
	if err != nil {
		return nil, errors.Wrap(err, "unable to connect to database")
	}

	return &PGStorage{
		conn: conn,
	}, nil
}
