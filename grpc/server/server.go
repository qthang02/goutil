package server

import (
	"context"
	"fmt"
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

	if s.healthServer != nil {
		fmt.Println("GrpcServerStop: healthcheck to NOT_SERVING")
		s.healthServer.Shutdown()
	}

	if s.server != nil {
		s.server.GracefulStop()
	}

	return nil
}
