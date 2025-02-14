package goapp

import (
	"context"
	"fmt"
	gosrv "github.com/qthang02/goutil/application/service"
	gocfg "github.com/qthang02/goutil/config"
)

func Run(serviceName string, cfg gocfg.ServiceConfig, services ...gosrv.GoService) {

	// Init config
	err := gocfg.Init(cfg, serviceName)
	if err != nil {
		fmt.Println("Run application: load config error: ", err)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, service := range services {
		if err := service.Init(ctx, cfg); err != nil {
			fmt.Println("Run application: init service error: ", err)
			return
		}
	}

	for _, service := range services {
		if err := service.Start(ctx); err != nil {
			fmt.Println("Run application: start service error: ", err)
			return
		}
	}

}
