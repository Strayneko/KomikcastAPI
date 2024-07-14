package comic

import (
	"context"
	"encoding/json"
	"github.com/Strayneko/KomikcastAPI/configs"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/Strayneko/KomikcastAPI/helpers"
	"github.com/Strayneko/KomikcastAPI/interfaces"
	"github.com/Strayneko/KomikcastAPI/services/scraper"
	"github.com/Strayneko/KomikcastAPI/types"
	"github.com/gofiber/fiber/v2"
)

var Doc *goquery.Document
var Helper interfaces.Helper
var RdbCtx context.Context = context.Background()

type comic struct {
	service interfaces.ComicListService
}

func New() interfaces.ComicListService {
	Helper = helpers.New()
	return &comic{}
}

// ExtractComicList extract a list of comics by extracting details from the provided context using goquery.
func (service *comic) ExtractComicList(ctx *fiber.Ctx) ([]types.ComicListInfoType, *fiber.Error) {
	var comicList []types.ComicListInfoType
	if Doc == nil {
		return nil, fiber.NewError(http.StatusServiceUnavailable, "Service Unavailable")
	}

	items := Doc.Find("#content .wrapper .listupd .bs .bsx")
	if items.Length() == 0 {
		return nil, fiber.NewError(http.StatusNotFound, "Page not found.")
	}
	// Find the comic details
	items.Each(func(i int, selector *goquery.Selection) {
		// For each item found, get the information
		comicList = append(comicList, service.ExtractComicDetail(selector))
	})
	return comicList, nil
}

// ExtractComicDetail extracts detailed information about a comic from the provided goquery selector.
// The function gathers various attributes such as the comic's URL, cover image, type, last chapter details,
// and rating information, and returns a ComicListInfoType struct populated with this data.
func (service *comic) ExtractComicDetail(selector *goquery.Selection) types.ComicListInfoType {
	comicUrl, _ := selector.Find("a").Attr("href")
	coverImage, _ := selector.Find("img.ts-post-image").Attr("src")
	comicType, _ := selector.Find("span.type").Attr("class")
	splitType := strings.Split(comicType, " ")
	lastChapter := selector.Find(".adds .epxs").Text()
	title := selector.Find("div.tt").Text()
	title = strings.TrimSpace(title)
	starRating, _ := selector.Find(".rt .rating .rtp span").Attr("style")
	ratingScore := selector.Find(".rt .rating div.numscore").Text()
	slug := Helper.ExtractSlug(comicUrl)

	if len(splitType) > 0 {
		comicType = splitType[len(splitType)-1]
	}

	return types.ComicListInfoType{
		Title:       title,
		CoverImage:  coverImage,
		ComicType:   types.ComicType(comicType),
		Url:         comicUrl,
		LastChapter: strings.TrimSpace(lastChapter),
		Slug:        slug,

		ComicRating: &types.ComicRatingType{
			StarRating: Helper.ExtractStarRatingValue(starRating),
			Rating:     ratingScore,
		},
	}
}

// GetComicList handles fetching a list of comics from a given path and returns it in the response context.
// It uses a scraper service to scrape the comics from the specified path and then extracts the list of comics.
func (service *comic) GetComicList(ctx *fiber.Ctx, path string, currentPage int16) error {
	var err *fiber.Error
	var comicList []types.ComicListInfoType
	var cachedComicList []types.ComicListInfoType

	cached, _ := configs.Cache.Get(RdbCtx, path).Result()
	cachedErr := json.Unmarshal([]byte(cached), &cachedComicList)
	if len(cached) > 0 && cachedErr == nil {
		comicList = cachedComicList
	} else {
		scraperService := scraper.New()
		Doc, err = scraperService.Scrape(path)
		if err != nil {
			return Helper.ResponseError(ctx, err)
		}

		comicList, err = service.ExtractComicList(ctx)

		if err != nil {
			return Helper.ResponseError(ctx, err)
		}

		cachedData, _ := json.Marshal(comicList)
		configs.Cache.Set(RdbCtx, path, cachedData, 3*time.Hour)
	}
	return ctx.Status(http.StatusOK).JSON(&types.ResponseType{
		Status:      true,
		Code:        http.StatusOK,
		CurrentPage: currentPage,
		Total:       int16(len(comicList)),
		Message:     "List of comics successfully fetched.",
		Data:        &comicList,
	})
}
