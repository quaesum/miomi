package main

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"log"
	"net/url"
	"os"
	"strconv"
	"time"
)

// Item represents the structure to hold the parsed data
type Item struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Photos      []string `json:"photos"`
	Link        string   `json:"link"`
}

func main() {
	baseURL := "https://garfield.by"
	subLinks := []string{
		"/catalog/cats.html",
		"/catalog/dogs.html",
	}

	var items []Item

	// Create a new collector
	c := colly.NewCollector(
		colly.Async(true),
		colly.CacheDir("./cache"),
	)

	err := c.Limit(&colly.LimitRule{
		DomainGlob:  "*garfield.*",
		Parallelism: 4,               // Increase parallelism
		Delay:       1 * time.Second, // Delay between requests
	})

	if err != nil {
		log.Fatalln(err)
	}

	// Create another collector for pagination
	paginationCollector := c.Clone()

	err = paginationCollector.Limit(&colly.LimitRule{
		DomainGlob:  "*garfield.*",
		Parallelism: 4,
		Delay:       1 * time.Second,
	})

	if err != nil {
		log.Fatalln(err)
	}

	// On every div with class="js-accordion-wrap active"
	c.OnHTML("div.catalog__side", func(e *colly.HTMLElement) {
		e.ForEach("div.js-accordion-content a", func(_ int, el *colly.HTMLElement) {
			href := el.Attr("href")
			fullLink := baseURL + href
			paginationCollector.Visit(fullLink)
		})
	})

	// On every div with class="catalog__content"
	paginationCollector.OnHTML("div.catalog__content", func(e *colly.HTMLElement) {
		e.ForEach("div.snippet", func(_ int, el *colly.HTMLElement) {
			var item Item

			item.Link = baseURL + el.ChildAttr("a.snippet-image", "href")
			item.Photos = append(item.Photos, baseURL+el.ChildAttr("a.snippet-image img", "src"))
			item.Name = el.ChildText("a.h2.link.snippet-desc__name")
			item.Description = el.ChildText("p.snippet-desc__additional.p2.gray")

			items = append(items, item)
		})

		// Handle pagination
		e.DOM.Find("div.catalog-pagination a.catalog-pagination__item.h-border").Each(func(_ int, s *goquery.Selection) {
			href, exists := s.Attr("href")
			if exists {
				u, err := url.Parse(href)
				if err != nil {
					log.Println("Failed to parse URL:", err)
					return
				}

				query := u.Query()
				pageStr := query.Get("PAGEN_1")
				if pageStr != "" {
					page, err := strconv.Atoi(pageStr)
					if err != nil {
						log.Println("Failed to convert page number:", err)
						return
					}

					// Update the page number and visit the next page
					query.Set("PAGEN_1", strconv.Itoa(page+1))
					u.RawQuery = query.Encode()
					nextPage := baseURL + u.String()

					paginationCollector.Visit(nextPage)
				}
			}
		})
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	paginationCollector.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	// Start the scraping
	for _, link := range subLinks {
		c.Visit(baseURL + link)
	}

	// Wait until all threads are finished
	c.Wait()
	paginationCollector.Wait()

	// Print the result
	file, err := os.Create("result.json")
	if err != nil {
		log.Fatalf("Could not create file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(items); err != nil {
		log.Fatalf("Could not encode items: %v", err)
	}
}
