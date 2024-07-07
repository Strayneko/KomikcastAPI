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
	comicRoute := route.Group("/comic", middlewares.CacheMiddleware())
	comicRoute.Get("/list", comic.GetComicList)
}
