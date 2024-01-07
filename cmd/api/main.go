package main

import (
	"log"
	"log/slog"
	"tests/internal/config"
	"tests/internal/repository"
	"tests/internal/services"
	"tests/internal/transport"
	"tests/pkg/database"
)

// @title           Swagger Example API
// @version         2.0
// @description     This is a server for getting tests
// @termsOfService  http://swagger.io/terms/

// @host      localhost:8080
// @BasePath  /api/

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

	srv := transport.NewHandler(s)

	err = srv.Init().Run()
	if err != nil {
		log.Fatalln("server running error:%w", err)
	}
}
