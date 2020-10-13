package main

import (
	"userApi/controller"
	"userApi/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.NewDB()
	defer database.DB.Close()

	server := gin.Default()

	server.GET("/users", controller.FindAll)
	server.GET("/users/:id",controller.FindUserbyID)

	server.Run()

}
