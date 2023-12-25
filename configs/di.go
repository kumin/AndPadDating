package configs

import "github.com/google/wire"

var ConfigGraphSet = wire.NewSet(
	NewServerConfiguration,
	NewMetricServerCfg,
)
