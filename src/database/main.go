package database

import "errors"

type DB[T any] struct {
	data map[string]T
}

func (db DB[T]) Create(data T, key string) (T, error) {
	var t T
	if _, found := db.data[key]; found {
		return t, errors.New("key not found")
	}
	db.data[key] = data
	return data, nil
}

func CreateDB[T any]() DB[T] {
	data := make(map[string]T, 0)
	return DB[T]{data}
}
