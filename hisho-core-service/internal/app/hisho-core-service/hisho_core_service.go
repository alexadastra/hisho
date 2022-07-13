package service

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"

	service2 "github.com/alexadastra/hisho/hisho-core-service/internal/app/service"
	"github.com/alexadastra/hisho/hisho-core-service/pkg/api"
)

// HishoCoreService handles stuff
type HishoCoreService struct {
	service *service2.Service
}

// NewHishoCoreService creates new server
func NewHishoCoreService(serv *service2.Service) api.HishoCoreServiceServer {
	return &HishoCoreService{
		service: serv,
	}
}

// Ping returns "pong" if ping is pinged
func (hs *HishoCoreService) Ping(ctx context.Context, request *api.PingRequest) (*api.PingResponse, error) {
	if request.Value == "ping" {
		return &api.PingResponse{
			Code:  200,
			Value: "pong",
		}, nil
	}
	return nil, fmt.Errorf("unknown request message: %s", request.Value)
}

// GetTasksByTerm gets tasks for terms (today, week, other etc.)
func (hs *HishoCoreService) GetTasksByTerm(ctx context.Context, request *api.TaskRequest) (*api.TasksResponse, error) {
	tasks, err := hs.service.GetTasksByTerm(ctx, request.Term.String())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get tasks by term")
	}
	tasksToReturn := make([]*api.Task, 0, len(tasks))
	for _, task := range tasks {
		tasksToReturn = append(tasksToReturn, &api.Task{
			Title:     task.Title,
			Term:      api.Term_name[int32(task.Term)],
			CreatedAt: timestamppb.New(*task.CreatedAt),
			UpdatedAt: timestamppb.New(*task.UpdatedAt),
			DoneAt:    timestamppb.New(*task.DoneAt),
		})
	}
	return &api.TasksResponse{
		Tasks: tasksToReturn,
	}, nil
}
