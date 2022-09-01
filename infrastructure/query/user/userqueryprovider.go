package user

import (
	"ddd-by-go/application/usecase/user/query"
	"ddd-by-go/infrastructure/db"
	"errors"

	"gorm.io/gorm"
)

type UserQueryProvider struct {
}

func NewUserQueryProvider() query.UserQueryProvider {
	return &UserQueryProvider{}
}

type User struct {
	Id   string `gorm:"primaryKey"`
	Name string
}

func (uqr *UserQueryProvider) SearchUserInfo(criteria query.UserSearchCriteria) (*[]query.UserSearchResult, error) {
	db := db.GetDBConnection()

	if criteria.Id != "" {
		db.Where("id = ?", criteria.Id)
	}
	if criteria.Name != "" {
		db.Where("name = ?", criteria.Name)
	}

	results := []query.UserSearchResult{}
	err := db.Table("users").Find(&results).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &results, nil
}
