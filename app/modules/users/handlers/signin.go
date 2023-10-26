package handlers

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vkhoa145/go-chat-next-ts/app/middlewares"
	"github.com/vkhoa145/go-chat-next-ts/app/models"
	"github.com/vkhoa145/go-chat-next-ts/config"
)

func (h *UserHandler) SignIn(config *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := models.SignInInput{}

		if err := c.BodyParser(&payload); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
		}

		user, err := h.userUsecase.SignIn(&payload)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
		}

		t, err := middlewares.CreateAccessToken(user, config.SIGNED_STRING, 24)
		if err != nil {
			return c.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
		}

		c.Set("Authorization", t)
		c.Cookie(&fiber.Cookie{
			Name:    "access_token",
			Value:   t,
			Expires: time.Now().Add(time.Hour * 1),
		})
		c.Status(http.StatusOK)
		return c.JSON(&fiber.Map{"status": http.StatusOK, "token": t, "error": nil})
	}
}
