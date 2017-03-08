package entity

import (
	"time"
)

type Repository struct {
	Id          string
	Name        string
	OwnerName   string
	Source      string
	Description string
	HomePage    string
	CreatedAt   string
}

func NewRepository(id, name, ownerName, source, description, homepage string) *Repository {
	return &Repository{
		Id:          id,
		Name:        name,
		OwnerName:   ownerName,
		Source:      source,
		Description: description,
		HomePage:    homepage,
		CreatedAt:   time.Now().Truncate(time.Second).Format(time.RFC3339),
	}
}
