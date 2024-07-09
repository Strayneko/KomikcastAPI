package comic

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/Strayneko/KomikcastAPI/helpers"
	"github.com/Strayneko/KomikcastAPI/interfaces"
	"github.com/Strayneko/KomikcastAPI/services/scraper"
	"github.com/Strayneko/KomikcastAPI/types"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var Doc *goquery.Document
var Helper interfaces.Helper

type comic struct {
	service interfaces.ComicService
}

func New() interfaces.ComicService {
	Helper = helpers.New()
	return &comic{}
}

// ExtractComicList extract a list of comics by extracting details from the provided context using goquery.
func (service *comic) ExtractComicList(ctx *fiber.Ctx) ([]types.ComicType, *fiber.Error) {
	var comicList []types.ComicType
	if Doc == nil {
		return nil, fiber.NewError(http.StatusServiceUnavailable, "Service Unavailable")
	}

	items := Doc.Find("#content .wrapper .list-update_items .list-update_items-wrapper .list-update_item")
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
// and rating information, and returns a ComicType struct populated with this data.

func (service *comic) ExtractComicDetail(selector *goquery.Selection) types.ComicType {
	comicUrl, _ := selector.Find("a.data-tooltip").Attr("href")
	listUpdateItem := selector.Find(".list-update_item-image")
	listUpdateItemInfo := selector.Find(".list-update_item-info")
	coverImage, _ := listUpdateItem.Find("img.ts-post-image").Attr("src")
	comicType := listUpdateItem.Find("span.type").Text()
	lastChapter := listUpdateItemInfo.Find(".other .chapter").Text()
	lastChapterUrl, _ := listUpdateItemInfo.Find(".other .chapter").Attr("href")
	title := listUpdateItemInfo.Find("h3.title").Text()
	starRating, _ := listUpdateItemInfo.Find(".other .rate .rating .rating-bintang span").Attr("style")
	ratingScore := listUpdateItemInfo.Find(".other .rate .rating .numscore").Text()

	return types.ComicType{
		Title:      title,
		CoverImage: coverImage,
		ComicType:  comicType,
		Url:        comicUrl,
		LastChapter: types.ComicChapterType{
			LastChapter:    strings.TrimSpace(lastChapter),
			LastChapterUrl: lastChapterUrl,
		},
		ComicRating: types.ComicRatingType{
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

// GetLastPageNumber retrieves the last page number from the document.
// It checks the pagination elements and extracts the highest page number.
func (service *comic) GetLastPageNumber() int16 {
	if Doc == nil {
		return 0
	}

	pageList := Doc.Find(".komiklist .pagination .page-numbers")
	if pageList.Length() == 0 {
		return 0
	}

	lastPageIsNumber := regexp.MustCompile(`\d`).MatchString(pageList.Last().Text())
	if !lastPageIsNumber {
		pageList = pageList.Slice(0, pageList.Length()-1)
	}

	lastPage, err := strconv.ParseInt(pageList.Last().Text(), 10, 16)
	if err != nil {
		return 0
	}

	return int16(lastPage)
}

// GetComicList handles fetching a list of comics from a given path and returns it in the response context.
// It uses a scraper service to scrape the comics from the specified path and then extracts the list of comics.
func (service *comic) GetComicList(ctx *fiber.Ctx, path string, currentPage int16) error {
	var comicList []types.ComicType
	var err *fiber.Error

	scraperService := scraper.New()

	Doc, err = scraperService.Scrape(path)
	if err != nil {
		return Helper.ResponseError(ctx, err)
	}

	comicList, err = service.ExtractComicList(ctx)
	lastPage := service.GetLastPageNumber()

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
