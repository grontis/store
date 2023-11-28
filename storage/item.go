package storage

import "github.com/google/uuid"

type Item struct {
	Id          uuid.UUID
	Name        string
	Price       float64
	Description string
	Tags        []string
}
