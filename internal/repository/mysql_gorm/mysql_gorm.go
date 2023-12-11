package mysql_gorm

import (
	"gorm.io/gorm"
	"tests/internal/models"
)

type Repository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) GetSubjectById(id int) (models.Subject, error) {
	var subject models.Subject

	r.db.Where("id = ?", id).First(&subject)
	return subject, nil
}
