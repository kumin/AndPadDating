package infras

import (
	"github.com/kumin/BityDating/configs"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinIOOption struct {
	Endpoint        string
	AccessKeyId     string
	SecretAccessKey string
	UseSSL          bool
}

var MinIODefaultOpt = &MinIOOption{
	Endpoint:        configs.MinioHost,
	AccessKeyId:     configs.AccessKeyId,
	SecretAccessKey: configs.SecretAccessKey,
}

type MinIOOptFunction func(opt *MinIOOption) *MinIOOption

func WithEndPoint(endPoint string) MinIOOptFunction {
	return func(opt *MinIOOption) *MinIOOption {
		opt.Endpoint = endPoint
		return opt
	}
}

func WithAccessKeyId(accessKeyId string) MinIOOptFunction {
	return func(opt *MinIOOption) *MinIOOption {
		opt.AccessKeyId = accessKeyId
		return opt
	}
}

func WithSecretAccessKey(secretAccessKey string) MinIOOptFunction {
	return func(opt *MinIOOption) *MinIOOption {
		opt.SecretAccessKey = secretAccessKey
		return opt
	}
}

func WithUseSSL(useSSL bool) MinIOOptFunction {
	return func(opt *MinIOOption) *MinIOOption {
		opt.UseSSL = useSSL
		return opt
	}
}

func NewMinioClientWithOptions(optFns ...MinIOOptFunction) (*minio.Client, error) {
	minIOOpt := MinIODefaultOpt
	for _, opt := range optFns {
		minIOOpt = opt(minIOOpt)
	}
	return minio.New(
		minIOOpt.Endpoint,
		&minio.Options{
			Creds:  credentials.NewStaticV4(minIOOpt.AccessKeyId, minIOOpt.SecretAccessKey, ""),
			Secure: minIOOpt.UseSSL,
		},
	)
}

func NewMinioClient() (*minio.Client, error) {
	return NewMinioClientWithOptions()
}
