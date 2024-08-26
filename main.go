package main

import (
	"fmt"
	"rest-workshop/src/router"
)

func main() {
	config := LoadConfig()

	r := router.CreateRouter()

	address := fmt.Sprintf("%s:%s", config.Host, config.Port)
	if err := r.Run(address); err != nil {
		panic(err.Error())
	}
}
