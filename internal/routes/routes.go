package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-user-api/internal/handler"
	"go-user-api/internal/middleware"
)

func Setup(app *fiber.App, userHandler *handler.UserHandler) {
	app.Use(middleware.RequestID())
	app.Use(middleware.RequestLogger())

	users := app.Group("/users")
	users.Post("/", userHandler.CreateUser)
	users.Get("/", userHandler.ListUsers)
	users.Get("/:id", userHandler.GetUserByID)
	users.Put("/:id", userHandler.UpdateUser)
	users.Delete("/:id", userHandler.DeleteUser)
}