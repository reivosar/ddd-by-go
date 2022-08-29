package user

import "errors"

var (
	ErrUserAlreadyExists = errors.New("ユーザは既に存在します。")
	ErrUserNotExists     = errors.New("ユーザが存在しません。")
)
