//go:build wireinject
// +build wireinject

package apps

import (
	"github.com/google/wire"
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
