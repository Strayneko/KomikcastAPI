package comic

import (
	"github.com/Strayneko/KomikcastAPI/configs"
	"github.com/Strayneko/KomikcastAPI/interfaces"
	"github.com/Strayneko/KomikcastAPI/services/comic"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"slices"
	"strconv"
	"strings"
)

type ListHandler struct {
	controller       interfaces.ComicListController
	Helper           interfaces.Helper
	ComicListService interfaces.ComicListService
}

func NewComicListController(helper interfaces.Helper, scraperService interfaces.ScraperService) interfaces.ComicListController {
	comicListService := comic.NewComicListService(helper, scraperService)
	return &ListHandler{
		Helper:           helper,
		ComicListService: comicListService,
	}
}

func (h *ListHandler) GetComicList(ctx *fiber.Ctx) error {
	path := "manga/"
	orderBy := ctx.Query("order")

	if isValidOrderBy := len(orderBy) > 0 && !slices.Contains(configs.ComicOrderParams, orderBy); isValidOrderBy {
		orderTypes := strings.Join(configs.ComicOrderParams, ",")
		return h.Helper.ResponseError(ctx, fiber.NewError(http.StatusBadRequest, "Order should be in "+orderTypes))
	}

	currentPage, err := h.Helper.ValidatePage(ctx)
	if err != nil {
		return h.Helper.ResponseError(ctx, err)
	}

	if currentPage > 0 {
		path += "?page=" + strconv.Itoa(int(currentPage))
	}
	if len(orderBy) > 0 {
		path += "&order=" + configs.GetComicOrderBy(orderBy)
	}

	return h.ComicListService.GetComicList(ctx, path, currentPage)
}

func (h *ListHandler) GetSearchedComics(ctx *fiber.Ctx) error {
	query := ctx.Query("query", "")
	path := "?s=" + query
	currentPage, err := h.Helper.ValidatePage(ctx)

	if err != nil {
		return h.Helper.ResponseError(ctx, err)
	}
	if currentPage > 0 {
		path = "page/" + strconv.Itoa(int(currentPage)) + "/?s=" + query
	}
	return h.ComicListService.GetComicList(ctx, path, currentPage)
}

func (h *ListHandler) GetProjectComics(ctx *fiber.Ctx) error {
	path := "project/"
	currentPage, err := h.Helper.ValidatePage(ctx)

	if err != nil {
		return h.Helper.ResponseError(ctx, err)
	}

	if currentPage > 0 {
		path += "page/" + strconv.Itoa(int(currentPage))
	}
	return h.ComicListService.GetComicList(ctx, path, currentPage)
}
