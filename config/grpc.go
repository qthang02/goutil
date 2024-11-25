package gocfg

type GrpcConfig struct {
	Listen string `json:"listen,omitempty"`
}

func (cfg GrpcConfig) GetListen() string {
	if cfg.Listen != "" {
		return cfg.Listen
	}

	return ""
}
