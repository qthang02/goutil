package gosrv

import (
	"context"
	gocfg "github.com/qthang02/goutil/config"
)

type GoService interface {
	Init(ctx context.Context, cfg gocfg.ServiceConfig) error
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}
