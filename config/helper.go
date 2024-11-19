package gocfg

import "fmt"

const (
	jsonConfig = "config.json"
)

func Init(cfg ServiceConfig, serverName string) error {

	loadInitConfig(cfg)

	return nil
}

func loadInitConfig(cfg ServiceConfig) {
	err := loadConfigFile(jsonConfig, cfg)
	if err == nil {
		fmt.Println("loadInitConfig: loaded config successfully")
		return
	}
}
