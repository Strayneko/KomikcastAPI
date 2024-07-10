package interfaces

import "github.com/gofiber/fiber/v2"

type HandlerInterface interface {
	ErrorHandler(ctx *fiber.Ctx, err error) error
}
