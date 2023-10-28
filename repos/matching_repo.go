package repos

import (
	"context"

	"github.com/kumin/AndPadDating/entities"
)

type MatchingRepo interface {
	CreateOne(ctx context.Context, matching *entities.UserMatching) (*entities.UserMatching, error)
	ListMatching(ctx context.Context, userId int64, page, limit int) ([]*entities.User, error)
	WhoLikeMe(ctx context.Context, partnerId int64) ([]*entities.User, error)
	WhoILike(ctx context.Context, userId int64) ([]*entities.User, error)
	DeleteOne(ctx context.Context, userId, partnerId int64) error
}
