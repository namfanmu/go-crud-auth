package utils

import (
	"go-auth/models"

	"github.com/gin-gonic/gin"
)

func CheckJSON(c *gin.Context, input interface{}) *models.APIResponse {
	if err := c.ShouldBindJSON(&input); err != nil {
		return &models.APIResponse{Status: 500, Message: err.Error()}
	}
	return &models.APIResponse{Status: 200, Message: "JSON in the right format"}
}
