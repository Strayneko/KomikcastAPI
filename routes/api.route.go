package routes

import (
	"github.com/Strayneko/KomikcastAPI/controllers/comic"
	"github.com/Strayneko/KomikcastAPI/helpers"
	"github.com/Strayneko/KomikcastAPI/services/scraper"
	"github.com/gofiber/fiber/v2"
)

func InitApiRoute(route fiber.Router) {
	initComicRoute(route)
}

func initComicRoute(route fiber.Router) {
	helper := helpers.NewHelper()

	scraperService := scraper.NewScraperService()
	list := comic.NewComicListController(helper, scraperService)
	detail := comic.NewComicDetailController(helper, scraperService)
	comicRoute := route.Group("/comic")

	comicRoute.Get("/list", list.GetComicList)
	comicRoute.Get("/search", list.GetSearchedComics)
	comicRoute.Get("/projects", list.GetProjectComics)

	comicRoute.Get("/detail/:slug", detail.GetComicDetail)
}
