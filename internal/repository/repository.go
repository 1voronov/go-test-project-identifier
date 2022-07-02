package repository

import (
	"test-project/internal/model"

	"gorm.io/gorm"
)

type Repository struct {
	Identifier
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Identifier: NewIdentifierPostgres(db),
	}
}

type Identifier interface {
	CreateRandom(identifier model.Identifier) (string, error)
	GetDescription(stringIdentifier string) (string, error)
}
