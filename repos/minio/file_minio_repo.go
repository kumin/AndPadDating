package minio

import (
	"context"
	"strings"

	"github.com/kumin/AndPadDating/configs"
	"github.com/kumin/AndPadDating/erroz"
	"github.com/kumin/AndPadDating/repos"
	"github.com/minio/minio-go/v7"
)

var _ repos.FileRepo = &FileMinioRepo{}

type FileMinioRepo struct {
	minioClient *minio.Client
}

func NewFileMinioRepo(
	minioClient *minio.Client,
) *FileMinioRepo {
	return &FileMinioRepo{
		minioClient: minioClient,
	}
}

func (f *FileMinioRepo) UploadFile(ctx context.Context, pathFile string) (fileUrl string, err error) {
	fileName := strings.Split(pathFile, "/")
	if len(fileName) < 2 {
		return "", erroz.ErrInvalidPathFile
	}
	fileInfo, err := f.minioClient.FPutObject(ctx, configs.BuketName, fileName[len(fileName)-1], pathFile,
		minio.PutObjectOptions{ContentType: "image/png"})
	if err != nil {
		return "", err
	}

	return fileInfo.Location, nil
}
