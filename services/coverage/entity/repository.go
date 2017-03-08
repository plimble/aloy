package entity

import (
	"time"
)

type Repository struct {
	Id          string
	Name        string
	Description string
	HomePage    string
	CreatedAt   string
}

func NewRepository(id, name, description, homepage string) *Repository {
	return &Repository{
		Id:          id,
		Name:        name,
		Description: description,
		HomePage:    homepage,
		CreatedAt:   time.Now().Truncate(time.Second).Format(time.RFC3339),
	}
}
