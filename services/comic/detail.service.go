package comic

import (
	"context"
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"github.com/Strayneko/KomikcastAPI/configs"
	"github.com/Strayneko/KomikcastAPI/interfaces"
	"github.com/Strayneko/KomikcastAPI/types"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
	"time"
)

type ComicDetailService struct {
	Service          interfaces.ComicDetailService
	RdbCtx           context.Context
	Helper           interfaces.Helper
	ScraperService   interfaces.ScraperService
	ComicListService interfaces.ComicListService
}

func NewComicDetailService(
	helper interfaces.Helper,
	scraperService interfaces.ScraperService,
	comicListService interfaces.ComicListService,
) interfaces.ComicDetailService {
	return &ComicDetailService{
		RdbCtx:           context.Background(),
		Helper:           helper,
		ScraperService:   scraperService,
		ComicListService: comicListService,
	}
}

// GetComicDetail retrieves the comic details for a given slug.
// It uses the scraperService to scrape the comic details from a webpage,
// extracts the relevant details using the ExtractComicDetail method,
// and sends the response back in JSON format.
func (service *ComicDetailService) GetComicDetail(ctx *fiber.Ctx, slug string) error {
	var cachedDetails *types.ComicDetailType
	var comicDetails *types.ComicDetailType
	path := "manga/" + slug

	cached, _ := configs.Cache.Get(service.RdbCtx, path).Result()
	cachedErr := json.Unmarshal([]byte(cached), &cachedDetails)

	if len(cached) > 0 && cachedErr == nil {
		comicDetails = cachedDetails
	} else {
		doc, err := service.ScraperService.Scrape(path)
		if err != nil {
			return service.Helper.ResponseError(ctx, err)
		}

		comicDetails = service.ExtractComicDetail(ctx, doc.Find("#content .wrapper"))
		cachedData, _ := json.Marshal(&comicDetails)
		configs.Cache.Set(service.RdbCtx, path, cachedData, 3*time.Hour)
	}

	return ctx.Status(http.StatusOK).JSON(&types.ResponseType{
		Status:  true,
		Code:    http.StatusOK,
		Message: "Comic details successfully fetched.",
		Data:    &comicDetails,
	})
}

// ExtractComicDetail extracts detailed information about a comic from a given HTML selector.
// It gathers data such as the title, alias, cover image, rating, summary, chapters, genres, and related series.
func (service *ComicDetailService) ExtractComicDetail(ctx *fiber.Ctx, selector *goquery.Selection) *types.ComicDetailType {
	var genres []types.GenreType
	var chapters []types.ChapterDetailType
	var relatedSeries []types.ComicListInfoType
	var comicInfo *types.ComicInfoType

	title := selector.Find(".seriestuheader h1.entry-title").Text()
	title = strings.TrimSpace(title)
	alias := selector.Find(".seriestuheader div.seriestualt").Text()
	alias = strings.TrimSpace(alias)
	coverImage, _ := selector.Find(".seriestucontent .thumb img.attachment-").Attr("src")
	starRating, _ := selector.Find(".rating .rtp .rtb span").Attr("style")
	ratingValue := selector.Find(".rating div.num").Text()

	summary, _ := selector.Find(".seriestucontentr .seriestuhead div.entry-content").Html()
	summary = strings.TrimSpace(summary)
	firstAndLastChapter := selector.Find(".seriestucontentr .seriestuhead .lastend div.inepcx")
	firstChapter := firstAndLastChapter.First()
	firstChapterText := firstChapter.Find(".epcur").Text()
	firstChapterUrl := firstChapter.Find("a").AttrOr("href", "")

	lastChapter := firstAndLastChapter.Last()
	lastChapterText := lastChapter.Find(".epcur").Text()
	lastChapterUrl := lastChapter.Find("a").AttrOr("href", "")

	comicInfo = service.ExtractComicInfo(selector)
	service.ExtractGenres(selector, &genres)
	service.ExtractChapters(selector, &chapters)
	service.ExtractRelatedSeries(selector, &relatedSeries)

	return &types.ComicDetailType{
		Title:      title,
		Alias:      alias,
		CoverImage: coverImage,
		ComicRating: &types.ComicRatingType{
			StarRating: service.Helper.ExtractStarRatingValue(starRating),
			Rating:     ratingValue,
		},
		Summary: summary,
		FirstChapter: &types.ChapterDetailType{
			Chapter:     firstChapterText,
			ChapterUrl:  firstChapterUrl,
			ChapterSlug: service.Helper.ExtractSlug(firstChapterUrl),
		},
		LatestChapter: &types.ChapterDetailType{
			Chapter:     lastChapterText,
			ChapterUrl:  lastChapterUrl,
			ChapterSlug: service.Helper.ExtractSlug(lastChapterUrl),
		},
		ComicInfo:     comicInfo,
		Genres:        genres,
		Chapters:      chapters,
		RelatedSeries: relatedSeries,
	}
}

func (service *ComicDetailService) ExtractComicInfo(selector *goquery.Selection) *types.ComicInfoType {
	comicInfo := make(map[string]string)
	selector.Find(".seriestucont .seriestucontr table.infotable tr").Each(func(i int, s *goquery.Selection) {
		info := s.Find("td").First().Text()
		info = strings.TrimSpace(info)
		info = strings.ToLower(info)
		info = strings.ReplaceAll(info, " ", "")
		infoData := s.Find("td").Last().Text()
		infoData = strings.TrimSpace(infoData)
		comicInfo[info] = infoData
	})

	return &types.ComicInfoType{
		Status:       types.ComicStatusType(comicInfo["status"]),
		ReleasedDate: comicInfo["released"],
		Artist:       comicInfo["artist"],
		PostedOn:     comicInfo["postedon"],
		Views:        comicInfo["views"],
		ComicType:    types.ComicType(comicInfo["type"]),
		Author:       comicInfo["author"],
		UpdatedAt:    comicInfo["updatedon"],
	}
}

// ExtractChapters extracts chapter details from a given HTML selector and appends them to the provided chapters slice.
func (service *ComicDetailService) ExtractChapters(selector *goquery.Selection, chapters *[]types.ChapterDetailType) {
	selector.Find(".bixbox #chapterlist ul li .chbox div.eph-num").Each(func(i int, s *goquery.Selection) {
		url := s.Find("a").AttrOr("href", "")
		*chapters = append(*chapters, types.ChapterDetailType{
			ChapterUrl:  url,
			Chapter:     s.Find("a span.chapternum").Text(),
			ChapterSlug: service.Helper.ExtractSlug(url),
			LastUpdated: s.Find("a span.chapterdate").Text(),
		})
	})
}

// ExtractGenres extracts genre details from a given HTML selector and appends them to the provided genres slice.
func (service *ComicDetailService) ExtractGenres(selector *goquery.Selection, genres *[]types.GenreType) {
	selector.Find(".seriestucont .seriestucontr div.seriestugenre a").Each(func(i int, s *goquery.Selection) {
		url := s.AttrOr("href", "")
		genre := s.Text()
		*genres = append(*genres, types.GenreType{
			Url:   url,
			Genre: genre,
			Slug:  service.Helper.ExtractSlug(url),
		})
	})
}

// ExtractRelatedSeries extracts related series details from a given HTML selector and appends them to the provided relatedSeries slice.
func (service *ComicDetailService) ExtractRelatedSeries(selector *goquery.Selection, relatedSeries *[]types.ComicListInfoType) {
	series := selector.Find(".listupd .bs .bsx")
	if series.Length() > 0 {
		series.Each(func(i int, s *goquery.Selection) {
			*relatedSeries = append(*relatedSeries, service.ComicListService.ExtractComicDetail(s))
		})
	}
}
