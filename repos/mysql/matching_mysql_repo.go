package mysql

import (
	"context"

	"github.com/kumin/BityDating/entities"
	"github.com/kumin/BityDating/infras"
	"github.com/kumin/BityDating/repos"
)

var _ repos.MatchingRepo = &MatchingMysqlRepo{}

type MatchingMysqlRepo struct {
	db *infras.MysqlConnector
}

func NewMatchingMysqlRepo(
	db *infras.MysqlConnector,
) *MatchingMysqlRepo {
	return &MatchingMysqlRepo{
		db: db,
	}
}

func (m *MatchingMysqlRepo) CreateOne(ctx context.Context, matching *entities.UserMatching) (*entities.UserMatching, error) {
	if err := m.db.Client.WithContext(ctx).Create(matching).Error; err != nil {
		return nil, err
	}

	return matching, nil
}

func (m *MatchingMysqlRepo) ListMatching(ctx context.Context, userId int64, page, limit int) ([]*entities.User, error) {
	var users []*entities.User
	if err := m.db.Client.WithContext(ctx).Model(&entities.User{}).
		Joins("INNER JOIN matching AS m1 ON user.id = m1.user_id").
		Joins("INNER JOIN matching AS m2 ON m2.user_id = m1.partner_id AND m2.partner_id = m1.user_id").
		Where("m1.partner_id=? AND m1.is_like=1 AND user.is_active = 1", userId).
		Offset(page * limit).
		Limit(limit).
		Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (m *MatchingMysqlRepo) WhoLikeMe(ctx context.Context, partnerId int64) ([]*entities.User, error) {
	var users []*entities.User
	if err := m.db.Client.WithContext(ctx).Model(&entities.User{}).
		Joins("INNER JOIN matching ON user.id = matching.user_id").
		Where("matching.partner_id=? AND matching.is_like=1 AND user.is_active = 1", partnerId).
		Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (m *MatchingMysqlRepo) WhoILike(ctx context.Context, userId int64) ([]*entities.User, error) {
	var users []*entities.User
	if err := m.db.Client.WithContext(ctx).Model(&entities.User{}).
		Joins("INNER JOIN matching ON user.id = matching.partner_id").
		Where("matching.user_id=? AND matching.is_like=1 AND user.is_active = 1", userId).
		Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (m *MatchingMysqlRepo) DeleteOne(ctx context.Context, userId, partnerId int64) error {
	return nil
}
