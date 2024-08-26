package models

type Item struct {
	Name  string  `json:"name" db:"name"`
	Price float64 `json:"price" db:"price"`
}

type Error struct {
	Message string `json:"message"`
}
