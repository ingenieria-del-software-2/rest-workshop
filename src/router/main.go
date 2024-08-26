package router

import (
	"github.com/gin-gonic/gin"
	"rest-workshop/src/controller"
	"rest-workshop/src/database"
	"rest-workshop/src/service"
)

func CreateRouter(db *database.DB) *gin.Engine {
	r := gin.Default()

	// Initialize services and controllers
	itemService := service.CreateServiceItem(db)
	itemController := controller.CreateControllerItem(itemService)

	// Define routes
	r.POST("/item", itemController.AddItem)
	r.GET("/item/:id", itemController.GetItem)

	return r
}
