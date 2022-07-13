package pgsql

import (
	"context"

	"github.com/alexadastra/hisho/hisho-core-service/internal/app/models"
)

// SetTaskAsDone updates task setting 'done_at' as now()
func (s *PGStorage) SetTaskAsDone(ctx context.Context, id int64) (*models.Task, error) {
	return nil, nil
}
