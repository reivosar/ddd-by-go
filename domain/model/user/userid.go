package user

import (
	"github.com/google/uuid"
)

type UserId struct {
	value string
}

func GenerateUserId() UserId {
	return UserId{
		value: uuid.NewString(),
	}
}

func ToUserId(value string) UserId {
	return UserId{
		value: value,
	}
}

func (u *UserId) ToNative() string {
	return u.value
}

func (u *UserId) ToString() string {
	return u.value
}
