package interfaces

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/Strayneko/KomikcastAPI/types"
	"github.com/gofiber/fiber/v2"
)

type ComicService interface {
	GetComicList(ctx *fiber.Ctx) ([]types.ComicType, *fiber.Error)
	ExtractComicDetail(selector *goquery.Selection) types.ComicType
	ExtractStarRatingValue(starRating string) int8
	GetLastPageNumber() int16
}

type ComicController interface {
	GetComicList(ctx *fiber.Ctx) error
	GetSearchedComics(ctx *fiber.Ctx) error
	BaseGetComicList(ctx *fiber.Ctx, path string, currentPage int16) error
}
