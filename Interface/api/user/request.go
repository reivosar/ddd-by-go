package user

type CreateUserRequest struct {
	UserName string `json:"userName" binding:"required"`
}

type ChangeUserRequest struct {
	UserId   string `json:"userId" binding:"required"`
	UserName string `json:"userName" binding:"required"`
}

type DeleteUserReuqest struct {
	UserId string `json:"userId" binding:"required"`
}
