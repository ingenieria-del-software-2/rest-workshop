package service

import (
	"rest-workshop/src/database"
	"rest-workshop/src/models"
)

type Item struct {
	db database.DB[models.Item]
}

func (i Item) CreateItem(data models.Item) (models.Item, error) {
	return i.db.Create(data, data.Id)
}

func CreateServiceItem(db database.DB[models.Item]) Item {
	return Item{db}
}
