package controller

import "rest-workshop/src/service"

type Item struct {
	service service.Item
}

func CreateControllerItem(service service.Item) Item {
	return Item{service}
}
