package helpers

import (
	"github.com/Strayneko/KomikcastAPI/interfaces"
	"github.com/Strayneko/KomikcastAPI/types"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type helper struct {
	controllerHelper interfaces.Helper
}

func New() interfaces.Helper {
	return &helper{}
}
func (h *helper) ValidatePage(ctx *fiber.Ctx, currentPage *int16) *fiber.Error {
	page := ctx.Query("page", "1")
	curPage, err := strconv.ParseInt(page, 10, 16)

	if err != nil || curPage <= 0 {
		return fiber.NewError(http.StatusBadRequest, "Bad Request: Invalid page")
	}

	cp := int16(curPage)
	currentPage = &cp
	return nil
}

func (h *helper) ResponseError(ctx *fiber.Ctx, err *fiber.Error) error {
	return ctx.Status(err.Code).JSON(types.ResponseType{
		Status:  false,
		Code:    int16(err.Code),
		Message: err.Message,
	})
}
