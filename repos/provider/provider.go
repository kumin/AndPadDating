package provider

import (
	"github.com/google/wire"
	"github.com/kumin/BityDating/infras"
	"github.com/kumin/BityDating/repos"
	"github.com/kumin/BityDating/repos/minio"
	"github.com/kumin/BityDating/repos/mysql"
)

var MysqlGraphSet = wire.NewSet(
	infras.InfaGraphSet,
	mysql.NewUserMysqlRepo,
	wire.Bind(new(repos.UserRepo), new(*mysql.UserMysqlRepo)),
	mysql.NewMatchingMysqlRepo,
	wire.Bind(new(repos.MatchingRepo), new(*mysql.MatchingMysqlRepo)),
	mysql.NewFeedMysqlRepo,
	wire.Bind(new(repos.FeedRepo), new(*mysql.FeedMysqlRepo)),
	minio.NewFileMinioRepo,
	wire.Bind(new(repos.FileRepo), new(*minio.FileMinioRepo)),
)
