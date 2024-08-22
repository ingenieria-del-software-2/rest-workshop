package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.POST("/item", func(context *gin.Context) {
		context.JSON(200, struct{}{})
	})
	if err := r.Run(); err != nil {
		panic(err.Error())
	}
}
