package user

import "errors"

var (
	ErrUserAlreadyExists = errors.New("User already exists.")
	ErrUserNotExists     = errors.New("User does not exist.")
)
