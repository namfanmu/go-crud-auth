package repositories

import (
	"go-auth/configs"
	"go-auth/models"
)

var ITEM_STATUS_AVAILABLE int = 1

type ItemRepository struct{}

func NewItemRepository() *ItemRepository {
	return &ItemRepository{}
}

func (ir *ItemRepository) CreateItem(item *models.Item) error {
	return configs.DB.Create(&item).Error
}

func (ir *ItemRepository) FindItemAvailable() []models.Item {
	var items []models.Item
	configs.DB.Where("status = ?", ITEM_STATUS_AVAILABLE).Find(&items)
	return items
}

func (ir *ItemRepository) FindItemById(id int) *models.Item {
	var item models.Item
	configs.DB.First(&item, id)
	return &item
}

func (ir *ItemRepository) UpdateItem(id int, item *models.Item) error {
	return configs.DB.Where("id = ?", id).Updates(&item).Error
}

func (ir *ItemRepository) DeleteItem(id int) error {
	return configs.DB.Model(&models.Item{}).Where("id = ?", id).Update("status", 0).Error
}
