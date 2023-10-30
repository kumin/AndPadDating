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

type FeedServiceTestSuite struct {
	suite.Suite
	feedService *FeedService
	feedRepo    *mocks_repo.FeedRepo
}

func (f *FeedServiceTestSuite) SetupTest() {
	f.feedRepo = mocks_repo.NewFeedRepo(f.T())
	f.feedService = NewFeedService(f.feedRepo)
}

func (f *FeedServiceTestSuite) TestFeedService_Get() {
	f.feedRepo.On("GetFeed", mock.AnythingOfType("context.backgroundCtx"), mock.AnythingOfType("int64"),
		mock.AnythingOfType("int"), mock.AnythingOfType("int")).
		Return(func(ctx context.Context, userId int64, page, limt int) ([]*entities.User, error) {
			return mocks_data.Users[2:], nil
		})

	users, err := f.feedService.GetFeed(context.Background(), int64(1), 0, 10)
	f.Nil(err)
	f.Len(users, 1)
	f.Equal(users[0].Id, int64(3))
}

func TestFeedService(t *testing.T) {
	suite.Run(t, new(FeedServiceTestSuite))
}
