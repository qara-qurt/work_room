package repository

import (
	"github.com/sirupsen/logrus"
	"projects_tasks/configs"
	"projects_tasks/internal/repository/postgres"
	pb "projects_tasks/proto/gen"
)

type IProjectRepository interface {
	CreateProject(project *pb.CreateProjectRequest) (uint64, error)
}

type ITaskRepository interface {
}

type Repository struct {
	Project IProjectRepository
	Task    ITaskRepository
}

func New(cfg *configs.Config) (*Repository, error) {
	postgresDB, err := postgres.NewDatabasePSQL(cfg)
	if err != nil {
		return nil, err
	}
	logrus.Info("postgres successfully connected")

	projectRepo := postgres.NewProject(postgresDB.DB)
	taskRepo := postgres.NewTask(postgresDB.DB)
	return &Repository{
		Project: projectRepo,
		Task:    taskRepo,
	}, nil
}
