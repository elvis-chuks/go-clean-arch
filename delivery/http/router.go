package http

import (
	"github.com/elvis-chuks/go-clean-arch/delivery/http/user"
	userU "github.com/elvis-chuks/go-clean-arch/usecases/user"
	"github.com/gofiber/fiber/v2"
)

func setupRouter(app *fiber.App, config Config) {
	app.Get("/api/v1/ping", ping)

	v1 := app.Group("/api/v1")

	userRouter := v1.Group("/user")

	userUsecase := userU.New(config.UserRepo)

	user.New(userRouter, userUsecase, config.UserRepo)
}
