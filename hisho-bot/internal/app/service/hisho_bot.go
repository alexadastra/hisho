package service

import (
	"context"
	"fmt"

	"github.com/alexadastra/hisho/hisho-bot/pkg/api"
)

// HishoBot handles stuff
type HishoBot struct{}

// NewHishoBot creates new server
func NewHishoBot() api.HishoBotServer {
	return &HishoBot{}
}

// Ping returns "pong" if ping in pinged
func (rt *HishoBot) Ping(ctx context.Context, request *api.PingRequest) (*api.PingResponse, error) {
	if request.Value == "ping" {
		return &api.PingResponse{
			Code:  200,
			Value: "pong",
		}, nil
	}
	return nil, fmt.Errorf("unknown request message: %s", request.Value)
}
