package mysql

import (
	"context"

	"github.com/kumin/BityDating/entities"
	"github.com/kumin/BityDating/infras"
	"github.com/kumin/BityDating/repos"
)

var _ repos.FeedRepo = &FeedMysqlRepo{}

type FeedMysqlRepo struct {
	db *infras.MysqlConnector
}

func NewFeedMysqlRepo(
	db *infras.MysqlConnector,
) *FeedMysqlRepo {
	return &FeedMysqlRepo{
		db: db,
	}
}

func (f *FeedMysqlRepo) GetFeed(ctx context.Context, userId int64, page, limit int) ([]*entities.User, error) {
	var users []*entities.User
	if err := f.db.Client.WithContext(ctx).
		Joins("LEFT JOIN matching ON user.id = matching.partner_id AND (matching.user_id = ? OR matching.partner_id = ?)", userId, userId).
		Where("user.id != ? AND user.is_active = 1 AND matching.user_id IS NULL", userId).
		Offset(page * limit).
		Limit(limit).
		Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
