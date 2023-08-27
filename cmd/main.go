package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"user_segmentation"
	"user_segmentation/pkg/handler"
	"user_segmentation/pkg/repository"
	"user_segmentation/pkg/service"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//	@title			User Segmentation App API
//	@version		1.0
//	@description	API Server for User Segmentation Application

//	@host		localhost:8000
//	@BasePath	/

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("Error occured while initializing configs: %s", err.Error())
	}

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
	go func() {
		if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("UserSegmentationApp STARTED")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("UserSegmentationApp is SHUTTING DOWN")

	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Errorf("Error occured while stopping http server: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("Error occured while closing db connection: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
