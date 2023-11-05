package configs

import (
	"time"

	"github.com/kumin/AndPadDating/pkg/envx"
)

var MaxPageLimit = envx.GetInt("MAX_PAGE_LIMIT", 100)

var (
	TokenExpiredTime = time.Duration(envx.GetInt("TOKEN_EXPRIED_TIME", 24)) * time.Hour
	SecretKey        = envx.GetString("SECRETKEY", "mysecretkey")
)
