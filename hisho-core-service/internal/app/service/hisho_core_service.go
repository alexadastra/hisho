package service

import (
	"context"
	"fmt"

	"github.com/alexadastra/hisho/hisho-core-service/pkg/api"
)

// HishoCoreService handles stuff
type HishoCoreService struct{}

// NewHishoCoreService creates new server
func NewHishoCoreService() api.HishoCoreServiceServer {
	return &HishoCoreService{}
}

// Ping returns "pong" if ping in pinged
func (rt *HishoCoreService) Ping(ctx context.Context, request *api.PingRequest) (*api.PingResponse, error) {
	if request.Value == "ping" {
		return &api.PingResponse{
			Code:  200,
			Value: "pong",
		}, nil
	}
	return nil, fmt.Errorf("unknown request message: %s", request.Value)
}
