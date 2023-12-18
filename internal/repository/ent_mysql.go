package repository

import (
	"gorm.io/gorm"
	"tests/internal/models"
)

type EntRepo struct {
	db *gorm.DB
}

func NewEntRepo(db *gorm.DB) *EntRepo {
	return &EntRepo{db}
}

func (r *EntRepo) GetMathLiteracyTest(subjectID int) (TestOutput, error) {
	var test TestOutput
	var questions []models.Question
	r.db.Preload("QuestionDetails").Find(&questions, "subject_id = ?", subjectID).Limit(10)
	test.Questions = questions
	return test, nil
}

func (r *EntRepo) GetKazHistoryTest(subjectID int) (TestOutput, error) {
	var test TestOutput
	var questions []models.Question
	r.db.Preload("QuestionDetails").Find(&questions, "subject_id = ?", subjectID).Limit(20)
	test.Questions = questions
	return test, nil
}

func (r *EntRepo) GetReadingLiteracyTest(subjectID int) (TestOutput, error) {
	var test TestOutput
	var questions []models.Question
	r.db.Preload("QuestionDetails").Find(&questions, "subject_id = ?", subjectID).Limit(10)
	test.Questions = questions
	return test, nil
}

func (r *EntRepo) GetProfileTest(subjectID int) (TestOutput, error) {
	var test TestOutput
	var questions []models.Question
	r.db.Preload("QuestionDetails").Find(&questions, "subject_id = ?", subjectID).Limit(40)
	test.Questions = questions
	return test, nil
}
