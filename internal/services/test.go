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

func (t *TestService) GetTestBySubjectId(subID uint) (models.Test, error) {
	var test models.Test

	subject, err := t.rep.GetSubjectById(subID)
	questions, err := t.rep.GetRandomQuestions(subID)
	if err != nil {
		return models.Test{}, err
	}
	//TODO handle error
	if err != nil {
		return test, err
	}
	test.Subject = subject
	test.Questions = questions
	return test, nil
}
