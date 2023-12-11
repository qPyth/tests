package models

import "errors"

type Question struct {
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
