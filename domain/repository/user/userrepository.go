package user

import "ddd-by-go/domain/model/user"

type UserRepository interface {
	FindById(userId user.UserId) (*user.User, error)
	FindByName(userName user.UserName) (*user.User, error)
	Save(user user.User) error
	Delete(userId user.UserId) error
}
