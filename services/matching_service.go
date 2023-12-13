package services

import (
	"context"

	"github.com/kumin/BityDating/entities"
	"github.com/kumin/BityDating/repos"
)

type MatchingService struct {
	matchingRepo repos.MatchingRepo
}

func NewMatchingService(
	matchingRepo repos.MatchingRepo,
) *MatchingService {
	return &MatchingService{
		matchingRepo: matchingRepo,
	}
}

func (m *MatchingService) CreateMatching(ctx context.Context, matching *entities.UserMatching) (*entities.UserMatching, error) {
	return m.matchingRepo.CreateOne(ctx, matching)
}

func (m *MatchingService) WhoILike(ctx context.Context, userId int64) ([]*entities.User, error) {
	return m.matchingRepo.WhoILike(ctx, userId)
}

func (m *MatchingService) WhoLikeMe(ctx context.Context, partnerId int64) ([]*entities.User, error) {
	return m.matchingRepo.WhoLikeMe(ctx, partnerId)
}

func (m *MatchingService) ListMatching(ctx context.Context, userId int64, page, limit int) ([]*entities.User, error) {
	return m.matchingRepo.ListMatching(ctx, userId, page, limit)
}
