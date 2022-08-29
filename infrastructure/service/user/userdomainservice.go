package user

import (
	model "ddd-by-go/domain/model/user"
	repository "ddd-by-go/domain/repository/user"
	service "ddd-by-go/domain/service/user"
)

type UserDomainService struct {
	r repository.UserRepository
}

func NewUserDomainService(r repository.UserRepository) service.UserDomainService {
	return &UserDomainService{
		r: r,
	}
}

func (s *UserDomainService) Get(userId model.UserId) (*model.User, error) {
	user, err := s.r.FindById(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserDomainService) Exists(userName model.UserName) (bool, error) {
	user, err := s.r.FindByName(userName)
	if err != nil {
		return false, err
	}
	return user != nil, nil
}

func (s *UserDomainService) CreateUser(user model.User) error {
	if err := s.r.Save(user); err != nil {
		return err
	}
	return nil
}

func (s *UserDomainService) ChangeUser(user model.User) error {
	if err := s.r.Save(user); err != nil {
		return err
	}
	return nil
}

func (s *UserDomainService) DeleteUser(userId model.UserId) error {
	if err := s.r.Delete(userId); err != nil {
		return err
	}
	return nil
}
