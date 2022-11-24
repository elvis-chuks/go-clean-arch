package http

import (
	"github.com/elvis-chuks/go-clean-arch/domain"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Config struct {
	UserRepo domain.UserRepository
}

func ping(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   "pong",
	})
}

func RunHttpServer(config Config) *fiber.App {
	app := fiber.New()
	app.Use(cors.New())

	setupRouter(app, config)

	return app
}
