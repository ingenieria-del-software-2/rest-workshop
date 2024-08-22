package database

type DB[T any] struct {
	data map[string]T
}
