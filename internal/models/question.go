package models

import (
	"errors"
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	ID            int
	SubjectID     int
	PartID        int
	TopicID       int
	TargetID      int
	Level         int
	Type          int
	CountVariants int
	CountAnswers  int
}

var (
	ErrQuestionNotFound = errors.New("question not found")
)
