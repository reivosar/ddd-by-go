package user

import (
	"ddd-by-go/application/usecase/user/query"
)

type UserQueryProvider struct {
	// Need Dao
}

func NewUserQueryProvider() query.UserQueryProvider {
	return &UserQueryProvider{}
}

func (uqr *UserQueryProvider) SearchUserInfo(criteria query.UserSearchCriteria) (*[]query.UserSearchResult, error) {
	return nil, nil
}
