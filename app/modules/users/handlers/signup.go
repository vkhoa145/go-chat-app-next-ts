package handlers

import (
	"net/http"
	"reflect"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/vkhoa145/go-chat-next-ts/app/middlewares"
	"github.com/vkhoa145/go-chat-next-ts/app/models"
	"github.com/vkhoa145/go-chat-next-ts/config"
)

func (h *UserHandler) SignUp(config *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := models.SignUpInput{}

		if err := c.BodyParser(&payload); err != nil {
			c.Status(http.StatusBadRequest)
		}

		validate := validator.New()
		if err := validate.Struct(payload); err != nil {
			errors := map[string]string{}
			for _, err := range err.(validator.ValidationErrors) {
				fieldName := err.Field()
				field, _ := reflect.TypeOf(payload).FieldByName(fieldName)
				jsonTag := field.Tag.Get("json")
				errors[jsonTag] = err.Tag()
			}
			c.Status(http.StatusUnprocessableEntity)
			return c.JSON(&fiber.Map{"status": http.StatusUnprocessableEntity, "message": "Unprocessable Content", "errors": errors})
		}
		createUser, err := h.userUsecase.SignUp(&payload)

		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
		}

		token, err := middlewares.CreateAccessToken(createUser, config.SIGNED_STRING, 24)

		if err != nil {
			return c.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
		}

		c.Status(http.StatusCreated)
		return c.JSON(&fiber.Map{"status": http.StatusCreated, "error": nil, "user": createUser, "token": token})
	}
}
