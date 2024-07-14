package interfaces

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/Strayneko/KomikcastAPI/types"
	"github.com/gofiber/fiber/v2"
)

type ComicListService interface {
	ExtractComicList(ctx *fiber.Ctx) ([]types.ComicListInfoType, *fiber.Error)
	ExtractComicDetail(selector *goquery.Selection) types.ComicListInfoType
	GetComicList(ctx *fiber.Ctx, path string, currentPage int16) error
}

type ComicListController interface {
	GetComicList(ctx *fiber.Ctx) error
	GetSearchedComics(ctx *fiber.Ctx) error
	GetProjectComics(ctx *fiber.Ctx) error
}

type DetailController interface {
	GetComicDetail(ctx *fiber.Ctx) error
}

type ComicDetailService interface {
	GetComicDetail(ctx *fiber.Ctx, slug string) error
	ExtractChapters(selector *goquery.Selection, chapters *[]types.ChapterDetailType)
	ExtractGenres(selector *goquery.Selection, genres *[]types.GenreType)
	ExtractComicDetail(ctx *fiber.Ctx, selector *goquery.Selection) (*types.ComicDetailType, error)
	ExtractRelatedSeries(selector *goquery.Selection, relatedSeries *[]types.ComicListInfoType)
}
