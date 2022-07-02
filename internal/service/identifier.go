package service

import (
	"test-project/internal/model"
	"test-project/internal/repository"
)

type IdentifierService struct {
	repo repository.Identifier
}

func NewIdentifierService(repo repository.Identifier) *IdentifierService {
	return &IdentifierService{repo: repo}
}

func (s *IdentifierService) CreateRandom(identifier model.Identifier) (string, error) {
	return s.repo.CreateRandom(identifier)
}

func (s *IdentifierService) GetDescription(stringIdentifier string) (string, error) {
	return s.repo.GetDescription(stringIdentifier)
}
