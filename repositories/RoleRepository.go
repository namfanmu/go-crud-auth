package repositories

import (
	"go-auth/configs"
	"go-auth/models"
)

type RoleRepository struct{}

func NewRoleRepository() *RoleRepository {
	return &RoleRepository{}
}

func (rr *RoleRepository) FindByID(id int) *models.Role {
	var role models.Role
	configs.DB.Where("id = ?", id).First(&role)
	return &role
}
