package infras

import (
	"log"
	"time"

	"github.com/kumin/go-tpc/pkg/envx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlConnector struct {
	Client *gorm.DB
}

func MysqlConnectionBuilder(fns ...optionFn) *MysqlConnector {
	opt := MysqlDefaultOption
	for _, f := range fns {
		f(opt)
	}
	client, err := gorm.Open(mysql.Open(opt.DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	pool, err := client.DB()
	if err != nil {
		log.Fatal(err)
	}
	pool.SetMaxOpenConns(opt.MaxConn)
	pool.SetConnMaxLifetime(opt.MaxLifetime)

	return &MysqlConnector{
		Client: client,
	}
}

func NewMysqlConnector() *MysqlConnector {
	return MysqlConnectionBuilder(
		WithDSN(envx.GetString("MYSQL_ADDRS",
			"root:12345678@tcp(localhost:3306)/kumin?charset=utf8&parseTime=True&loc=Local&multiStatements=true")),
		WithMaxConn(2),
		WithLifetime(1*time.Minute),
	)
}

type MysqlOption struct {
	DSN         string
	MaxConn     int
	MaxLifetime time.Duration
}

var MysqlDefaultOption = &MysqlOption{
	MaxConn:     2,
	MaxLifetime: 1 * time.Minute,
}

type optionFn func(opt *MysqlOption)

func WithDSN(dsn string) optionFn {
	return func(opt *MysqlOption) {
		opt.DSN = dsn
	}
}

func WithMaxConn(conns int) optionFn {
	return func(opt *MysqlOption) {
		opt.MaxConn = conns
	}
}

func WithLifetime(minus time.Duration) optionFn {
	return func(opt *MysqlOption) {
		opt.MaxLifetime = minus
	}
}
