package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"projects_tasks/configs"
	"projects_tasks/internal/service"
	"projects_tasks/internal/transport/grpc/handler"
	pb "projects_tasks/proto/gen"
)

type Server struct {
	GRPCServer *grpc.Server
	config     *configs.Config
}

func NewServer(cfg *configs.Config) *Server {
	s := grpc.NewServer()

	return &Server{
		GRPCServer: s,
		config:     cfg,
	}
}

func (g *Server) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(": %s", g.config.Server.Port))
	if err != nil {
		return err
	}
	if err := g.GRPCServer.Serve(lis); err != nil {
		return err
	}

	return nil
}

func (g *Server) RegisterServices(service *service.Service) {
	pb.RegisterProjectServiceServer(g.GRPCServer, handler.NewProject(service.Project))
	pb.RegisterTaskServiceServer(g.GRPCServer, handler.NewTask(service.Task))

}
