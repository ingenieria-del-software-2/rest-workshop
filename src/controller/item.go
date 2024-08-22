package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-workshop/src/models"
	"rest-workshop/src/service"
)

type Item struct {
	service service.Item
}

func (i Item) AddItem(context *gin.Context) {
	var item models.Item
	if err := context.Bind(&item); err != nil {
		context.JSON(http.StatusBadRequest, models.Error{err.Error()})
		return
	} else if item.Id == "" {
		context.JSON(http.StatusBadRequest, models.Error{"missing id"})
		return
	}
	if value, err := i.service.CreateItem(item); err != nil {
		context.JSON(http.StatusConflict, models.Error{"already existed"})
	} else {
		context.JSON(201, value)
	}
}

func CreateControllerItem(service service.Item) Item {
	return Item{service}
}
