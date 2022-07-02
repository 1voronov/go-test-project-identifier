package repository

import (
	"test-project/internal/model"

	"gorm.io/gorm"
)

type IdentifierPostgres struct {
	db *gorm.DB
}

func NewIdentifierPostgres(db *gorm.DB) *IdentifierPostgres {
	return &IdentifierPostgres{db: db}
}

func (r *IdentifierPostgres) CreateRandom(identifier model.Identifier) (string, error) {
	result := r.db.Select("string_identifier", "description").Create(&identifier)
	if result.Error != nil {
		return "", result.Error
	}

	string_identifier := identifier.StringIdentifier

	return string_identifier, result.Error
}

func (r *IdentifierPostgres) GetDescription(stringIdentifier string) (string, error) {
	var identifier model.Identifier

	result := r.db.Where("string_identifier = ?", stringIdentifier).First(&identifier)
	// .Where("string_identifier = ?", stringIdentifier)
	// .Model(&model.Identifier{StringIdentifier: stringIdentifier})
	if result.Error != nil {
		return "", result.Error
	}

	description := identifier.Description

	return description, result.Error
}
