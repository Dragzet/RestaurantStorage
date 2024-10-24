package main

import (
	"RestaurantStorage/internal/config"
	"RestaurantStorage/internal/repository"
	"RestaurantStorage/internal/storage/PostgreSQL"
	"context"
	"fmt"
	log "github.com/go-ozzo/ozzo-log"
	"github.com/joho/godotenv"
	"os"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func main() {
	cfg, err := config.LoadConfig()
	emailHost, emailPort, emailUser, emailPass := os.Getenv("EMAIL_HOST"), os.Getenv("EMAIL_PORT"), os.Getenv("EMAIL_USER"), os.Getenv("EMAIL_PASS")

	if err != nil || emailHost == "" || emailPort == "" || emailUser == "" || emailPass == "" {
		panic("Error loading config and secrets")
	}
	
	logger := setupLogger(cfg.LogsPath)
	logger.Open()
	defer logger.Close()

	logger.Info("Starting the application...")

	storage, err := PostgreSQL.NewStorage(context.TODO(), cfg.StoragePath)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	repo := repository.NewDB(storage)

	_ = repo

}

func setupLogger(loggsPath string) *log.Logger {
	logger := log.NewLogger()

	t1 := log.NewConsoleTarget()
	t2 := log.NewFileTarget()
	t2.FileName = loggsPath
	logger.Targets = append(logger.Targets, t1, t2)

	return logger
}
