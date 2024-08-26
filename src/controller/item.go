package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-workshop/src/models"
	"rest-workshop/src/service"
)

type Item struct {
	service *service.Item
}

func (i *Item) AddItem(context *gin.Context) {
	var item models.Item
	if err := context.Bind(&item); err != nil {
		context.JSON(http.StatusBadRequest, models.Error{Message: err.Error()})
		return
	}

	if value, err := i.service.CreateItem(item); err != nil {
		context.JSON(http.StatusConflict, models.Error{Message: "failed to create item"})
	} else {
		context.JSON(http.StatusCreated, value)
	}
}

func CreateControllerItem(service *service.Item) *Item {
	return &Item{service}
}
