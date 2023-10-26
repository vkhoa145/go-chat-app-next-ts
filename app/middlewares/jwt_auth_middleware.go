package middlewares

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/vkhoa145/go-chat-next-ts/app/models"
	"github.com/vkhoa145/go-chat-next-ts/config"
)

func JwtAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		authorized, err := isAuthorized(authHeader, config.LoadConfig().SIGNED_STRING)
		if authorized {
			userId, err := ExtractIdfromToken(authHeader, config.LoadConfig().SIGNED_STRING)
			if err != nil {
				return c.JSON(&fiber.Map{"status": http.StatusBadRequest, "error": err.Error()})
			}
			c.Request().Header.Set("user_id", userId)
			return c.Next()
		}
		return c.JSON(&fiber.Map{"status": http.StatusForbidden, "error": err.Error(), "message": "Unauthorized"})
	}
}

func CreateAccessToken(user *models.User, secret string, expiry int) (string, error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry)).Unix()
	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   exp,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return t, nil
}

func isAuthorized(accessToken string, secret string) (bool, error) {
	_, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %w", t.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}

	return true, nil
}

func ExtractIdfromToken(accessToken string, secret string) (string, error) {
	token, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %w", t.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return "", fmt.Errorf("Invalid Token")
	}

	id, ok := claims["id"].(string)
	if !ok {
		if floatId, ok := claims["id"].(float64); ok {
			id = strconv.FormatFloat(floatId, 'f', -1, 64)
		} else {
			return "", fmt.Errorf("Failed to extract 'id' from claims")
		}
	}
	return id, nil
}
