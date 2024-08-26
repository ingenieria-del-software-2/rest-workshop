package models

type Item struct {
	Id    int     `json:"id,omitempty" db:"id"`
	Name  string  `json:"name" db:"name"`
	Price float64 `json:"price" db:"price"`
}

type Error struct {
	Message string `json:"message"`
}
