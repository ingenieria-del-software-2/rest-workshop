package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-workshop/src/models"
	"rest-workshop/src/service"
	"strconv"
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

func (i *Item) GetItem(context *gin.Context) {
	idParam := context.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		context.JSON(http.StatusBadRequest, models.Error{Message: "invalid item ID"})
		return
	}

	item, err := i.service.GetItem(id)
	if err != nil {
		context.JSON(http.StatusNotFound, models.Error{Message: "item not found"})
		return
	}

	context.JSON(http.StatusOK, item)
}

func CreateControllerItem(service *service.Item) *Item {
	return &Item{service}
}
