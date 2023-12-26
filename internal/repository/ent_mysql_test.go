package repository

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"log/slog"
	"testing"
	"tests/internal/config"
	"tests/pkg/database"
)

func TestEntRepo_GetQuestions(t *testing.T) {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config reading error: %s", err)
	}

	//init database connection
	db, err := database.New(cfg.Storage.DBName, cfg.Storage.Username, cfg.Storage.Password)
	if err != nil {
		slog.Error("database conn failed", "error", err)
		return
	}

	rep := NewRepository(db)
	subjectID := 207
	count := 15
	typeOf := 0
	level := 1
	for i := 0; i < 1000; i++ {
		q, err := rep.Ent.GetQuestions(subjectID, count, typeOf, level)
		assert.NoError(t, err)
		assert.Equal(t, count, len(q), fmt.Sprintf("count of questions! expected %d, got %d", count, len(q)))
		for _, question := range q {
			assert.Equal(t, subjectID, question.SubjectID, fmt.Sprintf("subID of question! expected %d, got %d", subjectID, question.SubjectID))
			assert.Equal(t, typeOf, question.Type, fmt.Sprintf("type of questions! expected %d, got %d", typeOf, question.Type))
			assert.Equal(t, level, question.Level, fmt.Sprintf("level of questions! expected %d, got %d", level, question.Level))
			assert.Equal(t, 1, len(question.QuestionDetails), fmt.Sprintf("count of questions details! expected %d, got %d", 1, len(question.QuestionDetails)))
		}
	}

	subjectID = 207
	count = 5
	typeOf = 2
	level = 1
	for i := 0; i < 1000; i++ {
		q, err := rep.Ent.GetQuestions(subjectID, count, typeOf, level)
		assert.NoError(t, err)
		assert.Equal(t, count, len(q), fmt.Sprintf("count of questions! expected %d, got %d", count, len(q)))
		for _, question := range q {
			assert.Equal(t, subjectID, question.SubjectID, fmt.Sprintf("subID of question! expected %d, got %d", subjectID, question.SubjectID))
			assert.Equal(t, typeOf, question.Type, fmt.Sprintf("type of questions! expected %d, got %d", typeOf, question.Type))
			assert.Equal(t, level, question.Level, fmt.Sprintf("level of questions! expected %d, got %d", level, question.Level))
			assert.Equal(t, 5, len(question.QuestionDetails), fmt.Sprintf("count of questions details! expected %d, got %d", 1, len(question.QuestionDetails)))
		}
	}
}
