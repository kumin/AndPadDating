package repos

import (
	"context"

	"github.com/kumin/BityDating/entities"
)

type FeedRepo interface {
	GetFeed(ctx context.Context, userId int64, page, limit int) ([]*entities.User, error)
}
