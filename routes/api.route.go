package routes

import (
	"github.com/Strayneko/KomikcastAPI/controllers/comic"
	"github.com/gofiber/fiber/v2"
)

func InitApiRoute(route fiber.Router) {
	initComicRoute(route)
}

func initComicRoute(route fiber.Router) {
	comicRoute := route.Group("/comic")
	comicRoute.Get("/list", comic.GetComicList)
}
