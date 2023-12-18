package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
	"tests/internal/config"
	"tests/internal/repository"
	"tests/internal/services"
	"tests/pkg/database"
)

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
	test, err := s.GetTest(218, 220, 187, 207, 208)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("test", func(c *gin.Context) {
		c.JSON(200, test)
	})
	r.Run() // listen and serve on 0.0.0.0:8080

}
