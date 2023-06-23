package handler

import (
	"projects_tasks/internal/service"
	pb "projects_tasks/proto/gen"
)

type Task struct {
	service service.ITask
	pb.UnimplementedTaskServiceServer
}

func NewTask(service service.ITask) *Task {
	return &Task{
		service: service,
	}
}
