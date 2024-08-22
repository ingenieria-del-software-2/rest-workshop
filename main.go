package main

import (
	"rest-workshop/src/router"
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
	r := router.CreateRouter()
	if err := r.Run(); err != nil {
		panic(err.Error())
	}
}
