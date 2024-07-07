package database

import "gorm.io/gorm"

type Repository[T any] interface {
	GetAll() ([]T, error)
	Create(item *T) error
}

type GenericRepository[T any] struct {
	DB *gorm.DB
}

func (r *GenericRepository[T]) GetAll() ([]T, error) {
	var items []T
	result := r.DB.Find(&items)
	return items, result.Error
}

func (r *GenericRepository[T]) Create(item *T) error {
	result := r.DB.Create(item)
	return result.Error
}