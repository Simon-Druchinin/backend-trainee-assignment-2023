package main

import (
	"os"
	"user_segmentation"
	"user_segmentation/pkg/handler"
	"user_segmentation/pkg/repository"
	"user_segmentation/pkg/service"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := godotenv.Load(".env.dev"); err != nil {
		logrus.Fatalf("Error occured while loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	})

	if err != nil {
		logrus.Fatalf("Error occured while connecting to db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(user_segmentation.Server)
	if err := server.Run("8000", handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Error occured while running http server: %s", err.Error())
	}
}
