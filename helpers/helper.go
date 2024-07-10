package helpers

import (
	"net/http"
	"strconv"

	"github.com/Strayneko/KomikcastAPI/interfaces"
	"github.com/Strayneko/KomikcastAPI/types"
	"github.com/gofiber/fiber/v2"
)

type helper struct {
	controllerHelper interfaces.Helper
}

func New() interfaces.Helper {
	return &helper{}
}

// ValidatePage validates the "page" query parameter from the request context.
// It ensures that the page is a positive integer and converts it to int16.
func (h *helper) ValidatePage(ctx *fiber.Ctx) (int16, *fiber.Error) {
	page := ctx.Query("page", "1")
	curPage, err := strconv.ParseInt(page, 10, 16)

	if err != nil || curPage <= 0 {
		return 0, fiber.NewError(http.StatusBadRequest, "Bad Request: Invalid page")
	}

	return int16(curPage), nil
}

func (h *helper) ResponseError(ctx *fiber.Ctx, err *fiber.Error) error {
	return ctx.Status(err.Code).JSON(&types.ResponseType{
		Status:  false,
		Code:    int16(err.Code),
		Message: err.Message,
	})
}
