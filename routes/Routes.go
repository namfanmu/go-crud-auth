package routes

import (
	"go-auth/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {

	users := r.Group("/users")
	{
		users.POST("/signup", controllers.Signup)
		users.POST("/login", controllers.Login)
	}

	items := r.Group("/items")
	{
		items.POST("/create", controllers.CreateItem)
		items.GET("/get-all", controllers.GetAllItems)
		items.PUT("/update-item/:id", controllers.UpdateItem)
		items.DELETE("/delete-item/:id", controllers.DeleteItem)
	}

}
