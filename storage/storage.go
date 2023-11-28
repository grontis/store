package storage

import (
	"github.com/google/uuid"
)

type Storage interface {
	GetItems() ([]Item, error)
	GetItem(id uuid.UUID) (Item, error)
	CreateItem(item Item) (Item, error)
	UpdateItem(item Item) (Item, error)
	DeleteItem(id uuid.UUID) error
}
