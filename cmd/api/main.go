package main

import (
	"encoding/json"
	"log"
	"log/slog"
	"os"
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
	if err != nil {
		log.Fatal(err)
	}
	injson, err := json.Marshal(test)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.OpenFile("qqq", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	defer file.Close()
	_, err = file.Write(injson)
	if err != nil {
		log.Fatal(err)
	}

	//starting server
	//address := fmt.Sprintf("%s:%s", cfg.HTTPServer.Host, cfg.HTTPServer.Port)
	//r := gin.Default()
	//r.GET("/:name", func(c *gin.Context) {
	//	name := c.Param("name")
	//	c.JSON(200, gin.H{
	//		"Welcome": name,
	//	})
	//})
	//err = r.Run(address)
	//if err != nil {
	//	log.Fatalf("server starting error: %s", err.Error())
	//}
}
