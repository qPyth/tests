package models

import "errors"

type Topic struct {
	ID     int
	SubID  int
	PartID int
	Title  string
}

var (
	ErrTopicNotFound = errors.New("topic not found for this subject")
)
