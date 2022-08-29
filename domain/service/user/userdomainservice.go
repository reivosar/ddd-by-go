package user

import "ddd-by-go/domain/model/user"

type UserDomainService interface {
	Get(userId user.UserId) (*user.User, error)
	Exists(userName user.UserName) (bool, error)
	CreateUser(user user.User) error
	ChangeUser(user user.User) error
	DeleteUser(userId user.UserId) error
}
