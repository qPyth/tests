package services

import (
	"tests/internal/models"
	r "tests/internal/repository"
)

type EntTest interface {
	GetTest(mathLitID, kazHistoryID, readingLit, profile1, profile2 int) (EntTestOutput, error)
}

type EntTestOutput struct {
	Tests map[string][]models.Question
}

type Services struct {
	EntTest
}

func NewServices(rep *r.Repository) *Services {
	TestService := NewEntTestService(rep)

	return &Services{
		EntTest: TestService,
	}
}
