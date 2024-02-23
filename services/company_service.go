package services

import (
	"hustchihieu/todolist-golang/dtos"
	"hustchihieu/todolist-golang/repositories"
	// "fmt"
)

type CompanyService struct{}

func (svc CompanyService) FindAllCompany(repository repositories.CompanyRepository) dtos.Response {
	operationResult := repository.FindAll()

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result

	return dtos.Response{Success: true, Data: data}
}

// func (svc UserService) CreateUser(user *models.User, repository repositories.UserRepository) dtos.Response {
// 	uuidResult, err := uuid.NewRandom()

// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	user.ID = uuidResult.String()

// 	operationResult := repository.Save(user)

// 	if operationResult.Error != nil {
// 		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
// 	}

// 	var data = operationResult.Result

// 	return dtos.Response{Success: true, Data: data}
// }

// func (svc UserService) FindUserById(id string, repository repositories.UserRepository) dtos.Response {
// 	operationResult := repository.FindOneById(id)

// 	if operationResult.Error != nil {
// 		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
// 	}
// 	var data = operationResult.Result

// 	return dtos.Response{Success: true, Data: data}
// }

// func (svc UserService) DeleteUser(id string, userRepository repositories.UserRepository) dtos.Response {
// 	operationResult := userRepository.DeleteOneById(id)

// 	if operationResult.Error != nil {
// 		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
// 	}

// 	var data = operationResult.Result

// 	return dtos.Response{Success: true, Data: data}
// }

// func (svc UserService) UpdateUser(id string, user *models.User, userRepository repositories.UserRepository) dtos.Response {
// 	existingUser := userRepository.FindOneById(id)

// 	if existingUser.Error != nil {
// 		return dtos.Response{Success: false, Message: existingUser.Error.Error()}
// 	}

// 	var data = existingUser.Result.(*models.User)

// 	if user.Name != "" {
// 		data.Name = user.Name
// 	}

// 	if user.Email != "" {
// 		data.Email = user.Email
// 	}

// 	if user.Address != "" {
// 		data.Address = user.Address
// 	}

// 	operationResult := userRepository.Save(data)

// 	if operationResult.Error != nil {
// 		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
// 	}

// 	return dtos.Response{Success: true, Data: operationResult.Result}
// }

// // func (svc UserService) FindUserByName(name string, userRepository repositories.UserRepository) dtos.Response {

// // }
