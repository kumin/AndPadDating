package repos

import (
	"context"

	"github.com/kumin/AndPadDating/entities"
)

type UserRepo interface {
	// CURD API
	CreateOne(ctx context.Context, user *entities.User) (*entities.User, error)
	GetOne(ctx context.Context, id int64) (*entities.User, error)
	List(ctx context.Context, page, offset int) ([]*entities.User, error)
	UpdateOne(ctx context.Context, user *entities.User) (*entities.User, error)
	DeleteOne(ctx context.Context, id int64) error

	GetByPhone(ctx context.Context, phone string) (*entities.User, error)
	VerifyPhone(ctx context.Context, phone string) (bool, error)
}
