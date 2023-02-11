package services

import (
	"hustchihieu/todolist-golang/repositories"
	"hustchihieu/todolist-golang/dtos"
	"hustchihieu/todolist-golang/models"
	// "fmt"
)

type UserService struct {}

func (svc *UserService) FindAllUsers(repository repositories.UserRepository) dtos.Response {
	operationResult := repository.FindAll()

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*models.Users)

	return dtos.Response{Success: true, Data: data}
}