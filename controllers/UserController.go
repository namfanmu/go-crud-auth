package controllers

import (
	"go-auth/models"
	"go-auth/services"
	"go-auth/utils"

	"github.com/gin-gonic/gin"
)

var userService = services.NewUserService()

func Signup(c *gin.Context) {
	var user models.User
	if resultCheckJSON := utils.CheckJSON(c, &user); resultCheckJSON.Status != 200 {
		c.JSON(resultCheckJSON.Status, gin.H{"error": resultCheckJSON.Message})
		return
	}
	result := userService.Signup(c, &user)
	if result.Status != 200 {
		c.JSON(result.Status, gin.H{"error": result.Message})
		return
	}
	c.JSON(result.Status, gin.H{"success": result.Message})
}

func Login(c *gin.Context) {
	var userLogin models.UserLoginRequest
	if resultCheckJSON := utils.CheckJSON(c, &userLogin); resultCheckJSON.Status != 200 {
		c.JSON(resultCheckJSON.Status, gin.H{"error": resultCheckJSON.Message})
		return
	}
	token, err := userService.Login(c, &userLogin)
	if err != nil {
		c.JSON(err.Status, gin.H{"error": err.Message})
		return
	}
	c.JSON(200, gin.H{"token": token, "success": "login success"})
}
