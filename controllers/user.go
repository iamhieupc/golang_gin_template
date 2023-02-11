package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"hustchihieu/todolist-golang/services"
	"hustchihieu/todolist-golang/repositories"
	db "hustchihieu/todolist-golang/database"
)

//UserController ...
type UserController struct{}

var database = db.ConnectDB()
var userRepository = repositories.NewUserRepository(database)

func (ctrl UserController) FindAllUsers(context *gin.Context) {
	code := http.StatusOK
	
	userService := new(services.UserService)
	response := userService.FindAllUsers(*userRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}