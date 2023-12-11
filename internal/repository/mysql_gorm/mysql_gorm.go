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

func (r *Repository) GetSubjectById(id uint) (models.Subject, error) {
	var subject models.Subject

	r.db.Where("id = ?", id).First(&subject)
	return subject, nil
}

func (r *Repository) GetRandomQuestions(subjectID uint) ([]models.Question, error) {
	var questions []models.Question
	var topics []models.Topic

	if err := r.db.Where("subject_id = ?", subjectID).Find(&topics).Error; err != nil {
		return nil, err
	}

	if len(topics) == 0 {
		return nil, models.ErrTopicNotFound
	}

	levelDistribution := map[int]int{
		1: 15,
		2: 10,
		3: 5,
	}

	for level, totalQuestionsForLevel := range levelDistribution {
		questionsPerTopic := totalQuestionsForLevel / len(topics)
		remainingQuestions := totalQuestionsForLevel % len(topics)

		for i, topic := range topics {
			topicQuestions := questionsPerTopic
			if i < remainingQuestions {
				topicQuestions++
			}

			var levelQuestions []models.Question
			if err := r.db.Where("subject_id = ? AND topic_id = ? AND level = ?", subjectID, topic.ID, level).Order("RAND()").Limit(topicQuestions).Find(&levelQuestions).Error; err != nil {
				return nil, err
			}

			questions = append(questions, levelQuestions...)
		}
	}

	return questions, nil
}
