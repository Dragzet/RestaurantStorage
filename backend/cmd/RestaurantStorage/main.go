package main

import (
	"RestaurantStorage/internal/config"
	"RestaurantStorage/internal/http-server/handlers"
	"RestaurantStorage/internal/repository"
	"RestaurantStorage/internal/service"
	"RestaurantStorage/internal/storage/PostgreSQL"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/go-ozzo/ozzo-log"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	storage, err := PostgreSQL.NewStorage(context.TODO(), cfg.StoragePath)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	logger.Info("Storage connected successfully")
	repo := repository.NewDB(storage)
	router := gin.Default()
	serv := service.NewService(repo)
	handler := handlers.NewHandler(router, serv, logger)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	server := &http.Server{
		Addr:         cfg.Address,
		Handler:      handler,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			logger.Error("failed to start server: %s", err.Error())
		}
	}()

	logger.Info("Server started successfully")
	<-signalChan

	logger.Info("Shutting down the server")
}

func setupLogger(loggsPath string) *log.Logger {
	logger := log.NewLogger()

	t1 := log.NewConsoleTarget()
	t2 := log.NewFileTarget()
	t2.FileName = loggsPath
	logger.Targets = append(logger.Targets, t1, t2)

	return logger
}
