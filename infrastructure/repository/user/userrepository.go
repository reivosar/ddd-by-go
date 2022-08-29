package user

import (
	model "ddd-by-go/domain/model/user"
	repository "ddd-by-go/domain/repository/user"
)

type UserRepository struct {
	// Need Dao
}

func NewUserRepository() repository.UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) FindById(userId model.UserId) (*model.User, error) {
	return nil, nil
}

func (ur *UserRepository) FindByName(userName model.UserName) (*model.User, error) {
	return nil, nil
}

func (ur *UserRepository) Save(user model.User) error {
	return nil
}

func (ur *UserRepository) Delete(userId model.UserId) error {
	return nil
}
