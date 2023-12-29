package repositories

import (
	"go-auth/configs"
	"go-auth/models"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) FindByUsername(username string) *models.User {
	var user models.User
	configs.DB.Preload("Roles").Where("username = ?", username).First(&user)
	return &user
}

func (ur *UserRepository) CreateUser(user *models.User) {
	configs.DB.Create(&user)
}
