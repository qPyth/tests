package services

import (
	"tests/internal/models"
	r "tests/internal/repository"
)

type Test interface {
	GetTestBySubjectId(subID uint) (models.Test, error)
}

type Services struct {
	Test
}

func NewServices(rep r.Repository) *Services {
	TestService := NewTestService(rep)

	return &Services{
		Test: TestService,
	}
}
