package scraper

import (
	browser "github.com/EDDYCJY/fake-useragent"
	"github.com/PuerkitoBio/goquery"
	"github.com/Strayneko/KomikcastAPI/configs"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
)

var WebUrl string

type Service interface {
	Scrape(path string) (*goquery.Document, *fiber.Error)
}

type scrapper struct {
	service Service
}

func New() Service {
	WebUrl = configs.ViperEnv.GetString("WEB_URL")
	return &scrapper{}
}

// Scrape fetches and parses the HTML document from the given path on the Komikcast website.
// It handles HTTP requests and responses, and returns a goquery.Document for further processing.
func (s *scrapper) Scrape(path string) (*goquery.Document, *fiber.Error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", WebUrl+path, nil)

	if err != nil {
		log.Printf("Cannot create request to %s%s", WebUrl, path)
		return nil, fiber.NewError(http.StatusInternalServerError, "Cannot create request to "+WebUrl+". Reason: "+err.Error())
	}

	// Set randomw useragent
	userAgent := browser.Computer()
	req.Header.Set("User-Agent", userAgent)

	res, err := client.Do(req)
	if err != nil {
		log.Printf("Cannot connect to %s%s", WebUrl, path)
		return nil, fiber.NewError(http.StatusInternalServerError, "Cannot connect to "+WebUrl+". Reason: "+err.Error())
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Printf("status code error: %d %s", res.StatusCode, res.Status)
		return nil, fiber.NewError(http.StatusInternalServerError, "Cannot connect to "+WebUrl+". Reason: "+res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Printf("Cannot parse response from %s", WebUrl+path)
		return nil, fiber.NewError(http.StatusInternalServerError, "Cannot parse response from "+WebUrl+". Reason: "+err.Error())
	}
	return doc, nil
}
