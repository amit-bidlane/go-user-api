package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"go-user-api/config"
	"go-user-api/internal/handler"
	"go-user-api/internal/logger"
	"go-user-api/internal/repository"
	"go-user-api/internal/routes"
	"go-user-api/internal/service"
	"go.uber.org/zap"
)

func main() {
	logger.Init()
	defer logger.Log.Sync()

	cfg := config.Load()

	db, err := sql.Open("postgres", cfg.DSN())
	if err != nil {
		log.Fatal("failed to connect to database", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("failed to ping database", err)
	}

	logger.Log.Info("connected to database successfully")

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	app := fiber.New()

	routes.Setup(app, userHandler)

	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	logger.Log.Info("server starting", zap.String("port", cfg.ServerPort))

	if err := app.Listen(addr); err != nil {
		log.Fatal("server failed to start", err)
	}
}