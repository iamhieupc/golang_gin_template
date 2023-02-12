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
		api.GET("/user/:id", userController.FindUserById)
		// api.GET("/user/:name", userController.FindUserByName)
		api.POST("/user/", userController.CreateUser)
		api.DELETE("/user/:id", userController.DeleteUser)
		api.PUT("/user/:id", userController.UpdateUser)
	}

	return route
}