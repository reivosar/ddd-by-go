package user

import (
	usecase "ddd-by-go/application/usecase/user"
	query "ddd-by-go/application/usecase/user/query"
	model "ddd-by-go/domain/model/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	uc usecase.UserUsecase
}

func NewUserHandler(uc usecase.UserUsecase) UserHandler {
	return UserHandler{
		uc: uc,
	}
}

func (uh *UserHandler) SearchUserInfo(c *gin.Context) {
	req := query.UserSearchCriteria{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := uh.uc.SearchUserInfo(req)
	if err != nil {
		c.JSON(getHttpStatusCode(err), gin.H{"error": err.Error()})
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (uh *UserHandler) CreateUser(c *gin.Context) {
	req := CreateUserRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uh.uc.CreateUser(req.UserName); err != nil {
		c.JSON(getHttpStatusCode(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user created."})
}

func (uh *UserHandler) ChangeUserName(c *gin.Context) {
	req := ChangeUserRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uh.uc.ChangeUserName(req.UserId, req.UserName); err != nil {
		c.JSON(getHttpStatusCode(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user name changed."})
}

func (uh *UserHandler) DeleteUser(c *gin.Context) {
	req := DeleteUserReuqest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uh.uc.DeleteUser(req.UserId); err != nil {
		c.JSON(getHttpStatusCode(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted."})
}

func getHttpStatusCode(err error) int {
	var result = http.StatusInternalServerError
	switch err {
	case model.ErrUserAlreadyExists:
		result = http.StatusConflict
	case model.ErrUserNotExists:
		result = http.StatusNotFound
	}
	return result
}
