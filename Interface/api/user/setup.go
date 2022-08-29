package user

import (
	usecase "ddd-by-go/application/usecase/user"
	query "ddd-by-go/infrastructure/query/user"
	repository "ddd-by-go/infrastructure/repository/user"
	service "ddd-by-go/infrastructure/service/user"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	ur := repository.NewUserRepository()
	uqp := query.NewUserQueryProvider()
	us := service.NewUserDomainService(ur)
	uuc := usecase.NewUserUsecase(us, uqp)
	uh := NewUserHandler(uuc)

	router.GET("users", uh.SearchUserInfo)
	router.POST("users", uh.CreateUser)
	router.PATCH("users", uh.ChangeUserName)
	router.DELETE("users", uh.DeleteUser)
}
