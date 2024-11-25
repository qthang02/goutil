package gosrv

import (
	"context"
	"fmt"
	gocfg "github.com/qthang02/goutil/config"
	gocopier "github.com/qthang02/goutil/copier"
	grpcServer "github.com/qthang02/goutil/grpc/server"
)

type GrpcService struct {
	cfg    *gocfg.BaseConfig
	server *grpcServer.GrpcServer
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
	return nil
}

func (srv *GrpcService) Stop(ctx context.Context) error {
	if srv.server != nil {
		srv.server.Stop(ctx)
	}

	fmt.Println("GrpcService.Stop: done")
	return nil
}
