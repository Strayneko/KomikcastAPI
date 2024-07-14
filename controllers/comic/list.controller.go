package comic

import (
	"github.com/Strayneko/KomikcastAPI/configs"
	"github.com/Strayneko/KomikcastAPI/helpers"
	"github.com/Strayneko/KomikcastAPI/interfaces"
	"github.com/Strayneko/KomikcastAPI/services/comic"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"slices"
	"strconv"
	"strings"
)

var Helper interfaces.Helper
var ComicListService interfaces.ComicListService

type ListHandler struct {
	controller interfaces.ComicListController
}

func NewComicListController() interfaces.ComicListController {
	Helper = helpers.New()
	ComicListService = comic.New()
	return &ListHandler{}
}

func (h *ListHandler) GetComicList(ctx *fiber.Ctx) error {
	path := "manga/"
	orderBy := ctx.Query("order")

	if isValidOrderBy := len(orderBy) > 0 && !slices.Contains(configs.ComicOrderParams, orderBy); isValidOrderBy {
		orderTypes := strings.Join(configs.ComicOrderParams, ",")
		return Helper.ResponseError(ctx, fiber.NewError(http.StatusBadRequest, "Order should be in "+orderTypes))
	}

	currentPage, err := Helper.ValidatePage(ctx)
	if err != nil {
		return Helper.ResponseError(ctx, err)
	}

	if currentPage > 0 {
		path += "?page=" + strconv.Itoa(int(currentPage))
	}
	if len(orderBy) > 0 {
		path += "&order=" + configs.GetComicOrderBy(orderBy)
	}

	return ComicListService.GetComicList(ctx, path, currentPage)
}

func (h *ListHandler) GetSearchedComics(ctx *fiber.Ctx) error {
	query := ctx.Query("query", "")
	path := "?s=" + query
	currentPage, err := Helper.ValidatePage(ctx)

	if err != nil {
		return Helper.ResponseError(ctx, err)
	}
	if currentPage > 0 {
		path = "page/" + strconv.Itoa(int(currentPage)) + "/?s=" + query
	}
	return ComicListService.GetComicList(ctx, path, currentPage)
}

func (h *ListHandler) GetProjectComics(ctx *fiber.Ctx) error {
	path := "project/"
	currentPage, err := Helper.ValidatePage(ctx)

	if err != nil {
		return Helper.ResponseError(ctx, err)
	}

	if currentPage > 0 {
		path += "page/" + strconv.Itoa(int(currentPage))
	}
	return ComicListService.GetComicList(ctx, path, currentPage)
}
