package repository

import (
	"fmt"
	"gorm.io/gorm"
	"tests/internal/models"
)

type EntRepo struct {
	db *gorm.DB
}

func NewEntRepo(db *gorm.DB) *EntRepo {
	return &EntRepo{db}
}

func (e *EntRepo) GetQuestions(subjectID, count, typeOfQ, level int) ([]models.Question, error) {
	var questions []models.Question
	err := e.db.Where("subject_id = ? AND type=? AND level=?", subjectID, typeOfQ, level).
		Order("RAND()").
		Preload("QuestionDetails").Limit(count).Find(&questions)
	if err.Error != nil {
		return nil, err.Error
	}
	if len(questions) == 0 {
		fmt.Println(err)
		return nil, models.ErrQuestionNotFound
	}
	return questions, nil
}
