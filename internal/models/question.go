package models

import (
	"errors"
)

type Question struct {
	ID              int
	SubjectID       int
	PartID          int
	TopicID         int
	TargetID        int
	Level           int
	Type            int
	CountVariants   int
	CountAnswers    int
	Text            string
	QuestionDetails []QuestionDetails `gorm:"foreignKey:QuestionID"`
}

var (
	ErrQuestionNotFound = errors.New("question not found")
)
