//go:build unit
// +build unit

package services

import (
	"context"
	"testing"

	"github.com/kumin/AndPadDating/entities"
	mocks_data "github.com/kumin/AndPadDating/mocks/data"
	mocks_repo "github.com/kumin/AndPadDating/mocks/repos"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MatchingServiceTestSuite struct {
	suite.Suite
	matchingService *MatchingService
	matchingRepo    *mocks_repo.MatchingRepo
}

func (m *MatchingServiceTestSuite) SetupTest() {
	m.matchingRepo = mocks_repo.NewMatchingRepo(m.T())
	m.matchingService = NewMatchingService(m.matchingRepo)
}

func (m *MatchingServiceTestSuite) TestMatchingService_Create() {
	m.matchingRepo.On("CreateOne", mock.AnythingOfType("context.backgroundCtx"), mock.AnythingOfType("*entities.UserMatching")).
		Return(func(ctx context.Context, matching *entities.UserMatching) (*entities.UserMatching, error) {
			return matching, nil
		})

	_, err := m.matchingService.CreateMatching(context.Background(), mocks_data.Matchings[0])
	m.Nil(err)
}

func (m *MatchingServiceTestSuite) TestMatchingService_List() {
	m.matchingRepo.On("ListMatching", mock.AnythingOfType("context.backgroundCtx"), mock.AnythingOfType("int64"),
		mock.AnythingOfType("int"), mock.AnythingOfType("int")).
		Return(func(ctx context.Context, userId int64, page, limt int) ([]*entities.User, error) {
			return []*entities.User{mocks_data.Users[1]}, nil
		})

	matchers, err := m.matchingService.ListMatching(context.Background(), int64(1), 0, 10)
	m.Nil(err)
	m.Len(matchers, 1)
	m.Equal(matchers[0].Id, int64(2))
}

func (m *MatchingServiceTestSuite) TestMatchingService_WhoILike() {
	m.matchingRepo.On("WhoILike", mock.AnythingOfType("context.backgroundCtx"), mock.AnythingOfType("int64")).
		Return(func(ctx context.Context, userId int64) ([]*entities.User, error) {
			return []*entities.User{mocks_data.Users[1]}, nil
		})

	matchers, err := m.matchingService.WhoILike(context.Background(), 1)
	m.Nil(err)
	m.Len(matchers, 1)
	m.Equal(matchers[0].Id, int64(2))
}

func (m *MatchingServiceTestSuite) TestMatchingService_WhoLikeMe() {
	m.matchingRepo.On("WhoLikeMe", mock.AnythingOfType("context.backgroundCtx"), mock.AnythingOfType("int64")).
		Return(func(ctx context.Context, userId int64) ([]*entities.User, error) {
			return mocks_data.Users[1:3], nil
		})

	matchers, err := m.matchingService.WhoLikeMe(context.Background(), 1)
	m.Nil(err)
	m.Len(matchers, 2)
	m.Equal(matchers[1].Id, int64(3))
}

func TestMatchingService(t *testing.T) {
	suite.Run(t, new(MatchingServiceTestSuite))
}
