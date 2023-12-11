package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
	"os"
	"tests/internal/config"
	"tests/internal/repository/mysql_gorm"
	s "tests/internal/services"
	"tests/pkg/database"
)

func main() {
	// load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config reading error: %s", err)
	}
	// init logger
	slog := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)

	//init database connection
	db, err := database.New(cfg.Storage.DBName, cfg.Storage.Username, cfg.Storage.Password)
	if err != nil {
		slog.Error("database conn failed", "error", err)
		return
	}
	slog.Info("successful connection to database")

	rep := mysql_gorm.NewRepo(db)
	services := s.NewServices(rep)

	test, err := services.Test.GetTestBySubjectId(189)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(len(test.Questions))
	for _, q := range test.Questions {
		if q.Level == 1 {
			fmt.Printf("questionID: %d\n", q.ID)
			fmt.Printf("SubID: %d\n", q.SubjectID)
			fmt.Printf("Level: %d\n", q.Level)
			fmt.Printf("TotalVars: %d\n", q.CountVariants)
			fmt.Printf("TotalAnswers: %d\n", q.CountAnswers)
		}
	}
	//starting server
	address := fmt.Sprintf("%s:%s", cfg.HTTPServer.Host, cfg.HTTPServer.Port)

	r := gin.Default()
	r.GET("/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"Welcome": name,
		})
	})
	err = r.Run(address)
	if err != nil {
		log.Fatalf("server starting error: %s", err.Error())
	}
}
