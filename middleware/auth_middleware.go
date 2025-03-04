package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/trisatya23/go-restful-api-master/model/web"
)

type AuthMiddleware struct{}

func NewAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Get("X-API-Key") == "RAHASIA" {
			return c.Next()
		}

		return c.Status(fiber.StatusUnauthorized).JSON(web.WebResponse{
			Code:   fiber.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		})
	}
}
