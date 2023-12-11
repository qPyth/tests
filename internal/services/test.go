package services

import (
	"tests/internal/models"
	r "tests/internal/repository"
)

type TestService struct {
	rep r.Repository
}

func NewTestService(rep r.Repository) *TestService {
	return &TestService{
		rep,
	}
}

func (t *TestService) GetTestBySubjectId(subID int) (models.Test, error) {
	var test models.Test

	subject, err := t.rep.GetSubjectById(subID)
	//TODO handle error
	if err != nil {
		return test, err
	}
	test.Subject = subject
	return test, err
}
