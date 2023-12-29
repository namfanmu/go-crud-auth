package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name   string  `json:"name"`
	Price  float32 `json:"price"`
	Status int8    `json:"status"`
}
