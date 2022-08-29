package user

type UserName struct {
	value string
}

func ToUserName(value string) UserName {
	return UserName{
		value: value,
	}
}
