package repository

import "tests/internal/models"

type Repository interface {
	GetSubjectById(id uint) (models.Subject, error)
	GetRandomQuestions(subjectID uint) ([]models.Question, error)
}
