package scraper

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/Strayneko/KomikcastAPI/configs"
	"log"
	"net/http"
)

var KomikcastUrl string

type Service interface {
	Scrape(path string) *goquery.Document
}

type scrapper struct {
	service Service
}

func New() Service {
	KomikcastUrl = configs.ViperEnv.Get("KOMIKCAST_URL").(string)
	return &scrapper{}
}

func (s *scrapper) Scrape(path string) *goquery.Document {
	res, err := http.Get(KomikcastUrl + path)
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		return nil
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return doc
}
