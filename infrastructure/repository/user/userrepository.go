package user

import (
	model "ddd-by-go/domain/model/user"
	repository "ddd-by-go/domain/repository/user"
	db "ddd-by-go/infrastructure/db"
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
	result := db.Where("id = ?", userId.ToNative).First(&userDto)

	if result.Error != nil {
		return nil, result.Error
	}

	var user = model.ToUser(model.ToUserId(userDto.Id), model.ToUserName(userDto.Name))
	return &user, nil
}

func (ur *UserRepository) FindByName(userName model.UserName) (*model.User, error) {
	db := db.GetDBConnection()

	var userDto User
	result := db.Where("name = ?", userName.ToNative).First(&userDto)
	if result.Error != nil {
		return nil, result.Error
	}

	var user = model.ToUser(model.ToUserId(userDto.Id), model.ToUserName(userDto.Name))
	return &user, nil
}

func (ur *UserRepository) Save(user model.User) error {
	db := db.GetDBConnection()

	if result := db.Exec("INSERT INTO users (`id`, `name`) VALUES (?, ?) ON DUPLICATE KEY UPDATE name=VALUES(name)"); result.Error != nil {
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
