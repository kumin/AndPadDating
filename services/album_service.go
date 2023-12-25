package services

import (
	"context"
	"mime/multipart"

	"github.com/kumin/BityDating/entities"
	"github.com/kumin/BityDating/repos"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type AlbumService struct {
	albumRepo repos.AlbumRepo
	tracer    trace.Tracer
}

func NewAlbumService(
	albumRepo repos.AlbumRepo,
) *AlbumService {
	return &AlbumService{
		albumRepo: albumRepo,
		tracer:    otel.Tracer("AlbumService"),
	}
}

func (a *AlbumService) CreateOne(ctx context.Context, imageFile *multipart.FileHeader) (*entities.Image, error) {
	image, err := ConvertMultipartToFile(imageFile)
	if err != nil {
		return nil, err
	}

	return a.albumRepo.CreateOne(ctx, image)
}

func (a *AlbumService) CreateMany(ctx context.Context, imageFiles []*multipart.FileHeader) ([]*entities.Image, error) {
	images := make([]*entities.File, 0, len(imageFiles))
	for _, imageFile := range imageFiles {
		image, err := ConvertMultipartToFile(imageFile)
		if err != nil {
			return nil, err
		}
		images = append(images, image)
	}

	return a.albumRepo.CreateMany(ctx, images)
}

func (a *AlbumService) GetUserAlbum(ctx context.Context) ([]*entities.Image, error) {
	ctx, span := a.tracer.Start(ctx, "GetUserAlbum")
	defer span.End()
	return a.albumRepo.GetUserAlbum(ctx)
}
