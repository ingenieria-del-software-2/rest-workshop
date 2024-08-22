package service

import (
	"rest-workshop/src/database"
	"rest-workshop/src/models"
)

type Item struct {
	db database.DB[models.Item]
}
