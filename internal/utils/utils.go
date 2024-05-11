package utils

import (
	"errors"
	"github.com/samber/lo"
	"madmax/internal/entity"
	"math"
	"strings"
)

func GetRecordsOnCurrentPage(req entity.SearchRequest, length int) (int, int, error) {
	page := req.Page
	perPage := req.PerPage
	page -= 1
	if page < 0 {
		return 0, 0, errors.New("page must be greater than 0")
	}

	leftBorder := page * perPage
	rightBorder := page*perPage + perPage
	if rightBorder > length {
		rightBorder = length
	}

	return leftBorder, rightBorder, nil
}

func GetMaxPages(length int, perPage int) (int8, error) {
	pages := math.Ceil(float64(length) / float64(perPage))
	return int8(pages), nil
}

func CleanQuery(query string) string {
	words := strings.Fields(query)

	cleanedWords := lo.Filter(words, func(item string, _ int) bool {
		return len(item) > 2
	})

	cleanedQuery := strings.Join(cleanedWords, " ")
	query = strings.TrimSpace(cleanedQuery)

	return cleanedQuery
}

func ProcessPhotos(photos interface{}) ([]string, error) {
	var photosOut []string
	switch v := photos.(type) {
	case string:
		photosOut = append(photosOut, v)
		return photosOut, nil
	case []interface{}:
		for _, photo := range v {
			if str, ok := photo.(string); ok {
				photosOut = append(photosOut, str)
			}
		}
		return photosOut, nil
	default:
		return nil, errors.New("invalid photo type")
	}
}
