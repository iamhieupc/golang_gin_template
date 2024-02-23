package repositories

import (
	"hustchihieu/todolist-golang/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r UserRepository) FindAll() RepositoryResult {
	var users models.Users
	err := r.db.Find(&users).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: users}
}

func (r UserRepository) Save(user *models.User) RepositoryResult {
	err := r.db.Save(user).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &user}
}

func (r UserRepository) FindOneById(id string) RepositoryResult {
	var user *models.User
	err := r.db.Where(&models.User{ID: id}).Take(&user).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: user}
}

func (r UserRepository) findOneByName(name string) RepositoryResult {
	var user *models.User
	err := r.db.Where(&models.User{Name: name}).Take(&user).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &user}
}

func (r UserRepository) DeleteOneById(id string) RepositoryResult {
	err := r.db.Delete(&models.User{ID: id}).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: nil}
}
