package database

type DB[T any] struct {
	data map[string]T
}

func CreateDB[T any]() DB[T] {
	data := make(map[string]T, 0)
	return DB[T]{data}
}
