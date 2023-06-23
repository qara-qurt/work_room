package service

import (
	"projects_tasks/configs"
	"projects_tasks/internal/repository"
	pb "projects_tasks/proto/gen"
)

type IProject interface {
	CreateProject(project *pb.CreateProjectRequest) (uint64, error)
	GetProjects()
	GetProject()
	UpdateProject()
	DeleteProject()
}

type ITask interface {
	CreateTasks()
	GetTask()
	UpdateTask()
	DeleteTask()
}

type Service struct {
	Project IProject
	Task    ITask
}

func New(repo *repository.Repository, cfg *configs.Config) *Service {

	project := NewProject(repo.Project)
	task := NewTask(repo.Task)
	return &Service{
		Project: project,
		Task:    task,
	}
}
