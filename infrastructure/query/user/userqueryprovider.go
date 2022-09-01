package user

import (
	"ddd-by-go/application/usecase/user/query"
	"ddd-by-go/infrastructure/db"
)

type UserQueryProvider struct {
}

func NewUserQueryProvider() query.UserQueryProvider {
	return &UserQueryProvider{}
}

func (uqr *UserQueryProvider) SearchUserInfo(criteria query.UserSearchCriteria) (*[]query.UserSearchResult, error) {
	db := db.GetDBConnection()

	if criteria.UserId != "" {
		db.Where("id = ?", criteria.UserId)
	}
	if criteria.UserName != "" {
		db.Where("name = ?", criteria.UserName)
	}

	userSearchResults := []query.UserSearchResult{}
	if result := db.Find(&userSearchResults); result.Error != nil {
		return nil, result.Error
	}
	return &userSearchResults, nil
}
