package models

import (
	"errors"
	"gorm.io/gorm"
)

type Subject struct {
	gorm.Model
	Id            int
	Product       int
	Title         string
	Class         int
	IsNecessarily int
	Language      int
	QCount        int
}

var (
	ErrSubNotFound = errors.New("subject not found")
)
