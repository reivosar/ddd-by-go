package user

type User struct {
	id   UserId
	name UserName
}

func NewUser(username UserName) User {
	return User{
		id:   GenerateUserId(),
		name: username,
	}
}

func (u *User) ChangeUserName(username UserName) User {
	return User{
		id:   u.id,
		name: username,
	}
}

func (u *User) ToString() string {
	return "userId=" + u.id.value + ",userName=" + u.name.value
}
