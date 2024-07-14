package interfaces

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gofiber/fiber/v2"
)

type ScraperService interface {
	Scrape(path string) (*goquery.Document, *fiber.Error)
}
