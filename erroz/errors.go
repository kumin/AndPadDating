package erroz

import "errors"

var (
	ErrPhoneAlreadyUsed = errors.New("phone already used")
	ErrPhoneIsMissing   = errors.New("phone is missing")
	ErrBadToken         = errors.New("bad token")
)
