package service

import (
	"projects_tasks/internal/repository"
	pb "projects_tasks/proto/gen"
)

type Project struct {
	repo repository.IProjectRepository
}

func (p Project) CreateProject(project *pb.CreateProjectRequest) (uint64, error) {
	return p.repo.CreateProject(project)
}

func (p Project) GetProjects() {
}

func (p Project) GetProject() {
}

func (p Project) UpdateProject() {

}

func (p Project) DeleteProject() {

}

func NewProject(repo repository.IProjectRepository) *Project {
	return &Project{
		repo: repo,
	}
}
