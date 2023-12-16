package mysql

import (
	"context"

	"github.com/kumin/BityDating/entities"
	"github.com/kumin/BityDating/infras"
	"github.com/kumin/BityDating/repos"
)

var _ repos.AlbumRepo = &AlbumMysqlRepo{}

type AlbumMysqlRepo struct {
	mysqlClient *infras.MysqlConnector
	fileRepo    repos.FileRepo
}

func NewAlbumMysqlRepo(
	mysqlClient *infras.MysqlConnector,
	fileRepo repos.FileRepo,
) *AlbumMysqlRepo {
	return &AlbumMysqlRepo{
		mysqlClient: mysqlClient,
		fileRepo:    fileRepo,
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

func (a *AlbumMysqlRepo) CreateMany(ctx context.Context, imageFile []*entities.File) ([]*entities.Image, error) {
	return nil, nil
}

func (a *AlbumMysqlRepo) GetUserAlbum(ctx context.Context) ([]*entities.Image, error) {
	var images []*entities.Image
	userId := ctx.Value(entities.CtxUserIdKey).(int64)
	if err := a.mysqlClient.Client.WithContext(ctx).Where("user_id=?", userId).Find(&images).Error; err != nil {
		return nil, err
	}

	return images, nil
}
