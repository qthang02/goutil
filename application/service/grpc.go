package gosrv

import (
	"context"
	"fmt"
	gocfg "github.com/qthang02/goutil/config"
	gocopier "github.com/qthang02/goutil/copier"
	grpcServer "github.com/qthang02/goutil/grpc/goserver"
)

type GrpcService struct {
	cfg    *gocfg.BaseConfig
	server *grpcServer.GrpcServer
	opts   []grpcServer.ServerOption
}

func NewGrpcService() *GrpcService {
	return &GrpcService{}
}

func (srv *GrpcService) Init(ctx context.Context, cfg gocfg.ServiceConfig) error {
	srv.cfg = &gocfg.BaseConfig{}

	err := gocopier.JsonCopy(srv.cfg, cfg)
	if err != nil {
		fmt.Println("err: ", err)
		return err
	}

	if len(cfg.GetGrpc().GetListen()) == 0 {
		srv.cfg.Grpc.Listen = ":"
	}

	fmt.Println("grpc listen: ", srv.cfg.Grpc.Listen)

	return nil
}

func (srv *GrpcService) Start(ctx context.Context) error {
	if len(srv.opts) > 0 {
		srv.opts = append(srv.opts, grpcServer.WithServiceConfig(srv.cfg))
	} else {
		srv.opts = []grpcServer.ServerOption{grpcServer.WithServiceConfig(srv.cfg)}
	}

	server, err := grpcServer.CreateGrpcServer(ctx, srv.opts...)
	if err != nil {
		return err
	}

	srv.server = server
	fmt.Println("grpc server start: done")

	return nil
}

func (srv *GrpcService) Stop(ctx context.Context) error {
	if srv.server != nil {
		srv.server.Stop(ctx)
	}

	fmt.Println("GrpcService.Stop: done")
	return nil
}
