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
	Get(subjectID int) models.QuestionDetails
}

type TestOutput struct {
	Questions []QuestionWithDetails `json:"questions"`
}

type QuestionWithDetails struct {
	Question        models.Question        `json:"question"`
	QuestionDetails models.QuestionDetails `json:"question_details"`
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
