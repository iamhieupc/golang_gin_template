package controllers

import (
	"hustchihieu/todolist-golang/repositories"
	"hustchihieu/todolist-golang/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserController ...
type CompanyController struct{}

// var database = db.ConnectDB()
var companyRepository = repositories.NewCompanyRepository(database)

func (ctrl CompanyController) FindAllCompany(context *gin.Context) {
	code := http.StatusOK

	companyService := new(services.CompanyService)
	response := companyService.FindAllCompany(*companyRepository)

	if !response.Success {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

// func (ctrl UserController) CreateUser(context *gin.Context) {
// 	var user models.User

// 	err := context.ShouldBindJSON(&user)

// 	if err != nil {
// 		response := helpers.GenerateValidationResponse(err)
// 		context.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	code := http.StatusOK

// 	userService := new(services.UserService)
// 	response := userService.CreateUser(&user, *userRepository)

// 	if !response.Success {
// 		code = http.StatusBadRequest
// 	}

// 	context.JSON(code, response)
// }

// func (ctrl UserController) FindUserById(context *gin.Context) {
// 	id := context.Param("id")
// 	code := http.StatusOK

// 	userService := new(services.UserService)
// 	response := userService.FindUserById(id, *userRepository)

// 	if !response.Success {
// 		code := http.StatusBadRequest
// 		context.JSON(code, response)
// 		return
// 	}

// 	context.JSON(code, response)
// }

// func (ctrl UserController) DeleteUser(context *gin.Context) {
// 	id := context.Param("id")

// 	code := http.StatusOK

// 	userService := new(services.UserService)
// 	response := userService.DeleteUser(id, *userRepository)

// 	if !response.Success {
// 		code := http.StatusBadRequest
// 		context.JSON(code, response)
// 		return
// 	}

// 	context.JSON(code, response)
// }

// func (ctrl UserController) UpdateUser(context *gin.Context) {
// 	id := context.Param("id")

// 	var user models.User

// 	err := context.ShouldBindJSON(&user)

// 	if err != nil {
// 		response := helpers.GenerateValidationResponse(err)
// 		context.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	code := http.StatusOK

// 	userService := new(services.UserService)
// 	response := userService.UpdateUser(id, &user, *userRepository)

// 	if !response.Success {
// 		code := http.StatusBadRequest
// 		context.JSON(code, response)
// 	}

// 	context.JSON(code, response)
// }

// // func (ctrl UserController) FindUserByName(context *gin.Context) {
// // 	name := context.Param("name")

// // 	code := http.StatusOK

// // 	userService := new(services.UserService)
// // 	response := userService.FindUserByName(name, *userRepository)

// // 	if !response.Success {
// // 		code := http.StatusBadRequest
// // 		context.JSON(code, response)
// // 		return
// // 	}

// // 	return context.JSON(code, response)
// // }
