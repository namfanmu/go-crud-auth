package controllers

import (
	"go-auth/models"
	"go-auth/services"
	"go-auth/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

var itemService = services.NewItemService()

func CreateItem(c *gin.Context) {
	var item models.Item
	if resultCheckJSON := utils.CheckJSON(c, &item); resultCheckJSON.Status != 200 {
		c.JSON(resultCheckJSON.Status, gin.H{"error": resultCheckJSON.Message})
		return
	}
	result := itemService.CreateItem(c, &item)
	c.JSON(result.Status, gin.H{"message": result.Message})
}

func GetAllItems(c *gin.Context) {
	items := itemService.FindAllItem()
	c.JSON(200, gin.H{"items": items})
}

func UpdateItem(c *gin.Context) {
	var item models.Item
	if resultCheckJSON := utils.CheckJSON(c, &item); resultCheckJSON.Status != 200 {
		c.JSON(resultCheckJSON.Status, gin.H{"error": resultCheckJSON.Message})
		return
	}
	result := itemService.UpdateItem(c, &item)
	c.JSON(result.Status, gin.H{"message": result.Message})
}

func DeleteItem(c *gin.Context) {
	itemId := c.Param("id")
	intItemId, err := strconv.Atoi(itemId)
	if err != nil {
		c.JSON(500, gin.H{"error": "incorrect id"})
	}
	result := itemService.DeleteItem(c, intItemId)
	c.JSON(result.Status, gin.H{"message": result.Message})
}
