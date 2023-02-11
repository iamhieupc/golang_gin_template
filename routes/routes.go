package configs

import (
	"github.com/gin-gonic/gin"
	"hustchihieu/todolist-golang/controllers"
)


func SetupRoutes() *gin.Engine {
	route := gin.Default()

	api := route.Group("/api")
	
	userController := new(controllers.UserController)
	{
		api.GET("/user/", userController.FindAllUsers)

	}

	return route
}