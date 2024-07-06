package comic

import (
	"github.com/Strayneko/KomikcastAPI/services/comic"
	"github.com/Strayneko/KomikcastAPI/types"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Controller interface {
	GetComicList(ctx *fiber.Ctx) error
}

func GetComicList(ctx *fiber.Ctx) error {
	var message string

	comicService := comic.New()
	comicList := comicService.GetComicList(ctx)
	if comicList == nil || len(comicList) == 0 {
		message = "Cannot fetch comic list."
	} else {
		message = "List of comics successfully fetched."
	}

	return ctx.JSON(types.ResponseType{
		Status:  true,
		Code:    http.StatusOK,
		Message: message,
		Data:    &comicList,
	})
}
