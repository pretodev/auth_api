package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"main/controller"
)

func main() {
	server := gin.Default()
	controllers := controller.Controllers{}
	server.Use(cors.Default())
	server.POST("/api/users", controllers.CreateUser)
	server.POST("/api/auth/token", controllers.CreateToken)
	server.GET("/api/users/profile", controller.AuthMiddleware(), controllers.GetProfile)
	server.PUT("/api/users/profile", controller.AuthMiddleware(), controllers.UpdateProfile)
	err := server.Run()
	if err != nil {
		return
	}
}
