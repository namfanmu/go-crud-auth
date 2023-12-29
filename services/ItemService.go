package services

import (
	"go-auth/models"
	"go-auth/repositories"
	"go-auth/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ItemService struct {
	itemRepository repositories.ItemRepository
}

func NewItemService() *ItemService {
	return &ItemService{
		itemRepository: *repositories.NewItemRepository(),
	}
}

func (is *ItemService) FindAllItem() []models.Item {
	return is.itemRepository.FindItemAvailable()
}

func (is *ItemService) CreateItem(c *gin.Context, item *models.Item) *models.APIResponse {

	claims, errToken := utils.ExtractToken(c)
	if errToken != nil || claims.Role != "ROLE_ADMIN" {
		return &models.APIResponse{Status: 401, Message: "unauthorized, do not have permission"}
	}

	item.Status = 1
	if err := is.itemRepository.CreateItem(item); err != nil {
		return &models.APIResponse{Status: 400, Message: "failed to create item"}
	}

	return &models.APIResponse{Status: 200, Message: "item created"}
}

func (is *ItemService) UpdateItem(c *gin.Context, item *models.Item) *models.APIResponse {
	claims, errToken := utils.ExtractToken(c)
	if errToken != nil || (claims.Role != "ROLE_ADMIN" && claims.Role != "ROLE_EDITOR") {
		return &models.APIResponse{Status: 401, Message: "unauthorized, do not have permission"}
	}

	itemId := c.Param("id")
	intItemId, err := strconv.Atoi(itemId)
	if err != nil {
		return &models.APIResponse{Status: 500, Message: "incorrect id"}
	}
	checkedItem := *is.itemRepository.FindItemById(intItemId)
	if checkedItem.ID == 0 {
		return &models.APIResponse{Status: 500, Message: "item not found"}
	}
	if err := is.itemRepository.UpdateItem(intItemId, item); err != nil {
		return &models.APIResponse{Status: 500, Message: "failed to update item"}
	}
	return &models.APIResponse{Status: 200, Message: "item updated"}
}

func (is *ItemService) DeleteItem(c *gin.Context, id int) *models.APIResponse {
	claims, errToken := utils.ExtractToken(c)
	if errToken != nil || claims.Role != "ROLE_ADMIN" {
		return &models.APIResponse{Status: 401, Message: "unauthorized, do not have permission"}
	}

	checkedItem := is.itemRepository.FindItemById(id)
	if checkedItem.ID == 0 {
		return &models.APIResponse{Status: 500, Message: "item not found"}
	}
	if err := is.itemRepository.DeleteItem(id); err != nil {
		return &models.APIResponse{Status: 500, Message: "failed to delete item"}
	}
	return &models.APIResponse{Status: 200, Message: "item deleted"}
}
