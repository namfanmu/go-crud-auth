package utils

import (
	"errors"
	"go-auth/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func ExtractToken(c *gin.Context) (*models.Claims, error) {
	token := c.GetHeader("Authorization")
	if token == "" {
		return nil, errors.New("unauthorized, please login")
	}

	actualToken := strings.TrimPrefix(token, "Bearer ")
	claims, err := ParseToken(actualToken)
	if err != nil {
		return nil, errors.New("unauthorized, wrong token")
	}

	return claims, nil
}
