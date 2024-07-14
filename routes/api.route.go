package routes

import (
	"github.com/Strayneko/KomikcastAPI/controllers/comic"
	"github.com/gofiber/fiber/v2"
)

func InitApiRoute(route fiber.Router) {
	initComicRoute(route)
}

func initComicRoute(route fiber.Router) {
	list := comic.NewComicListController()
	detail := comic.NewComicDetailController()
	comicRoute := route.Group("/comic")

	comicRoute.Get("/list", list.GetComicList)
	comicRoute.Get("/search", list.GetSearchedComics)
	comicRoute.Get("/projects", list.GetProjectComics)

	comicRoute.Get("/detail/:slug", detail.GetComicDetail)
}
