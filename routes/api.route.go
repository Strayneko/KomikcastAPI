package routes

import (
	"github.com/Strayneko/KomikcastAPI/controllers/comic"
	"github.com/Strayneko/KomikcastAPI/middlewares"
	"github.com/gofiber/fiber/v2"
)

func InitApiRoute(route fiber.Router) {
	initComicRoute(route)
}

func initComicRoute(route fiber.Router) {
	controller := comic.NewController()
	comicRoute := route.Group("/comic")

	comicRoute.Get("/list", middlewares.CacheMiddleware("list"), controller.GetComicList)
	comicRoute.Get("/search", middlewares.CacheMiddleware("search"), controller.GetSearchedComics)
	comicRoute.Get("/projects", middlewares.CacheMiddleware("list"), controller.GetProjectComics)
}
