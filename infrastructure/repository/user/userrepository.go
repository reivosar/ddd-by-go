package user

import (
	model "ddd-by-go/domain/model/user"
	repository "ddd-by-go/domain/repository/user"
	db "ddd-by-go/infrastructure/db"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository struct {
}

func NewUserRepository() repository.UserRepository {
	return &UserRepository{}
}

type User struct {
	Id   string `gorm:"primaryKey"`
	Name string
}

func (ur *UserRepository) FindById(userId model.UserId) (*model.User, error) {
	db := db.GetDBConnection()

	var userDto User
	err := db.Where("id = ?", userId.ToNative()).First(&userDto).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &model.User{
		Id:   model.ToUserId(userDto.Id),
		Name: model.ToUserName(userDto.Name),
	}, nil
}

func (ur *UserRepository) FindByName(userName model.UserName) (*model.User, error) {
	db := db.GetDBConnection()

	var userDto User
	err := db.Where("name = ?", userName.ToNative()).First(&userDto).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &model.User{
		Id:   model.ToUserId(userDto.Id),
		Name: model.ToUserName(userDto.Name),
	}, nil
}

func (ur *UserRepository) Save(user model.User) error {
	db := db.GetDBConnection()

	users := User{
		Id:   user.Id.ToNative(),
		Name: user.Name.ToNative(),
	}
	if result := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"name": user.Name.ToNative()}),
	}).Create(&users); result.Error != nil {
		return result.Error
	}

	return nil
}

func (ur *UserRepository) Delete(userId model.UserId) error {
	db := db.GetDBConnection()

	user := User{Id: userId.ToNative()}
	if result := db.Delete(&user); result.Error != nil {
		return result.Error
	}

	return nil
}
