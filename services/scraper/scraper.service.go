package scraper

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/Strayneko/KomikcastAPI/configs"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
)

var KomikcastUrl string

type Service interface {
	Scrape(path string) (*goquery.Document, *fiber.Error)
}

type scrapper struct {
	service Service
}

func New() Service {
	KomikcastUrl = configs.ViperEnv.Get("KOMIKCAST_URL").(string)
	return &scrapper{}
}

func (s *scrapper) Scrape(path string) (*goquery.Document, *fiber.Error) {
	res, err := http.Get(KomikcastUrl + path)

	if err != nil {
		log.Printf("Cannot conenct to " + KomikcastUrl + path)
		return nil, fiber.NewError(http.StatusInternalServerError, "Cannot connect to "+KomikcastUrl+". Reason: "+err.Error())
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Printf("status code error: %d %s", res.StatusCode, res.Status)
		return nil, fiber.NewError(http.StatusInternalServerError, "Cannot connect to "+KomikcastUrl+". Reason: "+res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Printf("Cannot parse response from %s", KomikcastUrl+path)
		return nil, fiber.NewError(http.StatusInternalServerError, "Cannot parse response from "+KomikcastUrl+". Reason: "+err.Error())
	}
	return doc, nil
}
