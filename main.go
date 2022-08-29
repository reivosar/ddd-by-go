package main

import (
	"ddd-by-go/Interface/api/user"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	user.SetupRouter(router)
	router.Run("8081")
}
