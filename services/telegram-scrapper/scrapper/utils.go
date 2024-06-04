package scrapper

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// writeToFile writes target amount of Animals in file
func writeToFile(animals []Animal) error {
	jsonData, err := json.MarshalIndent(animals, "", "  ")
	if err != nil {
		log.Fatalf("Marshaling data error: %v", err)
		return err
	}

	err = os.WriteFile("animals.json", jsonData, 0666)

	return nil
}

// ProcessHTML reads html code and parse animals
func ProcessHTML(html string, limit int, records map[int]struct{}, count int) ([]Animal, error) {
	animals := make([]Animal, 0)

	urlRe := regexp.MustCompile(`url\((.*?)\)`)

	reader := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("error creating goquery document: %w", err)
	}

	doc.Find(".tgme_widget_message").Each(func(i int, s *goquery.Selection) {
		dataPost, ok := s.Attr("data-post")
		var id int
		if ok {
			id, err = strconv.Atoi(regexp.MustCompile(`\d+`).FindString(dataPost))
			if err != nil {
				log.Println("Error getting data-post:", err)
				return
			}
		}
		if _, found := records[id]; found {
			return
		}
		records[id] = struct{}{}

		text := s.Find(".tgme_widget_message_text").Text()

		var photos []string
		s.Find(".tgme_widget_message_photo_wrap").Each(func(i int, s *goquery.Selection) {
			style, ok := s.Attr("style")
			matches := urlRe.FindStringSubmatch(style)
			if ok {
				href := strings.Trim(matches[1], `"'`)
				photos = append(photos, href)
			}
		})

		animal := Animal{
			ID:     id,
			Text:   text,
			Photos: photos,
		}
		photos = nil
		count++
		animals = append(animals, animal)
	})

	return animals, nil
}

func remove(slice []Animal, s int) []Animal {
	return append(slice[:s], slice[s+1:]...)
}
