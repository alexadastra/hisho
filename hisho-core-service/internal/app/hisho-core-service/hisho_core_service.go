package service

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/alexadastra/hisho/hisho-core-service/internal/app/models"
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
	term, err := models.NewTermFromString(request.Term.String())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to create term entity: %s", err)
	}

	paging, err := models.NewPaging(request.Page)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to create paging entity: %s", err)
	}

	tasks, err := hs.service.GetTasksByTerm(ctx, term, paging)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get tasks by term")
	}
	tasksToReturn := make([]*api.Task, 0, len(tasks))
	for _, task := range tasks {
		tasksToReturn = append(tasksToReturn, task.ToProto())
	}
	return &api.TasksResponse{
		Tasks: tasksToReturn,
	}, nil
}

// AddTask adds new task
func (hs *HishoCoreService) AddTask(ctx context.Context, request *api.AddTaskRequest) (*api.Task, error) {
	taskRequest, err := models.NewTaskRequest(request)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create new task")
	}
	task, err := hs.service.AddTask(ctx, []*models.TaskRequest{taskRequest})
	if err != nil {
		return nil, errors.Wrap(err, "failed to add task")
	}
	return task.ToProto(), nil
}

// ChangeTaskStatus implements api.HishoCoreServiceServer
func (hs *HishoCoreService) ChangeTaskStatus(ctx context.Context, request *api.ChangeTaskStatusRequest) (*api.Task, error) {
	panic("unimplemented")
}

// EditTask implements api.HishoCoreServiceServer
func (hs *HishoCoreService) EditTask(context.Context, *api.Task) (*api.Task, error) {
	panic("unimplemented")
}
