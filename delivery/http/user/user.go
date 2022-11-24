package user

import (
	"encoding/json"
	"github.com/elvis-chuks/go-clean-arch/domain"
	"github.com/elvis-chuks/go-clean-arch/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type handler struct {
	usecase domain.UserUseCase
	repo    domain.UserRepository
	logger  *zap.Logger
}

func New(router fiber.Router, usecase domain.UserUseCase, repo domain.UserRepository) {
	handler := handler{
		usecase: usecase,
		repo:    repo,
	}

	l, _ := logger.Init()

	handler.logger = l

	router.Get("/echo", handler.Echo)

}

func (h *handler) Echo(c *fiber.Ctx) error {
	var user domain.User

	if err := json.Unmarshal(c.Body(), &user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	user = h.usecase.Echo(user)

	return c.JSON(fiber.Map{
		"error": false,
		"data":  user,
	})
}
