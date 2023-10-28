package mocks

import (
	"time"

	"github.com/kumin/AndPadDating/entities"
)

var (
	Users = []*entities.User{
		{
			Id:       1,
			Username: "phuong thao",
			Phone:    "03670187576",
			Email:    "test9@.com",
			BirthDay: time.Now(),
			Gender:   "female",
			IsActive: 1,
		},
		{
			Id:       2,
			Username: "nhat linh",
			Phone:    "0367018751",
			Email:    "test1@.com",
			BirthDay: time.Now(),
			Gender:   "female",
			IsActive: 1,
		},
		{
			Id:       3,
			Username: "kumin",
			Phone:    "0367018752",
			Email:    "test2@.com",
			BirthDay: time.Now(),
			Gender:   "female",
			IsActive: 1,
		},
	}
	Matchings = []*entities.UserMatching{}
)
