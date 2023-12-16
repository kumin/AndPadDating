package configs

import (
	"time"

	"github.com/kumin/BityDating/pkg/envx"
)

var MaxPageLimit = envx.GetInt("MAX_PAGE_LIMIT", 100)

var (
	TokenExpiredTime = time.Duration(envx.GetInt("TOKEN_EXPRIED_TIME", 24)) * time.Hour
	SecretKey        = envx.GetString("SECRETKEY", "mysecretkey")
)

var (
	MinioHost       = envx.GetString("MINIO_HOST", "localhost:9000")
	MinioProtocol   = envx.GetString("MINIO_PROTOCOL", "http")
	AccessKeyId     = envx.GetString("ACCESS_KEY_ID", "I91rtieW4WuptqlNR6pn")
	SecretAccessKey = envx.GetString("SECRET_ACCESS_KEY", "xl1gMawfhfjQdXhXW7B2Yy3Bm4ZAHZcZoRjvkeGe")
	BuketName       = envx.GetString("MINIO_BUCKET_NAME", "bity")
)
