package models

import "errors"

type Subject struct {
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
