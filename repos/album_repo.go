package repos

import (
	"context"

	"github.com/kumin/BityDating/entities"
)

type AlbumRepo interface {
	CreateOne(ctx context.Context, imageFile *entities.File) (*entities.Image, error)
	CreateMany(ctx context.Context, imageFiles []*entities.File) ([]*entities.Image, error)
	GetUserAlbum(ctx context.Context) ([]*entities.Image, error)
}
