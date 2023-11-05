package mysql

import (
	"context"
	"errors"

	"github.com/kumin/AndPadDating/entities"
	"github.com/kumin/AndPadDating/infras"
	"github.com/kumin/AndPadDating/repos"
	"gorm.io/gorm"
)

var _ repos.UserRepo = &UserMysqlRepo{}
var immutableColumns = []string{"username", "birthday"}

type UserMysqlRepo struct {
	db *infras.MysqlConnector
}

func NewUserMysqlRepo(
	client *infras.MysqlConnector,
) *UserMysqlRepo {
	return &UserMysqlRepo{
		db: client,
	}
}
func (u *UserMysqlRepo) CreateOne(ctx context.Context, user *entities.User) (*entities.User, error) {
	if err := u.db.Client.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserMysqlRepo) GetOne(ctx context.Context, id int64) (*entities.User, error) {
	var user entities.User
	if err := u.db.Client.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserMysqlRepo) List(ctx context.Context, page, limit int) ([]*entities.User, error) {
	var users []*entities.User
	if err := u.db.Client.WithContext(ctx).Offset(page * limit).Limit(limit).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserMysqlRepo) UpdateOne(ctx context.Context, user *entities.User) (*entities.User, error) {
	if err := u.db.Client.
		WithContext(ctx).
		Model(&user).
		Omit(immutableColumns...).
		Updates(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserMysqlRepo) DeleteOne(ctx context.Context, id int64) error {
	return u.db.Client.WithContext(ctx).Delete(&entities.User{}, id).Error
}

func (u *UserMysqlRepo) GetByPhone(ctx context.Context, phone string) (*entities.User, error) {
	var user entities.User
	if err := u.db.Client.WithContext(ctx).Where("phone=?", phone).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserMysqlRepo) VerifyPhone(ctx context.Context, phone string) (bool, error) {
	var user entities.User
	if err := u.db.Client.WithContext(ctx).WithContext(ctx).Where("phone=?", phone).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return true, nil
		}
		return false, err
	}

	return false, nil
}
