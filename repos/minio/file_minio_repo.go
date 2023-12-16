package minio

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/kumin/BityDating/configs"
	"github.com/kumin/BityDating/entities"
	"github.com/kumin/BityDating/pkg/strings"
	"github.com/kumin/BityDating/repos"
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

func (f *FileMinioRepo) UploadFile(
	ctx context.Context,
	file *entities.File) (fileUrl string, err error) {
	fileName := fmt.Sprintf("%s-%s", uuid.NewString(), file.Name)
	fileInfo, err := f.minioClient.PutObject(ctx,
		configs.BuketName,
		fileName,
		file.Buffer,
		file.Size,
		minio.PutObjectOptions{ContentType: file.ContentType})
	if err != nil {
		return "", err
	}
	if strings.IsEmpty(fileInfo.Key) {
		return "", nil
	}
	//should use Nginx for proxy image server
	fileUrl = fmt.Sprintf("%s://%s/%s/%s", configs.MinioProtocol, configs.MinioHost, configs.BuketName, fileInfo.Key)
	err = nil
	return
}
