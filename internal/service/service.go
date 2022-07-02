package service

import (
	"test-project/internal/model"
	"test-project/internal/repository"
)

type Service struct {
	Identifier
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Identifier: NewIdentifierService(repos.Identifier),
	}
}

type Identifier interface {
	CreateRandom(identifier model.Identifier) (string, error)
	GetDescription(stringIdentifier string) (string, error)
}
