package storage

import (
	"grontis/store/models"

	"github.com/google/uuid"
)

type Storage interface {
	GetItems() ([]models.Item, error)
	GetItem(id uuid.UUID) (models.Item, error)
	CreateItem(item models.Item) (models.Item, error)
	UpdateItem(item models.Item) (models.Item, error)
	DeleteItem(id uuid.UUID) error

	GetUsers() ([]models.User, error)
	GetUser(id uuid.UUID) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(id uuid.UUID) error
}
