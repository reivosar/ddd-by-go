package user

import (
	query "ddd-by-go/application/usecase/user/query"
	model "ddd-by-go/domain/model/user"
	service "ddd-by-go/domain/service/user"
)

type UserUsecase struct {
	uds service.UserDomainService
	uqr query.UserQueryProvider
}

func NewUserUsecase(uds service.UserDomainService, uqr query.UserQueryProvider) UserUsecase {
	return UserUsecase{
		uds: uds,
		uqr: uqr,
	}
}

func (uc *UserUsecase) SearchUserInfo(criteria query.UserSearchCriteria) (*[]query.UserSearchResult, error) {
	results, error := uc.uqr.SearchUserInfo(criteria)
	if error != nil {
		return nil, error
	}
	return results, nil
}

func (uc *UserUsecase) CreateUser(name string) error {
	userName := model.ToUserName(name)
	b, error := uc.uds.Exists(userName)
	if error != nil {
		return error
	}
	if b {
		return model.ErrUserAlreadyExists
	}

	if error := uc.uds.CreateUser(model.NewUser(userName)); error != nil {
		return error
	}

	return nil
}

func (uc *UserUsecase) ChangeUserName(id string, name string) error {
	user, error := uc.uds.Get(model.ToUserId(id))
	if error != nil {
		return error
	}
	if user == nil {
		return model.ErrUserNotExists
	}

	if error := uc.uds.ChangeUser(user.ChangeUserName(model.ToUserName(name))); error != nil {
		return error
	}

	return nil
}

func (uc *UserUsecase) DeleteUser(id string) error {
	userId := model.ToUserId(id)
	user, error := uc.uds.Get(userId)
	if error != nil {
		return error
	}
	if user == nil {
		return model.ErrUserNotExists
	}

	if error := uc.uds.DeleteUser(userId); error != nil {
		return error
	}

	return nil
}
