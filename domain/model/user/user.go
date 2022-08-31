package user

type User struct {
	Id   UserId
	Name UserName
}

func ToUser(userId UserId, username UserName) User {
	return User{
		Id:   userId,
		Name: username,
	}
}

func NewUser(username UserName) User {
	return User{
		Id:   GenerateUserId(),
		Name: username,
	}
}

func (u *User) ChangeUserName(username UserName) User {
	return User{
		Id:   u.Id,
		Name: username,
	}
}

func (u *User) ToString() string {
	return "userId=" + u.Id.ToString() + ",userName=" + u.Name.ToString()
}
