package configs

import (
	"hustchihieu/todolist-golang/controllers"

	"github.com/gin-gonic/gin"
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

	companyController := new(controllers.CompanyController)
	{
		api.GET("/company/", companyController.FindAllCompany)
	}

	return route
}
