package main

import (
	"instagram/user-service/internal/config"
	"instagram/user-service/internal/database"
	"instagram/user-service/internal/handler"
	"instagram/user-service/internal/repository"
	"instagram/user-service/internal/server"
	"instagram/user-service/internal/service"
	"log"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.Connect(cfg.MongoURL, cfg.DBName)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewUserRepository(db)
	svc := service.NewUserService(repo)
	h := handler.NewUserHandler(svc)

	log.Println("User service running on", cfg.GRPCPort)

	if err := server.StartGRPC(cfg.GRPCPort, h); err != nil {
		log.Fatal(err)
	}
}
