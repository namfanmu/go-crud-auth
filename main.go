package main

import (
	"go-auth/configs"
	"go-auth/models"
	"go-auth/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config := models.DBConnection{
		Connection: os.Getenv("MYSQL_CONNECTION"),
	}
	modelsToMigrate := []interface{}{&models.User{}, &models.Role{}, &models.Item{}}
	configs.InitDB(config, modelsToMigrate...)
	routes.Routes(r)

	r.Run(":8080")
}
