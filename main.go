package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Item struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Error struct {
	Message string `json:"message"`
}

func main() {
	r := gin.Default()
	database := make(map[string]Item, 0)
	r.POST("/item", func(context *gin.Context) {
		var item Item
		if err := context.Bind(&item); err != nil {
			context.JSON(http.StatusBadRequest, Error{err.Error()})
			return
		} else if item.Id == "" {
			context.JSON(http.StatusBadRequest, Error{"missing id"})
			return
		}
		if _, found := database[item.Id]; found {
			context.JSON(http.StatusConflict, Error{"already existed"})
			return
		}
		database[item.Id] = item
		context.JSON(201, item)
	})
	if err := r.Run(); err != nil {
		panic(err.Error())
	}
}
