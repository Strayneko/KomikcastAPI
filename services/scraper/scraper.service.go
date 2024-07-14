package scraper

import (
	browser "github.com/EDDYCJY/fake-useragent"
	"github.com/PuerkitoBio/goquery"
	"github.com/Strayneko/KomikcastAPI/configs"
	"github.com/Strayneko/KomikcastAPI/interfaces"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"time"
)

type scrapper struct {
	service interfaces.ScraperService
	WebUrl  string
}

func NewScraperService() interfaces.ScraperService {
	webUrl := configs.ViperEnv.GetString("WEB_URL")
	return &scrapper{
		WebUrl: webUrl,
	}
}

// Scrape fetches and parses the HTML document from the given path on the Komikcast website.
// It handles HTTP requests and responses, and returns a goquery.Document for further processing.
func (s *scrapper) Scrape(path string) (*goquery.Document, *fiber.Error) {
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", s.WebUrl+path, nil)

	if err != nil {
		log.Printf("Cannot create request to %s%s", s.WebUrl, path)
		return nil, fiber.NewError(http.StatusInternalServerError, "Cannot create request to "+s.WebUrl+". Reason: "+err.Error())
	}

	// Set randomw useragent
	userAgent := browser.Computer()
	req.Header.Set("User-Agent", userAgent)

	res, err := client.Do(req)

	if err != nil {
		log.Printf("Cannot connect to %s%s", s.WebUrl, path)
		return nil, fiber.NewError(http.StatusInternalServerError, "Cannot connect to "+s.WebUrl+path+". Reason: "+err.Error())
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Printf("status code error: %d %s", res.StatusCode, res.Status)
		return nil, fiber.NewError(res.StatusCode, "Cannot connect to "+s.WebUrl+path+". Reason: "+res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Printf("Cannot parse response from %s", s.WebUrl+path)
		return nil, fiber.NewError(http.StatusInternalServerError, "Cannot parse response from "+s.WebUrl+". Reason: "+err.Error())
	}

	return doc, nil
}
