package services

import (
	"go-auth/models"
	"go-auth/repositories"
	"go-auth/utils"

	"github.com/gin-gonic/gin"
)

var ROLE_USER int = 3

type UserService struct {
	userRepository repositories.UserRepository
	roleRepository repositories.RoleRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepository: *repositories.NewUserRepository(),
		roleRepository: *repositories.NewRoleRepository(),
	}
}

func (us *UserService) Login(c *gin.Context, userLogin *models.UserLoginRequest) (string, *models.APIResponse) {
	existingUser := us.userRepository.FindByUsername(userLogin.Username)
	if existingUser.ID == 0 {
		return "", &models.APIResponse{Status: 400, Message: "username invalid"}
	}

	errHash := utils.CompareHashPassword(userLogin.Password, existingUser.Password)
	if !errHash {
		return "", &models.APIResponse{Status: 400, Message: "password invalid"}
	}
	token, err := utils.GenerateJWTToken(*existingUser)
	if err != nil {
		return "", &models.APIResponse{Status: 500, Message: "could not generate token"}
	}
	return token, nil
}

func (us *UserService) Signup(c *gin.Context, user *models.User) *models.APIResponse {
	if len(user.Username) == 0 {
		return &models.APIResponse{Status: 500, Message: "username can not be empty"}
	}

	if len(user.Password) <= 5 {
		return &models.APIResponse{Status: 500, Message: "password must be length than 5 charactors"}
	}

	existingUser := us.userRepository.FindByUsername(user.Username)
	if existingUser.ID != 0 {
		return &models.APIResponse{Status: 400, Message: "user already exist"}
	}
	var errHash error
	user.Password, errHash = utils.GenerateHashPassword(user.Password)
	if errHash != nil {
		return &models.APIResponse{Status: 500, Message: "could not generate password hash"}
	}
	//Add role
	role := us.roleRepository.FindByID(ROLE_USER)
	user.Roles = append(user.Roles, role)

	us.userRepository.CreateUser(user)
	return &models.APIResponse{Status: 200, Message: "user created"}
}
