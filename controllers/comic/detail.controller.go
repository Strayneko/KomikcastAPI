package comic

import (
	"github.com/Strayneko/KomikcastAPI/interfaces"
	"github.com/Strayneko/KomikcastAPI/services/comic"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type DetailHandler struct {
	DetailController   interfaces.DetailController
	ComicDetailService interfaces.ComicDetailService
	Helper             interfaces.Helper
	ScraperService     interfaces.ScraperService
	ComicListService   interfaces.ComicListService
}

func NewComicDetailController(
	helper interfaces.Helper,
	scraperService interfaces.ScraperService,
) interfaces.DetailController {
	comicListService := comic.NewComicListService(helper, scraperService)
	service := comic.NewComicDetailService(helper, scraperService, comicListService)
	return &DetailHandler{
		ComicDetailService: service,
		Helper:             helper,
		ScraperService:     scraperService,
		ComicListService:   comicListService,
	}
}

func (h *DetailHandler) GetComicDetail(ctx *fiber.Ctx) error {
	slug := ctx.Params("slug")

	if len(slug) == 0 {
		return h.Helper.ResponseError(ctx, fiber.NewError(http.StatusBadRequest, "Invalid slug"))
	}

	return h.ComicDetailService.GetComicDetail(ctx, slug)
}
