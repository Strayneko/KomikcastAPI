package comic

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/Strayneko/KomikcastAPI/configs"
	"github.com/Strayneko/KomikcastAPI/services/scraper"
	"github.com/Strayneko/KomikcastAPI/types"
	"github.com/gofiber/fiber/v2"
	"time"
)

type Service interface {
	GetComicList(ctx *fiber.Ctx) []types.ComicType
	GetComicDetail(selector *goquery.Selection) types.ComicType
}

type comic struct {
	service Service
}

func New() Service {
	return &comic{}
}

func (s *comic) GetComicList(ctx *fiber.Ctx) []types.ComicType {
	cacheKey := "comicList"
	if cached, found := configs.Cache.Get(cacheKey); found {
		return cached.([]types.ComicType)
	}

	scraperService := scraper.New()
	doc := scraperService.Scrape("?s=3")
	if doc == nil {
		return nil
	}

	var comicList []types.ComicType
	// Find the comic details
	doc.Find(".list-update .list-update_items .list-update_items-wrapper .list-update_item").Each(func(i int, selector *goquery.Selection) {
		// For each item found, get the information
		comicList = append(comicList, s.GetComicDetail(selector))
	})

	configs.Cache.Set(cacheKey, comicList, 30*time.Minute)
	return comicList
}

func (s *comic) GetComicDetail(selector *goquery.Selection) types.ComicType {
	title := selector.Find(".list-update_item-info h3.title").Text()
	listUpdateItem := selector.Find(".list-update_item-image")
	coverImage, _ := listUpdateItem.Find("img.ts-post-image").Attr("src")
	comicType := listUpdateItem.Find("span.type").Text()
	return types.ComicType{
		Title:      title,
		CoverImage: coverImage,
		ComicType:  comicType,
	}
}
