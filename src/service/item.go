package service

import (
	"fmt"
	"rest-workshop/src/database"
	"rest-workshop/src/models"
)

type Item struct {
	db *database.DB
}

func (i *Item) CreateItem(data models.Item) (models.Item, error) {

	err := i.db.Create(data)
	if err != nil {
		return models.Item{}, fmt.Errorf("failed to create item: %w", err)
	}
	return data, nil
}

func (i *Item) GetItem(id int) (models.Item, error) {
	var item models.Item
	err := i.db.GetItem(id, &item)
	if err != nil {
		return models.Item{}, fmt.Errorf("item not found: %w", err)
	}
	return item, nil
}

func CreateServiceItem(db *database.DB) *Item {
	return &Item{db}
}
