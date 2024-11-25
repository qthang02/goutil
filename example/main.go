package main

import (
	goapp "github.com/qthang02/goutil/application"
	gosrv "github.com/qthang02/goutil/application/service"
	"github.com/qthang02/goutil/example/config"
)

func main() {
	goapp.Run("console", config.GetConfig(), gosrv.NewGrpcService())
}
