package models

type Item struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Error struct {
	Message string `json:"message"`
}
