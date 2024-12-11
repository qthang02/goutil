package goserver

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
)

type GrpcServer struct {
	server       *grpc.Server
	healthServer *health.Server
}

func (s *GrpcServer) Stop(context context.Context) error {
	if s == nil {
		return nil
	}

	if s.server != nil {
		s.server.GracefulStop()
	}

	return nil
}
