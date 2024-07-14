package interfaces

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/Strayneko/KomikcastAPI/types"
	"github.com/gofiber/fiber/v2"
)

type ComicListService interface {
	ExtractComicList(ctx *fiber.Ctx) ([]types.ComicListInfoType, *fiber.Error)
	ExtractComicDetail(selector *goquery.Selection) types.ComicListInfoType
	ExtractStarRatingValue(starRating string) int8
	GetComicList(ctx *fiber.Ctx, path string, currentPage int16) error
}

type ComicController interface {
	GetComicList(ctx *fiber.Ctx) error
	GetSearchedComics(ctx *fiber.Ctx) error
	GetProjectComics(ctx *fiber.Ctx) error
}
