package service

import (
	"github.com/alexadastra/hisho/hisho-core-service/internal/app/storage"
)

const timeFormat = "2006/01/02"

// Service handles everything
type Service struct {
	storage storage.Storage
}

// NewService constructs new Service
func NewService(storage storage.Storage) *Service {
	return &Service{
		storage: storage,
	}
}
