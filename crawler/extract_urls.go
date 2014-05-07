package crawler

import (
	"io"
	"log"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

func ExtractURLs(body io.Reader, host string) ([]string, error) {
	urls := []string{}

	document, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return urls, err
	}

	urlElementMatches := [][]string{
		[]string{"a", "href"},
		[]string{"img", "src"},
		[]string{"link", "href"},
		[]string{"script", "src"},
	}

	for _, attr := range urlElementMatches {
		element, attr := attr[0], attr[1]
		urls = append(urls, findByElementAttribute(document, host, element, attr)...)
	}

	return urls, err
}

func findByElementAttribute(document *goquery.Document, host string, element string, attr string) []string {
	urls := []string{}

	document.Find(element).Each(func(_ int, element *goquery.Selection) {
		href, exists := element.Attr(attr)

		u, err := url.Parse(href)
		if err != nil {
			log.Fatal(err)
		}

		if exists && u.Host == host {
			urls = append(urls, href)
		}
	})

	return urls
}