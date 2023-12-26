package repository

import (
	"gorm.io/gorm"
	"tests/internal/models"
)

const (
	type0 = iota
	type1
	type2
)

const (
	level1 = iota + 1
	level2
	level3
)

type Ent interface {
	GetQuestions(subjectID, count, typeOfQ, level int) ([]models.Question, error)
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
