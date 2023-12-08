package admin

import "github.com/gofiber/fiber/v2"

func New(config ...fiber.Config) fiber.Handler {
	return func(fiberCtx *fiber.Ctx) error {
		return fiberCtx.Next()
	}
}
