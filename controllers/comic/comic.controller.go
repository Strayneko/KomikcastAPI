package comic

import (
	"github.com/Strayneko/KomikcastAPI/services/comic"
	"strconv"

	"github.com/Strayneko/KomikcastAPI/helpers"
	"github.com/Strayneko/KomikcastAPI/interfaces"
	"github.com/gofiber/fiber/v2"
)

var Helper interfaces.Helper
var ComicService interfaces.ComicService

type handler struct {
	controller interfaces.ComicController
}

func NewController() interfaces.ComicController {
	Helper = helpers.New()
	ComicService = comic.New()
	return &handler{}
}

func (h *handler) GetComicList(ctx *fiber.Ctx) error {
	path := "daftar-komik/"

	currentPage, err := Helper.ValidatePage(ctx)
	if err != nil {
		return Helper.ResponseError(ctx, err)
	}

	if currentPage > 0 {
		path += "page/" + strconv.Itoa(int(currentPage))
	}
	return ComicService.GetComicList(ctx, path, currentPage)
}

func (h *handler) GetSearchedComics(ctx *fiber.Ctx) error {
	query := ctx.Query("query", "")
	path := "?s=" + query
	currentPage, err := Helper.ValidatePage(ctx)

	if err != nil {
		return Helper.ResponseError(ctx, err)
	}
	if currentPage > 0 {
		path = "page/" + strconv.Itoa(int(currentPage)) + "/?s=" + query
	}
	return ComicService.GetComicList(ctx, path, currentPage)
}

func (h *handler) GetProjectComics(ctx *fiber.Ctx) error {
	path := "project-list/"
	currentPage, err := Helper.ValidatePage(ctx)

	if err != nil {
		return Helper.ResponseError(ctx, err)
	}

	if currentPage > 0 {
		path += "page/" + strconv.Itoa(int(currentPage))
	}
	return ComicService.GetComicList(ctx, path, currentPage)
}
