package main

import (
	goapp "github.com/qthang02/goutil/application"
	"github.com/qthang02/goutil/example/config"
)

func main() {
	goapp.Run("sdsds", config.GetConfig())
}
