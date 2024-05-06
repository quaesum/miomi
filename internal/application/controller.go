package application

import (
	"errors"
	"madmax/internal/entity"
	"math"
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
