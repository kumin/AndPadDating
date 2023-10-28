package configs

import "github.com/kumin/go-tpc/pkg/envx"

type ServerConfiguration struct {
	Port int
}

func NewServerConfiguration() *ServerConfiguration {
	return &ServerConfiguration{
		Port: envx.GetInt("API_SERVING_PORT", 8080),
	}
}
