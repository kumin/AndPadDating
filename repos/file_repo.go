package repos

import (
	"context"

	"github.com/kumin/BityDating/entities"
)

type FileRepo interface {
	UploadFile(
		ctx context.Context,
		file *entities.File,
	) (fileUrl string, err error)
}
