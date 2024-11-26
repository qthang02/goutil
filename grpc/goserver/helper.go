package goserver

import (
	"context"
	"errors"
	"fmt"
	gocfg "github.com/qthang02/goutil/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"net"
)

type GrpcRegisterHandler func(ctx context.Context, server *grpc.Server)

type serverOptions struct {
	cfg gocfg.ServiceConfig

	unaryInterceptors  []grpc.UnaryServerInterceptor
	streamInterceptors []grpc.StreamServerInterceptor
	GrpcServerOpts     []grpc.ServerOption
	registerHandler    GrpcRegisterHandler
}

type ServerOption func(*serverOptions)

func defaultOpts() *serverOptions {
	opts := &serverOptions{}

	return opts
}

func CreateGrpcServer(ctx context.Context, opts ...ServerOption) (*GrpcServer, error) {
	server, healthServer, err := createGrpcServerConn(ctx, opts...)
	if err != nil {
		fmt.Println("CreateGrpcServer: create grpc server err:", err)
		return nil, err
	}

	return &GrpcServer{
		server:       server,
		healthServer: healthServer,
	}, nil
}

func createGrpcServerConn(ctx context.Context, opts ...ServerOption) (*grpc.Server, *health.Server, error) {
	serverOpts := defaultOpts()
	for _, opt := range opts {
		opt(serverOpts)
	}

	if serverOpts.cfg == nil {
		return nil, nil, errors.New("CreateGrpcServerConn: grpc server config is nil")
	}

	serverOpts.GrpcServerOpts = append(serverOpts.GrpcServerOpts,
		grpc.ChainUnaryInterceptor(serverOpts.unaryInterceptors...),
		grpc.ChainStreamInterceptor(serverOpts.streamInterceptors...),
		grpc.MaxHeaderListSize(2147483647),
	)

	grpcServer := grpc.NewServer(serverOpts.GrpcServerOpts...)

	if serverOpts.registerHandler != nil {
		serverOpts.registerHandler(ctx, grpcServer)
	}

	reflection.Register(grpcServer)
	lis, err := net.Listen("tcp", serverOpts.cfg.GetGrpc().GetListen())
	if err != nil {
		fmt.Println("CreateGrpcServerConn: listen error: ", err)
		return nil, nil, err
	}

	var hcsv *health.Server
	hcsv = health.NewServer()
	healthgrpc.RegisterHealthServer(grpcServer, hcsv)
	fmt.Println("CreateGrpcServerConn: registerd healthcheck /grpc.health.v1.Health/Check")

	if hcsv != nil {
		hcsv.SetServingStatus("", healthgrpc.HealthCheckResponse_SERVING)
		fmt.Println("CreateGrpcServerConn: healthcheck to SERVING")
	}

	if err := grpcServer.Serve(lis); err != nil {
		fmt.Println("CreateGrpcServerConn: grpc server serve error: ", err)
	}

	return grpcServer, hcsv, nil
}

func WithServiceConfig(cfg gocfg.ServiceConfig) ServerOption {
	return func(opt *serverOptions) {
		opt.cfg = cfg
	}
}
