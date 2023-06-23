package handler

import (
	"context"
	log "github.com/sirupsen/logrus"
	"projects_tasks/internal/service"
	pb "projects_tasks/proto/gen"
)

type Project struct {
	service service.IProject
	pb.UnimplementedProjectServiceServer
}

func NewProject(service service.IProject) *Project {
	return &Project{
		service: service,
	}
}

func (p Project) CreateProject(ctx context.Context, request *pb.CreateProjectRequest) (*pb.CreateProjectResponse, error) {

	projectID, err := p.service.CreateProject(request)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	res := &pb.CreateProjectResponse{
		ProjectId: projectID,
	}
	return res, nil
}

func (p Project) GetProjects(ctx context.Context, request *pb.GetProjectsRequest) (*pb.GetProjectsResponse, error) {
	p.service.GetProjects()
	return &pb.GetProjectsResponse{}, nil
}

func (p Project) GetProject(ctx context.Context, request *pb.GetProjectRequest) (*pb.GetProjectResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p Project) UpdateProject(ctx context.Context, request *pb.UpdateProjectRequest) (*pb.GetProjectResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p Project) DeleteProject(ctx context.Context, request *pb.DeleteProjectRequest) (*pb.GetProjectResponse, error) {
	//TODO implement me
	panic("implement me")
}
