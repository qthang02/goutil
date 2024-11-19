package config

import gocfg "github.com/qthang02/goutil/config"

var (
	cfg = &Config{}
)

type Config struct {
	gocfg.BaseConfig
	TestConfig string `json:"testConfig"`
}

func GetConfig() *Config {
	return cfg
}
