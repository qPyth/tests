package models

import (
	"errors"
	"gorm.io/gorm"
)

type Topic struct {
	gorm.Model
	ID     int
	SubID  int
	PartID int
	Title  string
}

var (
	ErrTopicNotFound = errors.New("topic not found for this subject")
)
