package main

import (
	"context"
	"log"

	"github.com/arizdn234/go-clean-architecture-1/internal/api"
	"github.com/arizdn234/go-clean-architecture-1/internal/config"
	"github.com/arizdn234/go-clean-architecture-1/internal/database"
	"github.com/arizdn234/go-clean-architecture-1/internal/repository"
	"github.com/arizdn234/go-clean-architecture-1/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	conn, err := database.NewPostgres(cfg)

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close(context.Background())

	repo := repository.NewUserRepository(conn)
	svc := service.NewUserService(repo)
	handler := api.NewUserHandler(svc)

	r := gin.Default()

	r.POST("/users", handler.CreateUser)

	log.Println("server running on :8080")

	r.Run(":8080")
}