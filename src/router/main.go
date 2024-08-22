package router

import (
	"github.com/gin-gonic/gin"
	"rest-workshop/src/controller"
	"rest-workshop/src/database"
	"rest-workshop/src/models"
	"rest-workshop/src/service"
)

func CreateRouter() *gin.Engine {
	r := gin.Default()
	db := database.CreateDB[models.Item]()
	c := controller.CreateControllerItem(service.CreateServiceItem(db))
	r.POST("/item", c.AddItem)
	return r
}
