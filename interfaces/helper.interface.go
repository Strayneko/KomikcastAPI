package interfaces

import "github.com/gofiber/fiber/v2"

type Helper interface {
	ResponseError(ctx *fiber.Ctx, err *fiber.Error) error
	ValidatePage(ctx *fiber.Ctx, currentPage *int16) *fiber.Error
}
