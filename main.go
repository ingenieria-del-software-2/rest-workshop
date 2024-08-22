package main

import (
	"rest-workshop/src/router"
)

func main() {
	r := router.CreateRouter()
	if err := r.Run(); err != nil {
		panic(err.Error())
	}
}
