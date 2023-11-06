package repos

import "context"

type FileRepo interface {
	UploadFile(ctx context.Context, pathFile string) (fileUrl string, err error)
}
