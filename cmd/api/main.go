package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
	"net/http"
	"tests/internal/config"
	"tests/internal/models"
	"tests/internal/repository"
	"tests/internal/services"
	"tests/pkg/database"
)

type ErrorResponse struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
}

func main() {
	// load config
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
	slog.Info("successful connection to database")
	rep := repository.NewRepository(db)
	s := services.NewServices(rep)

	r := gin.Default()
	r.GET("ent-test", func(c *gin.Context) {
		test, err := s.GetTest(1200, 220, 187, 199, 194)
		if err != nil {
			if errors.Is(err, models.ErrQuestionNotFound) {
				log.Println(err)
				c.JSON(400, ErrorResponse{ErrorCode: 400, Message: http.StatusText(400)})
			}
			log.Println(err)
			c.JSON(500, ErrorResponse{ErrorCode: 500, Message: http.StatusText(500)})
			return
		}
		c.JSON(200, test)
	})
	err = r.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
