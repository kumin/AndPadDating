package mysql

import (
	"context"
	"fmt"

	"github.com/kumin/BityDating/entities"
	"github.com/kumin/BityDating/infras"
	"github.com/kumin/BityDating/repos"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

var _ repos.AlbumRepo = &AlbumMysqlRepo{}

type AlbumMysqlRepo struct {
	mysqlClient *infras.MysqlConnector
	fileRepo    repos.FileRepo
	tracer      trace.Tracer
}

func NewAlbumMysqlRepo(
	mysqlClient *infras.MysqlConnector,
	fileRepo repos.FileRepo,
) *AlbumMysqlRepo {
	return &AlbumMysqlRepo{
		mysqlClient: mysqlClient,
		fileRepo:    fileRepo,
		tracer:      otel.Tracer("AlbumMysqlRepo"),
	}
}

func (a *AlbumMysqlRepo) CreateOne(ctx context.Context, imageFile *entities.File) (*entities.Image, error) {
	imageUrl, err := a.fileRepo.UploadFile(ctx, imageFile)
	if err != nil {
		return nil, err
	}
	image := &entities.Image{
		UserId: ctx.Value(entities.CtxUserIdKey).(int64),
		Url:    imageUrl,
	}
	if err := a.mysqlClient.Client.WithContext(ctx).Create(&image).Error; err != nil {
		return nil, err
	}

	return image, nil
}

func (a *AlbumMysqlRepo) CreateMany(ctx context.Context, imageFiles []*entities.File) ([]*entities.Image, error) {
	imageUrls := make([]string, 0, len(imageFiles))
	for _, file := range imageFiles {
		imageUrl, err := a.fileRepo.UploadFile(ctx, file)
		if err != nil {
			return nil, err
		}
		imageUrls = append(imageUrls, imageUrl)
	}
	fmt.Println(len(imageFiles), len(imageUrls))
	images := make([]*entities.Image, 0, len(imageUrls))
	userId := ctx.Value(entities.CtxUserIdKey).(int64)
	for _, url := range imageUrls {
		image := &entities.Image{
			Url:    url,
			UserId: userId,
		}
		images = append(images, image)
	}
	if err := a.mysqlClient.Client.WithContext(ctx).CreateInBatches(images, 1000).Error; err != nil {
		return nil, err
	}

	return images, nil
}

func (a *AlbumMysqlRepo) GetUserAlbum(ctx context.Context) ([]*entities.Image, error) {
	ctx, span := a.tracer.Start(ctx, "GetUserAlbum")
	defer span.End()
	var images []*entities.Image
	userId := ctx.Value(entities.CtxUserIdKey).(int64)
	if err := a.mysqlClient.Client.WithContext(ctx).Where("user_id=?", userId).Find(&images).Error; err != nil {
		return nil, err
	}

	return images, nil
}
