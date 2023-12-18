package repository

import (
	"gorm.io/gorm"
	"tests/internal/models"
)

type Ent interface {
	GetMathLiteracyTest(subjectID int) (TestOutput, error)
	GetKazHistoryTest(subjectID int) (TestOutput, error)
	GetReadingLiteracyTest(subjectID int) (TestOutput, error)
	GetProfileTest(subjectID int) (TestOutput, error)
}

type TestOutput struct {
	Questions []models.Question `json:"questions"`
}

type Repository struct {
	Ent Ent
}

func NewRepository(db *gorm.DB) *Repository {
	ent := NewEntRepo(db)
	return &Repository{
		Ent: ent,
	}
}
