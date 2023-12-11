package repository

import "tests/internal/models"

type Repository interface {
	GetSubjectById(id int) (models.Subject, error)
}
