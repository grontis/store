package models

import "github.com/google/uuid"

type IdGetter interface {
	Item | User
	GetId() uuid.UUID
}
