package handler

import (
	"errors"

	"github.com/kumin/BityDating/entities"
	"github.com/kumin/BityDating/pkg/stringx"
)

func ValidateCreateUser(user *entities.User) error {
	if user == nil {
		return errors.New("user is nil")
	}
	if stringx.IsEmpty(user.Username) {
		return errors.New("Username is missing")
	}
	if stringx.IsEmpty(user.Phone) {
		return errors.New("Phone is missing")
	}
	return nil
}

func ValidateCreateMatching(matching *entities.UserMatching) error {
	if matching == nil {
		return errors.New("matching is nil")
	}
	return nil
}
