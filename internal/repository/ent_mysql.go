package repository

import (
	"gorm.io/gorm"
	"log"
	"tests/internal/models"
)

type EntRepo struct {
	db *gorm.DB
}

func NewEntRepo(db *gorm.DB) *EntRepo {
	return &EntRepo{db}
}

func (r *EntRepo) Get(subjectID int) models.QuestionDetails {
	var qzxcd models.QuestionDetails

	res := r.db.Where("question_id=?", subjectID).First(&qzxcd)
	if res.Error != nil {
		log.Fatal(res.Error)
	}

	return qzxcd
}

func (r *EntRepo) GetMathLiteracyTest(subjectID int) (TestOutput, error) {
	var test TestOutput
	var selectedQuestionsWithDetails []QuestionWithDetails
	err := r.db.Table("question_details").
		Select("question_details.*, questions.*").
		Joins("INNER JOIN questions ON question_details.question_id = questions.ID").
		Find(&selectedQuestionsWithDetails).Error
	if err != nil {
		return TestOutput{}, err
	}
	test.Questions = selectedQuestionsWithDetails

	return test, nil
}

func (r *EntRepo) GetKazHistoryTest(subjectID int) (TestOutput, error) {
	var test TestOutput
	var selectedQuestionsWithDetails []QuestionWithDetails
	err := r.db.Table("questions").Select("questions.*, question_details.*").
		Joins("INNER JOIN question_details ON questions.id = question_details.question_id").Where("questions.subject_id=?", subjectID).Limit(40).Find(&selectedQuestionsWithDetails).Error
	if err != nil {
		return TestOutput{}, err
	}
	test.Questions = selectedQuestionsWithDetails
	return test, nil
}

func (r *EntRepo) GetReadingLiteracyTest(subjectID int) (TestOutput, error) {
	var test TestOutput
	var selectedQuestionsWithDetails []QuestionWithDetails
	err := r.db.Table("questions").Select("questions.*, question_details.*").
		Joins("INNER JOIN question_details ON questions.id = question_details.question_id").Where("questions.subject_id=?", subjectID).Limit(40).Find(&selectedQuestionsWithDetails).Error
	if err != nil {
		return TestOutput{}, err
	}
	test.Questions = selectedQuestionsWithDetails
	return test, nil
}

func (r *EntRepo) GetProfileTest(subjectID int) (TestOutput, error) {
	var test TestOutput
	var selectedQuestionsWithDetails []QuestionWithDetails
	err := r.db.Table("questions").Select("questions.*, question_details.*").
		Joins("INNER JOIN question_details ON questions.id = question_details.question_id").Where("questions.subject_id=?", subjectID).Limit(40).Find(&selectedQuestionsWithDetails).Error
	if err != nil {
		return TestOutput{}, err
	}
	test.Questions = selectedQuestionsWithDetails
	return test, nil
}
