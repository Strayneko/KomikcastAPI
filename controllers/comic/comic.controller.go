package comic

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/Strayneko/KomikcastAPI/helpers"
	"github.com/Strayneko/KomikcastAPI/interfaces"
	"github.com/Strayneko/KomikcastAPI/services/comic"
	"github.com/Strayneko/KomikcastAPI/services/scraper"
	"github.com/Strayneko/KomikcastAPI/types"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

var Helper interfaces.Helper

type handler struct {
	controller interfaces.ComicController
}

func NewController() interfaces.ComicController {
	Helper = helpers.New()
	return &handler{}
}

func (h *handler) GetComicList(ctx *fiber.Ctx) error {
	var currentPage int16
	path := "daftar-komik/"

	if err := Helper.ValidatePage(ctx, &currentPage); err != nil {
		return Helper.ResponseError(ctx, err)
	}

	if currentPage > 0 {
		path += "page/" + string(currentPage)
	}
	return h.BaseGetComicList(ctx, path, currentPage)
}

func (h *handler) GetSearchedComics(ctx *fiber.Ctx) error {
	var currentPage int16
	query := ctx.Query("query", "")
	path := "?s=" + query

	if err := Helper.ValidatePage(ctx, &currentPage); err != nil {
		return Helper.ResponseError(ctx, err)
	}
	if currentPage > 0 {
		path = "page/" + string(currentPage) + "/?s=" + query
	}
	return h.BaseGetComicList(ctx, path, 1)
}

func (h *handler) BaseGetComicList(ctx *fiber.Ctx, path string, currentPage int16) error {
	var doc *goquery.Document
	var comicList []types.ComicType
	var err *fiber.Error

	scraperService := scraper.New()

	doc, err = scraperService.Scrape(path)
	if err != nil {
		return Helper.ResponseError(ctx, err)
	}

	comicService := comic.New(doc)
	comicList, err = comicService.GetComicList(ctx)
	lastPage := comicService.GetLastPageNumber()

	if err != nil {
		return Helper.ResponseError(ctx, err)
	}

	return ctx.Status(http.StatusOK).JSON(types.ResponseType{
		Status:      true,
		Code:        http.StatusOK,
		LastPage:    lastPage,
		CurrentPage: currentPage,
		Total:       int16(len(comicList)),
		Message:     "List of comics successfully fetched.",
		Data:        &comicList,
	})
}
