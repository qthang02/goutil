package gocfg

import "strings"

type ServiceConfig interface {
	GetEnvironment() string
	GetServiceName() string
	SetServiceName(string)
}

type BaseConfig struct {
	Environment string `json:"environment,omitempty"`
	ServiceName string `json:"serviceName,omitempty"`
}

func (bc *BaseConfig) GetEnvironment() string {
	if bc == nil {
		return ""
	}

	return bc.Environment
}

func (bc *BaseConfig) GetServiceName() string {
	if bc == nil {
		return ""
	}

	return bc.ServiceName
}

func (bc *BaseConfig) SetServiceName(name string) {
	if bc == nil {
		return
	}

	if strings.HasPrefix(name, bc.Environment) {
		bc.ServiceName = name
		return
	}

	bc.ServiceName = bc.Environment + "." + name
}