package repositories

import (
	"gorm.io/gorm"
	"hustchihieu/todolist-golang/models"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindAll() RepositoryResult {
	var users models.Users
	err := r.db.Find(&users).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &users}
}