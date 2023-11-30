package models

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID
	Username string
	Password string
}

func (u User) GetId() uuid.UUID {
	return u.Id
}
