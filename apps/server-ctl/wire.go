//go:build wireinject
// +build wireinject

package apps

import (
	"github.com/google/wire"
	"github.com/kumin/BityDating/monitor/instrumentation"
)

var SuperGraphSet = wire.NewSet(
	ServerGraphSet,
)

func BuildServer() (*HttpServer, error) {
	wire.Build(
		SuperGraphSet,
	)

	return nil, nil
}

func BuildMetricServer() (*instrumentation.MetricServer, error) {
	wire.Build(
		instrumentation.GraphSet,
	)

	return nil, nil
}
