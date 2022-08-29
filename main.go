package main

import (
	"ddd-by-go/Interface/api/user"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	router := gin.Default()
	user.SetupRouter(router)
	router.Run(":" + os.Getenv("PORT_IN_ADDRESS"))
}

func loadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
}
