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

type userDto struct {
	id   string
	name string
}

func (ur *UserRepository) FindById(userId model.UserId) (*model.User, error) {
	db := db.GetDBConnection()
	defer db.Close()

	rows, err := db.Query("SELECT id, username FROM users WHERE id = ?", userId.ToNative)
	if err != nil {
		panic(err.Error())
	}

	var userDto userDto
	for rows.Next() {
		err = rows.Scan(&userDto.id, &userDto.name)
		if err != nil {
			panic(err.Error())
		}
	}

	var result = model.ToUser(model.ToUserId(userDto.id), model.ToUserName(userDto.name))
	return &result, nil
}

func (ur *UserRepository) FindByName(userName model.UserName) (*model.User, error) {
	db := db.GetDBConnection()
	defer db.Close()

	rows, err := db.Query("SELECT id, username FROM users WHERE name = ?", userName.ToNative)
	if err != nil {
		panic(err.Error())
	}

	var userDto userDto
	for rows.Next() {
		err = rows.Scan(&userDto.id, &userDto.name)
		if err != nil {
			panic(err.Error())
		}
	}

	var result = model.ToUser(model.ToUserId(userDto.id), model.ToUserName(userDto.name))
	return &result, nil
}

func (ur *UserRepository) Save(user model.User) error {
	db := db.GetDBConnection()
	defer db.Close()

	ins, err := db.Prepare("INSERT INTO users (`id`, `name`) VALUES (?, ?) ON DUPLICATE KEY UPDATE name=VALUES(name)")
	if err != nil {
		return err
	}

	ins.Exec(user.Id.ToNative, user.Name.ToNative)

	return nil
}

func (ur *UserRepository) Delete(userId model.UserId) error {
	db := db.GetDBConnection()
	defer db.Close()

	ins, err := db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}

	ins.Exec(userId.ToNative())

	return nil
}
