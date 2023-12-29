package configs

import (
	"fmt"
	"go-auth/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(cfg models.DBConnection, modelsToMigrate ...interface{}) {
	dsn := fmt.Sprintf(cfg.Connection)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	for _, model := range modelsToMigrate {
		if err := db.AutoMigrate(model); err != nil {
			panic(err)
		}
	}
	fmt.Println("Migrated database")
	DB = db
}
