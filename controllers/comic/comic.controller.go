package comic

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/Strayneko/KomikcastAPI/services/comic"
	"github.com/Strayneko/KomikcastAPI/services/scraper"
	"github.com/Strayneko/KomikcastAPI/types"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type Controller interface {
	GetComicList(ctx *fiber.Ctx) error
}

func GetComicList(ctx *fiber.Ctx) error {
	path := "daftar-komik/"
	page := ctx.Query("page", "1")
	currentPage, err := strconv.ParseInt(page, 10, 16)

	if err != nil || currentPage <= 0 {
		return responseError(ctx, fiber.NewError(http.StatusBadRequest, "Bad Request: Invalid page"))
	}

	if len(page) > 0 {
		path += "page/" + page
	}

	return baseGetComicList(ctx, path, int16(currentPage))
}

func baseGetComicList(ctx *fiber.Ctx, path string, page int16) error {
	var doc *goquery.Document
	var comicList []types.ComicType
	var err *fiber.Error

	scraperService := scraper.New()

	doc, err = scraperService.Scrape(path)
	if err != nil {
		return responseError(ctx, err)
	}

	comicService := comic.New(doc)
	comicList, err = comicService.GetComicList(ctx)
	lastPage := comicService.GetLastPageNumber()

	if err != nil {
		return responseError(ctx, err)
	}

	return ctx.Status(http.StatusOK).JSON(types.ResponseType{
		Status:      true,
		Code:        http.StatusOK,
		LastPage:    lastPage,
		CurrentPage: page,
		Total:       int16(len(comicList)),
		Message:     "List of comics successfully fetched.",
		Data:        &comicList,
	})
}

func responseError(ctx *fiber.Ctx, err *fiber.Error) error {
	return ctx.Status(err.Code).JSON(types.ResponseType{
		Status:  false,
		Code:    int16(err.Code),
		Message: err.Message,
	})
}
