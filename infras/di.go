package infras

import "github.com/google/wire"

var InfaGraphSet = wire.NewSet(
	NewMysqlConnector,
	NewMinioClient,
)
