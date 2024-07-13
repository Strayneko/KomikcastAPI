package comic

import (
	"context"
	"encoding/json"
	"github.com/Strayneko/KomikcastAPI/configs"
	"net/http"
	"regexp"
	"strconv"
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
	service interfaces.ComicService
}

func New() interfaces.ComicService {
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

	if len(splitType) > 0 {
		comicType = splitType[len(splitType)-1]
	}

	return types.ComicListInfoType{
		Title:       title,
		CoverImage:  coverImage,
		ComicType:   types.ComicType(comicType),
		Url:         comicUrl,
		LastChapter: strings.TrimSpace(lastChapter),

		ComicRating: &types.ComicRatingType{
			StarRating: service.ExtractStarRatingValue(starRating),
			Rating:     ratingScore,
		},
	}
}

// ExtractStarRatingValue extracts the star rating value from a width css attribute Ex: width: 70%, will result 3.5.
// It uses a regular expression to find and return the number in the string.
func (service *comic) ExtractStarRatingValue(starRating string) int8 {
	// Compile the regex pattern to extract the number
	re := regexp.MustCompile(`\d+(\.\d+)?`)

	// Find the match
	match := re.FindString(starRating)

	if len(match) == 0 {
		return 0
	}
	res, err := strconv.Atoi(match)
	if err != nil {
		return 0
	}

	return int8(res / 20)

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
