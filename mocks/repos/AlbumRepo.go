// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	entities "github.com/kumin/BityDating/entities"
	mock "github.com/stretchr/testify/mock"
)

// AlbumRepo is an autogenerated mock type for the AlbumRepo type
type AlbumRepo struct {
	mock.Mock
}

// CreateMany provides a mock function with given fields: ctx, imageFiles
func (_m *AlbumRepo) CreateMany(ctx context.Context, imageFiles []*entities.File) ([]*entities.Image, error) {
	ret := _m.Called(ctx, imageFiles)

	var r0 []*entities.Image
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []*entities.File) ([]*entities.Image, error)); ok {
		return rf(ctx, imageFiles)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []*entities.File) []*entities.Image); ok {
		r0 = rf(ctx, imageFiles)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.Image)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []*entities.File) error); ok {
		r1 = rf(ctx, imageFiles)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateOne provides a mock function with given fields: ctx, imageFile
func (_m *AlbumRepo) CreateOne(ctx context.Context, imageFile *entities.File) (*entities.Image, error) {
	ret := _m.Called(ctx, imageFile)

	var r0 *entities.Image
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.File) (*entities.Image, error)); ok {
		return rf(ctx, imageFile)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entities.File) *entities.Image); ok {
		r0 = rf(ctx, imageFile)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Image)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entities.File) error); ok {
		r1 = rf(ctx, imageFile)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserAlbum provides a mock function with given fields: ctx
func (_m *AlbumRepo) GetUserAlbum(ctx context.Context) ([]*entities.Image, error) {
	ret := _m.Called(ctx)

	var r0 []*entities.Image
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*entities.Image, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*entities.Image); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.Image)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAlbumRepo creates a new instance of AlbumRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAlbumRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *AlbumRepo {
	mock := &AlbumRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
