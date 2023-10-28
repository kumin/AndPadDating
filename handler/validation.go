package handler

import (
	"errors"

	"github.com/kumin/AndPadDating/entities"
	"github.com/kumin/AndPadDating/pkg/strings"
)

func ValidateCreateUser(user *entities.User) error {
	if user == nil {
		return errors.New("user is nil")
	}
	if strings.IsEmpty(user.Username) {
		return errors.New("Username is missing")
	}
	return nil
}

func ValidateCreateMatching(matching *entities.UserMatching) error {
	if matching == nil {
		return errors.New("matching is nil")
	}
	return nil
}
