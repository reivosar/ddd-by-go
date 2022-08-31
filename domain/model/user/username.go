package user

type UserName struct {
	value string
}

func ToUserName(value string) UserName {
	return UserName{
		value: value,
	}
}

func (u *UserName) ToNative() string {
	return u.value
}

func (u *UserName) ToString() string {
	return u.value
}
